import React, { Component } from 'react';
import LoginController from './Login'
import './App.css'


class App extends Component {
  render() {
    return <LoginController>
      <p>Logged in!</p>
    </LoginController>
  }
}

export default App;
