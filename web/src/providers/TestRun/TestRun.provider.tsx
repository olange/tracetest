import {noop} from 'lodash';
import {createContext, useCallback, useContext, useEffect, useMemo, useState} from 'react';
import {useGetRunByIdQuery, useGetRunEventsQuery, useStopRunMutation} from 'redux/apis/TraceTest.api';
import {TLintBySpan} from 'models/LinterResult.model';
import TestRun, {isRunStateFinished} from 'models/TestRun.model';
import TestRunEvent from 'models/TestRunEvent.model';
import SpanService from 'services/Span.service';
import TestProvider from '../Test';

interface IContext {
  run: TestRun;
  isError: boolean;
  isLoadingStop: boolean;
  runEvents: TestRunEvent[];
  stopRun(): void;
  runLinterResultsBySpan: TLintBySpan;
}

export const Context = createContext<IContext>({
  run: {} as TestRun,
  isError: false,
  isLoadingStop: false,
  runEvents: [],
  stopRun: noop,
  runLinterResultsBySpan: {},
});

interface IProps {
  testId: string;
  runId?: string;
  children: React.ReactNode;
}

export const useTestRun = () => useContext(Context);

const POLLING_INTERVAL = 5000;

const TestRunProvider = ({children, testId, runId = ''}: IProps) => {
  const [pollingInterval, setPollingInterval] = useState<number | undefined>(POLLING_INTERVAL);
  const {data: run, isError} = useGetRunByIdQuery({testId, runId}, {skip: !runId, pollingInterval});
  const {data: runEvents = []} = useGetRunEventsQuery({testId, runId}, {skip: !runId});
  const [stopRunAction, {isLoading: isLoadingStop}] = useStopRunMutation();
  const runLinterResultsBySpan = useMemo(
    () => (run?.linter ? SpanService.getLintBySpan(run.linter) : {}),
    [run?.linter]
  );

  const stopRun = useCallback(async () => {
    await stopRunAction({runId, testId});
  }, [runId, stopRunAction, testId]);

  const value = useMemo<IContext>(
    () => ({run: run!, isError, isLoadingStop, runEvents, stopRun, runLinterResultsBySpan}),
    [run, isError, isLoadingStop, runEvents, stopRun, runLinterResultsBySpan]
  );

  useEffect(() => {
    const shouldStopPolling = run?.state && isRunStateFinished(run.state);
    setPollingInterval(shouldStopPolling ? undefined : POLLING_INTERVAL);
  }, [run?.state]);

  return run ? (
    <Context.Provider value={value}>
      <TestProvider testId={testId} version={run.testVersion}>
        {children}
      </TestProvider>
    </Context.Provider>
  ) : (
    <div data-cy="loading_test_run" />
  );
};

export default TestRunProvider;
