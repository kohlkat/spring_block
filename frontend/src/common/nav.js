import _ from 'lodash';
import React, { Component } from 'react';
import { Link } from 'react-router-dom';

export class Sidebar extends Component {
    state = {
        navItems: [
            { pathname: '/', label: 'Home', icon: 'home' },
            { pathname: '/history', label: 'History', icon: 'list-ul' },
            { pathname: '/network', label: 'Network', icon: 'dot-circle-o' },
            { pathname: '/about', label: 'About', icon: 'info' },
        ],
    }

    isSelected(navItem) {
        return this.props.history.location.pathname === navItem.pathname ? 'selected' : '';
    }

    renderLinks() {
        return _.map(this.state.navItems, (navItem) => {
            return (
                <li style={{top:"0"}} className={`al-sidebar-list-item ${this.isSelected(navItem)}`} key={navItem.pathname} >
                    <Link className="al-sidebar-list-link" to={{ pathname: navItem.pathname, query: navItem.query }}>
                        <i className={`fa fa-${navItem.icon}`}></i>
                        <span>{navItem.label}</span>
                    </Link>
                </li>
            );
        });
    }

    render() {
        return (
            <aside style={{top:"0px"}} className="al-sidebar">
                <div className="panel-title text-center" style={{marginTop:"24px"}}>SpringBlock</div>
                <ul className="al-sidebar-list">
                    {this.renderLinks()}
                </ul>
            </aside>
        );
    }
}