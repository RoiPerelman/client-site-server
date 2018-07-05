import React from 'react';
import { connect } from 'react-redux';
import { Route } from 'react-router-dom';
import Loader from 'react-loader';
import ChooseSectionPage from '../pages/UserPage/ChooseSectionPage';
const DYUserRoute = ({
  isDYLoading,
  isDYLoaded,
  isMulti,
  sections,
  DYRequestError,
  component: Component,
  ...rest
}) => {
  return (
    <Loader loaded={!isDYLoading}>
      <Route
        {...rest}
        render={props => {
          return !isDYLoaded && isMulti ? (
            <ChooseSectionPage />
          ) : (
            <Component {...props} {...rest} />
          );
        }}
      />
    </Loader>
  );
};

function mapStateToProps(state) {
  return {
    isDYLoading: state.user.isDYLoading,
    isDYLoaded: state.user.isDYLoaded,
    isMulti: state.user.isMulti
  };
}

export default connect(mapStateToProps)(DYUserRoute);
