import React from 'react';
import TopNavigation from '../navigation/TopNavigation';
import SettingsPage from './SettingPage';
import { Route, Switch } from 'react-router-dom';
import { connect } from 'react-redux';
import { loadDynamicYieldRequestAction } from '../../store/user/actions';
import DYUserRoute from '../routes/DYUserRoute';
import HomePage from '../pages/HomePage';

class UserPage extends React.Component {
  render() {
    return (
      <div className="container-fluid">
        <TopNavigation />
        {Array.from(Array(3)).map((value, idx) => <hr key={idx} />)}
        <Switch>
          <Route
            path="/settings"
            exact
            render={() => {
              return <SettingsPage />;
            }}
          />
          <DYUserRoute path="/" exact component={HomePage} />
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
