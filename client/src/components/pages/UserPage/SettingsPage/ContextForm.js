import React from 'react';
import ContextForm2 from './ContextForm2';
import {
  addProductContextRequestAction,
  delProductContextRequestAction,
  addCartContextRequestAction,
  delCartContextRequestAction,
  addCategoryContextRequestAction,
  delCategoryContextRequestAction
} from '../../../../store/user/actions';
import { connect } from 'react-redux';

class ContextsForm extends React.Component {
  state = {
    productItemInput: '',
    cartItemInput: '',
    categoryItemInput: ''
  };

  onChangeProduct = e => this.setState({ productItemInput: e.target.value });
  onChangeCart = e => this.setState({ cartItemInput: e.target.value });
  onChangeCategory = e => this.setState({ categoryItemInput: e.target.value });

  addProductItem = () =>
    this.props.addProductContextRequestAction(this.state.inputSection);
  delProductItem = () =>
    this.props.delProductContextRequestAction(this.state.inputSection);

  addCartItem = () =>
    this.props.addCartContextRequestAction(this.state.inputSection);
  delCartItem = () =>
    this.props.delCartContextRequestAction(this.state.inputSection);

  addCategoryItem = () =>
    this.props.addCategoryContextRequestAction(this.state.inputSection);
  delCategoryItem = () =>
    this.props.delCategoryContextRequestAction(this.state.inputSection);

  render() {
    const { productContext, cartContext, categoryContext } = this.props;
    return (
      <div className="row">
        <div className="col-6">
          <h3>Contexts</h3>
          <br />
          <h3>Product Context</h3>
          <ContextForm2
            state={this.state.productItemInput}
            onChange={this.onChangeProduct}
            context={productContext}
            addItem={this.addProductItem}
            delItem={this.delProductItem}
          />
          <h3>Cart Context</h3>
          <ContextForm2
            state={this.state.cartItemInput}
            onChange={this.onChangeCart}
            context={cartContext}
            addItem={this.addCartItem}
            delItem={this.delCartItem}
          />
          <h3>Category Context</h3>
          <ContextForm2
            state={this.state.categoryItemInput}
            onChange={this.onChangeCategory}
            context={categoryContext}
            addItem={this.addCategoryItem}
            delItem={this.delCategoryItem}
          />
        </div>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
    productContext: state.user.contexts.productContext,
    cartContext: state.user.contexts.cartContext,
    categoryContext: state.user.contexts.categoryContext
  };
}

export default connect(
  mapStateToProps,
  {
    addProductContextRequestAction: addProductContextRequestAction,
    delProductContextRequestAction: delProductContextRequestAction,
    addCartContextRequestAction: addCartContextRequestAction,
    delCartContextRequestAction: delCartContextRequestAction,
    addCategoryContextRequestAction: addCategoryContextRequestAction,
    delCategoryContextRequestAction: delCategoryContextRequestAction
  }
)(ContextsForm);
