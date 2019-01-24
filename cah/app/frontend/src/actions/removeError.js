export const REMOVE_ERROR = 'REMOVE_ERROR'

export const removeErrorReduce = (state, action) => {
  let { index } = action.payload
  console.log("Removing error at index", index)
  let newValue = [...state.errors]
  newValue.splice(index, 1)
  return { ...state, errors: newValue }
}

export default (index) => ({
  type: REMOVE_ERROR,
  payload: { index: index }
})