import React, { Component } from 'react';
import { connect } from 'react-redux'
import { useNavigate, useLocation, useParams } from "react-router-dom";
import {NavLink} from 'react-router-dom';
import i18n from '../../../locales/config'
import App from '@nokia-csf-uxr/ccfk/App';
import {AppContent, AppContentWrapper, AppBody } from '@nokia-csf-uxr/ccfk/App';
import TabHeader from '../../../components/tabHeader'
import DataGrid from '@nokia-csf-uxr/ccfk/DataGrid';
import { AllCommunityModules } from '@ag-grid-community/all-modules';
import Button, { ButtonText } from '@nokia-csf-uxr/ccfk/Button';
import Dialog, { DialogContent, DialogFooter, DialogTitle } from '@nokia-csf-uxr/ccfk/Dialog';
import Label from '@nokia-csf-uxr/ccfk/Label';
import TextInput, { TextInputLabelContent } from '@nokia-csf-uxr/ccfk/TextInput';
import Typography from '@nokia-csf-uxr/ccfk/Typography';

import Select from '../../../components/widget/select';

import AXIOS from "../../../axios/"
import {API_ProvisionService} from '../../../global/API'
import utils from '../../../global/utils';
import GLOBAL from "../../../global"

const EditServiceWidget = (props) => {
  const navigate = useNavigate();
  const params = useParams();
  const path = "/service/edit"
  const handleClick = value => (action) => {
    // console.log("handleClick   value =  ",value)
    // console.log("handleClick   params =  ",params)
    // navigate(path,{ replace: true })
    // navigate('/system',{ replace: true })
  };
  return (
      <NavLink onClick={handleClick(props.info)} to={path} style={{ cursor:"pointer",textDecoration: 'none', color:'#0a224d' }} >{props.info.name}</NavLink>
  );
}

const gridColumnDefs = (isRTL) => [
  {
    headerName: i18n.t('table.name'),
    field: 'name',
    pinned: isRTL ? 'right' : 'left',
    filter: 'agTextColumnFilter',
    cellRendererFramework: params => {
    //   console.log("params = ",params.data)
    //   // return "<a style='text-decoration: none; color:#0a224d' href='/service/edit'> "+params.value+"</a>"; 
      return <div>
        <EditServiceWidget info={params.data}></EditServiceWidget>
      </div>
    }
  },
  { headerName: i18n.t('table.version'), field: 'version', filter: 'agTextColumnFilter' },
  {
    headerName: i18n.t('table.status'),
    field: 'status',
    filter: 'agTextColumnFilter',
  },
  {
    headerName: i18n.t('table.description'),
    field: 'description',
    filter: 'agTextColumnFilter',
  },
];

const ServiceType = [
  {label: 'service', isHeader: true },
  {value:"CCTVv1",label:"CCTVv1"},
  {value:"DIGISIGNv1",label:"DIGISIGNv1"},
  {value:"HSIv1",label:"HSIv1"},
  {value:"HSIv1",label:"HSIv1"},
  {value:"IPTVv1",label:"IPTVv1"},
  {value:"PUBLICANNOUNCEv1",label:"PUBLICANNOUNCEv1"},
  {value:"SECURITYv1",label:"SECURITYv1"},
  {value:"VOIPv1",label:"VOIPv1"},
  {value:"WIFIAPv1",label:"WIFIAPv1"},
]

class ServiceMainPage extends Component {
  constructor(props) {
    super(props);
    this.state = {
      gridApi:"",
      gridColumnApi:"",
      showCreatation:false,
      newName:"",
      newDes:"",
      newServiceType:"",
      services: [],
    };
  }

  componentDidMount (){
    if(utils.isOltSelected(this.props.oltInfo)){
      this.getProvisionService()
    }
  }

  componentWillUnmount() {
  }

  getProvisionService(){
    AXIOS
    .get(API_ProvisionService)
    .then((res) => {
      if(res.data.status == 200){
        let list = []
        res.data.service_list.map((item, index)=>{
          let service = {};
          service.name = item.Name
          service.version = item.Version
          service.status = item.Status=== "active"? "Active":"Inactive"
          service.description = item.Description
          service.id = item.ID
          list.push(service)
        })
        this.setState({
          services: list,
        })
      }
    })
    .catch((err) => {
    });
  }

