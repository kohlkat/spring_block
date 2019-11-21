import React from 'react';
import { Link } from 'react-router-dom';

export class PageTop extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            isMenuOpen: false,
            appName: "SpringBlock Labs"
        }
    }

    renderLogo() {
        return (
            <Link to={{ pathname: '/' }} className="al-logo clearfix">{this.state.appName}</Link>
        );
    }

    render() {
        return (
            <div className="page-top clearfix" max-height="50">
                {this.renderLogo()}
            </div>
        );
    }
}