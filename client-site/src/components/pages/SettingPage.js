import React from 'react';
import { Collapse, Card, CardBody } from 'reactstrap';
import { setIsMultipleSectionsUserRequestAction } from '../../store/user/actions';
import { connect } from 'react-redux';

class SettingsPage extends React.Component {
  constructor(props) {
    super(props);
    this.toggle = this.toggle.bind(this);
    this.state = {
      collapse: false,
      errors: ''
    };
  }

  componentDidMount() {
    const { isMulti } = this.props;
    this.setState({ collapse: isMulti });
  }

  componentWillReceiveProps(nextProps) {
    const { isMulti } = nextProps;
    this.setState({ collapse: isMulti });
  }

  toggle() {
    this.props.setIsMultipleSectionsUserRequest(!this.state.collapse);
  }

  render() {
    const { collapse, errors } = this.state;

    return (
      <div>
        <label className="switch">
          {errors && <div className="alert alert-danger">{errors}</div>}
          <input type="checkbox" checked={collapse} onClick={this.toggle} />
          <span className="slider round"> use multiple sections</span>
          <Collapse isOpen={collapse}>
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
    errors: state.user.errors.isMulti
  };
}

export default connect(
  mapStateToProps,
  { setIsMultipleSectionsUserRequest: setIsMultipleSectionsUserRequestAction }
)(SettingsPage);
