import React from 'react'
import Card from './Card'
import axios from 'axios'
import Typography from '@material-ui/core/Typography'
import { withStyles } from '@material-ui/core/styles'

const styles = theme => ({
  playerPlay: {
    marginLeft: theme.spacing.unit * 2,
    marginBottom: theme.spacing.unit * 2, // add the negative bottom margin from cards
    display: "inline-block",
  },
});

const handleOnClick = (id) => {
  axios.post('rest/test/GiveBlackCardToWinner', {
    winner: id
  }).catch(r => window.alert(r.response.data)) // We'll need prettier things
}

const PlayerWhiteCardsPlayed = ({ play, isCzar, classes }) => {
  const { whiteCards } = play
  if (whiteCards == null || whiteCards.length === 0) {
    return null // The Czar will have an empty play
  }
  return (<div className={classes.playerPlay}>
    {whiteCards.map(whiteCard =>
      <Card {...whiteCard} onClick={() => isCzar && handleOnClick(play.id)} />)}
  </div>)
}

const WhiteCardsPlayed = ({ state, classes }) => {
  const isCzar = state.myPlayer.id === state.currentCzarID
  if (state.sinnerPlays.length > 0) {
    return <React.Fragment>
      <Typography variant='h6' gutterBottom>
        Czar choosing winner...
      </Typography>
      {state.sinnerPlays.map((sp) =>
        <PlayerWhiteCardsPlayed play={sp} isCzar={isCzar} classes={classes} />)}
    </React.Fragment>
  } else {
    return <Typography variant='h6' gutterBottom>
      Waiting for players...
    </Typography>
  }
}

export default withStyles(styles)(WhiteCardsPlayed)
