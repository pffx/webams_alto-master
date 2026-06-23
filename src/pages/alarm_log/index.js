import React, { Component } from 'react';
import { ToastContainer } from 'react-toastify';
import App from '@nokia-csf-uxr/ccfk/App';
import {AppContent, AppContentWrapper, AppBody } from '@nokia-csf-uxr/ccfk/App';
//import Tabs from '@nokia-csf-uxr/ccfk/HorizontalNavigation';
import { connect } from 'react-redux'
import TabHeader from '../../components/tabHeader'
import SubTabLink from '../../components/subTabLink'
import AlarmLogRoutes from '../../router/alarm_log'
import GLOBAL from '../../global'
import MENU from '../../global/menu';

class AlarmLogPage extends Component {
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
                MENU.ALARMLOG_TAB_INFO.map((info)=>{
                  return (<SubTabLink key={info.index} tab_id={info.index} parent="alarm_log" selected={this.props.subTabIndex === info.index} > {info.label}</SubTabLink>)
                })
              }
            </ul>
              <div >
              {<AlarmLogRoutes/>}
              </div>
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
  }
}
const dispatchToProps = (dispatch) => {
  return {

  }
}
export default  connect(stateToProps, dispatchToProps)(AlarmLogPage)