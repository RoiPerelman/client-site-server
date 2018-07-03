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
  activeSection: '',
  sectionId: '',
  token: '',
  isMulti: false,
  sections: [],
  errors: {
    email: '',
    username: '',
    password: '',
    server: '',
    isMultipleSection: '',
    addSection: '',
    DYRequest: '',
    contexts: ''
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
        isMulti: action.isMulti,
        errors: {
          ...state.errors,
          isMultipleSection: ''
        }
      };
    case types.SET_MULTIPLE_SECTION_USER_FAILURE:
      return {
        ...state,
        errors: {
          ...state.errors,
          isMultipleSection: action.error
        }
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
    case types.UPDATE_ACTIVE_SECTION:
      return {
        ...state,
        activeSection: action.section
      };
    case types.ADD_USER_SECTION_SUCCESS:
      return {
        ...state,
        sections: [...state.sections, { sectionId: action.section, name: '' }],
        errors: {
          ...state.errors,
          addSection: ''
        }
      };
    case types.ADD_USER_SECTION_FAILURE:
      return {
        ...state,
        errors: {
          ...state.errors,
          addSection: action.error
        }
      };
    case types.DEL_USER_SECTION_SUCCESS:
      return {
        ...state,
        sections: state.sections.filter(section => section !== action.section),
        errors: {
          ...state.errors,
          addSection: ''
        }
      };
    case types.DEL_USER_SECTION_FAILURE:
      return {
        ...state,
        errors: {
          ...state.errors,
          addSection: action.error
        }
      };
    case types.ADD_CONTEXT_ITEM_SUCCESS:
      return {
        ...state,
        sections: [...state.sections]
      };
    case types.ADD_CONTEXT_ITEM_FAILURE:
      return {
        ...state,
        errors: {
          ...state.errors,
          contexts: action.error || 'error NOT WORKING'
        }
      };
    default:
      return state;
  }
};
