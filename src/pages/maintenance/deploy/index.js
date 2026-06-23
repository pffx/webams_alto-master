import React, { Component } from 'react';


import { toast } from 'react-toastify';
import { connect } from 'react-redux'

import Typography from '@nokia-csf-uxr/ccfk/Typography';
import Select from '../../../components/widget/select';
import SimpleCard from '@nokia-csf-uxr/ccfk/Card';
import i18n from '../../../locales/config';

import AXIOS from '../../../axios';
import Button from '../../../components/widget/button';
import ConfirmationDialog from '../../../components/widget/dialog/confirm';
import WarningDialog from '../../../components/widget/dialog/warning';

import {API_ProvisionOltGuiDeploy,API_ProvisionOltResetAll, API_ProvisionOltGuiUndeploy} from '../../../global/API'
import GLOBAL,{TOAST_CONF,COLOR} from '../../../global';
// import { OLT_PORT_NAME } from '../../../global/index'
import utils from '../../../global/utils';
class DeployMaintenPage extends Component {
  constructor(props) {
    super(props);
    this.state = {
      showConfirm:false,
      confirmMessage:"",
      showWarning:false,
      wrongMessage:"",
      card:"",
      nt:"unknown",
      standby_nt:"unknown",
      lt_1:"unknown",
      lt_2:"unknown",
    };
  }

  componentDidMount (){
    this.getWebGuiStatus()
  }

  getWebGuiStatus(){
    AXIOS
      .get(API_ProvisionOltGuiDeploy)
      .then((res) => {
        if(res.data.status!== GLOBAL.ERROR_NUM.Success){
          toast.error(res.data.message,TOAST_CONF)
        }else{
          let data = res.data.data
          // data = "active nt : not-installed,lt-1 : not-installed,lt-2 : not-installed"
          let tmp = data.split(",")
          console.log("getWebGuiStatus  tmp =  ",tmp)
          if(tmp.length>=3){
            this.setState({
              nt:tmp[0].split(":")[1],
              // standby_nt:"unknown",
              lt_1:tmp[1].split(":")[1],
              lt_2:tmp[2].split(":")[1],
            })
          }
        }

        // this.setState({
        //   nt:"installed version=Fiber-0.52a valid=Yes",
        //   standby_nt:"unknown",
        //   lt_1:"installed version=Fiber-0.52a valid=Yes",
        //   lt_2:"installed version=Fiber-0.52a valid=Yes",
        // })
      })
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
    resetData.append("oltPort", "830")
    AXIOS
    .postFormData(API_ProvisionOltResetAll,resetData)
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

  handlUndeploy = ()=>{
    AXIOS
    .post(API_ProvisionOltGuiUndeploy)
    .then(resp => {
        console.log("API_ProvisionOltGuiUndeploy  resp=",resp)
        if(resp.data.status!== GLOBAL.ERROR_NUM.Success){
          toast.error(resp.data.message,TOAST_CONF)
        }else{
          // this.getWebGuiStatus()
          let data = resp.data.data
          // data = "active nt : not-installed,lt-1 : not-installed,lt-2 : not-installed"
          let tmp = data.split(",")
          console.log("handle unDeploy  tmp =  ",tmp)
          if(tmp.length>=3){
            this.setState({
              nt:tmp[0].split(":")[1],
              lt_1:tmp[1].split(":")[1],
              lt_2:tmp[2].split(":")[1],
            })
            toast.success("Undeploy success!",TOAST_CONF)
          }
        }
    });
  }

  handleDeploy = ()=>{
    // const deployData = new FormData();
    // deployData.append("oltId", this.props.oltInfor.ip)
    // deployData.append("oltPort", "830")
    AXIOS
    .post(API_ProvisionOltGuiDeploy)
    .then(resp => {
        // console.log("test file uploading  resp=",resp)
        if(resp.data.status!== GLOBAL.ERROR_NUM.Success){
          toast.error(resp.data.message,TOAST_CONF)
        }else{
          // this.getWebGuiStatus()
          let data = resp.data.data
          // data = "active nt : not-installed,lt-1 : not-installed,lt-2 : not-installed"
          let tmp = data.split(",")
          console.log("handleDeploy  tmp =  ",tmp)
          if(tmp.length>=3){
            this.setState({
              nt:tmp[0].split(":")[1],
              lt_1:tmp[1].split(":")[1],
              lt_2:tmp[2].split(":")[1],
            })
            toast.success("Deploy success!",TOAST_CONF)
          }
        }
    });
  }

  handlFactoryReset = ()=>{
    // if(this.props.oltInfor.ip === ""){
    //   this.setState({
    //     showWarning:true,
    //     wrongMessage:"Please select one OLT in Home page!"
    //   })
    // }else{
      // if( this.state.card === ""){
      //   this.setState({
      //     showWarning:true,
      //     wrongMessage:"Please select the OLT card!",
      //   })
      //   // return false
      // }else{
        this.setState({
          showConfirm:true,
          confirmMessage:"Factory reset will delete all configuration, are you sure??"
        })
      // }
    // }
  }

  handlRefresh = ()=>{
    console.log("refresh     !")
    this.setState({
      nt:"unknown",
      lt_1:"unknown",
      lt_2:"unknown",
    })
    this.getWebGuiStatus()
  }

  render() {
    return (
      <>
        <div style={{margin:"0.625rem"}} className='card'>
          <div className='card-header'>
            <div className='row3' >
              <div style={{width:"70%"}}>WebGui Summary</div>
              <div>
              <Button onClick={this.handlRefresh} title={i18n.t('button.refresh')} ></Button>
              </div>
            </div>
          </div>
          <div className="card-body">
            <SimpleCard style={{padding: "1rem",}} variant="stack">
                <table style={{ width: "100%"}}>
                  <tbody id="alarm" >
                    <tr>
                      <th>Active NT: </th>
                      <td>{this.state.nt}</td>
                    </tr>
                    {/* <tr>
                      <th>Standby NT</th>
                      <td>{this.state.standby_nt}</td>
                    </tr> */}
                    <tr>
                      <th>LT-1: </th>
                      <td>{this.state.lt_1}</td>
                    </tr>
                    <tr>
                      <th>LT-2: </th>
                      <td>{this.state.lt_2}</td>
                    </tr>
                  </tbody>
                </table>
            </SimpleCard>
            <div  className='row2'>
              <Button onClick={this.handleDeploy} title={"Deploy"} ></Button>
              <Button onClick={this.handlFactoryReset} title={"Factory Reset"} ></Button>
              <Button onClick={this.handlUndeploy} title={"Undeploy"} ></Button>
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
export default  connect(stateToProps, dispatchToProps)(DeployMaintenPage)