import React from 'react';
import { Route, Switch } from 'react-router-dom';
import Selectors from './Selectors';

class Page extends React.Component {
  // TODO find a way to re render on each route without adding the div's
  render() {
    return (
      <Switch>
        <Route path="/homepage" exact render={() => <Selectors />} />
        <Route
          path="/productpage"
          exact
          render={() => (
            <div>
              <div />
              <Selectors />
            </div>
          )}
        />
        <Route
          path="/categorypage"
          exact
          render={() => (
            <div>
              <div />
              <div />
              <Selectors />
            </div>
          )}
        />
        <Route
          path="/cartpage"
          exact
          render={() => (
            <div>
              <div />
              <div />
              <div />
              <Selectors />
            </div>
          )}
        />
        <Route
          path="/"
          exact
          render={() => (
            <div>
              this is where in the future we will add monitoring features
            </div>
          )}
        />
      </Switch>
    );
  }
}

export default Page;
