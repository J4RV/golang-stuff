import React from 'react'

const PlayerInfo = ({ player, isCzar }) => (
  <div className={`cah-playerinfo hovering ${isCzar ? ' cah-czarinfo' : ''}`}>
    <div>{player.name}</div>
    <div>{player.points.length} points</div>
    <div>{isCzar ? 'Current Czar' : `${player.whiteCardsInPlay.length} card(s) in play`}</div>
  </div>
)

const PlayersInfo = ({ state }) => (
  <div className='cah-playersinfo'>
    {state.players.map((p, i) =>
      <PlayerInfo player={p} isCzar={i === parseInt(state.currentCzarIndex, 10)} />
    )}
  </div>
)

export default PlayersInfo
