import React from 'react';
import { Switch } from 'react-router-dom';
import GuestPage from './components/pages/GuestPage';
import UserRoute from './components/routes/UserRoute';
import GuestRoute from './components/routes/GuestRoute';
import UserPage from './components/pages/UserPage';

const App = () => {
  return (
    <div>
      <Switch>
        <GuestRoute path="/guest" component={GuestPage} />
        <UserRoute path="/" component={UserPage} />
      </Switch>
    </div>
  );
};

export default App;
