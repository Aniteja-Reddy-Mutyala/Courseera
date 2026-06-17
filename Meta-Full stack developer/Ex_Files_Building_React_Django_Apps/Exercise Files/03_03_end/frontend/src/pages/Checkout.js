import React from 'react';
import './Checkout.css';
import { AppContext } from '../AppContext';
import { FormField } from '../components/form/FormField';
import { FaMinusCircle } from 'react-icons/fa';

export default class Checkout extends React.Component {
  static contextType = AppContext;

  render() {
    const { booking, item, updateField, clearOrderItem } = this.context;
    const inputFields = [
      { label: 'Name', name: 'name' },
      { label: 'Email Address', name: 'email_address' },
      { label: 'Street Address', name: 'street_address' },
      { label: 'City', name: 'city' }
    ];
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
            <button>
              Place order
            </button>
          </div>
        </section>
      </section>
    );
  }
}
