import React, { Component } from 'react';

import { connect } from 'react-redux'
// import { FileUploadListSubComponentStory } from '@nokia-csf-uxr/ccfk/FileUpload';

import ExpansionPanels from './expansionPanel';
import WarningDialog from '../../../components/widget/dialog/warning';
import Button from '../../../components/widget/button';
import {API_OltSoftwareInfo, API_ServerSoftware, API_OltList} from "../../../global/API"
import GLOBAL,{TOAST_CONF,COLOR} from '../../../global';
import utils from '../../../global/utils';
import AXIOS from '../../../axios';
import i18n from '../../../locales/config';
import Typography from '@nokia-csf-uxr/ccfk/Typography';
import Label, {LabelHelpIcon} from '@nokia-csf-uxr/ccfk/Label';
import Tooltip from '@nokia-csf-uxr/ccfk/Tooltip';
import HelpCircleIcon from '@nokia-csf-uxr/ccfk-assets/HelpCircleIcon';
import { toast } from 'react-toastify';
import { toast } from 'react-toastify';
import { updateOltInfor } from '../../../actions/global'
import {
  parseOltCardSoftwareInfo,
  parseOltSoftwareInfo,
  getHostName as resolveHostName,
} from '../../../utils/softwareUpgrade';
const RESULT2 = {
  "XMLName": {
    "Space": "",
    "Local": "data"
  },
  "HardwareState": {
    "XMLName": {
      "Space": "urn:ietf:params:xml:ns:yang:ietf-hardware",
      "Local": "hardware-state"
    },
    "Component": {
      "XMLName": {
        "Space": "urn:ietf:params:xml:ns:yang:ietf-hardware",
        "Local": "component"
      },
      "Name": "Chassis",
      "Software": {
        "XMLName": {
          "Space": "urn:bbf:yang:bbf-software-image-management-one-dot-one",
          "Local": "software"
        },
        "Software": {
          "XMLName": {
            "Space": "urn:bbf:yang:bbf-software-image-management-one-dot-one",
            "Local": "software"
          },
          "Name": "application_software",
          "Download": {
            "XMLName": {
              "Space": "urn:bbf:yang:bbf-software-image-management-one-dot-one",
              "Local": "download"
            },
            "CurrentState": {
              "XMLName": {
                "Space": "urn:bbf:yang:bbf-software-image-management-one-dot-one",
                "Local": "current-state"
              },
              "State": "idle",
              "Timestamp": "1970-01-30T01:46:32+00:00"
            },
            "LastDownloadState": {
              "XMLName": {
                "Space": "urn:bbf:yang:bbf-software-image-management-one-dot-one",
                "Local": "last-download-state"
              },
              "State": "successful",
              "Timestamp": "1970-01-30T01:46:32+00:00",
              "SoftwareName": "L6GQAG2209.421"
            }
          },
          "Revisions": {
            "XMLName": {
              "Space": "urn:bbf:yang:bbf-software-image-management-one-dot-one",
              "Local": "revisions"
            },
            "Revision": [{
                "XMLName": {
                  "Space": "urn:bbf:yang:bbf-software-image-management-one-dot-one",
                  "Local": "revision"
                },
                "Name": "L6GQAG2209.420",
                "DownloadTimestamp": "1970-01-30T00:37:49+00:00",
                "Version": "2209.420",
                "IsValid": "true",
                "IsCommitted": "false",
                "IsActive": "false"
              },
              {
                "XMLName": {
                  "Space": "urn:bbf:yang:bbf-software-image-management-one-dot-one",
                  "Local": "revision"
                },
                "Name": "L6GQAG2209.421",
                "DownloadTimestamp": "1970-01-30T01:46:32+00:00",
                "Version": "2209.421",
                "IsValid": "true",
                "IsCommitted": "true",
                "IsActive": "true"
              }
            ]
          }
        }
      }
    }
  },
  "Ip": "192.168.1.1",
  "Port": "845"
}
const RESULT = {
  "0": {
  "XMLName":{
      "Space":"",
      "Local":"data"
  },
  "HardwareState":{
      "XMLName":{
          "Space":"urn:ietf:params:xml:ns:yang:ietf-hardware",
          "Local":"hardware-state"
      },
      "Component":{
          "XMLName":{
              "Space":"urn:ietf:params:xml:ns:yang:ietf-hardware",
              "Local":"component"
          },
          "Name":"Chassis",
          "Software":{
              "XMLName":{
                  "Space":"urn:bbf:yang:bbf-software-image-management-one-dot-one",
                  "Local":"software"
              },
              "Software":{
                  "XMLName":{
                      "Space":"urn:bbf:yang:bbf-software-image-management-one-dot-one",
                      "Local":"software"
                  },
                  "Name":"application_software",
                  "Download":{
                      "XMLName":{
                          "Space":"urn:bbf:yang:bbf-software-image-management-one-dot-one",
                          "Local":"download"
                      },
                      "CurrentState":{
                          "XMLName":{
                              "Space":"urn:bbf:yang:bbf-software-image-management-one-dot-one",
                              "Local":"current-state"
                          },
                          "State":"idle",
                          "Timestamp":"1970-01-30T01:46:32+00:00"
                      },
                      "LastDownloadState":{
                          "XMLName":{
                              "Space":"urn:bbf:yang:bbf-software-image-management-one-dot-one",
                              "Local":"last-download-state"
                          },
                          "State":"successful",
                          "Timestamp":"1970-01-30T01:46:32+00:00",
                          "SoftwareName":"L6GQAG2209.421"
                      }
                  },
                  "Revisions":{
                      "XMLName":{
                          "Space":"urn:bbf:yang:bbf-software-image-management-one-dot-one",
                          "Local":"revisions"
                      },
                      "Revision":[
                          {
                              "XMLName":{
                                  "Space":"urn:bbf:yang:bbf-software-image-management-one-dot-one",
                                  "Local":"revision"
                              },
                              "Name":"L6GQAG2209.420",
                              "DownloadTimestamp":"1970-01-30T00:37:49+00:00",
                              "Version":"2209.420",
                              "IsValid":"true",
                              "IsCommitted":"false",
                              "IsActive":"false"
                          },
                          {
                              "XMLName":{
                                  "Space":"urn:bbf:yang:bbf-software-image-management-one-dot-one",
                                  "Local":"revision"
                              },
                              "Name":"L6GQAG2209.421",
                              "DownloadTimestamp":"1970-01-30T01:46:32+00:00",
                              "Version":"2209.421",
                              "IsValid":"true",
                              "IsCommitted":"true",
                              "IsActive":"true"
                          }
                      ]
                  }
              }
          }
      }
  },
  "Ip":"192.168.1.1",
  "HostName":"DF-16",
  "Port": "845",
},
  "1": {
  "XMLName":{
      "Space":"",
      "Local":"data"
  },
  "HardwareState":{
      "XMLName":{
          "Space":"urn:ietf:params:xml:ns:yang:ietf-hardware",
          "Local":"hardware-state"
      },
      "Component":{
          "XMLName":{
              "Space":"urn:ietf:params:xml:ns:yang:ietf-hardware",
              "Local":"component"
          },
          "Name":"Chassis",
          "Software":{
              "XMLName":{
                  "Space":"urn:bbf:yang:bbf-software-image-management-one-dot-one",
                  "Local":"software"
              },
              "Software":{
                  "XMLName":{
                      "Space":"urn:bbf:yang:bbf-software-image-management-one-dot-one",
                      "Local":"software"
                  },
                  "Name":"application_software",
                  "Download":{
                      "XMLName":{
                          "Space":"urn:bbf:yang:bbf-software-image-management-one-dot-one",
                          "Local":"download"
                      },
                      "CurrentState":{
                          "XMLName":{
                              "Space":"urn:bbf:yang:bbf-software-image-management-one-dot-one",
                              "Local":"current-state"
                          },
                          "State":"idle",
                          "Timestamp":"1970-01-30T01:46:32+00:00"
                      },
                      "LastDownloadState":{
                          "XMLName":{
                              "Space":"urn:bbf:yang:bbf-software-image-management-one-dot-one",
                              "Local":"last-download-state"
                          },
                          "State":"successful",
                          "Timestamp":"1970-01-30T01:46:32+00:00",
                          "SoftwareName":"L6GQAG2209.421"
                      }
                  },
                  "Revisions":{
                      "XMLName":{
                          "Space":"urn:bbf:yang:bbf-software-image-management-one-dot-one",
                          "Local":"revisions"
                      },
                      "Revision":[
                          {
                              "XMLName":{
                                  "Space":"urn:bbf:yang:bbf-software-image-management-one-dot-one",
                                  "Local":"revision"
                              },
                              "Name":"L6GQAG2209.420",
                              "DownloadTimestamp":"1970-01-30T00:37:49+00:00",
                              "Version":"2209.420",
                              "IsValid":"true",
                              "IsCommitted":"false",
                              "IsActive":"false"
                          }
                      ]
                  }
              }
          }
      }
  },
  "Ip":"10.10.1.1",
  "HostName":"DF-16",
  "Port": "840"
}
}

