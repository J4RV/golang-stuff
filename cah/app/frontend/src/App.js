import './App.css'

import { MuiThemeProvider, createMuiTheme } from '@material-ui/core/styles'
import React, { Component } from 'react'
import {Route, BrowserRouter as Router, Switch} from 'react-router-dom'

import AppBar from './AppBar'
import CssBaseline from '@material-ui/core/CssBaseline'
import Game from './pages/Game'
import GameCreate from './pages/GameCreate'
import GameList from './pages/GameList'
import Login from './pages/Login'
import cyan from '@material-ui/core/colors/cyan'
import grey from '@material-ui/core/colors/grey'
import red from '@material-ui/core/colors/red'

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
        <CssBaseline />
        <Switch>
          <Route path="/" component={Login} exact />
          <Route path="/" component={AppBar} />
        </Switch>
        <Route exact path="/game/new" component={GameCreate} />
        <Route exact path="/game/open" component={GameList} />
        <Route path="/ingame/:stateID" component={Game} />
      </MuiThemeProvider>
    </Router>
  }
}

export default App
