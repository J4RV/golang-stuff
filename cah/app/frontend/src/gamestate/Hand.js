import React, { Component } from 'react'

import Button from '@material-ui/core/Button'
import Card from './Card'
import Check from '@material-ui/icons/Check'
import Fab from '@material-ui/core/Fab'
import Typography from '@material-ui/core/Typography'
import axios from 'axios'
import { connect } from "react-redux"
import { playCardsUrl } from '../restUrls'
import pushError from "../actions/pushError"
import withWidth from '@material-ui/core/withWidth'

let PlayCardsButton = ({ isCzar, width, playCards }) => {
  if (isCzar) {
    return null
  }
  if (width === "sm" || width === "xs") {
    return <Fab
      aria-label="Play selected cards"
      color="primary"
      onClick={playCards}
      style={{ position: "fixed", right: 8, bottom: 8 }}>
      <Check />
    </Fab>
  } else {
    return <Button
      variant="contained"
      color="primary"
      onClick={playCards}
    >
      Play cards
    </Button>
  }
}
PlayCardsButton = withWidth()(PlayCardsButton)

const CardsToPlay = ({ state }) => {
  const isCzar = state.myPlayer.id === state.currentCzarID
  if (isCzar) {
    return null
  }
  const cardsToPlay = state.blackCardInPlay.blanksAmount - state.myPlayer.whiteCardsInPlay.length
  if (cardsToPlay === 0) {
    return null
  }
  return <Typography variant='h6' gutterBottom>
    Play {cardsToPlay} cards
  </Typography>
}

class Hand extends Component {
  state = { cardIndexes: [], errormsg: null }

  render() {
    const gamestate = this.props.state
    return (
      <div className="cah-hand">
        <CardsToPlay state={gamestate} />
        <div className="cah-hand-cards">
          {gamestate.myPlayer.hand.map((c, i) =>
            <Card
              {...c}
              handIndex={i}
              elevated
              inHand
              glowing={this.state.cardIndexes.includes(i)}
              onClick={() => this.handleCardClick(i)}
            />
          )}
        </div>
        <div style={{ marginTop: "2rem" }}>
          <PlayCardsButton playCards={this.playCards} isCzar={this.isCzar()} />
        </div>
      </div>
    )
  }

  isCzar = () => {
    return this.props.state.currentCzarID === this.props.state.myPlayer.id
  }

  handleCardClick = (i) => {
    if (this.isCzar()) {
      return
    }
    let newList = this.state.cardIndexes.slice()
    if (newList.includes(i)) {
      newList.splice(newList.indexOf(i), 1)
    } else {
      newList.push(i)
    }
    this.setState({ cardIndexes: newList })
  }

  playCards = () => {
    axios.post(playCardsUrl(this.props.state.id), {
      cardIndexes: this.state.cardIndexes
    }).then(r => {
      this.setState({ cardIndexes: [] })
    }).catch(r => this.props.pushError(r))
  }
}

export default connect(
  () => { },
  { pushError }
)(withWidth()(Hand))