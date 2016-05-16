import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import { Router, Route, browserHistory } from 'react-router';
import { syncHistoryWithStore } from 'react-router-redux';

import ExampleComponent from '../ExampleComponent/ExampleComponent';
import configureStore from './configureStore';

const store = configureStore();
const history = syncHistoryWithStore(browserHistory, store);

const RootComponent = () => {
    return (
        <Provider store={store}>
            <Router history={history}>
                    <Route path='/' component={ExampleComponent}>
                    </Route>
            </Router>
        </Provider>
    );
};

ReactDOM.render(
    <RootComponent />,
    document.getElementById('mount-point')
);
