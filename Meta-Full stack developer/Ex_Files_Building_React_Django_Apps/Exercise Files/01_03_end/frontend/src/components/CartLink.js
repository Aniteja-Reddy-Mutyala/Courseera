import React from 'react';
import { FaShoppingCart } from 'react-icons/fa';

export default function CartLink(props) {
  const { item } = props;
  return (
    <span className="Cart">
      <FaShoppingCart />
      {item && item.name}
    </span>
  );
}
