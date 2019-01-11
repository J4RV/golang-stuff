import React, { Component } from 'react';
import Card from './Card'
import './App.css'

class App extends Component {
  state = undefined;
  render() {    
    return (
      <div className="App">
        <header className="cah-hand">
          {this.state ? this.state.Players[0].Hand.map(c =>
            <Card text={c.text} isBlack={false} />            
          ) : <Card text={"Loading hand..."} isBlack={false} />}
        </header>
      </div>
    );
  } 
  componentWillMount() {
    fetch("rest/test/0/State")
      .then(r => console.log(r.json()
      .then(j => console.log(j) & this.setState(j))
    ))
  }
}

export default App;
