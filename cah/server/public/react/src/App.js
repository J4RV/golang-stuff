import React, { Component } from 'react'
import { MuiThemeProvider, createMuiTheme } from '@material-ui/core/styles'
import LoginController from './Login'
import Game from './Game'
import AppBar from './AppBar'
import red from '@material-ui/core/colors/red';
import blueGrey from '@material-ui/core/colors/blueGrey';
import './App.css'

const theme = createMuiTheme({
  palette: {
    primary: red,
    secondary: blueGrey,
    type: 'dark',
  }
})

class App extends Component {
  render() {
    return <MuiThemeProvider theme={theme}>
      <AppBar title="Cards Against Humanity" />
      <LoginController>
        <Game gameid="test" />
      </LoginController>
    </MuiThemeProvider>
  }
}

export default App
