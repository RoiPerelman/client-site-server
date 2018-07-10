import React from 'react';
import { Button, Form, FormGroup, Input, FormText } from 'reactstrap';
import { connect } from 'react-redux';
import { updateJSCodeRequestAction } from '../../../../store/user/actions';

class JSForm extends React.Component {
  state = {
    jsCode: ''
  };

  onChange = e => this.setState({ jsCode: e.target.value });

  componentDidMount() {
    this.setState({ jsCode: this.props.jsCode });
  }

  render() {
    const { updateJSCodeRequestAction } = this.props;
    const { jsCode } = this.state;
    const placeholder = `// we can even manipulate DYExps example:
    DYExps.hooks.beforeSmartExecution = (tagId, tagName) => {
      console.log("beforeTagExecuted", tagId, tagName)
    }
    `;
    return (
      <div>
        <Form>
          <h3>JSCode</h3>
          <br />
          <FormGroup>
            <Input
              type="textarea"
              name="text"
              id="exampleText"
              placeholder={placeholder}
              style={{ height: '200px' }}
              onChange={this.onChange}
              value={jsCode}
            />
            <FormText color="muted">
              JS code to run after dynamic is loaded and before static is loaded
            </FormText>
          </FormGroup>
          <Button onClick={() => updateJSCodeRequestAction(jsCode)}>
            Submit
          </Button>
        </Form>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
    jsCodeError: state.user.errors.jsCode,
    jsCode: state.user.jsCode
  };
}

export default connect(
  mapStateToProps,
  {
    updateJSCodeRequestAction: updateJSCodeRequestAction
  }
)(JSForm);
