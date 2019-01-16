import React, { Component } from 'react';
import YourCardsInPlay from './YourCardsInPlay'
import Card from './Card'
import Button from '@material-ui/core/Button'
import Fab from '@material-ui/core/Fab';
import Check from '@material-ui/icons/Check';
import axios from 'axios'
import withWidth from '@material-ui/core/withWidth'

let PlayCardsButton = ({width, playCards}) => {
  if(width === "sm" || width === "xs"){
    return <Fab 
      aria-label="Play selected cards"
      color="primary"
      onClick={playCards}
      style={{position: "fixed", right: 8, bottom: 8}}>
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

class Hand extends Component {
  state = { cardIndexes: [] }

  render() {
    const gamestate = this.props.state
    return (
      <div className="cah-hand">
        <YourCardsInPlay state={gamestate} />
        <div className="cah-hand-cards">
          {gamestate.myPlayer.hand.map((c, i) =>
            <Card
              {...c}
              handIndex={i}
              className={`hovering ${this.state.cardIndexes.includes(i) ? 'selected' : ''}`}
              onClick={() => this.handleCardClick(i)}
            />
          )}
        </div>
        <div style={{marginTop: "2rem"}}>
          <PlayCardsButton playCards={this.playCards} />
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
    axios.post('rest/test/PlayCards', {
      cardIndexes: this.state.cardIndexes
    }).then(r => {
      this.setState({ cardIndexes: [] })
    }).catch(r => window.alert(r.response.data)); // We'll need prettier things
  }
}

export default withWidth()(Hand)