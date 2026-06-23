import React, { Component } from 'react';
import { connect } from 'react-redux'
import App,{ AppContent, AppContentWrapper, AppBody } from '@nokia-csf-uxr/ccfk/App';
import DataGrid from '@nokia-csf-uxr/ccfk/DataGrid';
import SearchwChips from '@nokia-csf-uxr/ccfk/SearchwChips';
import Label from '@nokia-csf-uxr/ccfk/Label';
import TextInput, { TextInputLabelContent } from '@nokia-csf-uxr/ccfk/TextInput';
import IconButton from "@nokia-csf-uxr/ccfk/IconButton"
import HorizontalDivider from '@nokia-csf-uxr/ccfk/HorizontalDivider';
import { AllCommunityModules } from '@ag-grid-community/all-modules';
import Typography from '@nokia-csf-uxr/ccfk/Typography';
import TabHeader from '../../components/tabHeader'
import Select from '../../components/widget/select';
import WarningDialog from '../../components/widget/dialog/warning';
import "../../css/index.css"
import utils from './../../global/utils';
import GLOBAL,{TableColorMap,TableIconMap } from "./../../global"
import {API_OltList} from "../../global/API"
import AXIOS from '../../axios';
import { updateOltInfor } from '../../actions/global'

const data=[
  {status:"UP",name:"A",description:"aa",serial:"ALCLFAD9CD88",template:"G240GC", location:"111/111", last:"2022-07-12 07:05:32"},
  {status:"Down",name:"B",description:"bb",serial:"ALCLFAD9CD99",template:"G2425GA",location:"222/111", last:"2022-07-10 07:05:32"},
  {status:"Down",name:"C",description:"cc",serial:"ALCLFAD9CD33",template:"G2426GA",location:"334/111", last:"2022-07-11 06:05:32"},
  {status:"Down",name:"D",description:"dd",serial:"ALCLFAD9CD22",template:"G140WG",location:"111/333", last:"2022-07-11 07:05:32"},
  {status:"Down",name:"E",description:"ee",serial:"ALCLFAD9CD12",template:"G240GC",location:"111/111", last:"2022-07-11 08:05:32"},
  {status:"UP",name:"F",description:"ff",serial:"ALCLFAD9CD45",template:"G2426GB",location:"111/222", last:"2022-07-11 07:05:32"},
  {status:"UP",name:"G",description:"gg",serial:"ALCLFAD9CD64",template:"G240GC",location:"555/111", last:"2022-07-11 07:05:32"},
  {status:"UP",name:"H",description:"hh",serial:"ALCLFAD9CD23",template:"G240GC",location:"111/555", last:"2022-07-15 07:02:32"},
  {status:"UP",name:"I",description:"ii",serial:"ALCLFAD9CD89",template:"G240GC",location:"111/111", last:"2022-07-14 07:03:32"},
]
const gridColumnDefs = isRTL => [
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
  { headerName: 'Name', field: 'name', filter: 'agTextColumnFilter' },
  {
    headerName: 'Description',
    field: 'description',
    filter: 'agTextColumnFilter',
    type: 'editableCell'
  },
  {
    headerName: 'Serial',
    field: 'serial',
    filter: 'agTextColumnFilter',
  },
  { headerName: 'Template', field: 'template', filter: 'agTextColumnFilter' },
  { headerName: 'Location', field: 'location', filter: 'agTextColumnFilter' },
  //{ headerName: 'Location', field: 'location', filter: 'agNumberColumnFilter' },
  { headerName: 'Last Change', field: 'last', filter: 'agTextColumnFilter' },
];
class OntsPage extends Component {
  constructor(props) {
    super(props);
    this.state = {
      onts:[],
      showWarning:false,
      wrongMessage:"",
    };
  }

  componentDidMount (){
    if(utils.isOltSelected(this.props.oltInfo)){
      this.setState({
        onts:data,
      })
    }
  }

  onWarningClose=()=>{
    this.setState({
      showWarning:false,
      wrongMessage:""
    })
  }

  componentWillUnmount() {
  }

  renderOnts(){
    return(
      <>
        <DataGrid
          modules={AllCommunityModules}
          wrapperProps={{
            style: {
              height: '100%',
              width: '100%',
              // ...style
            }
          }}
          rowData={this.state.onts}
          suppressContextMenu
          defaultColDef={{
            menuTabs: ['filterMenuTab']
          }}
          columnDefs={gridColumnDefs(false)}
          onGridReady={(params) => {
          }}
        />
      </>
    );
  }

  render() {
    return (
      <App>
        <TabHeader/>
        <AppBody>
          <AppContentWrapper>
            <AppContent className="row2" style={{background: GLOBAL.COLOR.Background}}>
                <div className="card" style={{height: "98%",width:"100%"}}>
                    <div className='card-header'>
                        <div>ONTs </div>
                    </div>
                    <div className="card-body">
                        {utils.isOltSelected(this.props.oltInfo)?this.renderOnts():utils.renderNullSelectedOlt()}
                    </div>
                </div>
              <WarningDialog open={this.state.showWarning} body={this.state.wrongMessage} onRightClick={this.onWarningClose}></WarningDialog>
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
    updateOltInfor(data) {
      dispatch(updateOltInfor(data))
    }
  }
}
export default  connect(stateToProps, dispatchToProps)(OntsPage)