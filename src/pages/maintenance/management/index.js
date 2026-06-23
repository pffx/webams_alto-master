import React, { Component } from 'react';
import { connect } from 'react-redux'

import Select from '../../../components/widget/select';
import CliWindow from '../../../components/widget/cli';
import Button from '../../../components/widget/button';
import HorizontalSlider from '../../../components/widget/horizontalSlider';
import WarningDialog from '../../../components/widget/dialog/warning';

import i18n from '../../../locales/config';
import {API_ProvisionOltPing, API_NTGATEWAY} from "../../../global/API"
import AXIOS from '../../../axios';
import UTILS from '../../../global/utils';
import { TOAST_CONF } from '../../../global';
import { toast } from 'react-toastify';
import GLOBAL from '../../../global';

import HorizontalDivider from '@nokia-csf-uxr/ccfk/HorizontalDivider';
import TextInput, { TextInputLabelContent } from '@nokia-csf-uxr/ccfk/TextInput';
import Typography from '@nokia-csf-uxr/ccfk/Typography';
import Label, {LabelHelpIcon} from '@nokia-csf-uxr/ccfk/Label';
const DESTINATION_TYPE=
[
  // {label: 'Type', isHeader: true },
  {value:"ip",label:"IP address"},
  {value:"router",label:"Default Router"},
  // {value:"ntp",label:"NTP Server"},
]
const SOURCE_TYPE=
[
  // {label: 'Type', isHeader: true },
  // {value:"alto",label:"Alto"},
  {value:"nt",label:"NT card"},
]

const ACTION_TYPE=
[
  // {label: 'Type', isHeader: true },
  {value:"ping",label:"Ping"},
  {value:"traceroute",label:"Trace Route"},
]
class ManagementMaintenPage extends Component {
  constructor(props) {
    super(props);
    this.state = {
      destination:"ip",
      source:"nt",
      mode:"ping",
      destinationIP:"",
      showWarning:false,
      wrongMessage:"",
      gateway:"",
      attempts:0,
    };
  }
  cliWindow = React.createRef();
  componentDidMount(){
    this.getNTGateway()
  }
  getNTGateway = ()=>{
    let data = {
      oltId: this.props.oltInfor.ip,
      oltPort:830,
    }
    AXIOS
      .get(API_NTGATEWAY,data)
      .then((res) => {
        console.log(res)
        if(res.data.status === GLOBAL.ERROR_NUM.Success){
          this.setState({
            gateway:res.data.data
          })
        }
        // this.setState({
        //   gateway:"10.10.1.1"
        // })
      })
      .catch((err) => {
    });
  }
  onWarningClose=()=>{
    this.setState({
      showWarning:false,
      wrongMessage:""
    })
  }

  onDestinationItemChange = (item)=>{
    this.setState({
      destination:item.value
    })
  }

  onSourceItemChange = (item)=>{
    this.setState({
      source:item.value
    })
    if(item.value === "nt"){
      if(this.props.oltInfor.ip == ""){
        this.setState({
          showWarning:true,
          wrongMessage:"Please select one OLT in Home Page!",
        })
      }else{
        this.getNTGateway()
      }
    }
  }
  onModeItemChange = (item)=>{
    this.setState({
      mode:item.value
    })
  }
  onAttamptsChange = (data)=>{
    console.log("onSliderChange=",data)
    this.setState({
      attempts:data
    })
  }

  handlStartPing = ()=>{
    let desIP
    if(this.state.destination === "ip"){
      desIP = this.state.destinationIP
    }else if(this.state.destination === "router"){
      desIP = this.state.gateway
    }else{
      desIP=""
    }
    if(desIP === ""){
      this.setState({
        showWarning:true,
        wrongMessage:"The destination IP address is empty!",
      })
      return false
    }
    if(!UTILS.isValidIpAddress(desIP)){
      this.setState({
        showWarning:true,
        wrongMessage:"Please input a valid destination IP address!",
      })
    }
    let info = this.state.mode + " " + desIP +" excuting.... " + "\n"
    this.cliWindow.current.write(info);

    let data = {
      oltId: this.props.oltInfor.ip,
      oltPort:830,
      mode:this.state.mode,
      source: this.state.source,
      dstIP: desIP,
      attempts:this.state.attempts
    }
    AXIOS
      .post(API_ProvisionOltPing,data)
      .then((res) => {
        console.log(res)
        if(res.data.status === GLOBAL.ERROR_NUM.Success){
          this.cliWindow.current.write(res.data.data);
        }else{
          this.cliWindow.current.write("Time out..." + "\n");
        }
      })
      .catch((err) => {
        this.cliWindow.current.write("Ping failed" + "\n");
    });
  }

  renderDestinationIP(){
    if(this.state.destination=== "ip"){
      return(
        <div style={{width:"50%"}}>
          <Label >
            <TextInputLabelContent >{"IP address"}</TextInputLabelContent>
          </Label>
          <TextInput value={this.state.destinationIP} onChange={(event) => {this.setState({destinationIP:event.target.value})}}   />
        </div>
      )
    }else if(this.state.destination=== "router"){
      return(
        <div style={{width:"50%"}}>
          <Typography variant="BODY">{this.state.gateway}</Typography>
        </div>
      )

    }else{
      return(
        <div></div>
      )
    }
  }

  render() {
    return (
      <>
        <div style={{margin:"0.625rem"}} className='card'>
          <div className='card-header'>
            <div>Reachability Test</div>
          </div>
          <div className="card-body">
          <Select  dataItems={ACTION_TYPE} title={"Mode"} onChange={this.onModeItemChange} selectedItem={this.state.mode}> </Select>
          <Select  dataItems={SOURCE_TYPE} title={"Source"} onChange={this.onSourceItemChange} selectedItem={this.state.source}> </Select>
          <div className='row' style={{width:"100%"}}>
            <Select style={{width:"50%"}} dataItems={DESTINATION_TYPE} title={"Destination"} onChange={this.onDestinationItemChange} selectedItem={this.state.destination}> </Select>
            {this.renderDestinationIP()}
          </div>
          <div className='row'>
            <HorizontalSlider horizonSliderChange={this.onAttamptsChange} style={{width:"94%"}} type="icon" label="Attempts" min={1} max={10}></HorizontalSlider>
            <Button onClick={this.handlStartPing} title={i18n.t('button.start')} ></Button>
          </div>
          
          <HorizontalDivider className="horizontal-divider" />
          <CliWindow style={{height:"15rem", width:"90%"}} ref={this.cliWindow}/>
          </div>
        </div>
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
export default  connect(stateToProps, dispatchToProps)(ManagementMaintenPage)