import React from 'react'

const Card = (props) => {
  const { text, isBlack, expansion, className, ...rest } = props
  return <div
    style={{ transform: `rotate(${Math.random() * 5 - 2.5}deg)` }}
    className={className + ` cah-card ${isBlack ? 'cah-card-black' : 'cah-card-white'}`}
    {...rest}
  >
    <span>{text}</span>
    <div className="cah-card-expansion">{expansion}</div>
  </div>
}

export default Card
