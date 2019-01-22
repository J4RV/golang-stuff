export const REMOVE_ERROR = 'REMOVE_ERROR'

export const removeErrorReduce = (state, action) => {
  let { index } = action.payload
  return { ...state, errors: state.errors.splice(index, 1) }
}

export default (index) => ({
  type: REMOVE_ERROR,
  payload: { index: index }
})