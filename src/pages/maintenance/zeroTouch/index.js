import React, { Component } from 'react';

import { connect } from 'react-redux'
class ZeroTouchMaintenPage extends Component {
  constructor(props) {
    super(props);
    this.state = {
    };
    if(this.props.isLogin){

    }
  }

  render() {
    return (
      <>
        <div style={{margin:"0.625rem"}} className='card'>
          <div className='card-header'>
            <div>Queued ONTs</div>
          </div>
          <div className="card-body">
          </div>
        </div>
        <div style={{margin:"0.625rem"}} className='card'>
          <div className='card-header'>
            <div>Blacklisted ONTs</div>
          </div>
          <div className="card-body">
          </div>
        </div>
        <div style={{margin:"0.625rem"}} className='card'>
          <div className='card-header'>
            <div>Blacklisted PONs</div>
          </div>
          <div className="card-body">
          </div>
        </div>
        <div style={{margin:"0.625rem"}} className='card'>
          <div className='card-header'>
            <div>Blacklisted LTs</div>
          </div>
          <div className="card-body">
          </div>
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
export default  connect(stateToProps, dispatchToProps)(ZeroTouchMaintenPage)