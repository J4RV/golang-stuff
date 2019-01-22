import React, { Component } from 'react'

import Button from '@material-ui/core/Button'
import Dialog from '@material-ui/core/Dialog'
import DialogActions from '@material-ui/core/DialogActions'
import DialogTitle from '@material-ui/core/DialogTitle'
import ErrorSnackbar from '../components/ErrorSnackbar'
import { Redirect } from 'react-router-dom'
import TextField from '@material-ui/core/TextField'
import axios from 'axios'
import { createGameUrl } from '../restUrls'
import { withStyles } from '@material-ui/core/styles'

const styles = theme => ({
  form: {
    maxWidth: 360,
    padding: theme.spacing.unit * 2,
    paddingTop: 0,
    display: "inline-block",
  },
})

class CreateGameDialog extends Component {
  state = { name: "", password: "", waitingResponse: false, closed: false }

  render() {
    const classes = this.props.classes
    if (this.state.closed) {
      return <Redirect to="/" />
    }
    return (
      <Dialog open={true} onClose={this.close} {...this.props} aria-labelledby="create-game-dialog-title">
        <DialogTitle id="create-game-dialog-title">Create new Game</DialogTitle>
        <form className={classes.form} onSubmit={this.handleSubmit} >
          <TextField required fullWidth margin="normal"
            label="Room name"
            autoComplete="roomName" s
            onChange={this.handleChangeName}
          />
          <TextField fullWidth margin="normal"
            label="Room password"
            autoComplete="roomPassword"
            onChange={this.handleChangePass}
          />
          <DialogActions>
            <Button margin="normal"
              onClick={this.close}
              disabled={this.state.waitingResponse}
            >Cancel</Button>
            <Button margin="normal" autoFocus
              type="submit"
              variant="contained"
              color="primary"
              disabled={this.state.waitingResponse}
            >Create Game</Button>
          </DialogActions>
          <ErrorSnackbar
            msg={this.state.errormsg}
            onClose={() => this.setState({ ...this.state, errormsg: null })}
          />
        </form>
      </Dialog>
    )
  }

  handleSubmit = () => {
    this.setState({ ...this.state, waitingResponse: true })
    let payload = {
      name: this.state.name,
      password: this.state.password,
    }
    axios.post(createGameUrl, payload)
      .then(this.close)
      .catch(r => {
        this.setState({
          ...this.state,
          errormsg: r.response.data,
          waitingResponse: false
        })
        return false
      })
  }

  close = () => {
    this.setState({ ...this.state, closed: true })
  }

  handleChangeName = (event) => {
    let newState = Object.assign({}, this.state)
    newState.name = event.target.value.trim()
    this.setState(newState)
  }

  handleChangePass = (event) => {
    let newState = Object.assign({}, this.state)
    newState.password = event.target.value.trim()
    this.setState(newState)
  }
}

export default withStyles(styles)(CreateGameDialog)