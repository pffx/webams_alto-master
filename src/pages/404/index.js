import React, { Component } from 'react';
import App,{ AppContent, AppContentWrapper, AppBody } from '@nokia-csf-uxr/ccfk/App';
import TabHeader from '../../components/tabHeader'

class NotFoundPage extends Component {
  constructor(props) {
    super(props);
    this.state = {
    };
  }

  render() {
    return (
      <App>
        <TabHeader/>
        <AppBody>
          <AppContentWrapper>
            <AppContent >
                <h1>Page Not Found!!</h1>
                <h2>Please back and retry!</h2>
            </AppContent>
          </AppContentWrapper>
        </AppBody>
    </App>
    );
  }
}

export default  NotFoundPage