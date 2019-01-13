import React from 'react';
import Card from './Card'
import axios from 'axios'
import LocalPlayerIndex from './LocalPlayerIndex'

const handleOnClick = ({id}) => {
  axios.post('rest/test/'+LocalPlayerIndex()+'/GiveBlackCardToWinner', {
    winner: id
  }).catch(r => window.alert(r.response.data)); // We'll need prettier things
}

const PlayerWhiteCardsPlayed = ({player, playerindex}) => {
  const {whiteCardsInPlay} = player
  return (<React.Fragment>
    {whiteCardsInPlay.map(whiteCard =>
      <Card {...whiteCard} className='in-table' onClick={() => handleOnClick(playerindex)} />)}
  </React.Fragment>)
}

const allSinnersPlayed = (state) => {
  for(let i = 0; i < state.players.length; i++){
    if(i == state.currentCzarIndex){
      continue
    }
    if(state.players[i].whiteCardsInPlay.length != state.blackCardInPlay.blanksAmount){
      return false
    }
  }
  return true
}

const WhiteCardsPlayed = ({state}) => {
  if (allSinnersPlayed(state)){
    return <React.Fragment>
      {state.players.map((p, i) => 
        <PlayerWhiteCardsPlayed player={p} playerindex={i} />)}      
    </React.Fragment>
  } else {
    return <h2>Waiting for players...</h2>
  }
}

export default WhiteCardsPlayed