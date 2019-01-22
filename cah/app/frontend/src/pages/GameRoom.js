import React, { Component } from 'react'

import Button from '@material-ui/core/Button'
import Game from './Game'
import { Link } from 'react-router-dom'
import Typography from '@material-ui/core/Typography'
import axios from 'axios'
import { connect } from 'react-redux'
import { roomStateUrl } from '../restUrls'
import { withStyles } from '@material-ui/core/styles'

const styles = theme => ({
  container: {
    textAlign: "center",
    marginTop: theme.spacing.unit * 2,
  },
});

class GameRoom extends Component {
  render() {
    if (this.state == null) {
      return <Typography variant="h5" gutterBottom>Now loading...</Typography>
    }
    const { room } = this.state
    if (room.phase !== "Not started") {
      return <Game stateID={room.stateID} />
    }
    const imOwner = this.props.username === room.owner
    return (
      <div className={this.props.classes.container}>
        <Typography variant="h4" gutterBottom>
          {room.name}
        </Typography>
        <Typography variant="h5" gutterBottom>
          Creator: {room.owner}.
        </Typography>
        <Typography gutterBottom>
          Players: {room.players.join(", ")}.
        </Typography>
        <Link to="../list/open">
          <Button>Back to games list</Button>
        </Link>
        {room.players.length > 2
          ? <EnoughPlayers imOwner={imOwner} />
          : <Typography gutterBottom>Waiting for more players to join</Typography>}
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

const EnoughPlayers = ({ imOwner }) => {
  if (imOwner) {
    return <Button variant="contained" color="primary">Start game</Button>
  } else {
    return <Typography gutterBottom>Waiting for the game creator to start the game</Typography>
  }
}

export default connect(
  state => ({ username: state.username })
)(withStyles(styles)(GameRoom))