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
  yield takeEvery('LOAD_DYNAMIC_YIELD_REQUEST', loadDynamicYieldRequestSaga);
  yield takeEvery('ADD_USER_SECTION_REQUEST', addSectionToUserRequestSaga);
  yield takeEvery('DEL_USER_SECTION_REQUEST', delSectionToUserRequestSaga);
  yield takeEvery('ADD_CONTEXT_ITEM_REQUEST', addContextItemRequestSaga);
  yield takeEvery('DEL_CONTEXT_ITEM_REQUEST', delContextItemRequestSaga);
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

export function* loginUserRequestSaga(action) {
  try {
    const user = yield call(api.login, action.user);
    localStorage.RPJWT = user.token;
    yield fetchCurrentUserRequestSaga();
    history.push('/');
  } catch (err) {
    yield put(actions.loginUserFailedAction(err.response.data.errors));
  }
}

export function* fetchCurrentUserRequestSaga() {
  try {
    if (localStorage.RPJWT) {
      setAuthorizationHeader(localStorage.RPJWT);
      const user = yield call(api.authorize);
      if (!user.isMulti) {
        yield loadDynamicYieldRequestSaga({
          section: user.defaultSection,
          contexts: user.sections[user.defaultSection].contexts
        });
      }
      yield put(actions.fetchCurrentUserSuccessAction(user));
    } else {
      yield put(actions.fetchCurrentUserFailedAction(false));
    }
  } catch (err) {
    yield put(actions.fetchCurrentUserFailedAction(false));
  }
}

export function* loadDynamicYieldRequestSaga(action) {
  const { section } = action;
  try {
    yield put(actions.changeActiveSectionAction(section));
    yield loadDynamicYield(action);
    yield put(actions.loadDynamicYieldSuccessAction(true));
  } catch (e) {
    yield put(actions.loadDynamicYieldFailureAction(e.stack));
  }
}

export function* setMultipleSectionUserSaga(action) {
  try {
    let isMulti = yield call(api.setMultipleSections, action.isMulti);
    yield put(actions.setIsMultipleSectionsUserSuccessAction(isMulti));
  } catch (err) {
    yield put(
      actions.setIsMultipleSectionsUserFailedAction(
        err.response.data.errors || 'failed'
      )
    );
  }
}

export function* addSectionToUserRequestSaga(action) {
  try {
    const sections = yield call(api.addSection, action.section);
    yield put(actions.addUserSectionSuccessAction(sections));
  } catch (err) {
    yield put(
      actions.addUserSectionFailureAction(
        err.response.data.errors || 'failed to add section'
      )
    );
  }
}

export function* delSectionToUserRequestSaga(action) {
  try {
    yield call(api.delSection, action.section);
    yield put(actions.delUserSectionSuccessAction(action.section));
  } catch (err) {
    yield put(
      actions.delUserSectionFailureAction(
        err.response.data.errors || err.response.data || 'failed to del section'
      )
    );
  }
}

export function* addContextItemRequestSaga(action) {
  try {
    const section = yield call(api.addContextItem, action.contextItem);
    yield put(actions.addContextItemSuccessAction(section));
  } catch (err) {
    yield put(
      actions.addContextItemFailureAction(
        err.response.data.errors ||
          err.response.data ||
          'failed to add contextItem'
      )
    );
  }
}

export function* delContextItemRequestSaga(action) {
  try {
    const section = yield call(api.delContextItem, action.contextItem);
    yield put(actions.delContextItemSuccessAction(section));
  } catch (err) {
    yield put(
      actions.delContextItemFailureAction(
        err.response.data.errors ||
          err.response.data ||
          'failed to del contextItem'
      )
    );
  }
}
