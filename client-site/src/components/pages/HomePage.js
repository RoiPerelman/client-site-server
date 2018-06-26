import React from 'react';

class HomePage extends React.Component {
  selectors = Array.from(Array(10)).map((value, idx) => {
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

  render() {
    return <div>{this.selectors}</div>;
  }
}

// export default connect(
//   mapStateToProps,
//   { loadDynamicYieldRequestAction: loadDynamicYieldRequestAction}
// )(UserPage);

export default HomePage;
