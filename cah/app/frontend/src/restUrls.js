// User
export const loginUrl = "user/login"
export const logoutUrl = "user/logout"
export const registerUrl = "user/register"
export const validCookieUrl = "user/validcookie"

// Game state
export const getGamestateUrl = (stateID) => `gamestate/${stateID}/State`
export const playCardsUrl = (stateID) => `gamestate/${stateID}/PlayCards`
export const chooseWinnerUrl = (stateID) => `gamestate/${stateID}/ChooseWinner`
