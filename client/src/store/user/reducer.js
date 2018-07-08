import * as types from './types';
import { history } from '../../index';

const clearErrors = {
  email: '',
  username: '',
  password: '',
  login: '',
  isMulti: '',
  section: '',
  DYRequest: '',
  addContext: ''
};

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
  sections: {},
  errors: clearErrors
};

export const userReducer = (state = initialUserState, action = {}) => {
  switch (action.type) {
    case types.SIGNUP_USER_FAILURE:
      return { ...state, errors: action.errors || state.errors };
    case types.LOGIN_USER_FAILURE:
      return {
        ...state,
        errors: {
          ...state.errors,
          login: action.error
        }
      };
    case types.FETCH_USER_FAILURE:
      return {
        ...state,
        isAuthenticated: false,
        isLoaded: true
      };
    case types.FETCH_USER_SUCCESS:
      return { ...state, ...action.user, errors: clearErrors, isLoaded: true };
    case types.LOGOUT_USER:
      delete localStorage.RPJWT;
      history.push('/');
      window.location.reload();
      return state;
    case types.SET_MULTI_SECTION_SUCCESS:
      return {
        ...state,
        isMulti: action.isMulti
      };
    case types.SET_MULTI_SECTION_FAILURE:
      return {
        ...state,
        errors: {
          ...state.errors,
          isMulti: action.error
        }
      };
    case types.LOAD_DYNAMIC_YIELD_REQUEST:
      return {
        ...state,
        isDYLoading: true
      };
    case types.LOAD_DYNAMIC_YIELD_SUCCESS:
      return {
        ...state,
        isDYLoaded: action.isDYLoaded,
        isDYLoading: false
      };
    case types.LOAD_DYNAMIC_YIELD_FAILURE:
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
    case types.ADD_USER_SECTION_FAILURE:
    case types.DEL_USER_SECTION_FAILURE:
      return {
        ...state,
        errors: {
          ...state.errors,
          section: action.error
        }
      };
    case types.ADD_USER_SECTION_SUCCESS:
      return {
        ...state,
        sections: {
          ...state.sections,
          [action.section.sectionId]: action.section
        },
        errors: {
          ...state.errors,
          section: ''
        }
      };
    case types.DEL_USER_SECTION_SUCCESS:
      let { [action.section.sectionId]: omit, ...rest } = state.sections;
      return {
        ...state,
        sections: rest,
        errors: {
          ...state.errors,
          addSection: ''
        }
      };
    case types.ADD_CONTEXT_ITEM_SUCCESS:
    case types.DEL_CONTEXT_ITEM_SUCCESS:
      return {
        ...state,
        sections: {
          ...state.sections,
          [action.section.sectionId]: action.section
        },
        errors: {
          ...state.errors,
          addContext: ''
        }
      };
    case types.ADD_CONTEXT_ITEM_FAILURE:
    case types.DEL_CONTEXT_ITEM_FAILURE:
      return {
        ...state,
        errors: {
          ...state.errors,
          addContext: action.error
        }
      };
    default:
      return state;
  }
};
