import { call, put } from 'redux-saga/effects';
import * as actions from './actions';
import api from './api';
import { history } from '../../index';
import setAuthorizationHeader from '../../utils/setAuthorizationHeader';
import loadDynamicYield from '../../utils/dynamicyield';
import { takeEvery } from 'redux-saga/effects';

export function* rootSaga() {
  yield takeEvery('FETCH_USER_REQUEST', fetchCurrentUserRequestSaga);
  yield takeEvery('SIGNUP_USER_REQUEST', signupUserRequestSaga);
  yield takeEvery('LOGIN_USER_REQUEST', loginUserRequestSaga);
  yield takeEvery(
    'SET_MULTIPLE_SECTION_USER_REQUEST',
    setMultipleSectionUserSaga
  );
  yield takeEvery(
    'LOAD_DYNAMIC_YIELD_REQUEST_ACTION',
    loadDynamicYieldRequestSaga
  );
}

export function* signupUserRequestSaga(action) {
  try {
    const user = yield call(api.signup, action.user);
    localStorage.RPJWT = user.token;
    yield fetchCurrentUserRequestSaga();
    history.push('/');
  } catch (err) {
    console.log('err is ' + err);
    yield put(actions.signupUserFailedAction(err.response.data.errors));
  }
}

export function* fetchCurrentUserRequestSaga() {
  try {
    if (localStorage.RPJWT) {
      setAuthorizationHeader(localStorage.RPJWT);
      const user = yield call(api.authorize);
      if (!user.isMulti) {
        yield loadDynamicYieldRequestSaga({ section: user.defaultSection });
      }
      yield put(actions.fetchCurrentUserSuccessAction(user));
    } else {
      yield put(actions.fetchCurrentUserFailedAction(false));
    }
  } catch (err) {
    yield put(actions.fetchCurrentUserFailedAction(false));
  }
}

export function* loginUserRequestSaga(action) {
  try {
    const user = yield call(api.login, action.user);
    console.log(user);
    localStorage.RPJWT = user.token;
    yield fetchCurrentUserRequestSaga();
    history.push('/');
  } catch (err) {
    yield put(actions.loginUserFailedAction(err.response.data.errors));
  }
}

export function* setMultipleSectionUserSaga(action) {
  try {
    let isMulti = yield call(api.setMultipleSections, action.isMulti);
    console.log(isMulti);
    isMulti = action.isMulti;
    yield put(actions.setIsMultipleSectionsUserSuccessAction(isMulti));
  } catch (err) {
    console.log(err.response.data.errors || 'failed');
    yield put(
      actions.setIsMultipleSectionsUserFailedAction(
        err.response.data.errors || 'failed'
      )
    );
  }
}

export function* loadDynamicYieldRequestSaga(action) {
  try {
    console.log('in loadDynamicYieldRequestSaga ' + action);
    yield loadDynamicYield(action.section);
    yield put(actions.loadDynamicYieldSuccessAction(true));
  } catch (e) {
    console.log('in catch block loadDynamicYieldRequestSaga');
    yield put(actions.loadDynamicYieldFailureAction(e.stack));
  }
}
