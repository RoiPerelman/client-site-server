import React from 'react';
import { Container, Row, Col } from 'reactstrap';
import Logo from '../../../Logo.svg';
import LoginPage from './LoginPage';
import SignupPage from './SignupPage';
import { Route } from 'react-router-dom';

const GuestPage = ({ isLogin }) => (
  <Container
    fluid
    style={{
      height: '100vh',
      color: 'black',
      background: 'linear-gradient(to right, #3A7C9C, #1b1819)'
    }}
  >
    <Row
      className="align-items-center justify-content-center text-center"
      style={{ height: '100%' }}
    >
      <Col xs={12} sm={6}>
        <img className="img-fluid" alt="Logo" src={Logo} />
      </Col>
      <Col
        xs={12}
        sm={6}
        style={{
          fontFamily: "'Open Sans', sans-serif"
        }}
      >
        <Route
          path="/guest/signup"
          exact
          render={() => {
            return <SignupPage />;
          }}
        />
        <Route
          path="/guest/login"
          exact
          render={() => {
            return <LoginPage />;
          }}
        />
      </Col>
    </Row>
  </Container>
);

export default GuestPage;
