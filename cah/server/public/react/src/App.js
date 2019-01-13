import React, { Component } from 'react';
import Game from './Game'
import Cookies from "js-cookie"
import './App.css'

class Login extends React.Component {
  state = {id: ''};

  handleChange = (event) => {
    if(isNaN(event.target.value)){
      return
    }
    this.setState({id: event.target.value});
  }

  handleSubmit = (event) => {
    this.props.onLogin(this.state)
    event.preventDefault();
  }

  render() {
    return (
      <form onSubmit={this.handleSubmit}>
        <label>
          Name:
          <input type="text" value={this.state.id} onChange={this.handleChange} />
        </label>
        <input type="submit" value="Submit" />
      </form>
    );
  }

}

class App extends Component {
  state = {id: undefined};

  render() {
    if(this.state.id == null){
      return (
        <Login onLogin={this.onLogin} />
      );
    } else {
      return (
        <Game pid={this.state.id} />
      );
    }
  }

  onLogin = (s) => {
    Cookies.set("cah-currplayer-id", s.id) // fast n dirty
    this.setState({id: s.id})
  }
}

export default App;
