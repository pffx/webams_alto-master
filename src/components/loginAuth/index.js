import React from 'react';
//import { useNavigate, useLocation } from "react-router-dom";
import {useSelector,useDispatch} from 'react-redux'
import Redirect from '../../router/redirect'
// import AXIOS from "../../axios";
// import { logoutAction } from '../../actions/login'
// import { clearAllTabIndex } from '../../actions/global'
// import GLOBAL from "../../global";

function LoginAuth({ children }) {
    //const dispatch = useDispatch()
    // const navigate = useNavigate();
    // const location = useLocation();
    // console.log("LoginAuth navigate:", navigate)
    // console.log("LoginAuth location:", location)

    const {isLogin} = useSelector((state) => state.LoginReducer)
    //const [login,updateLogin] = React.useState(isLogin)
    // const handleLogout = ()=>{
    //   dispatch(logoutAction())
    //   dispatch(clearAllTabIndex())
    // }
    // AXIOS
    //   .get(GLOBAL.API_ALTO + GLOBAL.API_GROUP.AUTH + "/token")
    //   .then((res) => {
    //     console.log("update the login status")
    //   })
    //   .catch((err) => {
    //     handleLogout()
    // });
  
    return isLogin
    ? ( children ) 
    : <Redirect to="/login"/>
  }
export default LoginAuth