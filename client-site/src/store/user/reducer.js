import * as types from './types';
import { history } from '../../index';

const initialUserState = {
  email: '',
  username: '',
  password: '',
  isAuthenticated: false,
  isLoaded: false, // after we load all user information from the server
  isDYLoaded: false,
  isDYLoading: false,
  defaultSection: '',
  sectionId: '',
  token: '',
  isMulti: false,
  sections: [],
  errors: {
    email: '',
    username: '',
    password: '',
    server: '',
    multipleSections: '',
    DYRequest: ''
  }
};

export const userReducer = (state = initialUserState, action = {}) => {
  switch (action.type) {
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
    case types.LOAD_DYNAMIC_YIELD_REQUEST_ACTION:
      return {
        ...state,
        isDYLoading: true
      };
    case types.LOAD_DYNAMIC_YIELD_SUCCESS_ACTION:
      return {
        ...state,
        isDYLoaded: action.isDYLoaded,
        isDYLoading: false
      };
    case types.LOAD_DYNAMIC_YIELD_FAILURE_ACTION:
      return {
        ...state,
        errors: {
          ...state.errors,
          DYRequest: action.error
        },
        isDYLoading: false
      };
    default:
      return state;
  }
};
