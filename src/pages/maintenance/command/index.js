import React, { Component } from 'react';
import CLIWindowWidget from "../../../components/widget/cli"
import {WS_CMD} from '../../../global/API';
import {COLOR} from '../../../global';

import { connect } from 'react-redux'
import { Typography } from '@nokia-csf-uxr/ccfk';
const enter = window.navigator.platform === 'win32' ? '\r\n' : '\n';
console.log("window.navigator =  ",window.navigator)
class CommandMaintenPage extends Component {
  constructor(props) {
    super(props);
    this.state = {
      socketStatus:false,
    };
    // if(this.props.isLogin){

    // }
  }
  cliWindow = React.createRef();
  componentDidMount (){
    // console.log("componentDidMount   GLOBAL = ",GLOBAL);
    this.initSocket();
    // this.cliWindow.current.focus();
  }
  initSocket() {
    //console.log("initSocket  socket = ",this.socket);
    // console.log("CommandMaintenPage   initSocket  WS_CMD = ",WS_CMD);
    if(this.socket == null){
      this.socket = new WebSocket(WS_CMD);
      this.socketOnClose();
      this.socketOnOpen();
      this.socketOnError();
    }
  }
  socketOnOpen() {
    this.socket.onopen = () => {
      console.log('CommandMaintenPage   open socket')
      //this.initTerm()
      // this.socket.send(enter);
      this.setState({
          socketStatus: true,
      })
    }
  }
  socketOnClose() {
    this.socket.onclose = () => {
      console.log('CommandMaintenPage   close socket')
        this.setState({
          socketStatus: false,
      })
    }
  }
  socketOnError() {
    this.socket.onerror = () => {
      console.log('CommandMaintenPage   socket failed')
      this.setState({
        socketStatus: false,
      })
    }
  }


  render() {
    return (
      <div style={{marginTop:"1rem", marginLeft:"1rem",marginRight:"1rem",width:"95%"}}>
        {this.state.socketStatus ?<CLIWindowWidget  ref={this.cliWindow} socket={this.socket}/> :<div><Typography style={{fontSize:"1rem", maxWidth:"40%",color: COLOR.Critical}}>Socket connection failed!</Typography></div>}
      </div>
    );
  }
}

const stateToProps = (state) => {
  return {
    // isLogin: state.LoginReducer.isLogin,
  }
}
const dispatchToProps = (dispatch) => {
  return {

  }
}
export default  connect(stateToProps, dispatchToProps)(CommandMaintenPage)