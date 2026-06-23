import React, { Component } from 'react';
import i18n from '../../../locales/config'
import App from '@nokia-csf-uxr/ccfk/App';
import {AppContent, AppContentWrapper, AppBody } from '@nokia-csf-uxr/ccfk/App';

import TabHeader from "../../../components/tabHeader"

import GLOBAL from "../../../global"

class ServiceEditPage extends Component {
  constructor(props) {
    super(props);
    console.log("ServiceEditPage   props = ",props)
    this.state = {
    };
  }

  componentDidMount (){
  }

  componentWillUnmount() {
  }


  render() {
    return (
      <App>
      <TabHeader/>
      <AppBody style={{background: GLOBAL.COLOR.Background,}}>
        <AppContentWrapper>
          <AppContent >
           <h1>test test </h1>
          </AppContent>
        </AppContentWrapper>
      </AppBody>
    </App>
    );
  }
}

export default ServiceEditPage