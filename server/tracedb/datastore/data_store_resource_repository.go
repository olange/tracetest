package datastore

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/kubeshop/tracetest/server/id"
	"github.com/kubeshop/tracetest/server/resourcemanager"
)

var Operations = []resourcemanager.Operation{
	// resourcemanager.OperationCreate,
	// resourcemanager.OperationDelete,
	resourcemanager.OperationGet,
	// resourcemanager.OperationList,
	resourcemanager.OperationUpdate,
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
	getQuery = `
		SELECT
			"name",
			"type",
			"values",
			"created_at"
		FROM data_stores`
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
