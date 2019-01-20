import React, { Component } from 'react'
import { MuiThemeProvider, createMuiTheme } from '@material-ui/core/styles'
import LoginController from './user/UserController'
import Game from './game/Game'
import AppBar from './AppBar'
import CssBaseline from '@material-ui/core/CssBaseline'
import red from '@material-ui/core/colors/red'
import grey from '@material-ui/core/colors/grey'
import cyan from '@material-ui/core/colors/cyan'
import './App.css'

const theme = createMuiTheme({
  palette: {
    primary: red,
    secondary: grey,
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
        <AppBar title='Cards Against Humanity' shortTitle="CAH" />
        <Game stateID='1' />
      </LoginController>
    </MuiThemeProvider>
  }
}

export default App
