import React, { Component } from 'react'
import { MuiThemeProvider, createMuiTheme, getContrastText } from '@material-ui/core/styles'
import LoginController from './Login'
import Game from './Game'
import AppBar from './AppBar'
import CssBaseline from '@material-ui/core/CssBaseline'
import red from '@material-ui/core/colors/red';
import brown from '@material-ui/core/colors/brown';
import './App.css'

const theme = createMuiTheme({
  palette: {
    primary: red,
    secondary: brown,
    type: 'dark',
  }
})

class App extends Component {
  render() {
    return <MuiThemeProvider theme={theme}>
      <CssBaseline />
      <LoginController>
        <AppBar title="Cards Against Humanity" />
        <Game gameid="test" />
      </LoginController>
    </MuiThemeProvider>
  }
}

export default App
