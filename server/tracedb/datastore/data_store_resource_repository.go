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
	resourcemanager.OperationCreate,
	resourcemanager.OperationDelete,
	resourcemanager.OperationGet,
	resourcemanager.OperationList,
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

const insertQuery = `
INSERT INTO data_stores (
	"id",
	"name",
	"type",
	"is_default",
	"values",
	"created_at"
) VALUES ($1, $2, $3, $4, $5, $6)`

func (r *Repository) Create(ctx context.Context, dataStore DataStore) (DataStore, error) {
	dataStore.ID = IDGen.ID().String()
	dataStore.CreatedAt = time.Now()

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return DataStore{}, err
	}
	defer tx.Rollback()

	valuesJSON, err := json.Marshal(dataStore.Values)
	if err != nil {
		return DataStore{}, fmt.Errorf("could not marshal values field configuration: %w", err)
	}

	_, err = tx.ExecContext(ctx, insertQuery,
		dataStore.ID,
		dataStore.Name,
		dataStore.Type,
		dataStore.Default,
		valuesJSON,
		dataStore.CreatedAt,
	)
	if err != nil {
		return DataStore{}, fmt.Errorf("sql exec insert: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return DataStore{}, fmt.Errorf("commit: %w", err)
	}

	return dataStore, nil
}

// on this query, we don't update created_at
const updateQuery = `
UPDATE data_stores SET
	"name" = $2,
	"type" = $3,
	"is_default" = $4,
	"values" = $5
WHERE id = $1
`

func (r *Repository) Update(ctx context.Context, dataStore DataStore) (DataStore, error) {
	oldDataStore, err := r.Get(ctx, dataStore.ID)
	if err != nil {
		return DataStore{}, err
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return DataStore{}, err
	}
	defer tx.Rollback()

	valuesJSON, err := json.Marshal(dataStore.Values)
	if err != nil {
		return DataStore{}, fmt.Errorf("could not marshal values field configuration: %w", err)
	}

	_, err = tx.ExecContext(ctx, updateQuery,
		oldDataStore.ID,
		dataStore.ID,
		dataStore.Name,
		dataStore.Type,
		dataStore.Default,
		valuesJSON,
	)
	if err != nil {
		return DataStore{}, fmt.Errorf("sql exec update: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return DataStore{}, fmt.Errorf("commit: %w", err)
	}

	return demo, nil
}

const deleteQuery = `DELETE FROM data_stores WHERE "id" = $1`

func (r *Repository) Delete(ctx context.Context, id id.ID) error {
	dataStore, err := r.Get(ctx, id)
	if err != nil {
		return err
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, deleteQuery, dataStore.ID)

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("commit: %w", err)
	}

	return nil
}

const baseGetQuery = `
SELECT
	"id",
	"name",
	"type",
	"values",
	"created_at"
FROM data_stores`

const getByIdQuery = baseGetQuery + `WHERE "id" = $1`

func (r *Repository) Get(ctx context.Context, id id.ID) (DataStore, error) {
	var valuesJSON []byte
	err := r.db.
		QueryRowContext(ctx, getByIdQuery, id).
		Scan(
			&dataStore.ID,
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

func (r Repository) SortingFields() []string {
	return []string{"id", "name", "type", "created_at"}
}

func (r *Repository) getParsedValues(dsType DataStoreType, valuesJSON []byte) (DataStoreValues, error) {
	// TODO: Marshal logic depending on type
	return DataStoreValues{}, nil
}

func (r *Repository) Provision(ctx context.Context, dataStore DataStore) error {
	_, err := r.Update(ctx, dataStore)
	return err
}
