import { LoginSets } from '../actions/constants'
const defaultState = {
  isLogin: false,
  token:'',
  userInfo:{
    account:"",
    role:"",
    department:"",
  },
}

const LoginReducer = (state = defaultState, action) => {
  let newState = JSON.parse(JSON.stringify(state))
  // console.log("login reducer    action= ",action)
  const { type, data } = action
  // console.log("login reducer    data= ",data)
  switch(type) {
    // case Add_user:
    //   return state.push('awen')
    case LoginSets.Login:
      newState.isLogin = true
      newState.userInfo.account = data.account
      newState.userInfo.role = data.userInfo.role
      newState.userInfo.department = data.userInfo.department
      newState.token = data.token
      return newState
    case LoginSets.Logout:
      newState.isLogin = false
      newState.token = ''
      return newState
    case LoginSets.updateT:
      newState.token = data
      return newState
    case LoginSets.CheckT:
      //newState.token = data
      return newState
    default:
        return newState
  }
}

export default LoginReducer