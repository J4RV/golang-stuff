import React from 'react';

const Card = (props) => {
  const {text, isBlack, playable, handIndex, className, ...rest} = props
  return <div
      style={{transform: `rotate(${Math.random()*5 - 2.5}deg)`}}
      className={className + ` cah-card ${isBlack ? 'cah-card-black' : 'cah-card-white'}`}
      {...rest}
    >
      <span>{text}</span>
  </div>
}

export default Card;
