import React from 'react';
import { Link } from 'react-router-dom';
import Validator from 'validator';
import { connect } from 'react-redux';
import { loginUserRequestAction } from '../../../../store/user/actions';

class LoginForm extends React.Component {
  state = {
    data: {
      email: '',
      password: ''
    },
    errors: {
      email: '',
      password: '',
      login: ''
    }
  };

  componentWillReceiveProps(nextProps) {
    this.setState({
      ...this.state,
      errors: { ...this.state.errors, login: nextProps.loginError }
    });
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
      this.props.login(this.state.data);
    }
  };

  validate = data => {
    const errors = {};
    if (!Validator.isEmail(data.email)) errors.email = 'Invalid email';
    if (!data.password) errors.password = "Can't be blank";
    return errors;
  };

  render() {
    const { data, errors } = this.state;

    return (
      <form onSubmit={this.onSubmit}>
        {(errors.login || errors.email) && (
          <div className="alert alert-danger">
            {errors.login || errors.email}
          </div>
        )}
        <div className="form-group">
          <label htmlFor="email">Email</label>
          <input
            type="email"
            id="email"
            name="email"
            value={data.email}
            onChange={this.onChange}
            className="form-control"
          />
        </div>
        {errors.password && (
          <div className="alert alert-danger">{errors.password}</div>
        )}
        <div className="form-group">
          <label htmlFor="password">Password</label>
          <input
            type="password"
            id="password"
            name="password"
            value={data.password}
            onChange={this.onChange}
            className="form-control"
          />
        </div>

        <button type="submit" className="btn btn-primary btn-block">
          Login
        </button>

        <small className="form-text text-center">
          <Link to="/guest/signup">Sign up</Link> if you don't have an account<br />
          <Link to="/forgot_password">Forgot Password?</Link>
        </small>
      </form>
    );
  }
}

function mapStateToProps(state) {
  return {
    loginError: state.user.errors.login
  };
}

export default connect(
  mapStateToProps,
  { login: loginUserRequestAction }
)(LoginForm);
