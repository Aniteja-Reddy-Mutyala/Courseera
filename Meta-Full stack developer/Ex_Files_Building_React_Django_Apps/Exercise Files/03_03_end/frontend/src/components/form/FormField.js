import React from 'react';
import './FormField.css';

export function FormField(props) {
  const { name } = props;
  return (
    <div key={name} className="FormField">
    </div>
  );
}
