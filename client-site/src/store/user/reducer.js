import * as types from './types';
import { history } from '../../index';

const initialUserState = {
  email: '',
  username: '',
  password: '',
  isAuthenticated: false,
  isLoaded: false,
  sectionId: '',
  token: '',
  isMulti: true,
  errors: {
    email: '',
    username: '',
    password: '',
    server: '',
    isMulti: ''
  }
};

export const userReducer = (state = initialUserState, action = {}) => {
  switch (action.type) {
    case types.SIGNUP_USER_SUCCESS:
    case types.LOGIN_USER_SUCCESS:
      return { ...state, ...action.user };
    case types.SIGNUP_USER_FAILURE:
    case types.LOGIN_USER_FAILURE:
      return { ...state, errors: action.errors || state.errors };
    case types.FETCH_USER_FAILURE:
      return {
        ...state,
        isAuthenticated: action.isAuthenticated,
        isLoaded: true
      };
    case types.FETCH_USER_SUCCESS:
      return { ...state, ...action.user, isLoaded: true };
    case types.LOGOUT_USER:
      delete localStorage.RPJWT;
      history.push('/');
      window.location.reload();
      return state;
    case types.SET_MULTIPLE_SECTION_USER_SUCCESS:
      return {
        ...state,
        isMulti: action.isMulti
      };
    case types.SET_MULTIPLE_SECTION_USER_FAILURE:
      return {
        ...state,
        errors: action.errors
      };
    default:
      return state;
  }
};