class SoftManagementSystemPage extends Component {
  constructor(props) {
    super(props);
    this.state = {
      olts:[],
      oltSoftwareInfo:[],
      oltCardSoftwareInfo:[],
      serverSoftware:{},
      // filePath:"",
      showWarning:false,
      wrongMessage:"",
    }
  }
  refreshOLTInfo(){
    AXIOS
      .get(API_OltList)
      .then((res) => {
        if(res.data.status == 200){
          let olt = this.props.oltInfo;
          res.data.olt_list.map((item, index)=>{
            // console.log("API_OltList   item = ",item)
            if(item.IP === this.props.oltInfo.ip && item.HostName === this.props.oltInfo.hostname ){
              olt.ltCardStatus = utils.generateLTCardPlanned(item)
            }
          })
          // let oltInfo = {
          //   ip:item.ip,
          //   hostname:item.hostname,
          //   status:item.status,
          //   software:item.software,
          //   type:item.type,
          //   ltNum:item.lt_num,
          //   ltCardStatus:item.ltCardStatus,
          //   controlled:true,
          // }
          this.props.updateOltInfor(olt)
        }
      })
      .catch((err) => {
    });
  }
  getOLTsSoftware(){
    console.log("getOLTsSoftware   oltInfo:",this.props.oltInfo)
    if(utils.isOltSelected(this.props.oltInfo)){
      this.getOLTsCardSoftware(0)
      if(this.props.oltInfo.type.startsWith("DF")){
        // no need to get other software information
      }else{
        //get the LTs of MF software information. 
        // for(let i = 0; i <= this.props.oltInfo.ltNum; i++) {
        for(let i = 1; i <= 14; i++) {
          if(this.props.oltInfo.ltCardStatus[i-1] == 1){
            this.getOLTsCardSoftware(i)
          }
        }
      }
    }
  }

