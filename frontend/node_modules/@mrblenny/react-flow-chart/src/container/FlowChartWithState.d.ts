import * as React from 'react';
import { IChart, IConfig, IFlowChartComponents } from '../';
export interface IFlowChartWithStateProps {
    initialValue: IChart;
    Components?: IFlowChartComponents;
    config?: IConfig;
}
/**
 * Flow Chart With State
 */
export declare class FlowChartWithState extends React.Component<IFlowChartWithStateProps, IChart> {
    state: IChart;
    private stateActions;
    constructor(props: IFlowChartWithStateProps);
    render(): JSX.Element;
}
