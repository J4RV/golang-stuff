import React, { Component } from 'react';
import Card from './Card'
import LocalPlayerIndex from './LocalPlayerIndex'
import Button from '@material-ui/core/Button';
import axios from 'axios'
import './App.css'

const PlayerWhiteCardsPlayed = ({player}) => {
  const {whiteCardsInPlay} = player
  return (<React.Fragment>
    {whiteCardsInPlay.map(whiteCard => <Card {...whiteCard} className='in-table' />)}
  </React.Fragment>)
}

const allSinnersPlayed = (state) => {
  for(let i = 0; i < state.players.length; i++){
    if(i == state.currentCzarIndex){
      continue
    }
    if(state.players[i].whiteCardsInPlay.length != state.blackCardInPlay.blanksAmount){
      return false
    }
  }
  return true
}

const WhiteCardsPlayed = ({state}) => {
  if (allSinnersPlayed(state)){
    return <React.Fragment>{state.players.map(p => <PlayerWhiteCardsPlayed player={p} />)}</React.Fragment>
  } else {
    return <h2>Waiting for players...</h2>
  }
}

const Table = ({state}) => {
  const card = state.blackCardInPlay
  if (card == null) return null
  return (
  <div style={{display: "flex"}}>
    <Card text={card.text} isBlack={true} className='in-table' />    
    <WhiteCardsPlayed state={state} />
  </div>)
}

const CardsInPlay = ({state, owner}) => {
  const cards = state.players[owner].whiteCardsInPlay.map(c =>
    <Card text={c.text} playable={false} className='in-table' />            
  )
  if(cards == null || cards.length === 0){
    return <h2>Play {state.blackCardInPlay.blanksAmount} cards</h2>
  }
  return <span>
    <h2>Your cards played this round:</h2>
    {cards}
  </span>
}

const PlayerInfo = ({player}) => (
  <div className="cah-playerinfo">
    <p>{player.name}</p>
    <p>{player.points.length} points</p>
    <p>{player.whiteCardsInPlay.length} cards in play</p>
  </div>
)

const PlayersInfo = ({state}) => (
  <div className="cah-playersinfo">
    {state.players.map(p => 
      <PlayerInfo player={p} />
    )}
  </div>
)

class Hand extends Component {
  state = {cardIndexes: []}

  render() {
    const gamestate = this.props.state
    return (
    <div className="cah-hand">
      <div className="cah-hand-cards">
        {gamestate.players[LocalPlayerIndex()].hand.map((c, i) =>
          <Card
            {...c}
            playable={this.isCzar() === false}
            handIndex={i}
            className={`hovering ${this.state.cardIndexes.includes(i) ? 'selected' : ''}`}
            onClick={() => this.handleCardClick(i)}
          />            
        )}
      </div>
      <Button variant="contained" color="primary" onClick={this.playCards}>
        Play cards
      </Button>
    </div>
    )
  }

  isCzar = () => {
    return this.props.state.currentCzarIndex == LocalPlayerIndex()
  }

  handleCardClick = (i) => {
    if(this.isCzar()){
      return
    }
    let newList = this.state.cardIndexes.slice()
    if(newList.includes(i)){
      newList.splice(newList.indexOf(i), 1)
    } else {
      newList.push(i)
    }
    this.setState({cardIndexes: newList})
  }

  playCards = () => {
    axios.post('rest/test/'+LocalPlayerIndex()+'/PlayCards', {
      cardIndexes: this.state.cardIndexes
    }).then(r => {
      this.setState({cardIndexes: []})
    }).catch(r => window.alert(r.response.data)); // We'll need prettier things
  }
}

class Game extends Component {
  render() {
    if(this.state == null) return null;
    return (
      <div className="Game">
        <PlayersInfo state={this.state} />
        <Table state={this.state} />
        <CardsInPlay state={this.state} owner={LocalPlayerIndex()} />
        <Hand state={this.state} />        
      </div>
    );
  } 
  componentWillMount() {
    this.updateState()
    // this would be much better with websockets
    window.setInterval(this.updateState, 500)
  }
  updateState = () => {
    fetch("rest/test/"+LocalPlayerIndex()+"/State")
      .then(r => r.json()
      .then(j => console.log(j) & this.setState(j))
    ) 
  }
}

export default Game