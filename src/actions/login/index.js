import {API_CheckToken} from "../../global/API"
import AXIOS from "../../axios"

const loginAction = data => ({type:"login",data:data})
const logoutAction  = () => ({type:"logout"})
const updateT  = (data) => ({type:"updateT",data:data})
const CheckT  = () => {
    AXIOS
      .get(API_CheckToken)
      .then((res) => {
        //console.log("update the login status")
      });
    return {type:"CheckT"}
}

export {
    loginAction ,
    logoutAction,
    updateT,
    CheckT
}