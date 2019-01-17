import React from 'react'
import Typography from '@material-ui/core/Typography'
import withWidth from '@material-ui/core/withWidth'

const PlayerInfo = ({ player, itsYou, isCzar }) => (
  <div className='cah-playerinfo hovering'>
    <div>{player.name} {itsYou ? <b>(You)</b> : null}</div>
    <div>{player.points.length} points</div>
    <div>{isCzar ? <b>Current Czar</b> : `${player.whiteCardsInPlay} card(s) in play`}</div>
  </div>
)

let PlayersInfo = ({ width, state }) => {
  if (width === 'sm' || width === 'xs') {
    return null
  }
  return <div className='cah-playersinfo'>
    <Typography>
      {state.players.map((p) =>
        <PlayerInfo
          player={p}
          itsYou={p.id === state.myPlayer.id}
          isCzar={p.id === state.currentCzarID}
        />
      )}
    </Typography>
  </div>
}
PlayersInfo = withWidth()(PlayersInfo)

export default PlayersInfo
