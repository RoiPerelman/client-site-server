import React from 'react';
import { Table } from 'reactstrap';
import {
  addContextItemRequestAction,
  delContextItemRequestAction
} from '../../../../store/user/actions';
import { connect } from 'react-redux';

class ContextForm extends React.Component {
  state = {
    contextItem: ''
  };

  onChange = e => this.setState({ contextItem: e.target.value });

  render() {
    const {
      section,
      contextType,
      context,
      addContextItemRequestAction,
      delContextItemRequestAction
    } = this.props;
    return (
      <div>
        <label className="switch">
          <input type="text" placeholder={`Item`} onChange={this.onChange} />
          <button
            type="button"
            onClick={() =>
              addContextItemRequestAction({
                sectionsId: section.id,
                sectionId: section.sectionId,
                contextType: contextType,
                item: this.state.contextItem
              })
            }
            className="small"
          >
            Add Item
          </button>
          <button
            type="button"
            onClick={() =>
              delContextItemRequestAction({
                sectionsId: section.id,
                sectionId: section.sectionId,
                contextType: contextType,
                item: this.state.contextItem
              })
            }
            className="small"
          >
            Del Item
          </button>
          <Table>
            <thead>
              <tr>
                <th>#</th>
                <th>Item</th>
              </tr>
            </thead>
            <tbody>
              {Object.keys(context).map(idx => (
                <tr key={idx}>
                  <th scope="row">{idx}</th>
                  <td>{context[idx]}</td>
                </tr>
              ))}
            </tbody>
          </Table>
        </label>
      </div>
    );
  }
}

export default connect(
  null,
  {
    addContextItemRequestAction: addContextItemRequestAction,
    delContextItemRequestAction: delContextItemRequestAction
  }
)(ContextForm);
