import React, { Fragment } from 'react';
import { Switch } from 'react-router-dom';
import GuestPage from './components/pages/GuestPage';
import UserRoute from './components/routes/UserRoute';
import GuestRoute from './components/routes/GuestRoute';
import UserPage from './components/pages/UserPage';

const App = () => {
  return (
    <Fragment>
      <Switch>
        <GuestRoute path="/guest" component={GuestPage} />
        <UserRoute path="/" component={UserPage} />
      </Switch>
    </Fragment>
  );
};

export default App;
