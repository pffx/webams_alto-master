import React, { Component } from 'react';


import { toast } from 'react-toastify';
import { connect } from 'react-redux'

import Typography from '@nokia-csf-uxr/ccfk/Typography';
import Select from '../../../components/widget/select';

import AXIOS from '../../../axios';
import Button from '../../../components/widget/button';
import ConfirmationDialog from '../../../components/widget/dialog/confirm';
import WarningDialog from '../../../components/widget/dialog/warning';

import {API_ProvisionOltReset} from '../../../global/API'
import GLOBAL,{TOAST_CONF,COLOR} from '../../../global';
// import { OLT_PORT_NAME } from '../../../global/index'
import utils from '../../../global/utils';
class FactoryResetMaintenPage extends Component {
  constructor(props) {
    super(props);
    this.state = {
      showConfirm:false,
      confirmMessage:"",
      showWarning:false,
      wrongMessage:"",
      card:"",
    };
  }
  onWarningClose=()=>{
    this.setState({
      showWarning:false,
      wrongMessage:""
    })
  }
  onConfirm=()=>{
    const resetData = new FormData();
    resetData.append("oltId", this.props.oltInfor.ip)
    resetData.append("oltPort", utils.getOltPort2(this.props.oltInfor,this.state.card))
    AXIOS
    .postFormData(API_ProvisionOltReset,resetData)
    .then(resp => {
        // console.log("test file uploading  resp=",resp)
        if(resp.data.status!== GLOBAL.ERROR_NUM.Success){
          toast.error(resp.data.message,TOAST_CONF)
        }else{
          toast.success("Reset success!",TOAST_CONF)
        }
    });
    this.setState({
      showConfirm:false,
      confirmMessage:""
    })
  }
  onCancel=()=>{
    this.setState({
      showConfirm:false,
      confirmMessage:""
    })
  }

  handlFactoryReset = ()=>{
    if(this.props.oltInfor.ip === ""){
      this.setState({
        showWarning:true,
        wrongMessage:"Please select one OLT in Home page!"
      })
    }else{
      if( this.state.card === ""){
        this.setState({
          showWarning:true,
          wrongMessage:"Please select the OLT card!",
        })
        // return false
      }else{
        this.setState({
          showConfirm:true,
          confirmMessage:"Factory reset will delete all configuration, are you sure??"
        })
      }
    }
  }

  onOltCardSelected = (item)=>{
    this.setState({
      card:item.value,
    })
  }

  render() {
    return (
      <>
        <div style={{margin:"0.625rem"}} className='card'>
          <div className='card-header'>
            <div>Reset OLT to Defaults</div>
          </div>
          <div className="card-body">
            <div  className='row2'>
              <div  style={{ width: "50%"}}>
                <Select  dataItems={utils.generateCardList(this.props.oltInfor)} title={"Card Selection"} onChange={this.onOltCardSelected} selectedItem={this.state.card}> </Select>
              </div>
              <Button onClick={this.handlFactoryReset} title={"Factory Reset"} ></Button>
            </div>
          </div>
        </div>
        <ConfirmationDialog open={this.state.showConfirm} body={this.state.confirmMessage} onConfirm={this.onConfirm}  onCancel={this.onCancel}></ConfirmationDialog>
        <WarningDialog open={this.state.showWarning} body={this.state.wrongMessage} onRightClick={this.onWarningClose}></WarningDialog>
      </>
    );
  }
}

const stateToProps = (state) => {
  return {
    oltInfor: state.GlobalReducer.oltInfor,
  }
}
const dispatchToProps = (dispatch) => {
  return {

  }
}
export default  connect(stateToProps, dispatchToProps)(FactoryResetMaintenPage)