import React, { Component } from 'react'
import { availableExpansionsUrl, startGameUrl } from '../restUrls'

import Button from '@material-ui/core/Button'
import DialogActions from '@material-ui/core/DialogActions'
import ErrorSnackbar from '../components/ErrorSnackbar'
import FormControl from '@material-ui/core/FormControl';
import Input from '@material-ui/core/Input';
import MenuItem from '@material-ui/core/MenuItem';
import { Redirect } from 'react-router-dom'
import Select from '@material-ui/core/Select';
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

class CreateGameDialog extends Component {
  state = { expansions: ["Loading..."] }

  render() {
    const { classes, gameID } = this.props
    const {expansions} = this.state
    if (this.state.closed) {
      return <Redirect to="/" />
    }
    return (
      <form className={classes.form} onSubmit={this.handleSubmit} >
        <input type="hidden" id="gameID" value={gameID} />>
        <FormControl>
          <Select
            multiple
            displayEmpty
            value={this.state.name}
            onChange={this.handleChange}
            input={<Input id="select-multiple-placeholder" />}
            renderValue={selected => {
              if (selected.length === 0) {
                return <em>Placeholder</em>;
              }

              return selected.join(', ');
            }}
          >
            {expansions.map(name => (
              <MenuItem key={name} value={name}>
                {name}
              </MenuItem>
            ))}
          </Select>
        </FormControl>
        <ErrorSnackbar
          msg={this.state.errormsg}
          onClose={() => this.setState({ ...this.state, errormsg: null })}
        />
      </form>
    )
  }

  componentWillMount() {
    axios.get(availableExpansionsUrl)
      .then(r => this.setState({ ...this.state, expansions: r.data }))
  }

  handleSubmit = (event) => {
    event.preventDefault() // necessary for firefox, or submitting the form will make the page reload!
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
      })
    return false
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