import React, { Component } from 'react';
import { connect } from 'react-redux'
// import { useNavigate, useLocation, useParams } from "react-router-dom";
import {NavLink} from 'react-router-dom';

import App,{ AppContent, AppContentWrapper, AppBody } from '@nokia-csf-uxr/ccfk/App';
import DataGrid from '@nokia-csf-uxr/ccfk/DataGrid';
// import SearchwChips from '@nokia-csf-uxr/ccfk/SearchwChips';
import Label, {LabelHelpIcon} from '@nokia-csf-uxr/ccfk/Label';
import TextInput, { TextInputLabelContent } from '@nokia-csf-uxr/ccfk/TextInput';
import TextArea, { TextAreaLabelContent } from '@nokia-csf-uxr/ccfk/TextArea';
import IconButton from "@nokia-csf-uxr/ccfk/IconButton"
// import FloatingActionButton, { FloatingActionButtonIcon } from '@nokia-csf-uxr/ccfk/FloatingActionButton';
// import AddIcon from '@nokia-csf-uxr/ccfk/node_modules/@nokia-csf-uxr/ccfk-assets/AddIcon';
import AddIcon from '@nokia-csf-uxr/ccfk-assets/AddIcon';
// import HorizontalDivider from '@nokia-csf-uxr/ccfk/HorizontalDivider';
import { AllCommunityModules } from '@ag-grid-community/all-modules';
// import SimpleCard from '@nokia-csf-uxr/ccfk/Card';
import Card from '@nokia-csf-uxr/ccfk/Card';
import Button, { ButtonText } from '@nokia-csf-uxr/ccfk/Button';
import Tooltip from '@nokia-csf-uxr/ccfk/Tooltip';
import HelpCircleIcon from '@nokia-csf-uxr/ccfk-assets/HelpCircleIcon';
import Dialog, { DialogContent, DialogFooter, DialogTitle } from '@nokia-csf-uxr/ccfk/Dialog';
import Typography from '@nokia-csf-uxr/ccfk/Typography';
import TabHeader from '../../components/tabHeader'
import Select from '../../components/widget/select';
import WarningDialog from '../../components/widget/dialog/warning';
import Signature from "../../components/signature"
import "../../css/index.css"
import i18n from '../../locales/config'
import GLOBAL,{TableColorMap,TableIconMap,OLT_TYPE,OLT_LT_NUM } from "./../../global"
import {API_OltList} from "../../global/API"
import AXIOS from '../../axios';
import { updateOltInfor, clearOltInfor } from '../../actions/global'
import utils from '../../global/utils';
import { featureListInitAction } from '../../actions/feature'

const EditOltWidget = (props) => {
  const {
      onClick,
      type,
  } = props;
  // const params = useParams();
  // const handleClick = value => (action) => {
  //   console.log("EditOltWidget   handleClick   value =  ",value)
  //   // console.log("handleClick   params =  ",params)
  //   // navigate(path,{ replace: true })
  //   // navigate('/system',{ replace: true })
  // };
  const getLabel = ()=>{
    if (type === "ip"){
      return props.info.ip
    }else if (type === "name"){
      return props.info.hostname
    }else if (type === "lt_num"){
      return props.info.lt_num
    }else if (type === "type"){
      return props.info.type
    }else{
      return ""
    }
  }
  return (
      <NavLink 
        onClick={ 
          () => {
             onClick();
            }
        } 
        to={""} 
        style={{ cursor:"pointer",textDecoration: 'none', color:'#0a224d' }}
        >
          {getLabel()}
        </NavLink>
  );
}

