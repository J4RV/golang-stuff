// User
export const loginUrl = "rest/user/login"
export const logoutUrl = "rest/user/logout"
export const registerUrl = "rest/user/register"
export const validCookieUrl = "rest/user/validcookie"

// Game state
export const getGamestateUrl = (stateID) => `rest/gamestate/${stateID}/State`
export const playCardsUrl = (stateID) => `rest/gamestate/${stateID}/PlayCards`
export const chooseWinnerUrl = (stateID) => `rest/gamestate/${stateID}/ChooseWinner`
