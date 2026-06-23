import React, { Component } from 'react';

import { connect } from 'react-redux'
import Select from '../../../components/widget/select';
import Button from '../../../components/widget/button';
import Label, {LabelHelpIcon} from '@nokia-csf-uxr/ccfk/Label';
import TextInput, { TextInputButton, TextInputLabelContent } from '@nokia-csf-uxr/ccfk/TextInput';
import Typography from '@nokia-csf-uxr/ccfk/Typography';
import Tooltip from '@nokia-csf-uxr/ccfk/Tooltip';
import HelpCircleIcon from '@nokia-csf-uxr/ccfk-assets/HelpCircleIcon';
import HorizontalDivider from '@nokia-csf-uxr/ccfk/HorizontalDivider';
import FileUploadComponent from '../../../components/widget/fileUpload';
import WarningDialog from '../../../components/widget/dialog/warning';
import AXIOS from '../../../axios';
import UTILS from '../../../global/utils';
import GLOBAL,{INITIALIAL_OLT_IP,OLT_TYPE,OLT_TYPE_PORT,TOAST_CONF,COLOR} from '../../../global';
import {API_ConfigFile} from '../../../global/API'
import { toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

class LoadCofnigMaintenPage extends Component {
  constructor(props) {
    super(props);
    this.state = {
      files:"",
      cardItem: "",
      IP:"",
      netmask:"",
      gateway:"",
      showWarning:false,
      wrongMessage:"",
    };
  }

  onWarningClose=()=>{
    this.setState({
      showWarning:false,
      wrongMessage:""
    })
  }
  canSubmitLoadConfig =()=>{
    if(this.state.files.length < 1 ){
      this.setState({
        showWarning:true,
        wrongMessage:"Please select one config file firstly!",
      })
      return false
    }
    // if(this.props.oltInfo.ip === ""){
    //   this.setState({
    //     showWarning:true,
    //     wrongMessage:"Please select one OLT in Home page firstly!",
    //   })
    // }
    if( this.state.cardItem === ""){
      this.setState({
        showWarning:true,
        wrongMessage:"Please select the OLT card!",
      })
      return false
    }
    if(this.state.IP === "" || this.state.netmask === "" || this.state.gateway === ""){
      this.setState({
        showWarning:true,
        wrongMessage:"Please complete the OAM information!",
      })
      return false
    }
    if(!UTILS.isValidIpAddress(this.state.IP)){
      this.setState({
        showWarning:true,
        wrongMessage:"Please input a valid IP address!",
      })
      return false
    }

    return true
  }
  handlLoadConfig = ()=>{
    if(!this.canSubmitLoadConfig()){
      return
    }
    const formData = new FormData();
    let flag = false
    this.state.files.forEach(file => {
      // console.log("handlLoadConfig   file.file= ",file.file.path)
      if(file.file.path.endsWith(".tpl") || file.file.path.endsWith(".xml")){
        formData.append("file", file.file);
        flag = true
      }
    });
    if(!flag){
      toast.error("Please select a correct template file.",TOAST_CONF)
      return
    }
    formData.append("oltId",INITIALIAL_OLT_IP)
    formData.append("dstPort",UTILS.getOltPort2(this.props.oltInfo,this.state.cardItem))
    if(this.state.IP !== ""){
      formData.append("ipAddr",this.state.IP)
    }
    if(this.state.netmask !== ""){
      formData.append("netMask",this.state.netmask)
    }
    if(this.state.gateway !== ""){
      formData.append("gateway",this.state.gateway)
    }
    
    AXIOS
    .postFormData(API_ConfigFile,formData)
    .then(resp => {
        // console.log("test file uploading  resp=",resp)
        if(resp.data.status!== GLOBAL.ERROR_NUM.Success){
          toast.error(resp.data.message,TOAST_CONF)
        }else{
          toast.success("Initial success!",TOAST_CONF)
        }
    });
  }

  onFilesChange = data=>{
    this.setState({
      files:data
    })
  }

  onCardChange = (item)=>{
    this.setState({
      cardItem:item.value
    })
  }

  render() {
    return (
      <>
      <div style={{margin:"0.625rem"}} className='card'>
        <div className='card-header row2'>
          <Typography style={{fontSize:"1rem", maxWidth:"20%",color: COLOR.Nokia_Blue}}>Load Configuration</Typography>
          <Tooltip placement="top" trigger={["hover", "focus" ]} fallbackPlacements={["bottom", "right" ]} modifiers={{offset: {offset: "[0, 8]"}}}
            tooltip="Initialize the OLT configuration through LEMI port when used for the first time. ">
            <LabelHelpIcon>
                <HelpCircleIcon />
            </LabelHelpIcon>
          </Tooltip>
        </div>
        <div className="card-body">
          <div  className='row2'>
            <div  style={{ width: "50%" ,marginBottom: "1rem"}}>
              {/* <Select required dataItems={OLT_TYPE} title={"OLT Type"} onChange={this.onCardChange} selectedItem={this.state.cardItem}> </Select> */}
              <Select  required dataItems={UTILS.generateCardList(this.props.oltInfo)} title={"Card Selection"} onChange={this.onCardChange} selectedItem={this.state.cardItem}> </Select>
              <Label maxWidth>
                <TextInputLabelContent required >{"OAM IP:"}</TextInputLabelContent>
              </Label>
              <TextInput
                variant="outlined"
                maxWidth
                type="text"
                disabled={false}
                value={this.state.IP}
                onChange={(event) => {this.setState({IP:event.target.value,});}}
                placeholder="Replaced the destination IP address for templete"
                inputProps={{ autoComplete: 'off' }}
                // error={this.state.feedbackMsg && this.state.feedbackMsg.variant === 'error'}
              />

              <Label maxWidth>
                <TextInputLabelContent required >{"OAM Netmask:"}</TextInputLabelContent>
              </Label>
              <TextInput
                variant="outlined"
                maxWidth
                type="text"
                disabled={false}
                value={this.state.netmask}
                onChange={(event) => {this.setState({netmask:event.target.value,});}}
                placeholder="Replaced the destination IP netmask for templete"
                inputProps={{ autoComplete: 'off' }}
                // error={this.state.feedbackMsg && this.state.feedbackMsg.variant === 'error'}
              />

              <Label maxWidth>
                <TextInputLabelContent required >{"OAM Gateway:"}</TextInputLabelContent>
              </Label>
              <TextInput
                variant="outlined"
                maxWidth
                type="text"
                disabled={false}
                value={this.state.gateway}
                onChange={(event) => {this.setState({gateway:event.target.value,});}}
                placeholder="Replaced the destination IP gateway for templete"
                inputProps={{ autoComplete: 'off' }}
                // error={this.state.feedbackMsg && this.state.feedbackMsg.variant === 'error'}
              />
            </div>
            <FileUploadComponent title={"Config File"} onFilesChange={this.onFilesChange}/>
          </div>
          <HorizontalDivider className="horizontal-divider" />
          <Button onClick={this.handlLoadConfig} title={"Load Config"} ></Button>
        </div>
      </div>
      <WarningDialog open={this.state.showWarning} body={this.state.wrongMessage} onRightClick={this.onWarningClose}></WarningDialog>
      </>
    );
  }
}

const stateToProps = (state) => {
  return {
    oltInfo: state.GlobalReducer.oltInfor,
  }
}
const dispatchToProps = (dispatch) => {
  return {

  }
}
export default  connect(stateToProps, dispatchToProps)(LoadCofnigMaintenPage)