import React, { Component } from 'react';
import Card from './Card'
import Button from '@material-ui/core/Button'
import Fab from '@material-ui/core/Fab';
import Check from '@material-ui/icons/Check';
import axios from 'axios'

class Hand extends Component {
  state = { cardIndexes: [] }

  render() {
    const gamestate = this.props.state
    return (
      <div className="cah-hand">
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
        <Fab aria-label="Play selected cards" color="primary" onClick={this.playCards}>
          <Check />
        </Fab>
      </div>
    )
  }

  isCzar = () => {
    return this.props.state.currentCzarID === this.props.state.myPlayer.ID
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

export default Hand