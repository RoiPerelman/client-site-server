import React from 'react';
import { Collapse, Card, CardBody } from 'reactstrap';
import { setIsMultipleSectionsUserRequestAction } from '../../store/user/actions';
import { connect } from 'react-redux';

class SettingsPage extends React.Component {
  toggle = () => {
    this.props.setIsMultipleSectionsUserRequest(!this.props.isMulti);
  };

  render() {
    const { isMulti, isMultiErrors } = this.props;

    return (
      <div>
        <label className="switch">
          {isMultiErrors && (
            <div className="alert alert-danger">{isMultiErrors}</div>
          )}
          <input type="checkbox" checked={isMulti} onChange={this.toggle} />
          <span className="slider round"> use multiple sections</span>
          <Collapse isOpen={isMulti}>
            <Card>
              <CardBody>Using Multiple Sections</CardBody>
            </Card>
          </Collapse>
        </label>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
    isMulti: state.user.isMulti,
    isMultiErrors: state.user.errors.isMulti
  };
}

export default connect(
  mapStateToProps,
  { setIsMultipleSectionsUserRequest: setIsMultipleSectionsUserRequestAction }
)(SettingsPage);
