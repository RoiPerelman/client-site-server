import React from 'react';
import { Collapse, Table } from 'reactstrap';
import {
  setIsMultipleSectionsUserRequestAction,
  addUserSectionRequestAction,
  delUserSectionRequestAction
} from '../../../../store/user/actions';
import { connect } from 'react-redux';

class MultipleSectionsForm extends React.Component {
  state = {
    inputSection: ''
  };

  toggle = () => {
    this.props.setIsMultipleSectionsUserRequest(!this.props.isMulti);
  };

  onChange = e => this.setState({ inputSection: e.target.value });

  addSection = () =>
    this.props.addUserSectionRequestAction({
      sectionId: this.state.inputSection,
      name: undefined
    });

  delSection = () =>
    this.props.delUserSectionRequestAction({
      sectionId: this.state.inputSection,
      name: undefined
    });

  render() {
    const { isMulti, sections, isMultiError, addSectionError } = this.props;

    return (
      <div>
        <h3>Sections</h3>
        <br />
        <label className="switch">
          {isMultiError && (
            <div className="alert alert-danger">{isMultiError}</div>
          )}
          <input type="checkbox" checked={isMulti} onChange={this.toggle} />
          <span className="slider round"> use multiple sections</span>
          <Collapse isOpen={isMulti}>
            {addSectionError && (
              <div className="alert alert-danger">{addSectionError}</div>
            )}
            <input
              type="text"
              placeholder={`SectionId`}
              value={this.state.inputSection}
              onChange={this.onChange}
            />
            <button type="button" onClick={this.addSection} className="small">
              Add Section
            </button>
            <button type="button" onClick={this.delSection} className="small">
              Del Section
            </button>
            <Table>
              <thead>
                <tr>
                  <th>id</th>
                  <th>SectionsId</th>
                  <th>SectionId</th>
                </tr>
              </thead>
              <tbody>
                {Object.keys(sections).map((key, idx) => (
                  <tr key={idx}>
                    <th scope="row">{idx}</th>
                    <td>{sections[key].id}</td>
                    <td>{sections[key].sectionId}</td>
                  </tr>
                ))}
              </tbody>
            </Table>
          </Collapse>
        </label>
        <hr />
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
    isMulti: state.user.isMulti,
    sections: state.user.sections,
    isMultiError: state.user.errors.isMultipleSection,
    addSectionError: state.user.errors.addSection
  };
}

export default connect(
  mapStateToProps,
  {
    setIsMultipleSectionsUserRequest: setIsMultipleSectionsUserRequestAction,
    addUserSectionRequestAction: addUserSectionRequestAction,
    delUserSectionRequestAction: delUserSectionRequestAction
  }
)(MultipleSectionsForm);