  componentDidMount (){
    // utils.isTestDataUsed()
    this.getOLTsSoftware()
    this.getServerSoftware()
  }

  onWarningClose=()=>{
    this.setState({
      showWarning:false,
      wrongMessage:""
    })
  }

  handlRefresh = ()=>{
    console.log("refresh     !")
    this.setState({
      oltSoftwareInfo:[]
    })
    //refresh the olt info
    this.refreshOLTInfo()
    this.getOLTsSoftware()
  }

  onFileChange = (panel,path,name)=>{
    // console.log("onFileChange    panel = ",panel)
    // console.log("onFileChange    path = ",path)
    // this.setState({
    //   filePath:path
    // })
  }

  getServerSoftware(){
    AXIOS
      .get(API_ServerSoftware)
      .then((res) => {
        if(res.data.status === GLOBAL.ERROR_NUM.Success){
          let folders = JSON.parse(res.data.software_list)
          // console.log("API_ServerSoftware   folders = ",folders)
          // Object.keys(folders).map(key => {
          //   console.log("API_ServerSoftware   key = ",key)
          // })

          this.setState({
            serverSoftware:folders,
          })
        }
      })
      .catch((err) => {
    });
  }
  getHostName(port){
    return resolveHostName(this.props.oltInfo.type, port)
  }
  resolveOltSoftwareInfo(data){
    return parseOltSoftwareInfo(data, this.props.oltInfo.type)
  }
  resolveOltCardSoftwareInfo(data){
    return parseOltCardSoftwareInfo(data, this.props.oltInfo.type)
  }
  getOLTsCardSoftware(index){
    console.log("getOLTsCardSoftware   index = ",index)
    AXIOS
      .get(API_OltSoftwareInfo,{
        oltId: this.props.oltInfo.ip,
        dstPort:utils.getOltPort(this.props.oltInfo,index),
      })
      .then((res) => {
        if(res.data.status === GLOBAL.ERROR_NUM.Success){
          // let oltSoftInfo = JSON.parse(res.data.olt_software_info)
          let list =this.state.oltSoftwareInfo
          let info = utils.isTestDataUsed() ? this.resolveOltCardSoftwareInfo(RESULT2) :this.resolveOltCardSoftwareInfo(JSON.parse(res.data.olt_software_info))
          list.push(info)
          this.setState({
            oltSoftwareInfo:list,
          })
        }else{
          console.log("server error, show test data in test mode")
          if(utils.isTestDataUsed()){
            let list =this.state.oltSoftwareInfo
            let info = this.resolveOltCardSoftwareInfo(RESULT2)
            list.push(info)
            this.setState({
              oltSoftwareInfo:list,
            })
          }
          // toast.error("Server error",TOAST_CONF)
        }
      })
      .catch((err) => {
    });
  }
  renderOltCardVersoin(){
    return(
      <>
      </>
    )
  }
  renderOltsVersion(){
    return(
      <>
        <div className='card-header row3'>
          <div className='row2' style={{width:"70%"}}>
          <Typography style={{fontSize:"1rem",maxWidth:"20%",color: COLOR.Nokia_Blue}}>OLT Software List</Typography>
          <Tooltip placement="top" trigger={["hover", "focus" ]} fallbackPlacements={["bottom", "right" ]} modifiers={{offset: {offset: "[0, 8]"}}}
            tooltip="Put the software in the folder 'software'. ">
            <LabelHelpIcon>
                <HelpCircleIcon />
            </LabelHelpIcon>
          </Tooltip>
          </div>
          <div>
          <Button onClick={this.handlRefresh} title={i18n.t('button.refresh')} ></Button>
          </div>
        </div>
        <div className="card-body">
        <ExpansionPanels data={this.state.oltSoftwareInfo} softwareServer={this.state.serverSoftware} onChange={this.onFileChange}>
        </ExpansionPanels>
        </div>
        
      </>
    )
  }

  render() {
    return (
      <>
        <div style={{margin:"0.625rem"}} className='card'>

        {this.renderOltsVersion()}

        </div>
        <WarningDialog open={this.state.showWarning} body={this.state.wrongMessage} onRightClick={this.onWarningClose}></WarningDialog>
      </>
    );
  }
}

const stateToProps = (state) => {
  return {
    isLogin: state.LoginReducer.isLogin,
    oltInfo: state.GlobalReducer.oltInfor,
  }
}
const dispatchToProps = (dispatch) => {
  return {
    updateOltInfor(data) {
      dispatch(updateOltInfor(data))
    },
    // clearOltInfor(){
    //   dispatch(clearOltInfor())
    // }
  }
}
export default  connect(stateToProps, dispatchToProps)(SoftManagementSystemPage)