import { FeatureListSets } from '../actions/constants'
const defaultState = {
  feature:{
  },
}

const FeatureListReducer = (state = defaultState, action) => {
  let newState = JSON.parse(JSON.stringify(state))
  // console.log("login reducer    action= ",action)
  const { type, data } = action
  // console.log("feature reducer    data= ",data)
  switch(type) {
    case FeatureListSets.Init:
      newState.feature = data
      return newState
    
    default:
        return newState
  }
}

export default FeatureListReducer