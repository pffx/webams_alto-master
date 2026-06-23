import React, { Component } from 'react';
import { connect } from 'react-redux'
import { toast } from 'react-toastify';

import Select from '../../../components/widget/select';
import Button from '../../../components/widget/button';
import FileUploadComponent from '../../../components/widget/fileUpload';
import WarningDialog from '../../../components/widget/dialog/warning'
//import InformationDialog from '../../../components/widget/dialog/information'
//import ConfirmationDialog from '../../../components/widget/dialog/confirm'
import AXIOS from '../../../axios';
import {API_ProvisionOltBackup, API_ProvisionOltRestore} from '../../../global/API'
import {TOAST_CONF} from '../../../global';
import utils from '../../../global/utils';

const backup2LocalItems = [
  {label: 'Card', isHeader: true },
  {value:"830",label:"NT"},
  {value:"931",label:"LT1"},
]

const restore2OltItems = [
  {label: 'Card', isHeader: true },
  {value:"830",label:"NT"},
  {value:"931",label:"LT1"},
]

class BackupRestoreMaintenPage extends Component {
  constructor(props) {
    super(props);
    this.state = {
      backup2LocalCardValue:"",
      restore2OltCardValue:"",
      backup2LocalTypeValue:"df",
      restore2OltTypeValue:"df",
      restoreFile:"",
      showWarning:false,
      wrongMessage:"",
    };
    this.handlBackup2Local = this.handlBackup2Local.bind(this)
    this.handlRestore2Olt = this.handlRestore2Olt.bind(this)
  }

  onBackup2LocalTypeChange = (item)=>{
    this.setState({
      backup2LocalTypeValue:item.value
    })
  }

  onBackup2LocalCardChange = (item)=>{
    this.setState({
      backup2LocalCardValue:item.value
    })
  }
  onRestore2OltCardChange = (item)=>{
    this.setState({
      restore2OltCardValue:item.value
    })
  }

  onRestore2OltTypeChange = (item)=>{
    this.setState({
      restore2OltTypeValue:item.value
    })
  }

  onRestoreFilesChange = data=>{
    this.setState({
      restoreFile:data
    })
  }

  checkLocalField(){
    return true
  }

  handlBackup2Local(){
    if(this.props.oltInfo.ip === ""){
      this.setState({
        showWarning:true,
        wrongMessage:"Please select one OLT in Home page!"
      })
      return false
    }
    if( this.state.backup2LocalCardValue === ""){
      this.setState({
        showWarning:true,
        wrongMessage:"Please select the OLT card!",
      })
      return false
    }
    let data = {
      oltId: this.props.oltInfo.ip,
      dstPort: utils.getOltPort2(this.props.oltInfo,this.state.backup2LocalCardValue)
    }
    AXIOS
      .post(API_ProvisionOltBackup, data)
      .then((res) => {
        if(res.status === 200){
          //let content = JSON.stringify(res.data);
          //let blob = new Blob([res.data], { type: 'application/octet-stream,charset=UTF-8' })
          const fileName = utils.getFileNameFromContentDisposition(res.headers["content-disposition"])

          const aLink = document.createElement('a')
          document.body.appendChild(aLink)
          aLink.style.display='none'
          const objectUrl = window.URL.createObjectURL(new Blob([res.data]))
          aLink.href = objectUrl
          aLink.download = fileName
          aLink.click()
          document.body.removeChild(aLink)
          // this.setState({
          //   showWarning:true,
          //   wrongMessage:"There is no file!"
          // })
        }else if(res.status === 20008){
          console.log("file not exsit")
          this.setState({
            showWarning:true,
            wrongMessage:"There is no file!"
          })
        }
      })
      .catch((err) => {
    });
  }
  handlRestore2Olt(){
    if( this.state.restore2OltCardValue === ""){
      this.setState({
        showWarning:true,
        wrongMessage:"Please select the OLT card!",
      })
      return false
    }
    const formData = new FormData();
    this.state.restoreFile.forEach(file => {
      formData.append("file", file.file);
    });
    formData.append("oltId", this.props.oltInfo.ip)
    formData.append("dstPort", utils.getOltPort2(this.props.oltInfo,this.state.restore2OltCardValue))
    AXIOS
    .postFormData(API_ProvisionOltRestore, formData)
    .then(resp => {
        console.log("test file uploading  resp=",resp)
        if(resp.data.status === 200 && resp.data.data === "Sent rpc success!"){
          toast.success("Restore success!",TOAST_CONF)
        }else{
          toast.error("Restore failed!",TOAST_CONF)
        }
    });
  }
  onWarningClose=()=>{
    this.setState({
      showWarning:false,
      wrongMessage:""
    })
  }

  render() {
    return (
      <>
        <div className='row2'>
          <div style={{margin:"0.625rem", width:"50%"}} className='card'>
            <div className='card-header'>
              <div>Backup from OLT to local PC</div>
            </div>
            <div className="card-body">
            <Select  maxWidth required dataItems={utils.generateCardList(this.props.oltInfo)} title={"Card Selection"} onChange={this.onBackup2LocalCardChange} selectedItem={this.state.backup2LocalCardValue}> </Select>
            <Button onClick={this.handlBackup2Local} title={"Backup"} ></Button>
            </div>
          </div>
          <div style={{margin:"0.625rem",width:"50%"}} className='card'>
            <div className='card-header'>
              <div>Restore from local PC to OLT</div>
            </div>
            <div className="card-body">
            <Select required dataItems={utils.generateCardList(this.props.oltInfo)} title={"Card Selection"} onChange={this.onRestore2OltCardChange} selectedItem={this.state.restore2OltCardValue}> </Select>
            <FileUploadComponent title={"Config File"} onFilesChange={this.onRestoreFilesChange}/>
            <Button onClick={this.handlRestore2Olt} title={"Restore"} ></Button>

            </div>
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
export default  connect(stateToProps, dispatchToProps)(BackupRestoreMaintenPage)