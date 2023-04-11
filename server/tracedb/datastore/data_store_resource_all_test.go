package datastore_test

import (
	"testing"

	"github.com/gorilla/mux"
	"github.com/kubeshop/tracetest/server/tracedb/datastore"
	"github.com/kubeshop/tracetest/server/id"
	"github.com/kubeshop/tracetest/server/resourcemanager"
	rmtests "github.com/kubeshop/tracetest/server/resourcemanager/testutil"
	"github.com/kubeshop/tracetest/server/testmock"
)

func TestDataStoreResource(t *testing.T) {
	db := testmock.MustGetRawTestingDatabase()

	testSpec := rmtests.ResourceTypeTest{
		ResourceType: "DataStore",
		RegisterManagerFn: func(router *mux.Router) resourcemanager.Manager {
			db := testmock.MustCreateRandomMigratedDatabase(db)
			dataStoreRepo := datastore.NewRepository(db)

			manager := resourcemanager.New[datastore.DataStore](
				"DataStore",
				dataStoreRepo,
				resourcemanager.WithOperations(datastore.Operations...),
				resourcemanager.WithIDGen(id.GenerateID),
			)
			manager.RegisterRoutes(router)

			return manager
		},
		SampleJSON: `{
			"type": "DataStore",
			"spec": {
				"id": "current",
				"name": "default",
				"default": true,
			}
		}`,
		SampleJSONUpdated: `{
			"type": "DataStore",
			"spec": {
				"id": "current",
				"name": "long test",
				"default": true,
			}
		}`,
	}

	excludedOperations := rmtests.ExcludeOperations(rmtests.OperationGetNotFound, rmtests.OperationUpdateNotFound)

	rmtests.TestResourceType(t, testSpec, excludedOperations)
}