class HomePage extends Component {
  constructor(props) {
    super(props);
    this.state = {
      critical:0,
      major:0,
      minor:0,
      warning:0,
      indeterminate:0,
      // version:"22.6(2206.215)",
      olts:[],
      selectOLT:"",
      oltCardExpanded:false,
      showOLTCreatation:false,
      newOLT:
      {
        newOLTIp:"",
        newOLTType:"",
        newOLTCardNum:0,
        newOLTAccount:"tmp",
        newOLTPwd:"tmp",
      },
      showOltModification:false,
      editOlt:"",
      showWarning:false,
      wrongMessage:"",
    };
    // var gridApi;
  }
  gridColumnDefsOLT = [
    {
      headerName: 'Status',
      field: 'status',
      type: 'colorBarColumn',
      // pinned: isRTL ? 'right' : 'left',
      cellRendererParams: {
        valueColorMap: TableColorMap,
        valueIconMap: TableIconMap
      }
    },
    { headerName: 'Controlled',
      field: 'controlled',
      filter: 'agTextColumnFilter'
    },
    {
      headerName: 'IP Address',
      field: 'value',
      filter: 'agTextColumnFilter',
      cellRendererFramework: params => {
        // console.log("params = ",params.data)
        return <div>
          <EditOltWidget type={"ip"} onClick={() => {this.oltModificationShowWithData(params.data); }} info={params.data}></EditOltWidget>
        </div>
      }
      // type: 'editableCell'
    },
    // {
    //   headerName: 'Name',
    //   field: 'hostname',
    //   filter: 'agTextColumnFilter',
    //   cellRendererFramework: params => {
    //   //   console.log("params = ",params.data)
    //     return <div>
    //       <EditOltWidget type={"name"} onClick={() => {this.oltModificationShowWithData(params.data); }} info={params.data}></EditOltWidget>
    //     </div>
    //   }
    // },
    // {
    //   headerName: 'Software',
    //   field: 'software',
    //   filter: 'agTextColumnFilter',
    // },
    {
      headerName: 'Type',
      field: 'type',
      filter: 'agTextColumnFilter',
      cellRendererFramework: params => {
        //   console.log("params = ",params.data)
          return <div>
            <EditOltWidget type={"type"} onClick={() => {this.oltModificationShowWithData(params.data); }} info={params.data}></EditOltWidget>
          </div>
        }
    },
    {
      headerName: 'Number of LT',
      field: 'lt_num',
      filter: 'agTextColumnFilter',
      cellRendererFramework: params => {
        //   console.log("params = ",params.data)
          return <div>
            <EditOltWidget type={"lt_num"} onClick={() => {this.oltModificationShowWithData(params.data); }} info={params.data}></EditOltWidget>
          </div>
        }
    },
  ];

  componentDidMount (){
    this.getOLTs()
    // this.setState({
    //   onts:data,
    // })
  }

  onWarningClose=()=>{
    this.setState({
      showWarning:false,
      wrongMessage:""
    })
  }

  handleOltDelete = ()=>{
    console.log("Todo ....")
  }

  editOlt = ()=>{
    if(this.state.newOLT.newOLTType === ""){
      this.setState({
        showWarning:true,
        wrongMessage:"Please indicate the OLT type!"
      })
      return
    }
    if(utils.isNum(this.state.newOLT.newOLTCardNum)){
      // if(this.state.newOLT.newOLTType.startsWith("DF") && this.state.newOLT.newOLTCardNum > 1){
      //   this.setState({
      //     showWarning:true,
      //     wrongMessage:"DF16 has only one LT card!"
      //   })
      //   return
      // }
      if(this.state.newOLT.newOLTType.startsWith("MF2") && this.state.newOLT.newOLTCardNum > 2){
        this.setState({
          showWarning:true,
          wrongMessage:"MF2 only has a maximum of 2 LTs!"
        })
        return
      }
      if(this.state.newOLT.newOLTType.startsWith("MF14") && this.state.newOLT.newOLTCardNum > 14){
        this.setState({
          showWarning:true,
          wrongMessage:"MF14 only has a maximum of14 LTs!"
        })
        return
      }
    }else{
      if(this.state.newOLT.newOLTType.startsWith("MF14") && this.state.newOLT.newOLTCardNum > 14){
        this.setState({
          showWarning:true,
          wrongMessage:"Please input a valid number!"
        })
        return
      }
    }
    AXIOS
    .put(API_OltList,{
      IP: this.state.editOlt.ip,
      Username: this.state.newOLT.newOLTAccount,
      Password: this.state.newOLT.newOLTPwd,
      OltType:this.state.newOLT.newOLTType,
      // OltLtNum:this.state.newOLT.newOLTCardNum,
    })
    .then((res)=>{
      let olt = {};
      let olts = this.state.olts
      this.state.olts.map((item, index)=>{
        if(item.value === this.state.editOlt.ip){
          olts[index].type = this.state.newOLT.newOLTType
          olts[index].lt_num = this.state.newOLT.newOLTCardNum ===undefined? 0 :this.state.newOLT.newOLTCardNum 
        }
      })
      this.setState({
        olts:olts,
        showOltModification:!this.state.showOltModification
      })
      if(this.state.editOlt.ip === this.props.oltInfor.ip){
        let oltInfo = {
          ip:this.state.editOlt.ip,
          hostname:"",
          status:"",
          software:"",
          type:this.state.newOLT.newOLTType,
          ltNum:this.state.newOLT.newOLTCardNum,
          controlled:true,
        }
        this.props.updateOltInfor(oltInfo)
      }
    })
    .catch((err) => {
    });

    
  }

