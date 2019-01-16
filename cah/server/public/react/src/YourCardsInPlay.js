import React from 'react'
import Card from './Card'

const YourCardsInPlay = ({ state }) => {
  if (state.currentCzarID == state.myPlayer.ID) {
    return <h2>You are the Czar!</h2>
  }
  const cards = state.myPlayer.whiteCardsInPlay.map(c =>
    <Card text={c.text} playable={false} className='in-table' />
  )
  if (cards == null || cards.length === 0) {
    return <h2>Play {state.blackCardInPlay.blanksAmount} cards</h2>
  }
  return <div>
    <h2>Your card(s) played this round:</h2>
    {cards != null && cards.length > 0
      ? <div className='cah-oneplayerwhitecards'>{cards}</div>
      : null}
  </div>
}

export default YourCardsInPlay
