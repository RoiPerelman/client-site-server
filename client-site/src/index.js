import React from 'react';
import ReactDOM from 'react-dom';
import { Router, Route } from 'react-router-dom';
import { Provider } from 'react-redux';
import App from './App';
import registerServiceWorker from './registerServiceWorker';
import { fetchCurrentUserRequestAction } from './store/user/actions';
import { configureStore } from './store';
import { createBrowserHistory } from 'history';
import 'bootstrap/dist/css/bootstrap.css';

export const history = createBrowserHistory();
const store = configureStore();

store.dispatch(fetchCurrentUserRequestAction());

ReactDOM.render(
  <Router history={history}>
    <Provider store={store}>
      <Route component={App} />
    </Provider>
  </Router>,
  document.getElementById('root')
);
registerServiceWorker();
