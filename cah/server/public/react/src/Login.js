import React, { Component } from 'react'
import axios from 'axios'
import FormControl from '@material-ui/core/FormControl'
import TextField from '@material-ui/core/TextField'
import Button from '@material-ui/core/Button'
import Typography from '@material-ui/core/Typography'
import Card from './Card'

/*class ErrorSnackbar extends Component {
  render() {
    return <Snackbar
      anchorOrigin={{
        vertical: 'bottom',
        horizontal: 'left',
      }}
      open={this.props.msg != null && this.props.msg !== ""}
      autoHideDuration={6000}
      ContentProps={{
        'aria-describedby': 'message-id',
      }}
      message={<span id="message-id">{this.props.msg}</span>}
    />
  }
}*/

class LoginForm extends Component {
  state = { username: "", password: "" };

  setErrorMsg = (msg) => {
    let newState = Object.assign({}, this.state)
    newState.errormsg = msg
    this.setState(newState);
  }

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

  handleSubmit = (event) => {
    event.preventDefault();
    let payload = {
      username: this.state.username,
      password: this.state.password
    }
    axios.post("user/login", payload)
      .then(this.props.onValidLogin)
      .catch(r => this.setErrorMsg(r.response.data))
  }

  render() {
    return <div className="cah-login-container">
      <Typography variant="h2" gutterBottom>
        Cards Against Humanity
      </Typography>
      <Typography variant="h4" gutterBottom>
        A party game for horrible people.
      </Typography>
      <form onSubmit={this.handleSubmit} className="cah-login-form" >
        <Card
          text="I'm _ and my password is _."
          isBlack
          expansion="Security questions"
          className="in-table"
        />
        <FormControl margin="normal" required fullWidth>
          <TextField
            label="Username"
            margin="normal"
            autoComplete="username"
            onChange={this.handleChangeUser}
          />
        </FormControl>
        <FormControl margin="normal" required fullWidth>
          <TextField
            label="Password"
            margin="normal"
            type="password"
            autoComplete="password"
            onChange={this.handleChangePass}
          />
        </FormControl>
        <FormControl margin="normal" required fullWidth>
          <Button
            type="submit"
            variant="contained"
            color="primary"
          >Sign in</Button>
        </FormControl>
        {this.state.errormsg ?
          <div className="cah-form-error">{this.state.errormsg}</div>
          : null}
      </form>
      <a href="https://github.com/J4RV"><h6>A J4RV production</h6></a>
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
    return <LoginForm onValidLogin={() => this.setValid(true)} />
  }

}

export default LoginController