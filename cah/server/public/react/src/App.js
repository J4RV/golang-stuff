import React, { Component } from 'react';
import Card from './Card'
import './App.css'

const BlackCard = ({card}) => {
  if (card == null) return null
  return <div>
    <Card text={card.text} isBlack={true} />
  </div>
}

const Hand = ({hand}) => (
  <div className="cah-hand">
    {hand.map(c =>
      <Card text={c.text} isBlack={false} />            
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
        <Hand {...this.state.players[0]} />        
      </div>
    );
  } 
  componentWillMount() {
    fetch("rest/test/0/State")
      .then(r => console.log(r.json()
      .then(j => console.log(j) & this.setState(j))
    ))
  }
}

export default Game;
