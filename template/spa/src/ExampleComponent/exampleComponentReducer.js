import {RUN_EXAMPLE_ACTION} from './ExampleComponentActions';

const initialState = {
    counter: 0
};

export default function exampleComponentReducer(state = initialState, action) {
    switch(action.type) {
    case RUN_EXAMPLE_ACTION:
        return Object.assign({}, state, {
            counter: state.counter + 1
        });
        
    default:
        return state;
    }
}
