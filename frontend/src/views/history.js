import React, { Component } from "react"
import { Page, Panel, Table, TableHead, TableBody, TableRow, Pagination, } from 'react-blur-admin';
import {Row, Col} from 'react-flex-proto';
import {Doughnut, Line} from 'react-chartjs-2';

export class History extends Component {
    constructor(props) {
        super(props);
        this.state = {
            chromeVisits: 1000,
            currentPage: 1,
        };
    }

    onSetCurrentPage(value) {
        this.setState({currentPage: value});
    }

    render() {
        return (
            <Page>
                <div className="row" style={{margin:"1em"}}>
                <Panel title='Recent Transactions'>
                    <Table>
                        <TableHead>
                            <th>Exchange Pair (OfferCreate)</th>
                            <th>Cycle Size</th>
                            <th>Volume</th>
                            <th>Gap Size</th>
                            <th>Timestamp</th>
                        </TableHead>
                        <TableBody>
                            <TableRow>
                                <td>XRP/ETH </td>
                                <td>3 </td>
                                <td>129.049033 XRP</td>
                                <td>0.030340 XRP</td>
                                <td>17/11/2019 13:29</td>
                            </TableRow>
                            <TableRow>
                                <td>XRP/USD </td>
                                <td>2 </td>
                                <td>123.049033 XRP</td>
                                <td>0.030340 XRP</td>
                                <td>17/11/2019 13:22</td>
                            </TableRow>
                            <TableRow>
                                <td>ETH/GBP </td>
                                <td>4 </td>
                                <td>0.2049033 ETH</td>
                                <td>0.0130340 ETH</td>
                                <td>17/11/2019 12:45</td>
                            </TableRow>
                            <TableRow>
                                <td>ETH/GBP </td>
                                <td>4 </td>
                                <td>0.2049033 ETH</td>
                                <td>0.0130340 ETH</td>
                                <td>17/11/2019 12:15</td>
                            </TableRow>
                            <TableRow>
                                <td>ETH/GBP </td>
                                <td>4 </td>
                                <td>0.2049033 ETH</td>
                                <td>0.0130340 ETH</td>
                                <td>17/11/2019 12:05</td>
                            </TableRow>
                            <TableRow>
                                <td>ETH/GBP </td>
                                <td>4 </td>
                                <td>0.2049033 ETH</td>
                                <td>0.0130340 ETH</td>
                                <td>17/11/2019 11:15</td>
                            </TableRow>
                            <TableRow>
                                <td>ETH/GBP </td>
                                <td>4 </td>
                                <td>0.2049033 ETH</td>
                                <td>0.0130340 ETH</td>
                                <td>17/11/2019 11:04</td>
                            </TableRow>
                        </TableBody>
                    </Table>
                    <Row>
                        <Col align='center'>
                            <Pagination currentPage={Number(this.state.currentPage)} totalPages={5} onChange={value => this.onSetCurrentPage(value)} />
                        </Col>
                    </Row>
                </Panel>
            </div>
                <div className="row" style={{margin:"0.1em"}}>
                    <div className="col-lg-6">
                        <Panel title="Freqently Mispriced">
                            <Doughnut data={data} options={{legend: {
                                labels: {
                                fontColor: 'white'
                            }}}}/>
                        </Panel>
                    </div>
                    <div className="col-lg-6">
                        <Panel title="Gap Size (profit per tx)">
                            <Line data={lineData} options={options}/>
                        </Panel>
                    </div>
                </div>
            </Page>

        )
    }
}


const data = {
    labels: [
        'XRP/USD',
        'XRP/BTC',
        'BTC/CNY'
    ],
    datasets: [{
        data: [300, 50, 100],
        backgroundColor: [
            '#0e8174',
            '#005562',
            '#ffffff'
        ],
        hoverBackgroundColor: [
            '#0b6d62',
            '#004853',
            '#d8d8d8'
        ]
    }],
    legend: {
        labels: {
            fontColor: 'white'
        }
    }
};

const lineData = {
    labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July'],
    datasets: [
        {
            label: 'Gap Size (XRP)',
            color: "white",
            fill: false,
            lineTension: 0.1,
            backgroundColor: 'rgba(75,192,192,0.4)',
            borderColor: 'rgba(75,192,192,1)',
            borderCapStyle: 'butt',
            borderDash: [],
            borderDashOffset: 0.0,
            borderJoinStyle: 'miter',
            pointBorderColor: 'rgba(75,192,192,1)',
            pointBackgroundColor: '#fff',
            pointBorderWidth: 1,
            pointHoverRadius: 5,
            pointHoverBackgroundColor: 'rgba(75,192,192,1)',
            pointHoverBorderColor: 'rgba(220,220,220,1)',
            pointHoverBorderWidth: 2,
            pointRadius: 1,
            pointHitRadius: 10,
            data: [0.0825, 0.0625, 0.0725, 0.080, 0.058, 0.0525, 0.0525]
        }
    ],

};

const options = {
    legend: {
        labels: {
            fontColor: 'white'
        }
    },
    scales: {
        xAxes: [{
            ticks: {
                fontColor: "white",
            }
        }],
        yAxes: [{
            ticks: {
                fontColor: "white",
            }
        }]
    }
}