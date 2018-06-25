import React from 'react';
import { Collapse, Card, CardBody } from 'reactstrap';
import { setIsMultipleSectionsUserRequestAction } from '../../store/user/actions';
import { connect } from 'react-redux';

class SettingsPage extends React.Component {
  constructor(props) {
    super(props);
    this.toggle = this.toggle.bind(this);
  }

  toggle() {
    this.props.setIsMultipleSectionsUserRequest(!this.props.user.isMulti);
  }

  render() {
    const { isMulti, errors } = this.props.user;

    return (
      <div>
        <label className="switch">
          {errors.isMulti && (
            <div className="alert alert-danger">{errors.isMulti}</div>
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
    user: state.user
  };
}

export default connect(
  mapStateToProps,
  { setIsMultipleSectionsUserRequest: setIsMultipleSectionsUserRequestAction }
)(SettingsPage);
