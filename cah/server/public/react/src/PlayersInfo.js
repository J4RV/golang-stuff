import React from 'react'

const PlayerInfo = ({ player, isCzar }) => (
  <div className={`cah-playerinfo hovering ${isCzar ? ' cah-czarinfo' : ''}`}>
    <div>{player.name}</div>
    <div>{player.points} points</div>
    <div>{isCzar ? 'Current Czar' : `${player.whiteCardsInPlay} card(s) in play`}</div>
  </div>
)

const PlayersInfo = ({ state }) => (
  <div className='cah-playersinfo'>
    {state.players.map((p) =>
      <PlayerInfo player={p} isCzar={p.id === state.currentCzarID} />
    )}
  </div>
)

export default PlayersInfo
