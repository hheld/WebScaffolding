import { combineReducers } from 'redux';
import { routerReducer } from 'react-router-redux';

import authReducer from '../auth/authReducer';
import editUserReducer from '../components/Admin/EditUsers/editUserReducer';
import changePwdReducer from '../components/User/changePwdReducer';

const rootReducer = combineReducers({
    user: authReducer,
    editUser: editUserReducer,
    routing: routerReducer,
    pwdChange: changePwdReducer
});

export default rootReducer;
