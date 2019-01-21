import React, { Component } from 'react'
import { createGameUrl } from '../restUrls'

import Button from '@material-ui/core/Button'
import ErrorSnackbar from '../components/ErrorSnackbar'
import DialogActions from '@material-ui/core/DialogActions'
import DialogTitle from '@material-ui/core/DialogTitle'
import Dialog from '@material-ui/core/Dialog'
import { Redirect } from 'react-router'
import TextField from '@material-ui/core/TextField'
import axios from 'axios'
import { withStyles } from '@material-ui/core/styles'

const styles = theme => ({
  form: {
    maxWidth: 360,
    padding: theme.spacing.unit * 2,
    paddingTop: 0,
    display: "inline-block",
  },
})

class CreateGameForm extends Component {
  state = { name: "", password: "", disabled: false, createdSuccessfully: false }

  render() {
    const classes = this.props.classes
    if (this.state.createdSuccessfully) {
      return <Redirect to="/game/open" />
    }
    return (
      <form className={classes.form} onSubmit={() => this.handleSubmit(createGameUrl)} >
        <TextField required fullWidth margin="normal"
          label="Room name"
          autoComplete="roomName"
          onChange={this.handleChangeName}
        />
        <TextField fullWidth margin="normal"
          label="Room password"
          autoComplete="roomPassword"
          onChange={this.handleChangePass}
        />
        <DialogActions>
          <Button margin="normal"
            onClick={this.props.onCancel}
            disabled={this.state.disabled}
          >Cancel</Button>
          <Button margin="normal" autoFocus
            type="submit"
            variant="contained"
            color="primary"
            onClick={() => this.handleSubmit(createGameUrl)}
            disabled={this.state.disabled}
          >Create Game</Button>
        </DialogActions>
        <ErrorSnackbar
          msg={this.state.errormsg}
          onClose={() => this.setState({ ...this.state, errormsg: null })}
        />
      </form>
    )
  }

  handleSubmit = (url) => {
    this.setState({ ...this.state, disabled: true })
    let payload = {
      name: this.state.name,
      password: this.state.password,
    }
    axios.post(url, payload)
      .then(this.props.onValidSubmit)
      .catch(r => {
        this.setState({
          ...this.state,
          errormsg: r.response.data,
          disabled: false
        })
        return false
      })
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

const CreateGameDialog = (props) => (
  <Dialog aria-labelledby="create-game-dialog-title" {...props}>
    <DialogTitle id="create-game-dialog-title">Create new Game</DialogTitle>
    <CreateGameForm classes={props.classes} onValidSubmit={props.onCreation} />
  </Dialog>
)

export default withStyles(styles)(CreateGameDialog)