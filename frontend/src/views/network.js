import React, { Component } from "react"
import { Page, Panel, Table, TableHead, TableBody, TableRow, } from 'react-blur-admin';
import {Row} from "react-flex-proto";
import { xrpLedger } from "./api_requests"
import * as am4core from "@amcharts/amcharts4/core";
import * as am4maps from "@amcharts/amcharts4/maps";
import am4geodata_worldLow from "@amcharts/amcharts4-geodata/worldLow";
import am4themes_animated from "@amcharts/amcharts4/themes/animated";
import {getName} from "country-list"

export class Network extends Component {
    constructor(props) {
        super(props);
        this.state = {
            offers: [],
            currencies: [],
            node_count: [],
            issuers: {
                rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B : "Bitstamp",
                rKiCet8SdvWxPXnAgYarFUXMh1zCPz432Y : "ripplefox",
                rhub8VRN55s94qWKDv6jmDy1pUykJzF3wq : "gatehub",
                rchGBxcD1A1C2tdxF6papQYZ8kjRKMYcL : "GateHub Fifth BTC (Cold)",
                razqQKzJRdB4UxFPWf5NEpEG3WMkmwgcXA : "RippleChina",
                rcA8X3TVMST1n3CJeAdGk1RdRCHii7N2h : "GateHub Fifth"
            },
            num_nodes: "x",
        };
    }

    async componentDidMount() {

        let offersInterval = setInterval(async () => {
            let offers = await xrpLedger.getOffers(9);
            this.setState({ offers : offers.data.transactions });
            console.log(this.state.offers)
        }, 3000);
        let currencies = await xrpLedger.getTopCurrencies();

        // Get topology data
        let nodes = await xrpLedger.getTopology();
        nodes = nodes.data.nodes;
        let node_count = {};

        nodes.forEach((node) => {
            if (node.country_code in node_count) {
                node_count[node.country_code] += 1
            } else {
                node_count[node.country_code] = 1
            }
        })
        let num_nodes = nodes.length;
        this.setState({offersInterval, currencies : currencies.data.currencies, node_count, num_nodes});

        /* Chart code */
        // Themes begin
        am4core.useTheme(am4themes_animated);
        // Themes end

        let chart = am4core.create("chartdiv", am4maps.MapChart);

        // Set map definition
        chart.geodata = am4geodata_worldLow;

        // Set projection
        chart.projection = new am4maps.projections.Orthographic();
        chart.panBehavior = "rotateLongLat";
        chart.deltaLatitude = -20;
        chart.padding(20,20,20,20);

        // Create map polygon series
        let polygonSeries = chart.series.push(new am4maps.MapPolygonSeries());

        // Make map load polygon (like country names) data from GeoJSON
        polygonSeries.useGeodata = true;

        // Configure series
        let polygonTemplate = polygonSeries.mapPolygons.template;
        polygonTemplate.tooltipText = "{name}";
        polygonTemplate.fill = am4core.color("#ffffff");
        polygonTemplate.stroke = am4core.color("#454a58");
        polygonTemplate.strokeWidth = 0.5;

        let graticuleSeries = chart.series.push(new am4maps.GraticuleSeries());
        graticuleSeries.mapLines.template.line.stroke = am4core.color("#ffffff");
        graticuleSeries.mapLines.template.line.strokeOpacity = 0.08;
        graticuleSeries.fitExtent = false;


        chart.backgroundSeries.mapPolygons.template.polygon.fillOpacity = 0.1;
        chart.backgroundSeries.mapPolygons.template.polygon.fill = am4core.color("#ffffff");

        // Create hover state and set alternative fill color
        let hs = polygonTemplate.states.create("hover");
        hs.properties.fill = chart.colors.getIndex(0).brighten(-0.5);

        let animation;
        setTimeout(function(){
            animation = chart.animate({property:"deltaLongitude", to:100000}, 20000000);
        }, 3000);

        chart.seriesContainer.events.on("down", function(){
            if(animation){
                animation.stop();
            }
        });

        let topography_data = [];

        for (const node in this.state.node_count) {
            if (node !== undefined) {
                topography_data.push({
                    "id" : node,
                    "name" : getName(node),
                    "value" : this.state.node_count[node],
                    "fill": am4core.color("#209e91")
                })
            }
        }

        polygonSeries.data = topography_data;
        polygonTemplate.tooltipText = "{name}: {value}";
        polygonTemplate.propertyFields.fill = "fill";

        this.chart = chart

    }

