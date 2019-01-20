import { Button, Typography } from '@material-ui/core';
import React, {Component} from 'react'

import {Link} from 'react-router-dom'
import Paper from '@material-ui/core/Paper'
import Table from '@material-ui/core/Table'
import TableBody from '@material-ui/core/TableBody'
import TableCell from '@material-ui/core/TableCell'
import TableHead from '@material-ui/core/TableHead'
import TableRow from '@material-ui/core/TableRow'
import axios from 'axios'
import {openGamesUrl} from '../restUrls'
import { withStyles } from '@material-ui/core/styles'
import withWidth from '@material-ui/core/withWidth';

const styles = theme => ({
  root: {
    maxWidth: 960 - theme.spacing.unit * 4,
    marginTop: theme.spacing.unit * 3,
    marginLeft: "auto",
    marginRight: "auto",
    overflowX: 'auto',
    padding: theme.spacing.unit * 2,
    [theme.breakpoints.down('sm')]: {
      padding: theme.spacing.unit,
    },
  },
  tableContainer: {
    minWidth: 800,
  },
  title: {
    textAlign: "center",
    marginBottom: theme.spacing.unit * 2,
  },
  createBtn: {
    float: "right",
    marginTop: theme.spacing.unit * 2,
  },
});

class GamesTable extends Component {
  state = {games: []};

  componentWillMount() {
    axios.get(openGamesUrl)
      .then(r => {
        this.setState({games: r.data})})
      .catch(e => console.log(e.response.data))
  }

  render() {
    const {classes} = this.props
    return <div className={classes.root}>
      <Typography variant="h5" className={classes.title}>
        Open games
      </Typography>
      <Paper className={classes.tableContainer}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell align="right">Name</TableCell>
              <TableCell align="right">Owner</TableCell>
              <TableCell align="right" >Has password</TableCell>
              <TableCell align="right">Expansions</TableCell>
              <TableCell />
            </TableRow>
          </TableHead>
          <TableBody>
            {this.state.games.map(game => (
              <TableRow key={game.id}>
                <TableCell align="right">{game.name}</TableCell>
                <TableCell align="right">{game.owner}</TableCell>
                <TableCell align="right">{game.hasPassword ? "Yes" : "No"}</TableCell>
                <TableCell align="right">{game.expansions.join(", ")}</TableCell>
                <TableCell align="right">
                  <Button color="primary" variant="contained">
                    Join
                  </Button>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </Paper>
      <Link to="/game/create">
        <Button
          type="button"
          //variant="contained"
          className={classes.createBtn}
        >
          Create new game
        </Button>
      </Link>
    </div>
  }
}

const GameList = (props) => (
  <React.Fragment>
    <GamesTable {...props} />
  </React.Fragment>
)

export default withWidth()(withStyles(styles)(GameList))