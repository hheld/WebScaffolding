import { combineReducers } from 'redux';
import { routerReducer } from 'react-router-redux';

import exampleComponentReducer from '../ExampleComponent/exampleComponentReducer';

const rootReducer = combineReducers({
    example: exampleComponentReducer,
    routing: routerReducer
});

export default rootReducer;
