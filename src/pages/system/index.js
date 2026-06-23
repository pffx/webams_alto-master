import React, { Component } from 'react';
import { connect } from 'react-redux'

import App from '@nokia-csf-uxr/ccfk/App';
import {AppContent, AppContentWrapper, AppBody } from '@nokia-csf-uxr/ccfk/App';

import TabHeader from '../../components/tabHeader'
import SubTabLink from '../../components/subTabLink'
import SystemRoutes from '../../router/system'
import GLOBAL from '../../global'
import MENU from '../../global/menu';

class SystemPage extends Component {
  constructor(props) {
    super(props);
    this.state = {
    };
    if(this.props.isLogin){

    }
  }

  componentDidMount (){
  }

  componentWillUnmount() {
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
                MENU.SYSTEM_TAB_INFO.map((info)=>{
                  return (<SubTabLink key={info.index} tab_id={info.index} parent="system" selected={this.props.subTabIndex === info.index} > {info.label}</SubTabLink>)
                })
              }
            </ul>
              <div >
              {<SystemRoutes/>}
              </div>
            </AppContent>
          </AppContentWrapper>
        </AppBody>
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
export default  connect(stateToProps, dispatchToProps)(SystemPage)