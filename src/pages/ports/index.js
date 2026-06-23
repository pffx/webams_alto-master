import React, { Component } from 'react';
import { connect } from 'react-redux'
import App,{ AppContent, AppContentWrapper, AppBody } from '@nokia-csf-uxr/ccfk/App';

import TabHeader from '../../components/tabHeader'
import GLOBAL,{TableColorMap,TableIconMap } from "./../../global"

class PortsPage extends Component {
    constructor(props) {
      super(props);
      this.state = {
        showWarning:false,
        wrongMessage:"",
      };
    }
  
    componentDidMount (){
    }
  
    onWarningClose=()=>{
      this.setState({
        showWarning:false,
        wrongMessage:""
      })
    }
  
    componentWillUnmount() {
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
                          <div>Ports </div>
                      </div>
                      <div className="card-body">
                          12345testst tsts
                      </div>
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
    }
  }
  const dispatchToProps = (dispatch) => {
    return {
    }
  }
  export default  connect(stateToProps, dispatchToProps)(PortsPage)