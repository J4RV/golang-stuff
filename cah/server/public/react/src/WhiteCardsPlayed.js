import React from 'react'
import Card from './Card'
import axios from 'axios'
import Typography from '@material-ui/core/Typography'

const handleOnClick = (id) => {
  console.log('Chose winner: ', id)
  axios.post('rest/test/GiveBlackCardToWinner', {
    winner: id
  }).catch(r => window.alert(r.response.data)) // We'll need prettier things
}

const PlayerWhiteCardsPlayed = ({ play }) => {
  const { whiteCards } = play
  if (whiteCards == null || whiteCards.length === 0) {
    return null // The Czar will have an empty play
  }
  return (<div className='cah-oneplayerwhitecards'>
    {whiteCards.map(whiteCard =>
      <Card {...whiteCard} className='in-table' onClick={() => handleOnClick(play.id)} />)}
  </div>)
}

const WhiteCardsPlayed = ({ state }) => {
  if (state.sinnerPlays.length > 0) {
    return <React.Fragment>
      {state.sinnerPlays.map((sp) =>
        <PlayerWhiteCardsPlayed play={sp} />)}
    </React.Fragment>
  } else {
    return <Typography variant='h4' gutterBottom>
      Waiting for players...
    </Typography>
  }
}

export default WhiteCardsPlayed
