import React from 'react'
import Card from './Card'
import axios from 'axios'
import LocalPlayerIndex from './LocalPlayerIndex'

const handleOnClick = (id) => {
  console.log('Chose winner: ', id)
  axios.post('rest/test/' + LocalPlayerIndex() + '/GiveBlackCardToWinner', {
    winner: id
  }).catch(r => window.alert(r.response.data)) // We'll need prettier things
}

const PlayerWhiteCardsPlayed = ({ player, playerindex }) => {
  const { whiteCardsInPlay } = player
  if (whiteCardsInPlay == null || whiteCardsInPlay.length === 0) return null
  return (<div className='cah-oneplayerwhitecards'>
    {whiteCardsInPlay.map(whiteCard =>
      <Card {...whiteCard} className='in-table' onClick={() => handleOnClick(playerindex)} />)}
  </div>)
}

const allSinnersPlayed = (state) => {
  for (let i = 0; i < state.players.length; i++) {
    if (i === parseInt(state.currentCzarIndex, 10)) {
      continue
    }
    if (state.players[i].whiteCardsInPlay.length !== parseInt(state.blackCardInPlay.blanksAmount, 10)) {
      return false
    }
  }
  return true
}

const WhiteCardsPlayed = ({ state }) => {
  if (allSinnersPlayed(state)) {
    return <React.Fragment>
      {state.players.map((p, i) =>
        <PlayerWhiteCardsPlayed player={p} playerindex={i} />)}
    </React.Fragment>
  } else {
    return <h2>Waiting for players...</h2>
  }
}

export default WhiteCardsPlayed
