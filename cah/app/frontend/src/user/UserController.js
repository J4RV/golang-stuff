import React, { Component } from 'react'
import axios from 'axios'
import FormControl from '@material-ui/core/FormControl'
import TextField from '@material-ui/core/TextField'
import Button from '@material-ui/core/Button'
import Typography from '@material-ui/core/Typography'
import { withStyles } from '@material-ui/core/styles'
import Card from '../gamestate/Card'
import Footer from '../Footer'
import ErrorSnackbar from '../components/ErrorSnackbar'
import {loginUrl, registerUrl, validCookieUrl} from '../restUrls'

const styles = theme => ({
  container: {
    textAlign: "center",
    marginTop: theme.spacing.unit * 2,
  },
  form: {
    maxWidth: 260,
    marginTop: theme.spacing.unit * 2,
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
    this.setState({...this.state, disabled: true})
    let payload = {
      username: this.state.username,
      password: this.state.password,
    }
    axios.post(url, payload)
      .then(this.props.onValidSubmit)
      .catch(r => {
        this.setState({...this.state,
          errormsg: r.response.data,
          disabled: false})
        return false})      
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
        <Card
          isBlack
          text="I'm _ and my password is _."
          expansion="Security questions"
        />
        <FormControl margin="normal" required fullWidth>
          <TextField
            label="Username"
            autoComplete="username"
            onChange={this.handleChangeUser}
          />
        </FormControl>
        <FormControl margin="normal" required fullWidth>
          <TextField
            label="Password"
            type="password"
            autoComplete="password"
            onChange={this.handleChangePass}
          />
        </FormControl>
        <FormControl margin="normal" fullWidth>
          <Button
            type="submit"
            variant="contained"
            color="primary"
            onClick={() => this.handleSubmit(loginUrl)}
            disabled={this.state.disabled}
          >Log in</Button>
        </FormControl>
        <FormControl margin="normal" fullWidth>
          <Button
            type="button"
            variant="outlined"
            color="primary"
            onClick={() => this.handleSubmit(registerUrl)}
            disabled={this.state.disabled}
          >Register</Button>
        </FormControl>
        <ErrorSnackbar
          msg={this.state.errormsg}
          onClose={() => this.setState({...this.state, errormsg: null})}
        />
        <Footer />
      </form>
    </div>
  }
}

class LoginController extends Component {
  state = {};
  setValid = (v) => { this.setState({ validcookie: v }) }

  componentWillMount() {
    axios.get(validCookieUrl)
      .then(r => {
        let v = (r.data === true) || (r.data === "true")
        this.setValid(v)
      })
  }

  render() {
    if (this.state.validcookie == null) {
      return <div>Loading...</div>
    }
    if (this.state.validcookie) {
      return this.props.children
    }
    return <LoginForm onValidSubmit={() => this.setValid(true)} classes={this.props.classes} />
  }

}

export default withStyles(styles)(LoginController)