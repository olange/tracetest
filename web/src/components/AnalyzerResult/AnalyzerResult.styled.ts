import {CheckCircleFilled, CloseCircleFilled, DownOutlined, UpOutlined} from '@ant-design/icons';
import {Button, Collapse, Progress, Typography} from 'antd';
import styled from 'styled-components';
import noResultsIcon from 'assets/SpanAssertionsEmptyState.svg';

export const StyledCollapse = styled(Collapse)`
  background-color: ${({theme}) => theme.color.white};
  border: 0;
`;

export const Container = styled.div`
  padding: 24px;
`;

export const Title = styled(Typography.Title)`
  && {
    margin-bottom: 8px;
    display: flex;
    align-items: center;
  }
`;

export const Description = styled(Typography.Paragraph).attrs({
  type: 'secondary',
})`
  && {
    margin-bottom: 30px;
  }
`;

export const GlobalResultWrapper = styled.div`
  display: grid;
  grid-template-columns: auto 1fr;
  margin-bottom: 28px;
  gap: 45px;
`;

export const GlobalScoreWrapper = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
`;

export const ScoreResultWrapper = styled(GlobalScoreWrapper)`
  align-items: flex-start;
`;

export const GlobalScoreContainer = styled.div`
  display: flex;
  flex-direction: row;
  justify-content: center;
`;

export const RuleContainer = styled.div`
  border-bottom: ${({theme}) => `1px dashed ${theme.color.borderLight}`};
  padding-bottom: 16px;
  margin-bottom: 16px;
  margin-left: 43px;
`;

export const RuleHeader = styled.div`
  display: flex;
  flex-direction: row;
  justify-content: space-between;
`;

export const Column = styled.div`
  display: flex;
  flex-direction: column;
  margin-bottom: 8px;
`;

export const RuleBody = styled(Column)`
  padding-left: 20px;
`;

export const Subtitle = styled(Typography.Title)`
  && {
    margin-bottom: 8px;
  }
`;

export const ResultText = styled(Typography.Text)<{$passed: boolean}>`
  && {
    color: ${({theme, $passed}) => ($passed ? theme.color.success : theme.color.error)};
  }
`;

export const ScoreProgress = styled(Progress)`
  .ant-progress-inner {
    height: 50px !important;
    width: 50px !important;
  }

  .ant-progress-circle-trail,
  .ant-progress-circle-path {
    stroke-width: 20px;
  }
`;

export const PluginPanel = styled(Collapse.Panel)`
  background-color: ${({theme}) => theme.color.white};
  border: ${({theme}) => `1px solid ${theme.color.border}`};
  margin-bottom: 12px;

  .ant-collapse-content {
    background-color: ${({theme}) => theme.color.background};
  }
`;

export const PassedIcon = styled(CheckCircleFilled)<{$small?: boolean}>`
  color: ${({theme}) => theme.color.success};
  font-size: ${({$small}) => ($small ? '14px' : '20px')};
`;

export const FailedIcon = styled(CloseCircleFilled)<{$small?: boolean}>`
  color: ${({theme}) => theme.color.error};
  font-size: ${({$small}) => ($small ? '14px' : '20px')};
`;

export const SpanButton = styled(Button)<{$error?: boolean}>`
  color: ${({theme, $error}) => ($error ? theme.color.error : theme.color.success)};
  padding-left: 0;
`;

export const CollapseIconContainer = styled.div`
  display: flex;
  position: absolute;
  top: 25%;
  right: 16px;
  border-left: 1px solid ${({theme}) => theme.color.borderLight};
  padding-left: 14px;
  height: 24px;
  align-items: center;
`;

export const DownCollapseIcon = styled(DownOutlined)`
  opacity: 0.5;
  font-size: ${({theme}) => theme.size.xs};
`;

export const UpCollapseIcon = styled(UpOutlined)`
  opacity: 0.5;
  font-size: ${({theme}) => theme.size.xs};
`;

export const EmptyContainer = styled.div`
  align-items: center;
  display: flex;
  flex-direction: column;
  height: calc(100% - 70px);
  justify-content: center;
  margin-top: 50px;
`;

export const EmptyIcon = styled.img.attrs({
  src: noResultsIcon,
})`
  height: auto;
  margin-bottom: 16px;
  width: 90px;
`;

export const EmptyText = styled(Typography.Text)`
  color: ${({theme}) => theme.color.textSecondary};
`;

export const EmptyTitle = styled(Typography.Title).attrs({level: 3})``;

export const ConfigureButtonContainer = styled.div`
  margin-top: 6px;
`;

export const SwitchContainer = styled.div`
  align-items: center;
  display: flex;
  gap: 8px;
  justify-content: flex-end;
  margin-bottom: 16px;
`;

export const List = styled.ul`
  padding-inline-start: 20px;
`;
