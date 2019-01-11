import React from 'react';
import axios from 'axios'
import LocalPlayerIndex from './LocalPlayerIndex'

const handleOnClick = (playable, handIndex) => {
  console.log("Clicked on card " + handIndex)
  if (playable) {
    // FIXME this wont work with cards that need more than one white
    axios.post('rest/test/'+LocalPlayerIndex()+'/PlayCards', {
      cardIndexes: [handIndex]
    }).then(r => console.log(r));
  }
}

const Card = ({text, isBlack, playable, handIndex, className}) => (
  <div 
    className={className + ` cah-card ${isBlack ? 'cah-card-black' : 'cah-card-white'}`}
    onClick={() => handleOnClick(playable, handIndex)}>
    <span>{text}</span>
  </div>
)

export default Card;