  onServiceTypeChange = (item)=>{
    this.setState({
      newServiceType:item.value
    })
  }

  serviceCreatationShow=()=>{
    this.setState({
      showCreatation:!this.state.showCreatation
    })
  }

  createNewService=()=>{
    this.setState({
      showCreatation:false
    })
  }

  renderServiceCreatationModal(){
    return(
      <Dialog
        isOpen={this.state.showCreatation}
      >
          <DialogTitle  title={"Create Service"} />
          <DialogContent
            isTopDividerVisible
            isBottomDividerVisible
            aria-hidden="true"
          >
            <div >
              <Label maxWidth>
                <TextInputLabelContent required>{"Name"}</TextInputLabelContent>
              </Label>
              <TextInput
                variant="outlined"
                maxWidth
                type="text"
                disabled={false}
                value={this.state.newName}
                onChange={(event) => {this.setState({newName:event.target.value,});}}
                inputProps={{ autoComplete: 'on' }}
              />
            </div>
            <div >
              <Label maxWidth>
                <TextInputLabelContent >{"Description"}</TextInputLabelContent>
              </Label>
              <TextInput
                variant="outlined"
                maxWidth
                type="text"
                disabled={false}
                value={this.state.newDes}
                onChange={(event) => {this.setState({newDes:event.target.value,});}}
                inputProps={{ autoComplete: 'off' }}
              />
            </div>
            <Select required dataItems={ServiceType} title={"Service Template"} onChange={this.onServiceTypeChange} > </Select>
          </DialogContent>
          <DialogFooter>
            <Button onClick={this.serviceCreatationShow}>{i18n.t('button.cancel')}</Button>
            <Button autoFocus variant="call-to-action" onClick={this.createNewService}>{i18n.t('button.create')}</Button>
          </DialogFooter>
      </Dialog>
    )

  }

  renderServiceList(){
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
        rowData={this.state.services}
        suppressContextMenu
        defaultColDef={{
          menuTabs: ['filterMenuTab']
        }}
        columnDefs={gridColumnDefs(false)}
        onGridReady={(params) => {
          this.setState({
            gridColumnApi:params.columnApi,
            gridApi:params.api
          })
        }}
      />
  );
  }

  renderServiceCard(){
    return(
      <div className="card" style={{ width: "70%",}}>
        <div className='card-header' style={{justifyContent: "flex-start", flexDirection: "row", display: "flex", alignItems: "center",justifyContent:"space-between"}}>
          <div>Service Instances</div>
          {
            utils.isOltSelected(this.props.oltInfo)
            ? <Button variant="outlined" onClick={this.serviceCreatationShow}>
              <ButtonText id="button-text-id">{i18n.t('button.create')}</ButtonText>
            </Button>
            :<></>
          }
        </div>
        <div className="card-body" style={{height:"100%"}}>
          {utils.isOltSelected(this.props.oltInfo)?this.renderServiceList():utils.renderNullSelectedOlt()}
        </div>
      </div>
    )
  }

  render() {
    return (
      // <>
      // <div  style={{background: GLOBAL.COLOR.Background,justifyContent:"center",display:"flex"}}>
      //   {this.renderServiceCard()}
      //   {this.renderServiceCreatationModal()}
      // </div>
      // </>
      <App>
        <TabHeader/>
        <AppBody style={{background: GLOBAL.COLOR.Background,}}>
          <AppContentWrapper>
            <AppContent style={{justifyContent:"center",display:"flex"}}>
              {this.renderServiceCard()}
              {this.renderServiceCreatationModal()}
            </AppContent>
          </AppContentWrapper>
        </AppBody>
      </App>
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
    // updateOltInfor(data) {
    //   dispatch(updateOltInfor(data))
    // }
  }
}
export default  connect(stateToProps, dispatchToProps)(ServiceMainPage)