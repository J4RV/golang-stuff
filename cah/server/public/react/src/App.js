import React, { Component } from 'react'
import { MuiThemeProvider, createMuiTheme } from '@material-ui/core/styles'
import LoginController from './Login'
import Game from './Game'
import AppBar from './AppBar'
import CssBaseline from '@material-ui/core/CssBaseline'
import red from '@material-ui/core/colors/red'
import brown from '@material-ui/core/colors/brown'
import cyan from '@material-ui/core/colors/cyan'
import './App.css'

const theme = createMuiTheme({
  palette: {
    primary: red,
    secondary: brown,
    whitecard: { text: "#161616", background: "#FAFAFA" },
    blackcard: { text: "#FAFAFA", background: "#161616" },
    expansion: "#888888",
    type: 'dark',
  },
  lights: {
    glow: `0 0 4px 2px ${cyan[100]}, 0 0 24px 2px ${cyan[500]}`,
  },
})

class App extends Component {
  render() {
    return <MuiThemeProvider theme={theme}>
      <CssBaseline />
      <LoginController>
        <AppBar title='Cards Against Humanity' />
        <Game gameid='test' />
      </LoginController>
    </MuiThemeProvider>
  }
}

export default App
