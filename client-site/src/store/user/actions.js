import * as types from './types';

export const signupUserRequestAction = user => ({
  type: types.SIGNUP_USER_REQUEST,
  user
});
export const signupUserSuccessAction = user => ({
  type: types.SIGNUP_USER_SUCCESS,
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
export const loginUserSuccessAction = user => ({
  type: types.LOGIN_USER_SUCCESS,
  user
});
export const loginUserFailedAction = errors => ({
  type: types.LOGIN_USER_FAILURE,
  errors
});

export const logoutUserAction = () => ({
  type: types.LOGOUT_USER
});
