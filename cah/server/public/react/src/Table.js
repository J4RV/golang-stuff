import React from 'react'
import Card from './Card'
import WhiteCardsPlayed from './WhiteCardsPlayed'

const Table = ({state}) => {
  const card = state.blackCardInPlay
  if (card == null) return null
  return (
  <div style={{display: "flex"}}>
    <Card text={card.text} isBlack={true} className='in-table' />    
    <WhiteCardsPlayed state={state} />
  </div>)
}

export default Table