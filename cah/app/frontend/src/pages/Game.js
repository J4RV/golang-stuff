import React, { Component } from 'react'
import PlayersInfo from '../game/PlayersInfo'
import Hand from '../game/Hand'
import Table from '../game/Table'
import axios from 'axios'
import './App.css'

class Game extends Component {
  render() {
    if (this.state == null) return null;
    return (
      <div className="cah-game">
        <PlayersInfo state={this.state} />
        <Table state={this.state} />
        <Hand state={this.state} />
      </div>
    );
  }
  componentWillMount() {
    this.updateState()
  }
  updateState = () => {
    axios.get(`rest/${this.props.gameID}/State`)
      .then(r => {
        console.log(r.data)
        this.setState(r.data)
        // this would be much better with websockets
        window.setTimeout(this.updateState, 1000)
      })
      .catch(e => console.log(e)
    )
  }
}

export default Game