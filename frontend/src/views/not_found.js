import React, { Component } from "react"

export class NotFound extends Component {
    render() {
        return (
                <div className="center-block">
                    <h1>404 Error</h1>
                    <p>Sorry, that page doesn't exist. <a href="/">Go to Home Page.</a></p>
                </div>
        )
    }
}