  addOlt = ()=>{
    // if(this.state.newOLT.newOLTIp === "" || this.state.newOLT.newOLTAccount === "" || this.state.newOLT.newOLTPwd === ""){
    if(this.state.newOLT.newOLTIp === ""){
      this.setState({
        showWarning:true,
        // wrongMessage:"Please input the completed information!"
        wrongMessage:"Please input the IP address!"
      })
      return
    }
    if(this.state.newOLT.newOLTType === ""){
      this.setState({
        showWarning:true,
        wrongMessage:"Please indicate the OLT type!"
      })
      return
    }
    if(utils.isNum(this.state.newOLT.newOLTCardNum)){
      // if(this.state.newOLT.newOLTType.startsWith("DF") && this.state.newOLT.newOLTCardNum > 1){
      //   this.setState({
      //     showWarning:true,
      //     wrongMessage:"DF16 has only one LT card!"
      //   })
      //   return
      // }
      if(this.state.newOLT.newOLTType.startsWith("MF2") && this.state.newOLT.newOLTCardNum > 2){
        this.setState({
          showWarning:true,
          wrongMessage:"MF2 only has a maximum of 2 LTs!"
        })
        return
      }
      if(this.state.newOLT.newOLTType.startsWith("MF14") && this.state.newOLT.newOLTCardNum > 14){
        this.setState({
          showWarning:true,
          wrongMessage:"MF14 only has a maximum of 14 LTs!"
        })
        return
      }
    }else{
      if(this.state.newOLT.newOLTType.startsWith("MF14") && this.state.newOLT.newOLTCardNum > 14){
        this.setState({
          showWarning:true,
          wrongMessage:"Please input a valid number!"
        })
        return
      }
    }
    AXIOS
    .put(API_OltList,{
      IP: this.state.newOLT.newOLTIp,
      Username: this.state.newOLT.newOLTAccount,
      Password: this.state.newOLT.newOLTPwd,
      OltType:this.state.newOLT.newOLTType,
      OltLtNum:this.state.newOLT.newOLTCardNum,
    })
    .then((res)=>{
      // console.log("add new olt")
      let olt = {};
      //read the ont info from server
      let olts = this.state.olts
      olt.value = this.state.newOLT.newOLTIp
      olt.ip = this.state.newOLT.newOLTIp
      olt.label = this.state.newOLT.newOLTIp
      // olt.account = this.state.newOLT.newOLTAccount
      // olt.pwd = this.state.newOLT.newOLTPwd

      olts.push(olt)
      this.setState({
        olts:olts,
        showOLTCreatation:!this.state.showOLTCreatation
      })
      // this.gridApi.setRowData(olts)

    })
    .catch((err) => {
    });
  }

  getSelectedOLT(){
    return 
  }

  getOLTs(){
    AXIOS
      .get(API_OltList)
      .then((res) => {
        if(res.data.status == 200){
          let featureList = res.data.featurelist
          this.props.featureListInit(featureList)
          let list = []
          let listForInitial = []
          let select
          let hasInital = false
          let initialOlt = {}
          res.data.olt_list.map((item, index)=>{
            let olt = {};
            // console.log("API_OltList   item = ",item)
            olt.ip = item.IP
            olt.value = item.IP
            olt.label = item.IP
            olt.account = item.Username
            olt.pwd = item.Password
            olt.status = item.Status === "Connected"?"UP":"Down"
            olt.hostname = item.HostName
            olt.software = item.Software
            olt.lt_num = item.OltLtNum
            olt.type=item.OltType
            if(olt.ip === this.props.oltInfor.ip && olt.hostname === this.props.oltInfor.hostname ){
                olt.controlled = "Yes"
                select = olt.ip
            }else{
              olt.controlled = "No"
            }
            olt.ltCardStatus = utils.generateLTCardPlanned(item)
            // console.log("API_OltList   olt = ",olt)
            list.push(olt)
            if(item.IP === "192.168.1.1"){
              hasInital = true
              initialOlt = olt
              initialOlt.controlled = "Yes"
              listForInitial.push(initialOlt)
            }else{
              listForInitial.push(olt)
            }
          })
          // this.setState({
          //   olts:list,
          //   selectOLT:select,
          // })
          console.log("getOLTs    select =  ",select)
          if(select === undefined || select === ""){
            if(hasInital){
              this.props.updateOltInfor(initialOlt)
              this.setState({
                olts:listForInitial,
                selectOLT:initialOlt.ip,
              })
            }else{
              this.props.clearOltInfor()
              this.setState({
                olts:list,
                selectOLT:select,
              })
            }
          }else{
            this.setState({
              olts:list,
              selectOLT:select,
            })
          }
        }
      })
      .catch((err) => {
    });
  }

