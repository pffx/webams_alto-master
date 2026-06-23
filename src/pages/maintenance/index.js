import React, { Component } from 'react';
import { ToastContainer } from 'react-toastify';
import App from '@nokia-csf-uxr/ccfk/App';
import {AppContent, AppContentWrapper, AppBody } from '@nokia-csf-uxr/ccfk/App';
//import Tabs from '@nokia-csf-uxr/ccfk/HorizontalNavigation';
import { connect } from 'react-redux'
import TabHeader from '../../components/tabHeader'
import SubTabLink from '../../components/subTabLink'
import MaintenanceRoutes from '../../router/maintenance'
// import Signature from '../../components/signature'
import GLOBAL from '../../global'
import MENU from '../../global/menu';

class MaintenancePage extends Component {
  constructor(props) {
    super(props);
    this.state = {
    };
    if(this.props.isLogin){

    }
  }

  componentDidMount (){
    //console.log(this.props.isLogin)
  }

  componentWillUnmount() {
  }

  renderIndicator(){
    return(
      <h4>
        123123
      </h4>

    )
  }

  render() {
    return (
      <App>
        <TabHeader/>
        <AppBody style={{background: GLOBAL.COLOR.Background}}>
          <AppContentWrapper>
            <AppContent >
            <ul className="nav nav-pills">
              {
                MENU.MAINTENANCE_TAB_INFO(this.props.featurelist).map((info)=>{
                  return (<SubTabLink key={info.index} tab_id={info.index} parent="maintenance" selected={this.props.subTabIndex === info.index} > {info.label}</SubTabLink>)
                })
              }
            </ul>
              <div >
              {<MaintenanceRoutes/>}
              </div>
              {/* <Signature></Signature> */}
            </AppContent>
          </AppContentWrapper>
        </AppBody>
        <ToastContainer/>
      </App>
    );
  }
}

const stateToProps = (state) => {
  return {
    subTabIndex: state.GlobalReducer.subTabIndex,
    featurelist: state.FeatureListReducer.feature,
  }
}
const dispatchToProps = (dispatch) => {
  return {

  }
}
export default  connect(stateToProps, dispatchToProps)(MaintenancePage)