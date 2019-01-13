import React, { Component } from 'react';
import YourCardsInPlay from './YourCardsInPlay'
import PlayersInfo from './PlayersInfo'
import Hand from './Hand'
import Table from './Table'
import LocalPlayerIndex from './LocalPlayerIndex'
import './App.css'

class Game extends Component {
  render() {
    if(this.state == null) return null;
    return (
      <div className="Game">
        <PlayersInfo state={this.state} />
        <Table state={this.state} />
        <YourCardsInPlay state={this.state} owner={LocalPlayerIndex()} />
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

export default Game