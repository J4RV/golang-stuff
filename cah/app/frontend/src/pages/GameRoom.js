import React, { Component } from 'react'

import Typography from '@material-ui/core/Typography'
import axios from 'axios'
import { roomStateUrl } from '../restUrls'

class GameRoom extends Component {
  render() {
    if (this.state == null) return null;
    const { room } = this.state
    //{"id":1,"name":"A long and descriptive game name","owner":"Green","hasPassword":false,"players":["Green"],"phase":"Not started"}
    return (
      <div className="cah-game">
        <Typography variant="h4" gutterBottom>
          {room.name}
        </Typography>
        <Typography variant="h5" gutterBottom>
          {room.phase}
        </Typography>
        <Typography gutterBottom>
          Players: {room.players.join(", ")}
        </Typography>
        <Typography gutterBottom>
          {room.players.length > 2
            ? "Waiting for owner to start the game"
            : "Waiting for more players to join"}
        </Typography>
      </div>
    );
  }
  componentWillMount() {
    this.updateState()
  }
  updateState = () => {
    const gameID = this.props.match.params.gameID
    axios.get(roomStateUrl(gameID))
      .then(r => {
        this.setState({ room: r.data })
        // this would be much better with websockets
        window.setTimeout(this.updateState, 5000)
      })
      .catch(e => window.alert(e)
      )
  }
}

export default GameRoom