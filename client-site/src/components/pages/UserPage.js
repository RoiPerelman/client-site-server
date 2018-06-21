import React from 'react';
import TopNavigation from '../navigation/TopNavigation';
import SettingsPage from './SettingPage';
import { Route } from 'react-router-dom';

const UserPage = () => {
  const selectors = Array.from(Array(10)).map((value, idx) => {
    return (
      <div className="d-md-flex flex-md-equal w-100 my-md-3 pl-md-3" key={idx}>
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
    <div className="container-fluid">
      <TopNavigation />
      <hr />
      <hr />
      <hr />
      <Route
        path="/settings"
        exact
        render={() => {
          return <SettingsPage />;
        }}
      />
      <Route
        path="/"
        exact
        render={() => {
          return <div className="">{selectors}</div>;
        }}
      />
    </div>
  );
};

export default UserPage;
