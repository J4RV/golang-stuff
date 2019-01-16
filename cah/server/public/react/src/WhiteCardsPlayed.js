import React from 'react'
import Card from './Card'
import axios from 'axios'

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
      <Card {...whiteCard} className='in-table' onClick={() => handleOnClick(play.ID)} />)}
  </div>)
}

const WhiteCardsPlayed = ({ state }) => {
  if (state.sinnerPlays.length > 0) {
    return <React.Fragment>
      {state.sinnerPlays.map((sp) =>
        <PlayerWhiteCardsPlayed play={sp} />)}
    </React.Fragment>
  } else {
    return <h2>Waiting for players...</h2>
  }
}

export default WhiteCardsPlayed
