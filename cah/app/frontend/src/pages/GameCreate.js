import React, { Component } from 'react'
import { createGameUrl } from '../restUrls'

import Button from '@material-ui/core/Button'
import ErrorSnackbar from '../components/ErrorSnackbar'
import Footer from '../Footer'
import FormControl from '@material-ui/core/FormControl'
import { Redirect } from 'react-router'
import TextField from '@material-ui/core/TextField'
import Typography from '@material-ui/core/Typography'
import axios from 'axios'
import { withStyles } from '@material-ui/core/styles'

const styles = theme => ({
  container: {
    textAlign: "center",
    marginTop: theme.spacing.unit * 2,
  },
  form: {
    maxWidth: 260,
    marginTop: theme.spacing.unit * 2,
    marginBottom: theme.spacing.unit * 2,
    padding: theme.spacing.unit * 2,
    display: "inline-block",
  },
});

class LoginForm extends Component {
  state = { username: "", password: "", disabled: false }

  handleChangeUser = (event) => {
    let newState = Object.assign({}, this.state)
    newState.username = event.target.value.trim()
    this.setState(newState)
  }

  handleChangePass = (event) => {
    let newState = Object.assign({}, this.state)
    newState.password = event.target.value.trim()
    this.setState(newState)
  }

  handleSubmit = (url) => {
    this.setState({ ...this.state, disabled: true })
    let payload = {
      username: this.state.username,
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

  render() {
    const classes = this.props.classes
    return <div className={classes.container}>
      <Typography variant="h2" gutterBottom>
        Cards Against Humanity
      </Typography>
      <Typography variant="h4" gutterBottom>
        A party game for horrible people.
      </Typography>
      <form className={classes.form} onSubmit={() => this.handleSubmit("login")} >
        <TextField required fullWidth margin="normal"
          label="Room name"
          autoComplete="room_name"
          onChange={this.handleChangeUser}
        />
        <TextField fullWidth margin="normal"
          label="Room password"
          type="password"
          autoComplete="room_password"
          onChange={this.handleChangePass}
        />
        <FormControl margin="normal" fullWidth>
          <Button margin="normal"
            type="submit"
            variant="contained"
            color="primary"
            onClick={() => this.handleSubmit(loginUrl)}
            disabled={this.state.disabled}
          >Create Game</Button>
        </FormControl>
        <ErrorSnackbar
          msg={this.state.errormsg}
          onClose={() => this.setState({ ...this.state, errormsg: null })}
        />
        <Footer />
      </form>
    </div>
  }
}

export default withStyles(styles)(LoginForm)