  componentWillUnmount() {
  }

  oltCreatationShow=()=>{
    this.setState({
      showOLTCreatation:!this.state.showOLTCreatation
    })
  }

  oltModificationShowWithData=(data)=>{
    // console.log("oltModificationShowWithData data = ",data)
    this.setState({
      showOltModification:!this.state.showOltModification,
      editOlt:data
    })
  }

  oltModificationShow=()=>{
    if(this.state.showOltModification){
      this.setState({
        showOltModification:!this.state.showOltModification,
        editOlt:""
      })
    }else{
      this.setState({
        showOltModification:!this.state.showOltModification
      })
    }
  }

  handleOltModification=()=>{
    this.setState({
      showOltModification:false,
      editOlt:""
    })
  }

  onOltChange=(item)=>{
    // console.log("onOLTchange   item: ",item)
    if(item.value === this.state.selectOLT){
      return
    }
    let newOlts = this.state.olts
    let mark
    this.state.olts.map((o, index)=>{
      if(o.controlled === "Yes"){
        mark = index
      }
      if(o.ip === item.ip && o.hostname === item.hostname){
        newOlts[index].controlled = "Yes"
      }
    })
    if(mark != undefined){
      newOlts[mark].controlled = "No"
    }
    // console.log("onOltChange  newOlts= ",newOlts)

    this.setState({
      selectOLT:item.value,
      olts:newOlts
    })
    console.log("onOltChange  this.gridApi= ",this.gridApi)
    // this.gridApi.setRowData(newOlts)
    let oltInfo = {
      ip:item.ip,
      hostname:item.hostname,
      status:item.status,
      software:item.software,
      type:item.type,
      ltNum:item.lt_num,
      ltCardStatus:item.ltCardStatus,
      controlled:true,
    }
    this.props.updateOltInfor(oltInfo)
  }

  onOltTypeChange = (item)=>{
    let obj = this.state.newOLT
    obj.newOLTType = item.value
    if(item.value ==="lt"){
      obj.newOLTCardNum = 1
    }
    this.setState({
      newOLT:obj,
    })
  }
  onOltLTNumChange = (event)=>{
    let obj = this.state.newOLT
    obj.newOLTCardNum = event.target.value
    this.setState({
      newOLT:obj,
    })
  }

  renderOltEditModal(){
    return(
      <Dialog
        isOpen={this.state.showOltModification}
      >
          <DialogTitle  title={"OLT Modification"} />
          <DialogContent
          isTopDividerVisible
          isBottomDividerVisible
          aria-hidden="true"
        >
        <div >
          <Label readOnly verticalLayout maxWidth>
            <TextAreaLabelContent>{"OLT:"}</TextAreaLabelContent>
            <TextArea
              readOnly
              // cols={cols}
              rows={1}
              value={this.state.editOlt==="" ? "" : this.state.editOlt.ip/* +"("+this.state.hostname+")"*/}
            />
          </Label>
          <div >
            <Select  dataItems={OLT_TYPE} title={"Type"} onChange={this.onOltTypeChange} selectedItem={this.state.newOLT.newOLTType}> </Select>
          </div>
          {/* <div >
            <Label maxWidth>
              <TextInputLabelContent >{"LT Number"}</TextInputLabelContent>
            </Label>
            <TextInput
              variant="outlined"
              maxWidth
              type="text"
              disabled={false}
              value={this.state.newOLT.newOLTCardNum}
              onChange={this.onOltLTNumChange}
              inputProps={{ autoComplete: 'on' }}
            />
          </div> */}
        </div>
        </DialogContent>
          <DialogFooter>
            <Button onClick={this.oltModificationShow}>{i18n.t('button.cancel')}</Button>
            <Button autoFocus variant="call-to-action" onClick={this.editOlt}>{i18n.t('button.edit')}</Button>
            <Button onClick={this.handleOltDelete}>{i18n.t('button.delete')}</Button>
          </DialogFooter>
      </Dialog>
    )

  }

