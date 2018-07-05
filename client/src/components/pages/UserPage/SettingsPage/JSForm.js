import React from 'react';
import { Button, Form, FormGroup, Input, FormText } from 'reactstrap';

export default class Example extends React.Component {
  render() {
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
            />
            <FormText color="muted">
              JS code to run after dynamic is loaded and before static is loaded
            </FormText>
          </FormGroup>
          <Button>Submit</Button>
        </Form>
      </div>
    );
  }
}
