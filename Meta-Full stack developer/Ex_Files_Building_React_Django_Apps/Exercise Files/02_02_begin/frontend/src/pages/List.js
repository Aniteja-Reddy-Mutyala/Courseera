import React from 'react';
import './List.css';

import { AppContext } from '../AppContext';
import Item from '../components/Item';
import ServiceApi from '../services/ServiceApi';

export default class OneItemList extends React.Component {
  static context = AppContext;

  render() {
    const { wishlist, toggleWishlist } = this.context;
    const route = "/details/1";
    const item = { id: 1, name: "Example", thumbnail_url: "/images/map_channel.gif", price: 777.77, rating: 'easy' , promo: 'Tour package description' };
    return (
      <section className="List" data-testid="list">
        <section className="List-items">
          <Item route={route} item={item}  wishlist={wishlist} toggleWishlist={toggleWishlist} />
        </section>
      </section>
    );
  }
}
