import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { Route, Redirect } from 'react-router-dom';
import Loader from 'react-loader';

const UserRoute = ({
  isLoaded,
  isAuthenticated,
  component: Component,
  ...rest
}) => {
  return (
    <Loader loaded={isLoaded}>
      <Route
        {...rest}
        render={props => {
          return isAuthenticated ? (
            <Component {...props} {...rest} />
          ) : (
            <Redirect to="/guest/login" />
          );
        }}
      />
    </Loader>
  );
};

UserRoute.propTypes = {
  component: PropTypes.func.isRequired,
  isAuthenticated: PropTypes.bool.isRequired,
  isLoaded: PropTypes.bool.isRequired
};

function mapStateToProps(state) {
  return {
    isAuthenticated: state.user.isAuthenticated,
    isLoaded: state.user.isLoaded
  };
}

export default connect(mapStateToProps)(UserRoute);
