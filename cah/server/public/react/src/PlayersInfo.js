import React, { Component } from 'react';

const PlayerInfo = ({player}) => (
  <div className="cah-playerinfo">
    <p>{player.name}</p>
    <p>{player.points.length} points</p>
    <p>{player.whiteCardsInPlay.length} cards in play</p>
  </div>
)

const PlayersInfo = ({state}) => (
  <div className="cah-playersinfo">
    {state.players.map(p => 
      <PlayerInfo player={p} />
    )}
  </div>
)

export default PlayersInfo