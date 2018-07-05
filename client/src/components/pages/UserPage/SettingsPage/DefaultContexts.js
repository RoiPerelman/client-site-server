import React from 'react';
import ContextForm from './ContextForm';
import { connect } from 'react-redux';

class DefaultContexts extends React.Component {
  render() {
    const { sections, activeSection, defaultSection } = this.props;
    const section = sections[activeSection] || sections[defaultSection];
    return (
      <div>
        <h3>Contexts for {activeSection || defaultSection}</h3>
        <br />
        <div className="row">
          <div className="col-4">
            <h3>Product Context</h3>
            <ContextForm
              section={section}
              contextType="PRODUCT"
              context={section.contexts.product}
            />
          </div>
          <div className="col-4">
            <h3>Cart Context</h3>
            <ContextForm
              section={section}
              contextType="CART"
              context={section.contexts.cart}
            />
          </div>
          <div className="col-4">
            <h3>Category Context</h3>
            <ContextForm
              section={section}
              contextType="CATEGORY"
              context={section.contexts.category}
            />
          </div>
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
