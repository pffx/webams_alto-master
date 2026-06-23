import React, { Component } from 'react';

import { connect } from 'react-redux'
class HardwareSystemPage extends Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <>
        <div style={{marginRight: "0.625rem"}}>
          <h3> System hardware </h3>
        </div>
        <div >
        <h4> System hardware sub  sub sub sub</h4>
        </div>
      </>
    );
  }
}

const stateToProps = (state) => {
  return {
    isLogin: state.LoginReducer.isLogin,
  }
}
const dispatchToProps = (dispatch) => {
  return {

  }
}
export default  connect(stateToProps, dispatchToProps)(HardwareSystemPage)