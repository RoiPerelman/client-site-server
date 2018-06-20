import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { Route, Redirect } from 'react-router-dom';
import Loader from 'react-loader';

const GuestRoute = ({
  isLoaded,
  isAuthenticated,
  component: Component,
  ...rest
}) => (
  <Loader loaded={isLoaded}>
    <Route
      {...rest}
      render={props => {
        return isAuthenticated ? (
          <Redirect to="/" />
        ) : (
          <Component {...props} {...rest} />
        );
      }}
    />
  </Loader>
);

GuestRoute.propTypes = {
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

export default connect(mapStateToProps)(GuestRoute);
