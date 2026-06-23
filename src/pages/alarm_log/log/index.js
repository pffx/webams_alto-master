import React, { Component } from 'react';

import { connect } from 'react-redux'
class LogPage extends Component {
  constructor(props) {
    super(props);
  }
  componentDidMount(){
    console.log("LogPage   oltinfo = ",[this.props.oltinfo])
  }

  render() {
    return (
      <>
        <div style={{marginRight: "0.625rem"}}>
          <h3>  Alarmn Log Page LogPage </h3>
        </div>
        <div >
        <h4>  Alarmn Log Page LogPage sub  sub sub sub</h4>
        </div>
      </>
    );
  }
}

const stateToProps = (state) => {
  return {
    oltinfo: state.GlobalReducer.oltInfor,
  }
}
const dispatchToProps = (dispatch) => {
  return {

  }
}
export default  connect(stateToProps, dispatchToProps)(LogPage)