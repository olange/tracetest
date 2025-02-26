package executor_test

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/kubeshop/tracetest/server/environment"
	"github.com/kubeshop/tracetest/server/executor"
	"github.com/kubeshop/tracetest/server/model"
	"github.com/kubeshop/tracetest/server/pkg/id"
	"github.com/kubeshop/tracetest/server/pkg/maps"
	"github.com/kubeshop/tracetest/server/subscription"
	"github.com/kubeshop/tracetest/server/testdb"
	"github.com/kubeshop/tracetest/server/testmock"
	"github.com/kubeshop/tracetest/server/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type fakeTestRunner struct {
	db                  model.Repository
	subscriptionManager *subscription.Manager
	returnErr           bool
	uid                 int
}

func (r *fakeTestRunner) Run(ctx context.Context, test model.Test, metadata model.RunMetadata, env environment.Environment) model.Run {
	run := model.NewRun()
	run.Environment = env
	run.State = model.RunStateCreated
	newRun, err := r.db.CreateRun(ctx, test, run)
	if err != nil {
		panic(err)
	}

	go func() {
		run := newRun                      // make a local copy to avoid race conditions
		time.Sleep(100 * time.Millisecond) // simulate some real work

		if r.returnErr {
			run.State = model.RunStateTriggerFailed
			run.LastError = fmt.Errorf("failed to do something")
		} else {
			run.State = model.RunStateFinished
		}

		r.uid++

		run.Outputs = (maps.Ordered[string, model.RunOutput]{}).MustAdd("USER_ID", model.RunOutput{
			Value: strconv.Itoa(r.uid),
		})

		err = r.db.UpdateRun(ctx, run)
		r.subscriptionManager.PublishUpdate(subscription.Message{
			ResourceID: newRun.ResourceID(),
			Type:       "result_update",
			Content:    run,
		})
	}()

	return newRun
}

func TestTransactionRunner(t *testing.T) {

	t.Run("NoErrors", func(t *testing.T) {
		runTransactionRunnerTest(t, false, func(t *testing.T, actual tests.TransactionRun) {
			assert.Equal(t, tests.TransactionRunStateFinished, actual.State)
			assert.Len(t, actual.Steps, 2)
			assert.Equal(t, actual.Steps[0].State, model.RunStateFinished)
			assert.Equal(t, actual.Steps[1].State, model.RunStateFinished)
			assert.Equal(t, "http://my-service.com", actual.Environment.Get("url"))

			assert.Equal(t, model.RunOutput{Name: "", Value: "1", SpanID: ""}, actual.Steps[0].Outputs.Get("USER_ID"))

			// this assertion is supposed to test that the output from the previous step
			// is injected in the env for the next. In practice, this depends
			// on the `fakeTestRunner` used here to actually save the environment
			// to the test run, like the real test runner would.
			// see line 27
			assert.Equal(t, "1", actual.Steps[1].Environment.Get("USER_ID"))
			assert.Equal(t, model.RunOutput{Name: "", Value: "2", SpanID: ""}, actual.Steps[1].Outputs.Get("USER_ID"))

			assert.Equal(t, "2", actual.Environment.Get("USER_ID"))

		})
	})

	t.Run("WithErrors", func(t *testing.T) {
		runTransactionRunnerTest(t, true, func(t *testing.T, actual tests.TransactionRun) {
			assert.Equal(t, tests.TransactionRunStateFailed, actual.State)
			require.Len(t, actual.Steps, 1)
			assert.Equal(t, model.RunStateTriggerFailed, actual.Steps[0].State)
		})
	})

}

func getDB() (model.Repository, *sql.DB) {
	rawDB := testmock.GetRawTestingDatabase()
	db := testmock.GetTestingDatabaseFromRawDB(rawDB)

	return db, rawDB
}

func runTransactionRunnerTest(t *testing.T, withErrors bool, assert func(t *testing.T, actual tests.TransactionRun)) {
	ctx := context.Background()
	db, rawDB := getDB()

	subscriptionManager := subscription.NewManager()

	testRunner := &fakeTestRunner{
		db,
		subscriptionManager,
		withErrors,
		0,
	}

	test1, err := db.CreateTest(ctx, model.Test{Name: "Test 1"})
	require.NoError(t, err)

	test2, err := db.CreateTest(ctx, model.Test{Name: "Test 2"})
	require.NoError(t, err)

	testsRepo, _ := testdb.Postgres(testdb.WithDB(rawDB))
	transactionsRepo := tests.NewTransactionsRepository(rawDB, testsRepo)
	transaction, err := transactionsRepo.Create(ctx, tests.Transaction{
		Name:    "transaction",
		StepIDs: []id.ID{test1.ID, test2.ID},
	})
	require.NoError(t, err)

	transaction, err = transactionsRepo.GetAugmented(context.TODO(), transaction.ID)
	require.NoError(t, err)

	metadata := model.RunMetadata{
		"environment": "production",
		"service":     "tracetest",
	}

	envRepository := environment.NewRepository(rawDB)
	env, err := envRepository.Create(ctx, environment.Environment{
		Name: "production",
		Values: []environment.EnvironmentValue{
			{
				Key:   "url",
				Value: "http://my-service.com",
			},
		},
	})
	require.NoError(t, err)

	runner := executor.NewTransactionRunner(testRunner, db, transactionsRepo, subscriptionManager)
	runner.Start(1)

	ctxWithTimeout, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	transactionRun := runner.Run(ctxWithTimeout, transaction, metadata, env)

	done := make(chan tests.TransactionRun, 1)
	sf := subscription.NewSubscriberFunction(func(m subscription.Message) error {
		tr := m.Content.(tests.TransactionRun)
		if tr.State.IsFinal() {
			done <- tr
		}

		return nil
	})
	subscriptionManager.Subscribe(transactionRun.ResourceID(), sf)

	var finalRun tests.TransactionRun
	select {
	case finalRun = <-done:
		subscriptionManager.Unsubscribe(transactionRun.ResourceID(), sf.ID()) //cleanup to avoid race conditions
		fmt.Println("run completed")
	case <-time.After(10 * time.Second):
		t.Log("timeout after 10 second")
		t.FailNow()
	}

	assert(t, finalRun)
}
