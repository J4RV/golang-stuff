import React from 'react';
import Card from './Card'

const PlayerWhiteCardsPlayed = ({player}) => {
  const {whiteCardsInPlay} = player
  return (<React.Fragment>
    {whiteCardsInPlay.map(whiteCard => <Card {...whiteCard} className='in-table' />)}
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
    return <React.Fragment>{state.players.map(p => <PlayerWhiteCardsPlayed player={p} />)}</React.Fragment>
  } else {
    return <h2>Waiting for players...</h2>
  }
}

export default WhiteCardsPlayed