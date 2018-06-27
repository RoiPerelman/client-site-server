import React from 'react';
import { connect } from 'react-redux';
import PropTypes from 'prop-types';
import { Link } from 'react-router-dom';
import isEmail from 'validator/lib/isEmail';
import { signupUserRequestAction } from '../../../../store/user/actions';

class SignupForm extends React.Component {
  state = {
    data: {
      email: '',
      username: '',
      password: '',
      defaultSection: ''
    },
    errors: {
      email: '',
      username: '',
      passwords: '',
      server: ''
    }
  };

  componentWillReceiveProps(nextProps) {
    this.setState({ ...this.state, errors: nextProps.errors });
  }

  onChange = e =>
    this.setState({
      data: { ...this.state.data, [e.target.name]: e.target.value }
    });

  onSubmit = e => {
    e.preventDefault();
    const errors = this.validate(this.state.data);
    this.setState({ errors });
    if (Object.keys(errors).length === 0) {
      this.props.submit(this.state.data);
    }
  };

  validate = data => {
    const errors = {};
    if (!isEmail(data.email)) {
      errors.email = 'Invalid email';
    }
    if (!data.password) {
      errors.password = "Can't be blank";
    }
    if (!data.username) {
      errors.username = "Can't be blank";
    }
    return errors;
  };

  render() {
    const { data, errors } = this.state;

    return (
      <form onSubmit={this.onSubmit}>
        <div className="form-group">
          <label htmlFor="email">Email</label>
          <input
            type="email"
            id="email"
            name="email"
            value={data.email}
            onChange={this.onChange}
            className={
              errors.email ? 'form-control is-invalid' : 'form-control'
            }
          />
          <div className="invalid-feedback">{errors.email}</div>
        </div>

        <div className="form-group">
          <label htmlFor="username">Username</label>
          <input
            type="text"
            id="username"
            name="username"
            value={data.username}
            onChange={this.onChange}
            className={
              errors.username ? 'form-control is-invalid' : 'form-control'
            }
          />
          <div className="invalid-feedback">{errors.username}</div>
        </div>

        <div className="form-group">
          <label htmlFor="defaultSection">defaultSection</label>
          <input
            type="text"
            id="defaultSection"
            name="defaultSection"
            value={data.defaultSection}
            onChange={this.onChange}
            className={
              errors.defaultSection ? 'form-control is-invalid' : 'form-control'
            }
          />
          <div className="invalid-feedback">{errors.defaultSection}</div>
        </div>

        <div className="form-group">
          <label htmlFor="password">Password</label>
          <input
            type="password"
            id="password"
            name="password"
            value={data.password}
            onChange={this.onChange}
            className={
              errors.password ? 'form-control is-invalid' : 'form-control'
            }
          />
          <div className="invalid-feedback">{errors.password}</div>
        </div>

        <button type="submit" className="btn btn-primary btn-block">
          Sign Up
        </button>

        <small className="form-text text-center">
          or <Link to="/guest/login">Login</Link> if you have an account
        </small>
      </form>
    );
  }
}

function mapStateToProps(state) {
  return {
    errors: state.user.errors
  };
}

SignupForm.propTypes = {
  submit: PropTypes.func.isRequired
};

export default connect(
  mapStateToProps,
  { submit: signupUserRequestAction }
)(SignupForm);
