import {useParams} from 'react-router-dom';

import Layout from 'components/Layout';
import TestSpecFormProvider from 'components/TestSpecForm/TestSpecForm.provider';
import withAnalytics from 'components/WithAnalytics/WithAnalytics';
import GuidedTourProvider from 'providers/GuidedTour/GuidedTour.provider';
import SpanProvider from 'providers/Span';
import TestSpecsProvider from 'providers/TestSpecs';
import TestRunProvider from 'providers/TestRun';
import Content from './Content';

const RunDetail = () => {
  const {testId = '', runId = ''} = useParams();

  return (
    <GuidedTourProvider>
      <Layout>
        <TestRunProvider testId={testId} runId={runId}>
          <TestSpecsProvider testId={testId} runId={runId}>
            <TestSpecFormProvider testId={testId}>
              <SpanProvider>
                <Content />
              </SpanProvider>
            </TestSpecFormProvider>
          </TestSpecsProvider>
        </TestRunProvider>
      </Layout>
    </GuidedTourProvider>
  );
};

export default withAnalytics(RunDetail, 'run-detail');