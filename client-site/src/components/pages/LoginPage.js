import React from 'react';
import LoginForm from '../forms/LoginForm';

const LoginPage = () => {
  return (
    <div>
      <div className="container" style={{ height: '100vh' }}>
        <div className="row align-items-center" style={{ height: '100vh' }}>
          <div className="col col-xs-12 col-sm-8 offset-sm-2 col-lg-6 offset-lg-3">
            <div className="card">
              <h2 className="card-header">Welcome Back!</h2>
              <div className="card-body">
                <LoginForm />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default LoginPage;
