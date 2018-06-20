import React from 'react';
import {
  Collapse,
  Navbar,
  NavbarToggler,
  NavbarBrand,
  Nav,
  NavItem,
  NavLink,
  UncontrolledDropdown,
  DropdownToggle,
  DropdownMenu,
  DropdownItem
} from 'reactstrap';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';
import { logoutUserAction } from '../../store/user/actions';

class Example extends React.Component {
  constructor(props) {
    super(props);

    this.toggle = this.toggle.bind(this);
    this.state = {
      isOpen: false
    };
  }
  toggle() {
    this.setState({
      isOpen: !this.state.isOpen
    });
  }
  render() {
    const { username, logout } = this.props;
    const selectors = Array.from(Array(5)).map((value, idx) => {
      return (
        <div
          className="d-md-flex flex-md-equal w-100 my-md-3 pl-md-3"
          key={idx}
        >
          <div className="col-sm bg-dark mr-md-3 pt-3 px-3 pt-md-5 px-md-5 text-center text-white overflow-hidden">
            <div className="my-3 py-3">
              <h2 className="display-5">{`selector${idx * 2 + 1}`}</h2>
              <p className="lead" id={`selector${idx * 2 + 1}`} />
            </div>
          </div>
          <div className="col-sm bg-light mr-md-3 pt-3 px-3 pt-md-5 px-md-5 text-center overflow-hidden">
            <div className="my-3 p-3">
              <h2 className="display-5">{`selector${idx * 2 + 2}`}</h2>
              <p className="lead" id={`selector${idx * 2 + 2}`} />
            </div>
          </div>
        </div>
      );
    });
    return (
      <div>
        <Navbar color="dark" dark expand="md">
          <NavbarBrand href="/">DynamicYield RP</NavbarBrand>
          <NavbarToggler onClick={this.toggle} />
          <Collapse isOpen={this.state.isOpen} navbar>
            <Nav className="ml-auto" navbar>
              <NavItem>
                <NavLink href="https://github.com/RoiPerelman/client-site">
                  GitHub
                </NavLink>
              </NavItem>
              <UncontrolledDropdown nav inNavbar>
                <DropdownToggle nav caret>
                  {username}
                </DropdownToggle>
                <DropdownMenu right>
                  <DropdownItem>
                    <Link to="/settings" style={{ textDecoration: 'none' }}>
                      settings
                    </Link>
                  </DropdownItem>
                  <DropdownItem divider />
                  <DropdownItem onClick={logout}>logout</DropdownItem>
                </DropdownMenu>
              </UncontrolledDropdown>
            </Nav>
          </Collapse>
        </Navbar>
        <div className="container-fluid">{selectors}</div>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
    username: state.user.username
  };
}

export default connect(
  mapStateToProps,
  { logout: logoutUserAction }
)(Example);
