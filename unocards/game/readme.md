
##1st iteration
###game: one player, one vanilla deck, one "table"  
 * can play cards if the table top card matches the colour or number (or playing wild card)
 * can change colors with wild cards
 * draws cards if they play draw cards
 * reverse and skip do nothing in this iter
 * game checks its hand size to finish the game

###gamestate:
```
State {
  Players []Player
  Currplayer Player
  Board Board
  Skip bool
  Orderreversed bool
  Drawacum int
  ?
}
Player {
  Name string
  Hand Card[]
}
Board {
  Topcard Card
  Discard Card[]
  Deck Deck
}
```
