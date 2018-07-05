import React from 'react';
import DefaultContexts from './DefaultContexts';
import MultiContexts from './MultiContexts';
import { connect } from 'react-redux';

class Contexts extends React.Component {
  render() {
    const { isMulti } = this.props;
    return isMulti ? <MultiContexts /> : <DefaultContexts />;
  }
}

function mapStateToProps(state) {
  return {
    isMulti: state.user.isMulti
  };
}

export default connect(mapStateToProps)(Contexts);
