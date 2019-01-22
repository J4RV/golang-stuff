export const PUSH_ERROR = 'PUSH_ERROR'

export const pushErrorReduce = (state, action) => {
  let { msg } = action.payload
  return { ...state, errors: [...state.errors, msg] }
}

export default (msg) => ({
  type: PUSH_ERROR,
  payload: { msg: msg }
})