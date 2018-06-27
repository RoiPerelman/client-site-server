import React from 'react';
import { loadDynamicYieldRequestAction } from '../../store/user/actions';
import { connect } from 'react-redux';

class ChooseSectionPage extends React.Component {
  constructor(props) {
    super(props);
    this.toggle = this.toggle.bind(this);
  }

  toggle() {
    this.props.setIsMultipleSectionsUserRequest(!this.props.isMulti);
  }

  onClick = section => {
    console.log(section);
    this.props.loadDynamicYieldRequestAction(section);
  };

  render() {
    const { sections, DYRequestError } = this.props;

    return (
      <div>
        {DYRequestError && (
          <div className="alert alert-danger">{DYRequestError}</div>
        )}
        {sections.map((section, id) => {
          return (
            <button
              key={id}
              section={section}
              color="primary"
              onClick={() => this.onClick(section)}
            >
              {section}
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
    DYRequestError: state.user.errors.DYRequest
  };
}

export default connect(
  mapStateToProps,
  { loadDynamicYieldRequestAction: loadDynamicYieldRequestAction }
)(ChooseSectionPage);