  renderOltAdd(){
    return(
      <Dialog
      isOpen={this.state.showOLTCreatation}
    >
        <DialogTitle  title={"Add New OLT"} />
        <DialogContent
          isTopDividerVisible
          isBottomDividerVisible
          aria-hidden="true"
        >
        <div >
          <div >
            <Select dataItems={OLT_TYPE} title={"Type"}  onChange={this.onOltTypeChange} selectedItem={this.state.newOLT.newOLTType}> </Select>
          </div>
          {/* <div >
            <Select dataItems={OLT_LT_NUM} title={"LT Number"}  onChange={this.onOltLTNumChange} selectedItem={this.state.newOLT.newOLTCardNum}> </Select>
          </div> */}
          {/* <div >
            <Label maxWidth>
              <TextInputLabelContent >{"LT Number"}</TextInputLabelContent>
            </Label>
            <TextInput
              variant="outlined"
              maxWidth
              type="number"
              disabled={false}
              value={this.state.newOLT.newOLTCardNum}
              onChange={this.onOltLTNumChange}
              inputProps={{ autoComplete: 'on' }}
            />
          </div> */}
          <div >
            <Label maxWidth>
              <TextInputLabelContent >{"OLT IP"}</TextInputLabelContent>
            </Label>
            <TextInput
              variant="outlined"
              maxWidth
              type="text"
              disabled={false}
              value={this.state.newOLT.newOLTIp}
              onChange={(event) => {let obj = this.state.newOLT; obj.newOLTIp = event.target.value;this.setState({newOLT:obj,});}}
              inputProps={{ autoComplete: 'on' }}
            />
          </div>
          {/* <div >
            <Label maxWidth>
              <TextInputLabelContent >{"Account"}</TextInputLabelContent>
            </Label>
            <TextInput
              variant="outlined"
              maxWidth
              type="text"
              disabled={false}
              value={this.state.newOLT.newOLTAccount}
              onChange={(event) => {let obj = this.state.newOLT; obj.newOLTAccount = event.target.value;this.setState({newOLT:obj,});}}
              inputProps={{ autoComplete: 'off' }}
            />
          </div>
          <div >
            <Label maxWidth>
              <TextInputLabelContent >{"Password"}</TextInputLabelContent>
            </Label>
            <TextInput
              variant="outlined"
              maxWidth
              type="text"
              disabled={false}
              value={this.state.newOLT.newOLTPwd}
              onChange={(event) => {let obj = this.state.newOLT; obj.newOLTPwd = event.target.value;this.setState({newOLT:obj,});}}
              inputProps={{ autoComplete: 'off' }}
            />
          </div> */}
        </div>
        </DialogContent>
        <DialogFooter>
          <Button onClick={this.oltCreatationShow}>{i18n.t('button.cancel')}</Button>
          <Button autoFocus variant="call-to-action" onClick={this.addOlt}>{i18n.t('button.create')}</Button>
        </DialogFooter>
    </Dialog>
    )

  }

  renderTeamSignature(){
    return(
      <div style={{marginLeft: 10,marginBottom: 10}}>
        <Card
          style={{padding: "1rem",width:"100%", height:"3rem"}}
          >
          <Typography variant="TITLE_16">Powered by APAC NPI</Typography>

        </Card>
      </div>
    )
  }

  renderOltSummary(){
    return(
      <div style={{marginLeft: 10,marginBottom: 10}}>
        <Card
          style={{padding: "1rem",width:"100%", }}
          // expandHeight={100}
          // onExpand={(e) => {
          //   this.setState({
          //     oltCardExpanded:!this.state.oltCardExpanded
          //   })
          // }}
          // expanded={this.state.oltCardExpanded}
          // expandedContent={
          //   <Select dataItems={this.state.olts} title={"OLT Selection"} onChange={this.onOltChange} selectedItem={this.state.selectOLT}> </Select>
          // }
          >
          <div className='row'>
            <Typography variant="TITLE_16">OLT Selection</Typography>
            <Tooltip placement="top" trigger={["hover", "focus" ]} fallbackPlacements={["bottom", "right" ]} modifiers={{offset: {offset: "[0, 8]"}}}
            tooltip="One OLT must be selected before configuration. ">
            <LabelHelpIcon>
                <HelpCircleIcon />
            </LabelHelpIcon>
          </Tooltip>
          </div>
          <div>
            <Select dataItems={this.state.olts} title={"An OLT must be selected before operation"} onChange={this.onOltChange} selectedItem={this.state.selectOLT}> </Select>
          </div>
        </Card>
      </div>
    )
  }

