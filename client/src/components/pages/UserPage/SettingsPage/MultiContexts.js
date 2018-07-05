import React from 'react';
import ContextForm from './ContextForm';
import { connect } from 'react-redux';

class DefaultContexts extends React.Component {
  render() {
    const { sections } = this.props;
    return (
      <div className="row">
        <div className="col-6">
          <h3>Contexts</h3>
          <br />
          {/*<h3>Product Context</h3>*/}
          {/*<ContextForm*/}
          {/*section={section}*/}
          {/*contextType="PRODUCT"*/}
          {/*context={section.contexts.product}*/}
          {/*/>*/}
          {/*<h3>Cart Context</h3>*/}
          {/*<ContextForm*/}
          {/*section={section}*/}
          {/*contextType="CART"*/}
          {/*context={section.contexts.cart}*/}
          {/*/>*/}
          {/*<h3>Category Context</h3>*/}
          {/*<ContextForm*/}
          {/*section={section}*/}
          {/*contextType="CATEGORY"*/}
          {/*context={section.contexts.category}*/}
          {/*/>*/}
        </div>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
    activeSection: state.user.activeSection,
    defaultSection: state.user.defaultSection,
    sections: state.user.sections
  };
}

export default connect(mapStateToProps)(DefaultContexts);
