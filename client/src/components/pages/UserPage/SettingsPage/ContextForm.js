import React from 'react';
import ContextForm2 from './ContextForm2';
import { connect } from 'react-redux';

class ContextsForm extends React.Component {
  render() {
    const { sections, activeSection, defaultSection } = this.props;
    const section = sections[activeSection] || sections[defaultSection];
    return (
      <div className="row">
        <div className="col-6">
          <h3>Contexts for {activeSection || defaultSection}</h3>
          <br />
          <h3>Product Context</h3>
          <ContextForm2
            section={section}
            contextType="PRODUCT"
            context={section.contexts.product}
          />
          <h3>Cart Context</h3>
          <ContextForm2
            section={section}
            contextType="CART"
            context={section.contexts.cart}
          />
          <h3>Category Context</h3>
          <ContextForm2
            section={section}
            contextType="CATEGORY"
            context={section.contexts.category}
          />
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

export default connect(mapStateToProps)(ContextsForm);
