import React from "react";
import {
    BrowserRouter as Router,
    Switch,
    Route,
    withRouter
} from "react-router-dom";
import { Sidebar } from "./common/nav"
import { Home } from "./views/home"
import { History } from "./views/history"
import { Network } from "./views/network"
import { About } from "./views/about"
import { NotFound } from "./views/not_found";

export default function AppRouter() {
    return (
        <Router >
            <div >
                {/*<PageTop/>*/}
                <Nav/>
                <div style={{marginLeft:"15em", marginRight:"1em", maxHeight:"100vh"}}>
                <Switch>
                    <Route exact path="/" component={withRouter(Home)}/>
                    <Route path="/network" component={withRouter(Network)}/>
                    <Route path="/about" component={withRouter(About)}/>
                    <Route path="/history" component={withRouter(History)}/>
                    <Route component={withRouter(NotFound)}/>
                </Switch>
                </div>
            </div>
        </Router>
    );
}

const Nav = withRouter(Sidebar);
