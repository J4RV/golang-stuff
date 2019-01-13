import React, { Component } from 'react';
import Card from './Card'
import Button from '@material-ui/core/Button'
import axios from 'axios'
import LocalPlayerIndex from './LocalPlayerIndex'

class Hand extends Component {
  state = {cardIndexes: []}

  render() {
    const gamestate = this.props.state
    return (
    <div className="cah-hand">
      <div className="cah-hand-cards">
        {gamestate.players[LocalPlayerIndex()].hand.map((c, i) =>
          <Card
            {...c}
            playable={this.isCzar() === false}
            handIndex={i}
            className={`hovering ${this.state.cardIndexes.includes(i) ? 'selected' : ''}`}
            onClick={() => this.handleCardClick(i)}
          />            
        )}
      </div>
      <Button variant="contained" color="primary" onClick={this.playCards}>
        Play cards
      </Button>
    </div>
    )
  }

  isCzar = () => {
    return this.props.state.currentCzarIndex == LocalPlayerIndex()
  }

  handleCardClick = (i) => {
    if(this.isCzar()){
      return
    }
    let newList = this.state.cardIndexes.slice()
    if(newList.includes(i)){
      newList.splice(newList.indexOf(i), 1)
    } else {
      newList.push(i)
    }
    this.setState({cardIndexes: newList})
  }

  playCards = () => {
    axios.post('rest/test/'+LocalPlayerIndex()+'/PlayCards', {
      cardIndexes: this.state.cardIndexes
    }).then(r => {
      this.setState({cardIndexes: []})
    }).catch(r => window.alert(r.response.data)); // We'll need prettier things
  }
}

export default Hand