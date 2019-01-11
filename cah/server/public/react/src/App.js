import React, { Component } from 'react';
import Card from './Card'
import LocalPlayerIndex from './LocalPlayerIndex'
import Button from '@material-ui/core/Button';
import axios from 'axios'
import './App.css'

const BlackCard = ({card}) => {
  if (card == null) return null
  return <div>
    <Card text={card.text} isBlack={true} className='in-table' />
    <p>{`Play ${card.blanksAmount} card${card.blanksAmount > 1 ? 's' : ''}`}</p>
  </div>
}

const CardsInPlay = ({state, owner}) => (
  <span>
    {state.players[owner].whiteCardsInPlay.map((c, i) =>
      <Card text={c.text} isBlack={false} playable={false} handIndex={i} className='in-table' />            
    )}
  </span>
)

const PlayerInfo = ({player}) => (
  <div className="cah-playerinfo">
    <p>{player.name}</p>
    <p>{player.points.length} points</p>
    <p>{player.whiteCardsInPlay.length} cards in play</p>
  </div>
)

const PlayersInfo = ({state}) => (
  <div style={{display: "flex"}}>
    {state.players.map(p => 
      <PlayerInfo player={p} />
    )}
  </div>
)



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
          isBlack={false}
          playable={gamestate.currentCzarIndex !== LocalPlayerIndex()}
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

  handleCardClick = (i) => {
    let newList = this.state.cardIndexes.slice()
    if(newList.includes(i)){
      newList.splice(newList.indexOf(i), 1)
    } else {
      newList.push(i)
    }
    this.setState({cardIndexes: newList})
  }

  playCards = () => {
    console.log(this.state)
    axios.post('rest/test/'+LocalPlayerIndex()+'/PlayCards', {
      cardIndexes: this.state.cardIndexes
    }).then(r => {
      if (r.status > 299){
        console.err(r)
      }
      this.setState({cardIndexes: []})
    }).catch(r => console.err(r));      
  }
}

class Game extends Component {
  render() {
    if(this.state == null) return null;
    return (
      <div className="Game">
        <PlayersInfo state={this.state} />
        <BlackCard card={this.state.blackCardInPlay} />
        <CardsInPlay state={this.state} owner={LocalPlayerIndex()} />
        <Hand state={this.state} />        
      </div>
    );
  } 
  componentWillMount() {
    this.updateState()
    // this would be much better with websockets
    window.setInterval(this.updateState, 500)
  }
  updateState = () => {
    fetch("rest/test/"+LocalPlayerIndex()+"/State")
      .then(r => r.json()
      .then(j => /*console.log(j) &*/ this.setState(j))
    ) 
  }
}

export default Game;
