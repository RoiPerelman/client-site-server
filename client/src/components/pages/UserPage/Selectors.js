import React from 'react';

const Selectors = () => {
  let bool = false;

  const selector = (idx, bool) => {
    return (
      <div
        className={
          'col-sm mr-md-3 pt-3 px-3 pt-md-5 px-md-5 text-center overflow-hidden' +
          (bool ? ' bg-dark text-white' : ' bg-light text-black')
        }
      >
        <div className="my-3 py-3">
          <h2 className="display-5">{`selector${idx}`}</h2>
          <p className="lead" id={`selector${idx}`} />
        </div>
      </div>
    );
  };

  return Array.from(Array(10)).map((value, idx) => {
    return (
      <div className="d-md-flex flex-md-equal w-100 my-md-3 pl-md-3" key={idx}>
        {selector(idx * 2 + 1, bool)}
        {(bool = !bool)}
        {selector(idx * 2 + 2, bool)}
      </div>
    );
  });
};

export default Selectors;
