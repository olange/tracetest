package datastore_test

import (
	"context"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/kubeshop/tracetest/server/id"
	"github.com/kubeshop/tracetest/server/resourcemanager"
	rmtests "github.com/kubeshop/tracetest/server/resourcemanager/testutil"
	"github.com/kubeshop/tracetest/server/testmock"
	"github.com/kubeshop/tracetest/server/tracedb/datastore"
)

func TestDataStoreResource(t *testing.T) {
	db := testmock.MustGetRawTestingDatabase()

	sampleDataStore := datastore.DataStore{
		ID:        "1",
		Name:      "default",
		Type:      datastore.DataStoreTypeOTLP,
		Default:   true,
		Values:    datastore.DataStoreValues{},
		CreatedAt: time.Now(),
	}

	secondSampleDataStore := datastore.DataStore{
		ID:        "2",
		Name:      "default 2",
		Type:      datastore.DataStoreTypeOTLP,
		Default:   true,
		Values:    datastore.DataStoreValues{},
		CreatedAt: time.Now().Add(2 * time.Minute),
	}

	thirdSampleDataStore := datastore.DataStore{
		ID:        "3",
		Name:      "default 3",
		Type:      datastore.DataStoreTypeOTLP,
		Default:   true,
		Values:    datastore.DataStoreValues{},
		CreatedAt: time.Now().Add(3 * time.Minute),
	}

	testSpec := rmtests.ResourceTypeTest{
		ResourceType: "DataStore",
		RegisterManagerFn: func(router *mux.Router) resourcemanager.Manager {
			db := testmock.MustCreateRandomMigratedDatabase(db)
			dataStoreRepository := datastore.NewRepository(db)

			manager := resourcemanager.New[datastore.DataStore](
				"DataStore",
				dataStoreRepository,
				resourcemanager.WithOperations(datastore.Operations...),
				resourcemanager.WithIDGen(id.GenerateID),
			)
			manager.RegisterRoutes(router)

			return manager
		},
		Prepare: func(t *testing.T, op rmtests.Operation, manager resourcemanager.Manager) {
			dataStoreRepository := manager.Handler().(*datastore.Repository)
			switch op {
			case rmtests.OperationGetSuccess,
				rmtests.OperationUpdateSuccess,
				rmtests.OperationDeleteSuccess,
				rmtests.OperationListSuccess:
				dataStoreRepository.Create(context.TODO(), sampleDataStore)
			case rmtests.OperationListPaginatedSuccess:
				dataStoreRepository.Create(context.TODO(), sampleDataStore)
				dataStoreRepository.Create(context.TODO(), secondSampleDataStore)
				dataStoreRepository.Create(context.TODO(), thirdSampleDataStore)
			}
		},
		SampleJSON: `{
			"type": "DataStore",
			"spec": {
				"id": "1",
				"name": "default",
				"type": "otlp",
				"default": true,
				"values": {}
			}
		}`,
		SampleJSONUpdated: `{
			"type": "DataStore",
			"spec": {
				"id": "1",
				"name": "another data store",
				"type": "otlp",
				"default": true,
				"values": {}
			}
		}`,
	}

	excludedOperations := rmtests.ExcludeOperations(rmtests.OperationGetNotFound, rmtests.OperationUpdateNotFound)

	rmtests.TestResourceType(t, testSpec, excludedOperations)
}
