import React, { Component } from 'react';
import Card from './Card'
import LocalPlayerIndex from './LocalPlayerIndex'
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

const Hand = ({state}) => (
  <div className="cah-hand">
    {state.players[LocalPlayerIndex()].hand.map((c, i) =>
      <Card text={c.text} isBlack={false} playable={true} handIndex={i} className='hovering' />            
    )}
  </div>
)

class Game extends Component {
  state = undefined;
  render() {
    if(this.state == null) return null;
    return (
      <div className="Game">
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
      .then(j => console.log(j) & this.setState(j))
    ) 
  }
}

export default Game;
