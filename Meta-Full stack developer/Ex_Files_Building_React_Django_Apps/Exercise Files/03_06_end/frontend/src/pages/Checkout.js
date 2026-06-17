import React from 'react';
import './Checkout.css';
import { AppContext } from '../AppContext';
import { FormField } from '../components/form/FormField';
import { FaMinusCircle } from 'react-icons/fa';

export default class Checkout extends React.Component {
  static contextType = AppContext;

  constructor(props) {
    super(props);
    this.state = { orderPlaced: false };
  }

  placeOrder() {
    this.context.placeOrder().then(() => {
      this.setState({ orderPlaced: true });
    }).catch((validationErrors) => {
      this.setState({ orderPlaced: false });
    })
  }

  render() {
    const { booking, item, updateField, clearOrderItem } = this.context;
    const { orderPlaced } = this.state;
    const inputFields = [
      { label: 'Name', name: 'name' },
      { label: 'Email Address', name: 'email_address' },
      { label: 'Street Address', name: 'street_address' },
      { label: 'City', name: 'city' }
    ];
    if (orderPlaced) {
      return (
        <section className="Checkout">
          <header className="Checkout-header">
            <h2>Checkout</h2>
            <h3>Thanks for buying an excursion with Explore California!</h3>
          </header>
        </section>
      );
    }
    const formFields = inputFields.map((fieldProps) => {
      return (
        <FormField
          key={fieldProps.name}
          value={booking[fieldProps.name]}
          onUpdate={updateField}
          {...fieldProps}
          />
      );
      }
    );
    let displayItem;
    if (item) {
      displayItem = (
        <div>
          <button className="Checkout-package-remove" onClick={() => clearOrderItem(item.id)}>
            <FaMinusCircle />
          </button>
          {item.name} - ${item.price} starts on {item.start} for {item.tour_length} days.
        </div>
      );
    }
    return (
      <section className="Checkout">
        <header className="Checkout-header">
          <h2>Checkout</h2>
        </header>
        <section className="Checkout-summary">
          {displayItem}
        </section>
        <section className="Checkout-form">
          <form>
            {formFields}
          </form>
        </section>
        <section className="Checkout-actions">
          <div className="Checkout-actions__next">
            <button onClick={this.placeOrder.bind(this)} disabled={item === null}>
              Place order
            </button>
          </div>
        </section>
      </section>
    );
  }
}
