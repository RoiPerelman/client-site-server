import { createStore, applyMiddleware, combineReducers } from 'redux';
import thunk from 'redux-thunk';
import { composeWithDevTools } from 'redux-devtools-extension';
import { userReducer } from './user/reducer';
import { rootSaga } from './user/saga';
import createSagaMiddleware from 'redux-saga';

// creates the rootReducer, rootSaga, saga and store
export const configureStore = () => {
  const rootReducer = combineReducers({
    user: userReducer
  });
  const sagaMiddleware = createSagaMiddleware();
  const store = createStore(
    rootReducer,
    composeWithDevTools(applyMiddleware(sagaMiddleware, thunk))
  );
  sagaMiddleware.run(rootSaga);
  return store;
};
