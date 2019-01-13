import React from 'react';
import Card from './Card'
import LocalPlayerIndex from './LocalPlayerIndex'

const YourCardsInPlay = ({state}) => {
  if(state.currentCzarIndex == LocalPlayerIndex()){
    return <h2>You are the Czar!</h2>
  }
  const cards = state.players[LocalPlayerIndex()].whiteCardsInPlay.map(c =>
    <Card text={c.text} playable={false} className='in-table' />            
  )
  if(cards == null || cards.length === 0){
    return <h2>Play {state.blackCardInPlay.blanksAmount} cards</h2>
  }
  return <span>
    <h2>Your card(s) played this round:</h2>
    {cards}
  </span>
}

export default YourCardsInPlay