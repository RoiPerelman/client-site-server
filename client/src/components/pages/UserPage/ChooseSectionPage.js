import React from 'react';
import { loadDynamicYieldRequestAction } from '../../../store/user/actions';
import { connect } from 'react-redux';

class ChooseSectionPage extends React.Component {
  constructor(props) {
    super(props);
    this.toggle = this.toggle.bind(this);
  }

  toggle() {
    this.props.setIsMultipleSectionsUserRequest(!this.props.isMulti);
  }

  onClick = (section, contexts, jsCode) => {
    this.props.loadDynamicYieldRequestAction({
      section: section,
      contexts: contexts,
      jsCode: jsCode
    });
  };

  render() {
    const { sections, DYRequestError, jsCode } = this.props;

    return (
      <div>
        {DYRequestError && (
          <div className="alert alert-danger">{DYRequestError}</div>
        )}
        {Object.keys(sections).map(idx => {
          return (
            <button
              key={idx}
              section={sections[idx].sectionId}
              color="primary"
              onClick={() =>
                this.onClick(
                  sections[idx].sectionId,
                  sections[idx].contexts,
                  jsCode
                )
              }
            >
              {sections[idx].sectionId}
            </button>
          );
        })}
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
    sections: state.user.sections,
    DYRequestError: state.user.errors.DYRequest,
    jsCode: state.user.jsCode
  };
}

export default connect(
  mapStateToProps,
  { loadDynamicYieldRequestAction: loadDynamicYieldRequestAction }
)(ChooseSectionPage);
