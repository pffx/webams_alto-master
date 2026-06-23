import React, { Component } from 'react';

import { connect } from 'react-redux'

import Select from '../../../components/widget/select';

const data = [
  {value:"card",label: 'Card', isHeader: true },
  {value:"nt",label:"NT"},
  {value:"lt",label:"LT"},
  {value:"shelf",label:"Shelf"},
]
class SoftManagementMaintenPage extends Component {
  constructor(props) {
    super(props);
    this.state = {
      cardItem: {value:"lt",label:"LT"},
    };
    if(this.props.isLogin){

    }
  }
  

  render() {
    return (
      <>
        <div style={{margin:"0.625rem"}} className='card'>
          <div className='card-header'>
            <div>ONT Software Overview</div>
          </div>
          <div className="card-body">
          </div>
        </div>
        <div style={{margin:"0.625rem"}} className='card'>
          <div className='card-header'>
            <div>ONT Preferred Software</div>
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
export default  connect(stateToProps, dispatchToProps)(SoftManagementMaintenPage)