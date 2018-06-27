import React from 'react';
import { Nav, NavItem, NavLink } from 'reactstrap';
import { connect } from 'react-redux';
import { history } from '../../../index';

class TabNavigation extends React.Component {
  state = {
    selectedIdx: null
  };

  pageTypes = {
    HomePage: { idx: 0, type: 'HOMEPAGE' },
    ProductPage: { idx: 1, type: 'PRODUCT' },
    CategoryPage: { idx: 2, type: 'CATEGORY' },
    CartPage: { idx: 3, type: 'CART' },
    OtherPage: { idx: 4, type: 'OTHER' }
  };

  onClick = (pageType, idx) => {
    this.setState({ selectedIdx: idx });
    history.push('/' + pageType);
    if (window.DY && window.DY.API) {
      window.DY.API('spa', {
        context: { type: this.pageTypes[pageType].type, data: [] },
        // context: { type: 'PRODUCT', data: ['1217282-400'] },
        url: window.location.href,
        countAsPageview: true
      });
    }
  };

  cutPathname = pathname => {
    if (pathname.charAt(0) === '/') return pathname.slice(1);
  };

  componentDidMount() {
    this.setState({
      selectedIdx: this.pageTypes[
        this.cutPathname(history.location.pathname).idx
      ]
    });
  }

  render() {
    return (
      <div>
        <Nav tabs>
          {Object.keys(this.pageTypes).map((pageType, idx) => (
            <NavItem key={idx}>
              <NavLink
                to={'/' + pageType}
                active={idx === this.state.selectedIdx}
                onClick={() => this.onClick(pageType, idx)}
              >
                {pageType}
              </NavLink>
            </NavItem>
          ))}
        </Nav>
        <br />
        {/*<hr />*/}
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
    username: state.user.username,
    activeSection: state.user.activeSection
  };
}

export default connect(mapStateToProps)(TabNavigation);
