import React, { Component } from "react"
import {Panel} from 'react-blur-admin';

export class About extends Component {
    render() {
        return(
            <div className="row-fluid">
                <div className="row" style={{margin:"0", marginTop:"1em"}}>
                <div className="col-lg-12 col-sm-12 col-xs-12">
                    <div ba-panel ba-panel-class="banner-column-panel">
                        <div className="banner">
                            <div className="large-banner-wrapper">
                                <img src="http://www.akveo.com/blur-admin/assets/img/app/typography/banner.png" alt="http://www.akveo.com/blur-admin/assets/img/app/typography/banner.png"/>
                            </div>
                            <div className="banner-text-wrapper">
                                <div className="banner-text">
                                    <h1>SpringBlock Labs</h1>
                                    <p>Block-Sprint Hackathon 2019</p>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                </div>
                <div className="row" style={{margin:"1em"}}>
                <Panel title="Market Optimisation for the XRP Ledger">
                <p>
                    <p>Arbitrage is the simultaneous purchase and sale of an asset to profit from an imbalance in price. It is a trade that profits by exploiting the price differences of identical or similar financial instruments on different markets or in different forms. In the conventional currency market, arbitrage is difficult to execute as a result of various market frictions such as excessive order book spread, high transaction fees and capital controls. On the XRP Ledger's decentralized exchange, users are able to conduct cross-currency payments with little friction thanks to its high speed and low transaction fees.</p>

                    <p>Taking advantage of the comparatively favourable conditions, Spring Block proposes a stylized arbitrage scheme on the XRP Ledger. The arbitrage methodology takes advantage of the unique XRP payment flags tfNoDirectRipple, tfPartialPayment, and tfLimitQuality alongside the ledger’s Currency Exchange Nodes, Trust Lines, Rippling, and IOUs, to close the price gap between exchanges.</p>

                    <p>Whilst we have already completed an illustrative case study to confirm that certain arbitrage schemes are already in use, we would like to conduct a more systematic investigation on the users of the XRP ledger to gauge the exact magnitude of arbitrage profit system-wide. We will also apply forensic data science to connect and group different arbitrageur addresses so that we know not only the system-wide profit, but the extent to which this method is employed by different network participants.</p>

                    <p>In addition, we would like to design an algorithm and accompanying arbitrage bot which uses the rippled API to automate arbitrage when opportunities present themselves, in a way which minimises losses and maximises profits.</p>

                    <p>Finally, we would like to discuss the economic effects of arbitrage on the XRP ledger. Specifically, we want to understand whether such activities contribute to increasing liquidity and reducing market inefficiencies by enabling assets on the decentralized exchange to adhere to the law of one price, or if the result is less meaningful given the associated cost and effort.</p>
                </p>
                </Panel>
                </div>
            </div>
        )
    }
}