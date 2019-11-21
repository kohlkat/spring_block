import React, { Component } from 'react';
import {Page, Panel, Table, TableHead, TableRow, TableBody} from 'react-blur-admin';
import { Row, Col } from 'react-flex-proto';
import * as am4core from "@amcharts/amcharts4/core";
import * as am4charts from "@amcharts/amcharts4/charts";
import am4themes_animated from "@amcharts/amcharts4/themes/animated";
import axios from 'axios';

export class Home extends Component {

    componentDidMount() {

        console.log("Connected")
        // TEST
        axios.get('http://localhost:8080/connect').then((res) => {
            console.log(res)
            
        })

        /* Chart code */
        // Themes begin
        am4core.useTheme(am4themes_animated);
        // Themes end

        let chart = am4core.create("chartdiv", am4charts.ChordDiagram);

        chart.data = [
            { from: "XRP", to: "ULT", value: 139 },
            { from: "ULT", to: "CNY", value: 1501 },
            { from: "CNY", to: "XRP", value: 310 },
            { from: "XRP", to: "ETH", value: 139 },
        ];

        chart.dataFields.fromName = "from";
        chart.dataFields.toName = "to";
        chart.dataFields.value = "value";

        // make nodes draggable
        let nodeTemplate = chart.nodes.template;
        nodeTemplate.readerTitle = "Click to show/hide or drag to rearrange";
        nodeTemplate.showSystemTooltip = true;

        let nodeLink = chart.links.template;
        let bullet = nodeLink.bullets.push(new am4charts.CircleBullet());
        bullet.fillOpacity = 1;
        bullet.circle.radius = 5;
        bullet.locationX = 0.5;

        // create animations
        chart.events.on("ready", function() {
            for (var i = 0; i < chart.links.length; i++) {
                let link = chart.links.getIndex(i);
                let bullet = link.bullets.getIndex(0);
                animateBullet(bullet);
            }
        })

        function animateBullet(bullet) {
            let duration = 1000 * Math.random() + 800;
            let animation = bullet.animate([{ property: "locationX", from: 0, to: 1 }], duration)
            animation.events.on("animationended", function(event) {
                animateBullet(event.target.object);
            })
        }

    }

    componentWillUnmount() {
        if (this.chart) {
            this.chart.dispose();
        }
    }

    render() {
        return (
            <Page>
                <Row>
                    <Col align="center" padding={5} >
                        <Panel title='Tx Completed'>
                            <div style={{margin:"1.8rem"}}>
                                    <p style={{fontSize:"18px"}}>24h: 18</p>
                                <p style={{fontSize:"18px"}}>All-time: 186</p>
                            </div>
                        </Panel>
                    </Col>
                    <Col align="center" padding={5}>
                        <Panel title='Latest Tx' >
                            <div style={{margin:"2rem"}}>
                                <p style={{fontSize:"18px"}}>17/11/2019</p>
                                <p style={{fontSize:"18px"}}>22:48</p>
                            </div>
                        </Panel>
                    </Col>
                    <Col align="center" padding={5}>
                        <Panel title='XRP Balance'>
                            <div style={{margin:"2rem"}}>
                                <p style={{fontSize:"18px"}}>1138 XRP</p>
                                <p style={{fontSize:"14px"}}>($13190 USD)</p>
                            </div>
                        </Panel>
                    </Col>
                    <Col align="center" padding={5}>
                        <Panel title='XRP Accumulated'>
                            <div style={{margin:"2rem"}}>
                                <p style={{fontSize:"18px"}}>113 XRP</p>
                                <p style={{fontSize:"14px"}}>($1319 USD)</p>
                            </div>
                        </Panel>
                    </Col>
                </Row>
                <Row>
                    <Col align="center">
                        <h2>Latest Transaction</h2>
                    </Col>
                </Row>
                <Row>
                    <div className="align-center" style={{width:"100%"}}>
                    <Panel title={`Cycle Details`} >

                    <div className="col-lg-6">
                            <div id="chartdiv" style={{ width: "100%", height: "620px" }}></div>
                    </div>
                        <div className="col-lg-6">
                            <Table>
                                <TableHead>
                                    <th>Step</th>
                                    <th>Sent</th>
                                    <th>Received</th>
                                    <th>Exchange Rate</th>
                                    <th>Issued By</th>
                                    <th>Timestamp</th>
                                </TableHead>
                                <TableBody>
                                    <TableRow>
                                        <td>1</td>
                                        <td>139.359 XRP</td>
                                        <td>1501.834 ULT</td>
                                        <td>10.443</td>
                                        <td><a style={{color:"#00d1b2"}} href="xrpscan.com" >DA011DAF4537EAE5E...58898566238CFF86</a></td>
                                        <td>2019-11-21 (11 : 42)</td>
                                    </TableRow>
                                    <TableRow>
                                        <td>2</td>
                                        <td>1501.834 ULT</td>
                                        <td>310.3489 CNY</td>
                                        <td>0.310</td>
                                        <td><a style={{color:"#00d1b2"}} href="xrpscan.com">DA011DAF4537EAE5E...58898566238CFF86</a></td>
                                        <td>2019-11-21 (11 : 42)</td>
                                    </TableRow>
                                    <TableRow>
                                        <td>3</td>
                                        <td>310.3489 CNY</td>
                                        <td>139.390 XRP</td>
                                        <td>0.54</td>
                                        <td><a style={{color:"#00d1b2"}} href="xrpscan.com">DA011DAF4537EAE5E...58898566238CFF86</a></td>
                                        <td>2019-11-21 (11 : 42)</td>
                                    </TableRow>
                                    <TableRow>
                                        <td>4</td>
                                        <td>139.390 XRP</td>
                                        <td>0.139654 ETH</td>
                                        <td>0.00140</td>
                                        <td><a style={{color:"#00d1b2"}} href="xrpscan.com">DA011DAF4537EAE5E...58898566238CFF86</a></td>
                                        <td>2019-11-21 (11 : 42)</td>
                                    </TableRow>
                                </TableBody>
                            </Table>
                            <div style={{marginTop:"12px"}}>Gap Size: 0.32 XRP</div>
                        </div>
                    </Panel>
                    </div>
                </Row>
            </Page>
        );
    }
}