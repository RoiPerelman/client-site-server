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
}

export function* signupUserRequestSaga(action) {
  try {
    const user = yield call(api.signup, action.user);
    localStorage.RPJWT = user.token;
    yield put(actions.signupUserSuccessAction(user));
    history.push('/');
    window.location.reload();
  } catch (err) {
    yield put(actions.signupUserFailedAction(err.response.data.errors));
  }
}

export function* fetchCurrentUserRequestSaga() {
  try {
    if (localStorage.RPJWT) {
      setAuthorizationHeader(localStorage.RPJWT);
      const user = yield call(api.authorize);
      yield loadDynamicYield(user.sectionId).catch(e => {
        console.log(e.stack);
      });
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
    localStorage.RPJWT = user.token;
    yield put(actions.loginUserSuccessAction(user));
    history.push('/');
    window.location.reload();
  } catch (err) {
    yield put(actions.loginUserFailedAction(err.response.data.errors));
  }
}

export function* setMultipleSectionUserSaga(action) {
  try {
    // const isMulti = yield call(api.setMultipleSections, action.isMulti);
    const isMulti = action.isMulti;
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
