import React from 'react';
import Card from './Card'

const YourCardsInPlay = ({state, owner}) => {
  const cards = state.players[owner].whiteCardsInPlay.map(c =>
    <Card text={c.text} playable={false} className='in-table' />            
  )
  if(cards == null || cards.length === 0){
    return <h2>Play {state.blackCardInPlay.blanksAmount} cards</h2>
  }
  return <span>
    <h2>Your cards played this round:</h2>
    {cards}
  </span>
}

export default YourCardsInPlay