import { GlobalSets } from '../actions/constants'
const defaultState = {
  tabIndex: 'home',
  subTabIndex: '',
  oltInfor:{
    ip:"",
    account:"",
    pwd:"",
    status:"",
    hostname:"",
    software:"",
    type:"",
    ltNum:0,
    ltCardStatus:[0,0,0,0,0,0,0,0,0,0,0,0,0,0],
    controlled:false,
  },
}
const GlobalReducer = (state = defaultState, action) => {
  let newState = JSON.parse(JSON.stringify(state))
  const { type, data } = action
  switch(type) {
    case GlobalSets.ChangeTabIndex:
      newState.tabIndex = data.index
      newState.subTabIndex = 'backup_restore'
      return newState
    case GlobalSets.ChangeSubTabIndex:
      newState.subTabIndex = data.index
      return newState
    case GlobalSets.ClearAllTabIndex:
      newState.tabIndex = "home"
      newState.subTabIndex = ""
      return newState
    case GlobalSets.UpdateOltInfor:
        newState.oltInfor = data
        return newState
    case GlobalSets.ClearOltInfor:
        newState.oltInfor = {
          ip:"",
          account:"",
          pwd:"",
          status:"",
          hostname:"",
          software:"",
          type:"",
          ltNum:0,
          controlled:false,
          ltCardStatus:[0,0,0,0,0,0,0,0,0,0,0,0,0,0],
        }
        return newState
    default:
        return newState
  }
}

export default GlobalReducer