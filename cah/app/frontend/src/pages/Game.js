import React, { Component } from 'react'

import Hand from '../gamestate/Hand'
import PlayersInfo from '../gamestate/PlayersInfo'
import Table from '../gamestate/Table'
import axios from 'axios'
import { gameStateUrl } from '../restUrls'

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
    const stateID = this.props.stateID
    axios.get(gameStateUrl(stateID))
      .then(r => {
        this.setState(r.data)
        // this would be much better with websockets
        window.setTimeout(this.updateState, 1000)
      })
      .catch(e => window.alert(e)
      )
  }
}

export default Game