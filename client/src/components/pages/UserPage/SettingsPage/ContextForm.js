import React from 'react';
import ContextForm2 from './ContextForm2';
import { connect } from 'react-redux';

class ContextsForm extends React.Component {
  render() {
    const { sections, activeSection, defaultSection } = this.props;
    const section = sections.filter(
      section =>
        section.sectionId === activeSection ||
        section.sectionId === defaultSection
    )[0];
    return (
      <div className="row">
        <div className="col-6">
          <h3>Contexts for {activeSection || defaultSection}</h3>
          <br />
          <h3>Product Context</h3>
          <ContextForm2
            sectionsId={section.id}
            contextType="PRODUCT"
            context={section.contexts.productContext}
          />
          <h3>Cart Context</h3>
          <ContextForm2
            sectionsId={section.id}
            contextType="CART"
            context={section.contexts.cartContext}
          />
          <h3>Category Context</h3>
          <ContextForm2
            sectionsId={section.id}
            contextType="CATEGORY"
            context={section.contexts.categoryContext}
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
