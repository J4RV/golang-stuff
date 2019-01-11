// for fast dirty dev
document.cookie = {
  currPlayer: 0
}
export default () => document.cookie.currPlayer;