import React , { Component } from 'react';
import { connect } from 'react-redux'
import HomePage from './pages/home';
import Redirect from './router/redirect'
import { clearAllTabIndex, clearOltInfor } from './actions/global'
import { CheckT } from './actions/login'

class AppHome extends Component {
  // constructor(props){
  //     super(props)
  // }
  componentDidMount(){
    this.props.initPage()
  }
  render(){
    return (
      this.props.isLogin
       ? <HomePage/>
       : <Redirect to="/login"/>
    );
  }
}
const stateToProps = (state) => {
  // console.log("state = ",state)
  return {
    isLogin: state.LoginReducer.isLogin,
    // tab: state.GlobalReducer.tabIndex,
    // subTab: state.GlobalReducer.subTabIndex,
  }
}
const dispatchToProps = (dispatch) => {
  return {
    initPage() {
      dispatch(CheckT())
      dispatch(clearAllTabIndex())
      dispatch(clearOltInfor())
    }
  }
}
export default  connect(stateToProps,dispatchToProps)(AppHome)
