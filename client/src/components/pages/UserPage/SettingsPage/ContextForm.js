import React from 'react';

export default class Example extends React.Component {
  render() {
    return (
      <div className="row">
        <div className="col-6">
          <h3>Contexts</h3>
          <br />
          <h5>Product Context</h5>
          <input
            type="text"
            placeholder={`['sku']`}
            value=""
            onChange={this.onChange}
          />
          <button type="button" onClick={console.log} className="small">
            Add Section
          </button>
          <h5>Cart Context</h5>
          <input
            type="text"
            placeholder={`['sku1', 'sku2', ...]`}
            value=""
            onChange={this.onChange}
          />
          <button type="button" onClick={console.log} className="small">
            Add Section
          </button>
          <h5>Category Context</h5>
          <input
            type="text"
            placeholder={`['category1', 'category2', ...]`}
            value=""
            onChange={this.onChange}
          />
          <button type="button" onClick={console.log} className="small">
            Add Section
          </button>
          <hr />
        </div>
      </div>
    );
  }
}
