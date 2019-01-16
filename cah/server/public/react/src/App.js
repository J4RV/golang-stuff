import React, { Component } from 'react'
import { MuiThemeProvider, createMuiTheme } from '@material-ui/core/styles'
import LoginController from './Login'
import './App.css'

const theme = createMuiTheme({
  palette: {
    type: 'dark'
  }
})

class App extends Component {
  render () {
    return <MuiThemeProvider theme={theme}>
      <LoginController>
        <p>Logged in!</p>
      </LoginController>
    </MuiThemeProvider>
  }
}

export default App
