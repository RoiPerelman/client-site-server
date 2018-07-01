import React from 'react';
import { Table } from 'reactstrap';

class ContextForm2 extends React.Component {
  render() {
    const { state, onChange, context, addItem, delItem } = this.props;
    return (
      <div>
        <label className="switch">
          <input
            type="text"
            placeholder={`Item`}
            value={state}
            onChange={onChange}
          />
          <button type="button" onClick={addItem} className="small">
            Add Item
          </button>
          <button type="button" onClick={delItem} className="small">
            Del Item
          </button>
          <Table>
            <thead>
              <tr>
                <th>#</th>
                <th>id</th>
                <th>Item</th>
              </tr>
            </thead>
            <tbody>
              {Object.keys(context).map(idx => (
                <tr key={idx}>
                  <th scope="row">{idx}</th>
                  <td>{context[idx].id}</td>
                  <td>{context[idx].sectionId}</td>
                  <td>{context[idx].name || 'noName'}</td>
                </tr>
              ))}
            </tbody>
          </Table>
        </label>
      </div>
    );
  }
}

export default ContextForm2;
