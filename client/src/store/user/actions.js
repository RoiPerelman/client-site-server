import * as types from './types';

export const signupUserRequestAction = user => ({
  type: types.SIGNUP_USER_REQUEST,
  user
});
export const signupUserFailedAction = errors => ({
  type: types.SIGNUP_USER_FAILURE,
  errors
});

export const fetchCurrentUserRequestAction = () => ({
  type: types.FETCH_USER_REQUEST
});
export const fetchCurrentUserFailedAction = isAuthenticated => ({
  type: types.FETCH_USER_FAILURE,
  isAuthenticated
});
export const fetchCurrentUserSuccessAction = user => ({
  type: types.FETCH_USER_SUCCESS,
  user
});

export const loginUserRequestAction = user => ({
  type: types.LOGIN_USER_REQUEST,
  user
});
export const loginUserFailedAction = error => ({
  type: types.LOGIN_USER_FAILURE,
  error
});

export const logoutUserAction = () => ({
  type: types.LOGOUT_USER
});

export const setIsMultiSectionsRequestAction = isMulti => ({
  type: types.SET_MULTI_SECTION_REQUEST,
  isMulti
});
export const setIsMultiSectionsSuccessAction = isMulti => ({
  type: types.SET_MULTI_SECTION_SUCCESS,
  isMulti
});
export const setIsMultiSectionsFailedAction = error => ({
  type: types.SET_MULTI_SECTION_FAILURE,
  error
});

export const loadDynamicYieldRequestAction = ({
  section,
  contexts,
  jsCode
}) => ({
  type: types.LOAD_DYNAMIC_YIELD_REQUEST,
  section,
  contexts,
  jsCode
});
export const loadDynamicYieldSuccessAction = isDYLoaded => ({
  type: types.LOAD_DYNAMIC_YIELD_SUCCESS,
  isDYLoaded
});
export const loadDynamicYieldFailureAction = error => ({
  type: types.LOAD_DYNAMIC_YIELD_FAILURE,
  error
});

export const changeActiveSectionAction = section => ({
  type: types.UPDATE_ACTIVE_SECTION,
  section
});

export const addUserSectionRequestAction = section => ({
  type: types.ADD_USER_SECTION_REQUEST,
  section
});
export const addUserSectionSuccessAction = section => ({
  type: types.ADD_USER_SECTION_SUCCESS,
  section
});
export const addUserSectionFailureAction = error => ({
  type: types.ADD_USER_SECTION_FAILURE,
  error
});

export const delUserSectionRequestAction = section => ({
  type: types.DEL_USER_SECTION_REQUEST,
  section
});
export const delUserSectionSuccessAction = section => ({
  type: types.DEL_USER_SECTION_SUCCESS,
  section
});
export const delUserSectionFailureAction = error => ({
  type: types.DEL_USER_SECTION_FAILURE,
  error
});

export const addContextItemRequestAction = contextItem => ({
  type: types.ADD_CONTEXT_ITEM_REQUEST,
  contextItem
});
export const addContextItemSuccessAction = section => ({
  type: types.ADD_CONTEXT_ITEM_SUCCESS,
  section
});
export const addContextItemFailureAction = error => ({
  type: types.ADD_CONTEXT_ITEM_FAILURE,
  error
});

export const delContextItemRequestAction = contextItem => ({
  type: types.DEL_CONTEXT_ITEM_REQUEST,
  contextItem
});
export const delContextItemSuccessAction = section => ({
  type: types.DEL_CONTEXT_ITEM_SUCCESS,
  section
});
export const delContextItemFailureAction = error => ({
  type: types.DEL_CONTEXT_ITEM_FAILURE,
  error
});

export const updateJSCodeRequestAction = jsCode => ({
  type: types.UPDATE_JSCODE_REQUEST,
  jsCode
});

export const updateJSCodeSuccessAction = jsCode => ({
  type: types.UPDATE_JSCODE_SUCCESS,
  jsCode
});

export const updateJSCodeFailureAction = error => ({
  type: types.UPDATE_JSCODE_FAILURE,
  error
});
