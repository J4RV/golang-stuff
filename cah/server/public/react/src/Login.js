import React, { Component } from 'react'
import axios from 'axios'

class LoginForm extends Component {
  state = {username: "", password: ""};

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
    let payload = {
      username: this.state.username,
      password: this.state.password
    }
    console.log(payload)
    axios.post("/rest/login", payload)
    .then(this.props.onValidLogin)
    .catch(r => this.setErrorMsg(r.response.data))
    event.preventDefault();
  }

  render() {
    return <form onSubmit={this.handleSubmit}>
      <label>
        Username:
        <input type="text" value={this.state.username} onChange={this.handleChangeUser} />
      </label>
      <label>
        Password:
        <input type="text" value={this.state.password} onChange={this.handleChangePass} />
      </label>
      <input type="submit" value="Submit" />
      <p>{this.state.errormsg}</p>
    </form>
  }
}

class LoginController extends Component {
  state = {};
  setValid = (v) => {this.setState({validcookie: v})}

  componentWillMount() {
    console.log(this.state.validcookie)
    fetch("rest/validcookie")
      .then(response => response.text())
      .then(value => this.setValid(value))
    console.log(this.state.validcookie)
  }

  render() {
    console.log(this.state.validcookie)
    if(this.state.validcookie == null){
      return <div>Loading...</div>
    }
    if(this.state.validcookie == true){
      return this.props.children
    }
    return <LoginForm onValidLogin={() => this.setValid(true)} />
  }

}

export default LoginController