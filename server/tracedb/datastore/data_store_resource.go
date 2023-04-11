package datastore

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/kubeshop/tracetest/server/id"
	"github.com/kubeshop/tracetest/server/resourcemanager"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/config/configtls"
	"golang.org/x/exp/slices"
)

const ResourceName = "DataStore"

var Operations = []resourcemanager.Operation{
	resourcemanager.OperationGet,
	resourcemanager.OperationUpdate,
}

var DefaultDataStore = DataStore{
	ID:       id.ID("current"),
	Name:     "default",
	Default:   true,
	CreatedAt: time.Now().UTC(),
}

type DataStoreType string

type DataStore struct {
	ID         id.ID                  `mapstructure:"id"`
	Name       string                 `mapstructure:"name"`
	Type       DataStoreType          `mapstructure:"type"`
	Default    bool                   `mapstructure:"default"`
	Values     DataStoreValues        `mapstructure:"values"`
	CreatedAt  time.Time              `mapstructure:"created_at"`
}

type DataStoreValues struct {
	Jaeger     *configgrpc.GRPCClientSettings
	Tempo      *BaseClientConfig
	OpenSearch *ElasticSearchDataStoreConfig
	ElasticApm *ElasticSearchDataStoreConfig
	SignalFx   *SignalFXDataStoreConfig
	AwsXRay    *AWSXRayDataStoreConfig
}

type BaseClientConfig struct {
	Type string
	Grpc configgrpc.GRPCClientSettings
	Http HttpClientConfig
}

type HttpClientConfig struct {
	Url        string
	Headers    map[string]string
	TLSSetting configtls.TLSClientSetting `mapstructure:"tls"`
}

type OTELCollectorConfig struct {
	Endpoint string
}

type ElasticSearchDataStoreConfig struct {
	Addresses          []string
	Username           string
	Password           string
	Index              string
	Certificate        string
	InsecureSkipVerify bool
}

type SignalFXDataStoreConfig struct {
	Realm string
	Token string
}

type AWSXRayDataStoreConfig struct {
	Region          string
	AccessKeyID     string
	SecretAccessKey string
	SessionToken    string
	UseDefaultAuth  bool
}

const (
	DataStoreTypeJaeger     DataStoreType = "jaeger"
	DataStoreTypeTempo      DataStoreType = "tempo"
	DataStoreTypeOpenSearch DataStoreType = "opensearch"
	DataStoreTypeSignalFX   DataStoreType = "signalfx"
	DataStoreTypeOTLP       DataStoreType = "otlp"
	DataStoreTypeNewRelic   DataStoreType = "newrelic"
	DataStoreTypeLighStep   DataStoreType = "lightstep"
	DataStoreTypeElasticAPM DataStoreType = "elasticapm"
	DataStoreTypeDataDog    DataStoreType = "datadog"
	DataStoreTypeAwsXRay    DataStoreType = "awsxray"
)

var validTypes = []DataStoreType{
	DataStoreTypeJaeger,
	DataStoreTypeTempo,
	DataStoreTypeOpenSearch,
	DataStoreTypeSignalFX,
	DataStoreTypeOTLP,
	DataStoreTypeNewRelic,
	DataStoreTypeLighStep,
	DataStoreTypeElasticAPM,
	DataStoreTypeDataDog,
	DataStoreTypeAwsXRay,
}

var otlpBasedDataStores = []DataStoreType{
	DataStoreTypeOTLP,
	DataStoreTypeNewRelic,
	DataStoreTypeLighStep,
	DataStoreTypeDataDog,
}

func (ds DataStore) Validate() error {
	if ds.Type != "" {
		return fmt.Errorf("data store should have a type")
	}

	if !slices.Contains(validTypes, ds.Type) {
		return fmt.Errorf("unsupported data store type %s", ds.Type)
	}

	if ds.Name != "" {
		return fmt.Errorf("data store should have a name")
	}

	return nil
}

func (ds DataStore) HasID() bool {
	return ds.ID.String() != ""
}

func (ds DataStore) IsOTLPBasedProvider() bool {
	return slices.Contains(otlpBasedDataStores, ds.Type)
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db}
}

type Repository struct {
	db *sql.DB
}

func (r *Repository) SetID(dataStore DataStore, id id.ID) DataStore {
	dataStore.ID = id
	return dataStore
}

const (
	insertQuery = `
		INSERT INTO data_stores (
			"id",
			"name",
			"type",
			"is_default",
			"values",
			"created_at"
		) VALUES ($1, $2, $3, $4, $5, $6)`
	deleteQuery = `DELETE FROM data_stores`
)

func (r *Repository) Update(ctx context.Context, updated DataStore) (DataStore, error) {
	// enforce ID and default
	updated.ID = "current"
	updated.Default = true

	tx, err := r.db.BeginTx(ctx, nil)
	defer tx.Rollback()
	if err != nil {
		return DataStore{}, err
	}

	_, err = tx.ExecContext(ctx, deleteQuery)
	if err != nil {
		return DataStore{}, fmt.Errorf("sql exec delete: %w", err)
	}

	valuesJSON, err := json.Marshal(updated.Values)
	if err != nil {
		return DataStore{}, fmt.Errorf("could not marshal values field configuration: %w", err)
	}

	_, err = tx.ExecContext(ctx, insertQuery,
		updated.ID,
		updated.Name,
		updated.Type,
		updated.Default,
		valuesJSON,
		updated.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return DefaultDataStore, nil
		}

		return DataStore{}, fmt.Errorf("sql exec insert: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return DataStore{}, fmt.Errorf("commit: %w", err)
	}

	return updated, nil

}

const (
	getQuery = `
		SELECT
			"name",
			"type",
			"values",
			"created_at"
		FROM data_stores`
)

func (r *Repository) GetDefault(ctx context.Context) DataStore {
	dataStore, _ := r.Get(ctx, id.ID("current"))
	return dataStore
}
func (r *Repository) Get(ctx context.Context, id id.ID) (DataStore, error) {
	dataStore := DataStore{
		ID:      "current",
		Default: true,
	}

	var valuesJSON []byte
	err := r.db.
		QueryRowContext(ctx, getQuery).
		Scan(
			&dataStore.Name,
			&dataStore.Type,
			&valuesJSON,
			&dataStore.CreatedAt,
		)

	if err != nil {
		return DataStore{}, fmt.Errorf("sql query: %w", err)
	}

	if string(valuesJSON) != "null" {
		values, err := r.getParsedValues(dataStore.Type, valuesJSON)
		if err != nil {
			return DataStore{}, fmt.Errorf("unable to parse data store values: %w", err)
		}

		dataStore.Values = values
	}

	return dataStore, nil
}

func (r *Repository) getParsedValues(dsType DataStoreType, valuesJSON []byte) (DataStoreValues, error) {
	// TODO: Marshal logic depending on type
	return DataStoreValues{}, nil
}

func (r *Repository) Provision(ctx context.Context, dataStore DataStore) error {
	_, err := r.Update(ctx, dataStore)
	return err
}
