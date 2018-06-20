import React from 'react';
import { Switch } from 'react-router-dom';
import GuestPage from './components/pages/GuestPage';
import UserRoute from './components/routes/UserRoute';
import GuestRoute from './components/routes/GuestRoute';
import UserPage from './components/pages/UserPage';

const App = ({ location }) => {
  return (
    <div>
      <Switch>
        <GuestRoute path="/guest" component={GuestPage} />
        {/*<GuestRoute path="/login" exact isLogin={true} component={GuestPage} />*/}
        {/*<GuestRoute path="/signup" exact isLogin={false} component={GuestPage} />*/}
        <UserRoute location={location} path="/" component={UserPage} />
      </Switch>
    </div>
  );
};

export default App;
