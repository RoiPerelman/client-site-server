import React from 'react';
import TopNavigation from './TopNavigation';
import TabNavigation from './TabNavigation';
import SettingsPage from './SettingsPage';
import { Route, Switch } from 'react-router-dom';
import { connect } from 'react-redux';
import { loadDynamicYieldRequestAction } from '../../../store/user/actions';
import DYUserRoute from '../../routes/DYUserRoute';
import HomePage from './HomePage';

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
          <DYUserRoute path="/" component={HomePage} />
        </Switch>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
    isDYLoaded: state.user.isDYLoaded,
    isMulti: state.user.isMulti,
    sections: state.user.sections,
    DYRequestError: state.user.errors.DYRequest
  };
}

export default connect(
  mapStateToProps,
  { loadDynamicYieldRequestAction: loadDynamicYieldRequestAction }
)(UserPage);
