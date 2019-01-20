import React from 'react'
import Card from './Card'
import axios from 'axios'
import Typography from '@material-ui/core/Typography'
import { withStyles } from '@material-ui/core/styles'
import {chooseWinnerUrl} from '../ServerUrls'

const styles = theme => ({
  playerPlay: {
    marginLeft: theme.spacing.unit * 2,
    marginBottom: theme.spacing.unit * 2, // add the negative bottom margin from cards
    display: "inline-block",
  },
});

const handleOnClick = (stateID, winnerID) => {
  axios.post(chooseWinnerUrl(stateID), {
    winner: winnerID
  }).catch(r => window.alert(r.response.data)) // We'll need prettier things
}

const PlayerWhiteCardsPlayed = ({ stateID, play, isCzar, classes }) => {
  const { whiteCards } = play
  if (whiteCards == null || whiteCards.length === 0) {
    return null // The Czar will have an empty play
  }
  return (<div className={classes.playerPlay}>
    {whiteCards.map(whiteCard =>
      <Card {...whiteCard} onClick={() => isCzar && handleOnClick(stateID, play.id)} />)}
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
        <PlayerWhiteCardsPlayed 
          stateID={state.id}
          play={sp} 
          isCzar={isCzar} 
          classes={classes} 
        />)}
    </React.Fragment>
  } else {
    return <Typography variant='h6' gutterBottom>
      Waiting for players...
    </Typography>
  }
}

export default withStyles(styles)(WhiteCardsPlayed)
