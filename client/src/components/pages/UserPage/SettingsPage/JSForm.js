import React from 'react';
import { Button, Form, FormGroup, Input, FormText } from 'reactstrap';

export default class Example extends React.Component {
  render() {
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
              placeholder="console.log('it works')"
              style={{ height: '200px' }}
            />
            <FormText color="muted">
              JS code to run before DY is loaded
            </FormText>
          </FormGroup>
          <Button>Submit</Button>
        </Form>
      </div>
    );
  }
}
