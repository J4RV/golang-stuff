import React, { Component } from 'react'
import axios from 'axios'
import FormControl from '@material-ui/core/FormControl'
import TextField from '@material-ui/core/TextField'
import Button from '@material-ui/core/Button'
import Typography from '@material-ui/core/Typography'
import ErrorIcon from '@material-ui/icons/Error'
import Snackbar from '@material-ui/core/Snackbar'
import SnackbarContent from '@material-ui/core/SnackbarContent'
import { withStyles } from '@material-ui/core/styles'
import Card from '../game/Card'
import Footer from '../Footer'

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
  error: {
    color: theme.palette.getContrastText(theme.palette.error.dark),
    background: theme.palette.error.dark,
    display: 'flex',
    alignItems: 'center',
  },
  icon: {
    marginRight: theme.spacing.unit,
  },
  message: {
    display: 'flex',
    alignItems: 'center',
  },
});

class ErrorSnackbar extends Component {
  render() {
    const classes = this.props.classes
    return <Snackbar
      anchorOrigin={{
        vertical: 'bottom',
        horizontal: 'left',
      }}
      open={this.props.msg != null && this.props.msg !== ""}
    >
      <SnackbarContent
        className={classes.error}
        aria-describedby="message-id"
        message={
          <span id="message-id" className={classes.message}>
            <ErrorIcon className={classes.icon} />
            {this.props.msg}
          </span>
        }
      />
    </Snackbar>
  }
}

class SignInForm extends Component {
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

  handleSubmit = (action) => {
    this.setState({...this.state, disabled: true})
    let payload = {
      username: this.state.username,
      password: this.state.password,
    }
    axios.post("user/" + action, payload)
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
            onClick={() => this.handleSubmit("login")}
            disabled={this.state.disabled}
          >Log in</Button>
        </FormControl>
        <FormControl margin="normal" fullWidth>
          <Button
            type="button"
            variant="outlined"
            color="primary"
            onClick={() => this.handleSubmit("register")}
            disabled={this.state.disabled}
          >Register</Button>
        </FormControl>
        <ErrorSnackbar msg={this.state.errormsg} classes={classes} />
        <Footer />
      </form>
    </div>
  }
}

class LoginController extends Component {
  state = {};
  setValid = (v) => { this.setState({ validcookie: v }) }

  componentWillMount() {
    axios.get("user/validcookie")
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
    return <SignInForm onValidSubmit={() => this.setValid(true)} classes={this.props.classes} />
  }

}

export default withStyles(styles)(LoginController)