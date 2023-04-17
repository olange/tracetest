import {PromiseWithKnownReason} from '@reduxjs/toolkit/dist/query/core/buildMiddleware/types';
import {noop} from 'lodash';
import webSocketGateway, {IListenerFunction} from 'gateways/WebSocket.gateway';

export type {IListenerFunction} from 'gateways/WebSocket.gateway';

interface IInitWebSocketSubscription {
  listener: IListenerFunction;
  resource: string;
  onInit?: () => void;
  waitToCleanSubscription: Promise<void>;
  waitToInitSubscription: PromiseWithKnownReason<any, any>;
}

const WebSocketService = () => ({
  async initWebSocketSubscription({
    listener,
    resource,
    waitToCleanSubscription,
    waitToInitSubscription,
    onInit = noop,
  }: IInitWebSocketSubscription) {
    try {
      await waitToInitSubscription;
      onInit();
      webSocketGateway.subscribe(resource, listener);
    } catch {
      // no-op in case `waitToCleanSubscription` resolves before `waitToInitSubscription`,
      // in which case `waitToInitSubscription` will throw
    }
    await waitToCleanSubscription;
    webSocketGateway.unsubscribe(resource);
  },
});

export default WebSocketService();
