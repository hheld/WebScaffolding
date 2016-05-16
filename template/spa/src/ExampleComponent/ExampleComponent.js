import React, { PropTypes } from 'react';
import { connect } from 'react-redux';
import * as actions from './ExampleComponentActions';

// component part #####################################################################################################
export class ExampleComponent extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
        return(
            <div>
                Counter value: {this.props.counter}
                <button onClick={() => this.runExampleAction()}>Run action</button>
            </div>
        );
    }
    
    runExampleAction() {
        this.props.runExampleAction();
    }
}

ExampleComponent.propTypes = {
    counter: PropTypes.number.isRequired,
    runExampleAction: PropTypes.func.isRequired
};

// container part #####################################################################################################
function mapStateToProps(state) {
    return {
        counter: state.example.counter
    };
}

function mapDispatchToProps(dispatch) {
    return {
        runExampleAction: () => dispatch(actions.exampleAction())
    };
}

export default connect(mapStateToProps, mapDispatchToProps)(ExampleComponent);
