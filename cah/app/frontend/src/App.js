import './App.css'

import { MuiThemeProvider, createMuiTheme } from '@material-ui/core/styles'
import React, { Component } from 'react'
import { Redirect, Route, BrowserRouter as Router } from 'react-router-dom'

import AppBar from './AppBar'
import CssBaseline from '@material-ui/core/CssBaseline'
import GameCreate from './pages/GameCreate'
import GameList from './pages/GameList'
import GameRoom from './pages/GameRoom'
import LoggedIn from './pages/LoggedIn'
import { Provider } from 'react-redux'
import { createStore } from 'redux'
import cyan from '@material-ui/core/colors/cyan'
import grey from '@material-ui/core/colors/grey'
import red from '@material-ui/core/colors/red'
import reducer from './reducer'

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
    return <Router>
      <MuiThemeProvider theme={theme}>
        <Provider store={createStore(reducer)}>
          <CssBaseline />
          <LoggedIn>
            <AppBar />
            <Route exact path="/" render={() => <Redirect to="/game/list" />} />
            <Route path="/game/list/create" component={GameCreate} />
            <Route path="/game/list" component={GameList} />
            <Route path="/game/room/:gameID" component={GameRoom} />
          </LoggedIn>
        </Provider>
      </MuiThemeProvider>
    </Router>
  }
}

export default App
