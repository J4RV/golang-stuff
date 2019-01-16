import React, { Component } from 'react'
import { MuiThemeProvider, createMuiTheme } from '@material-ui/core/styles'
import LoginController from './Login'
import Game from './Game'
import './App.css'

const theme = createMuiTheme({
  palette: {
    type: 'dark'
  }
})

class App extends Component {
  render() {
    return <MuiThemeProvider theme={theme}>
      <LoginController>
        <Game gameid="test" />
      </LoginController>
    </MuiThemeProvider>
  }
}

export default App
