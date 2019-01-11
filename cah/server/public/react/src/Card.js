import React from 'react';

const Card = ({text, isBlack}) => (
  <div className={`hovering cah-card ${isBlack ? 'cah-card-black' : 'cah-card-white'}`}>
    <span>{text}</span>
  </div>
)

export default Card;
