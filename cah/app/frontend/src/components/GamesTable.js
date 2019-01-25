import { Link, Redirect } from 'react-router-dom'
import React, { Component } from 'react'

import { Button } from '@material-ui/core';
import Table from '@material-ui/core/Table'
import TableBody from '@material-ui/core/TableBody'
import TableCell from '@material-ui/core/TableCell'
import TableHead from '@material-ui/core/TableHead'
import TableRow from '@material-ui/core/TableRow'
import axios from 'axios'
import { connect } from 'react-redux'
import { joinGameUrl } from '../restUrls'
import pushError from '../actions/pushError'
import { withStyles } from '@material-ui/core/styles'

const styles = theme => ({
})

const PrimaryButton = ({ username, game, joinGame }) => {
  if (game.players.includes(username)) {
    return <Link to={`/game/room/${game.id}`}>
      <Button
        color="primary"
        variant="contained">
        Enter
      </Button>
    </Link>
  } else {
    return <Button
      color="primary"
      variant="contained"
      onClick={() => joinGame(game.id)}>
      Join
    </Button>
  }
}

class GamesTable extends Component {
  state = { games: [], joinedGame: false }

  render() {
    if (this.state.joinedGame) {
      return <Redirect to={`/game/room/${this.state.joinedGame}`} />
    }
    const { username } = this.props
    const { games } = this.state
    return (
      <Table>
        <TableHead>
          <TableRow>
            <TableCell align="right">Name</TableCell>
            <TableCell align="right">Owner</TableCell>
            {/*<TableCell align="right">Has password</TableCell>*/}
            <TableCell align="right">Current players</TableCell>
            <TableCell />
          </TableRow>
        </TableHead>
        <TableBody>
          {games.map(game => (
            <TableRow key={game.id}>
              <TableCell align="right">{game.name}</TableCell>
              <TableCell align="right">{game.owner}</TableCell>
              {/*<TableCell align="right">{game.hasPassword ? "Yes" : "No"}</TableCell>*/}
              <TableCell align="right">{game.players.join(", ")}</TableCell>
              <TableCell align="right">
                <PrimaryButton game={game} joinGame={this.joinGame} username={username} />
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    )
  }

  componentWillMount() {
    this.refreshGames()
  }

  refreshGames = () => {
    axios.get(this.props.fetchGamesUrl)
      .then(r => {
        this.setState({ ...this.state, games: r.data })
      })
      .catch(e => this.props.pushError(e))
  }

  joinGame = (gameID) => {
    axios.post(joinGameUrl, { id: gameID })
      .then(this.setState({ ...this.state, joinedGame: gameID }))
      .catch(e => this.props.pushError(e))
  }
}

export default connect(
  state => ({ username: state.username }),
  { pushError }
)(/*withStyles(styles)*/(GamesTable))