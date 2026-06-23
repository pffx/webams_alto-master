import React, { Component } from 'react';
import App from '@nokia-csf-uxr/ccfk/App';
import {AppContent, AppContentWrapper, AppBody } from '@nokia-csf-uxr/ccfk/App';
import TabHeader from '../../components/tabHeader'
import ServiceRoutes from '../../router/service'
import GLOBAL from '../../global'

class ServicePage extends Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <>
      {<ServiceRoutes/>}
      </>
      // <App>
      //   <TabHeader/>
      //   <AppBody style={{background: GLOBAL.COLOR.Background}}>
      //     <AppContentWrapper>
      //       <AppContent >
      //         {<ServiceRoutes/>}
      //       </AppContent>
      //     </AppContentWrapper>
      //   </AppBody>
      // </App>
    );
  }
}

export default ServicePage