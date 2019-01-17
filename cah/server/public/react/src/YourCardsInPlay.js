import React from 'react'
import Card from './Card'
import Typography from '@material-ui/core/Typography'

const YourCardsInPlay = ({ state }) => {
  if (state.currentCzarID === state.myPlayer.id) {
    return <Typography variant='h4' gutterBottom>
      You are the Czar!
    </Typography>
  }
  const cards = state.myPlayer.whiteCardsInPlay.map(c =>
    <Card text={c.text} playable={false} className='in-table' />
  )
  if (cards == null || cards.length === 0) {
    return <Typography variant='h4' gutterBottom>
      Play {state.blackCardInPlay.blanksAmount} cards
    </Typography>
  }
  return <div className='cah-oneplayerwhitecards'>{cards}</div>
}

export default YourCardsInPlay
