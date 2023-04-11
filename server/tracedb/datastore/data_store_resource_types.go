package datastore

import (
	"fmt"
	"time"

	"github.com/kubeshop/tracetest/server/id"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/config/configtls"
	"golang.org/x/exp/slices"
)

const ResourceName = "DataStore"

type DataStoreType string

type DataStore struct {
	ID         id.ID                  `mapstructure:"id"`
	Name       string                 `mapstructure:"name"`
	Type       DataStoreType          `mapstructure:"type"`
	Default    bool                   `mapstructure:"default"`
	Values     DataStoreValues        `mapstructure:"values"`
	CreatedAt  time.Time              `mapstructure:"created_at"`
}

var DefaultDataStore = DataStore{
	ID:       id.ID("current"),
	Name:     "default",
	Type:     DataStoreTypeOTLP,
	Default:  true,
	Values: 	DataStoreValues{},
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
