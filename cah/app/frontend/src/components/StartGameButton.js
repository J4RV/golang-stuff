import React, { Component } from 'react'

import Button from '@material-ui/core/Button'
import axios from 'axios'
import { startGameUrl } from '../restUrls'

class StartGameButton extends Component {
  state = { disabled: false }

  render() {
    const { className } = this.props
    return <Button
      variant="contained"
      color="primary"
      className={className}
      onClick={this.handleClick}
      disabled={this.state.disabled}>
      Start game
    </Button>
  }

  handleClick = () => {
    this.setState({ disabled: true })
    const { gameID } = this.props
    axios.post(startGameUrl, { id: gameID })
      .catch(_ => this.setState({ disabled: false }))
  }
}

export default StartGameButton