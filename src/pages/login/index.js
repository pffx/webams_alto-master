import React, { Component } from 'react';
import i18n from '../../locales/config'
import TextInput from '@nokia-csf-uxr/ccfk/TextInput';
import App from '@nokia-csf-uxr/ccfk/App';
import CircularProgressIndicatorIndeterminate from '@nokia-csf-uxr/ccfk/CircularProgressIndicatorIndeterminate';
import Hyperlink from '@nokia-csf-uxr/ccfk/Hyperlink';
import Typography from '@nokia-csf-uxr/ccfk/Typography';
import InlineFeedbackMessage from '@nokia-csf-uxr/ccfk/InlineFeedbackMessage';
import SignIn, { SignInContent, SignInBackground, SignInLogo, SignInHeader, SignInFooter, SignInBody} from '@nokia-csf-uxr/ccfk/SignIn';
import { SPACING_12, SPACING_16, SPACING_32} from '@nokia-csf-uxr/freeform-design-tokens/tokens/spacing';
import Button from '../../components/widget/button';
import Checkbox from '../../components/widget/checkbox';
import AXIOS from "../../axios/"
import GLOBAL from '../../global'
import {API_Login} from '../../global/API'

import { loginAction, logoutAction } from '../../actions/login'
import { connect } from 'react-redux'
import LanguageSelector from '../../components/language'
import Redirect from '../../router/redirect'

//class Login extends React.Component {
class LoginPage extends Component {
  constructor(props) {
    super(props);

    this.state = {
      username: "admin",
      password: "123",
      feedbackMsg:{
        variant: 'warning',
        text: ''
      },
      remember: false,
      loginResult: "",
    };
  }

  componentDidMount (){

    this.handleLogin()// remove the login page, auto login by default account and password 
  }

  componentWillUnmount() {
  }

  handleRemenber=()=>{
    this.setState({remember:!this.state.remember});
    // console.log("handleRemenber")
  }

  handleLogin =()=>{
    // console.log("handleLogin")
    if(this.state.username === ""){
      let body = "Please input Username first!!";
      //this.showInforModal(body);
      return;
    }

    AXIOS
      .post(API_Login, {
        account: this.state.username,
        password: this.state.password,
      })
      .then((res) => {
        // console.log("login api res = ".res)
        if(res.data.status === 200){
          let userInfo = {
            account: this.state.username,
            userInfo: res.data.userInfo,
            token: res.data.token,
          }

          this.props.loginClick(userInfo)
          this.setState({
            feedbackMsg:{
              variant:'success',
              text:res.data.message
            },
          });
        }else{
          this.props.logoutClick();
          this.setState({
            feedbackMsg:{
              variant:'error',
              text:res.data.message
            },
          });
        }
        
      })
      .catch((err) => {
        this.props.logoutClick()
    });
  }

  renderTitlesandFeedback () {
    return(
      this.state.feedbackMsg.text && (
        <InlineFeedbackMessage variant={this.state.feedbackMsg.variant} style={{ marginBottom: SPACING_32 }}>
          {this.state.feedbackMsg.text}
        </InlineFeedbackMessage>
      )
    );
  }

  renderSignIn(){
    return (
      <>
        {this.renderTitlesandFeedback()}
        <div>
          <TextInput
            style={{ marginTop: "1.125rem" }}
            variant="underlined"
            disabled={false}
            value={this.state.username}
            onChange={(event) => {this.setState({username:event.target.value,});}}
            placeholder="Username or email"
            inputProps={{ autoComplete: 'off' }}
            error={this.state.feedbackMsg && this.state.feedbackMsg.variant === 'error'}
          />
          <TextInput
            style={{ marginTop: SPACING_12 }}
            variant="underlined"
            disabled={false}
            value={this.state.password}
            onChange={(event) => {this.setState({password:event.target.value,});}}
            type="password"
            placeholder="Password"
            inputProps={{ autoComplete: 'off', type: 'password' }}
            error={this.state.feedbackMsg && this.state.feedbackMsg.variant === 'error'}
          />
        </div>
        <Checkbox title={i18n.t('login.remember')} onChange={this.handleRemenber}  checked={this.state.remember}></Checkbox>
        <div>
          {false ? (
            <div style={{ display: 'flex', justifyContent: 'center' }}>
              <CircularProgressIndicatorIndeterminate style={{ width: 36 }} />
            </div>
          ) : (
            <Button onClick={this.handleLogin} fullWidth title={i18n.t('button.login')} ></Button>
          )}
        </div>
        <Hyperlink href="http://www.nokia.com" style={{ marginTop: SPACING_16 }}>
          {i18n.t('login.forget')}
        </Hyperlink>
      </>
    );
  };

  render() {
    return (
      this.props.isLogin
      ? <Redirect to="/home"/>
      : <App>
        <SignIn>
          <SignInContent>
            <SignInHeader>
              <SignInLogo />
            </SignInHeader>
            <SignInBody>
              {/* {this.renderSignIn()} */}
            </SignInBody>
            <SignInFooter>
              <Typography variant="CAPTION">
                <span style={{ marginRight: '0.125rem' }}>© {new Date().getFullYear()} NPI </span> |<LanguageSelector/>
              </Typography>
            </SignInFooter>
          </SignInContent>
          <SignInBackground />
        </SignIn>
    </App>
    );
  }
}

const stateToProps = (state) => {
  return {
    isLogin: state.LoginReducer.isLogin,
  }
}
const dispatchToProps = (dispatch) => {
  return {
    loginClick(data) {
      dispatch(loginAction(data))
    },
    logoutClick () {
      dispatch(logoutAction())
    }
  }
}

export default connect(stateToProps, dispatchToProps)(LoginPage)