  renderAlarmSummary(){
    return(<></>
    //   <div style={{margin: 10}}>
    //     <SimpleCard style={{padding: "1rem",}} variant="stack">
    //       <div>
    //         <Typography variant="TITLE_16">Alarm Summary</Typography>
    //       </div>
    //       <div>
    //         <table style={{ width: "100%"}}>
    //           <tbody id="alarm" >
    //             <tr>
    //               <th bgcolor={GLOBAL.COLOR.Critical}>Critical</th>
    //               <td bgcolor={GLOBAL.COLOR.Critical}>{this.state.critical}</td>
    //             </tr>
    //             <tr>
    //               <th bgcolor={GLOBAL.COLOR.Major}>Major</th>
    //               <td bgcolor={GLOBAL.COLOR.Major}>{this.state.major}</td>
    //             </tr>
    //             <tr>
    //               <th bgcolor={GLOBAL.COLOR.Minor}>Minor</th>
    //               <td bgcolor={GLOBAL.COLOR.Minor}>{this.state.minor}</td>
    //             </tr>
    //             <tr>
    //               <th bgcolor={GLOBAL.COLOR.Warning}>Warning</th>
    //               <td bgcolor={GLOBAL.COLOR.Warning}>{this.state.warning}</td>
    //             </tr>
    //             <tr>
    //               <th bgcolor={GLOBAL.COLOR.Indeterminate}>Indeterminate</th>
    //               <td bgcolor={GLOBAL.COLOR.Indeterminate}>{this.state.indeterminate}</td>
    //             </tr>
    //           </tbody>
    //         </table>
    //       </div>
    //     </SimpleCard>
    //   </div>
    )
  }
  renderOlts(){
    return(
        <DataGrid
          modules={AllCommunityModules}
          wrapperProps={{
            style: {
              height: '100%',
              width: '100%',
              // ...style
            }
          }}
          rowData={this.state.olts}
          suppressContextMenu
          defaultColDef={{
            menuTabs: ['filterMenuTab']
          }}
          columnDefs={this.gridColumnDefsOLT}
          onGridReady={(params) => {
            // console.log("DataGrid  params= ",params)
            // this.gridApi = params.api;
          }}
        />
    );
  }

  render() {
    return (
      <App>
        <TabHeader/>
        <AppBody>
          <AppContentWrapper>
            <AppContent style={{background: GLOBAL.COLOR.Background}}>
              <div style={{height: "92%"}}  className="row2">
                {/* <div style={{display:"flex",flexDirection: "column", justifyContent:"space-between",height: "100%",width: "22%",paddingRight: 10,}}> */}
                <div style={{height: "100%",width: "22%",paddingRight: 10,}}>
                  {this.renderOltSummary()}
                  {this.renderAlarmSummary()}
                  {/* {this.renderTeamSignature()} */}
                </div>
                <div style={{height: "99%", width: "74%"}}>
                  <div className="card" style={{height: "100%"}}>
                      <div className='card-header row3'>
                        <div>OLTs </div>
                        <IconButton aria-label="settings" onClick={this.oltCreatationShow}>
                            <AddIcon size={'2rem'}/>
                        </IconButton>
                      </div>
                      <div className="card-body">
                        {this.renderOlts()}
                      </div>
                  </div>
                </div>
                {this.renderOltAdd()}
                {this.renderOltEditModal()}
              </div>
              <Signature></Signature>
              <WarningDialog open={this.state.showWarning} body={this.state.wrongMessage} onRightClick={this.onWarningClose}></WarningDialog>
            </AppContent>
          </AppContentWrapper>
        </AppBody>
      </App>
    );
  }
}

const stateToProps = (state) => {
  // console.log("home page  state = ",state)
  return {
    oltInfor: state.GlobalReducer.oltInfor,
  }
}
const dispatchToProps = (dispatch) => {
  return {
    updateOltInfor(data) {
      dispatch(updateOltInfor(data))
    },
    featureListInit(data) {
      dispatch(featureListInitAction(data))
    },
    clearOltInfor(){
      dispatch(clearOltInfor())
    }
  }
}
export default  connect(stateToProps, dispatchToProps)(HomePage)