    componentWillUnmount() {
        if (this.chart) {
            this.chart.dispose();
        }

        clearInterval(this.state.offersInterval);
    }

    // async t1OnSetCurrentPage(value) {
    //     let offers = await xrpLedger.getOffers(100, "")
    //     this.setState({t1CurrentPage: value, marker : offers.data.marker, offers : offers.data.transactions});
    // }

    convertDate(date) {
        date = new Date(date);
        let time = date.getHours() + " : " + date.getMinutes();
        let year = date.getFullYear();
        let month = date.getMonth()+1;
        let dt = date.getDate();

        if (dt < 10) {
            dt = '0' + dt;
        }
        if (month < 10) {
            month = '0' + month;
        }
        return `${year}-${month}-${+dt} (${time})`
    }

    abbreviateHash(hash) {
        return hash.substr(0,10) + "..." + hash.substr(-10)
    }

    getTableRows(offer) {
        let baseCur = offer.tx.TakerPays.currency || "XRP";
        let topCur = offer.tx.TakerGets.currency || "XRP";
        let baseVal = offer.tx.TakerGets.value || offer.tx.TakerGets / 1E6;
        let topVal = offer.tx.TakerPays.value || offer.tx.TakerPays / 1E6;
        let rate = ( topVal / baseVal).toFixed(4);

        return (
            <TableRow key={offer.hash}>
                    <td>{topCur}/{baseCur} </td>
                    <td>{rate}</td>
                    <td>{parseFloat(topVal).toFixed(2)} {topCur}</td>
                    <td>{parseFloat(baseVal).toFixed(2)} {baseCur}</td>
                    <td><a target="_blank" href={`https://xrpscan.com/tx/${offer.hash}`} style={{color: "#00d1b2"}} rel="noopener noreferrer"><abbr>{this.abbreviateHash(offer.hash)}</abbr></a></td>
                    <td>{offer.meta.AffectedNodes.length}</td>
                    <td>{this.convertDate(offer.date)}</td>
            </TableRow>
        )
    }

    render() {
        return (
            <Page>
                <div className="row" style={{margin:"0.1em"}}>
                    <div className="col-md-6">
                    <Panel title="Top Currencies (Past 24 hours)">
                        <Table>
                            <TableHead>
                                <th>Rank</th>
                                <th>Currency</th>
                                <th>Issued Value</th>
                                <th>Volume</th>
                                <th>Issuer</th>
                            </TableHead>
                            <TableBody>
                                {this.state.currencies.map((currency, i) => {
                                    return (
                                        <TableRow key={currency.avg_payment_volume}>
                                            <td>{i + 1}</td>
                                            <td>{currency.currency}</td>
                                            <td>{parseFloat(currency.issued_value.substr(0,10)).toFixed(2)} XRP</td>
                                            <td>{parseFloat(currency.avg_exchange_volume).toFixed(2)}</td>
                                            <td>{this.state.issuers[currency.issuer] || this.abbreviateHash(currency.issuer)}</td>
                                        </TableRow>
                                    )
                                })}
                            </TableBody>
                        </Table>
                    </Panel>
                    </div>
                    <div className="col-md-6">
                        <Panel title={`Node Locality (${this.state.num_nodes} total)`}>
                            <div id="chartdiv" style={{ width: "100%", height: "388px" }}></div>
                        </Panel>
                    </div>
                </div>
                <div className="row" style={{margin:"1em"}}>
                    <Panel title='Latest Offers Created'>
                        <Table>
                            <TableHead>
                                <th>Exchange Pair</th>
                                <th>Exchange Rate</th>
                                <th>Buy</th>
                                <th>Sell</th>
                                <th>Hash</th>
                                <th>Affected Nodes</th>
                                <th>Timestamp</th>
                            </TableHead>
                            <TableBody>
                                {this.state.offers.map((offer) => {
                                    return this.getTableRows(offer)
                                })}
                            </TableBody>
                        </Table>
                        <Row>
                        </Row>
                    </Panel>
                </div>
            </Page>

        )
    }
}
