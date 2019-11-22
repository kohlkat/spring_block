import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './routes'
import * as serviceWorker from './serviceWorker';
import 'react-flex-proto/styles/flex.css';
import 'react-blur-admin/dist/assets/styles/react-blur-admin.min.css';

ReactDOM.render(<App />, document.getElementById('root'));

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
