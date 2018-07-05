import React from 'react';
import TopNavigation from './TopNavigation';
import TabNavigation from './TabNavigation';
import SettingsPage from './SettingsPage';
import { Route, Switch } from 'react-router-dom';
import DYUserRoute from '../../routes/DYUserRoute';
import Page from './Page';

class UserPage extends React.Component {
  render() {
    return (
      <div className="container-fluid">
        <TopNavigation />
        {Array.from(Array(3)).map((value, idx) => <hr key={idx} />)}
        <TabNavigation />
        <Switch>
          <Route
            path="/settings"
            exact
            render={() => {
              return <SettingsPage />;
            }}
          />
          <DYUserRoute path="/" component={Page} />
        </Switch>
      </div>
    );
  }
}

export default UserPage;
