import axios from "axios";
import qs from "qs";
import {useDispatch} from 'react-redux'
import store from '../store'
import { logoutAction, updateT } from '../actions/login'
import { clearAllTabIndex } from '../actions/global'
import {API_URI} from "../global/API"
import {TOAST_CONF} from "../global"
import { toast } from 'react-toastify';

// axios.defaults.baseURL = ''  //正式
axios.defaults.baseURL = API_URI //测试
//axios.defaults.responseType = "json" 


//post请求头
axios.defaults.headers.post["Content-Type"] = "application/x-www-form-urlencoded;charset=UTF-8";
//允许跨域携带cookie信息
axios.defaults.withCredentials = true;
//设置超时
axios.defaults.timeout = 30000;

axios.interceptors.request.use(
    config => {
        return config;
    },
    error => {
        return Promise.reject(error);
    }
);

axios.interceptors.response.use(
    response => {
        if (response.status == 200) {
            //common action when token time out
            if(response.data.status == 20002){
                store.dispatch(logoutAction())
                store.dispatch(clearAllTabIndex())
            }
            //update the token when deadline nearby
            if(response.headers.newtoken){
                store.dispatch(updateT(response.headers.newtoken))
            }
            return Promise.resolve(response);
        } else {
            return Promise.reject(response);
        }
    },
    error => {
        console.log("error  =",error)
        if(error === undefined){
            // alert("Server error, please wait and retry!!")
            toast.error("Server error, please wait and retry!!",TOAST_CONF)
            return
        }
        if(error.response.status == 500){
            // alert("Server error, please wait and retry!!")
            toast.error("Server error, please wait and retry!!",TOAST_CONF)
        }else{
            // alert(JSON.stringify(error), 'request error', {
            //     confirmButtonText: 'confirm',
            //     callback: (action) => {
            //         console.log(action)
            //     }
            // });
            console.log("request failed: ",JSON.stringify(error))
            toast.error(JSON.stringify(error),TOAST_CONF)
        }
    }
);

function setHeader(){
    let token = store.getState().LoginReducer.token
    //console.log("axios:   token= ",token)
    if(token != "" && token != undefined){
        return {"token": token}
    }else{
        return{}
    }
};

export default {
    /**
     * @param {String} url
     * @param {Object} data
     * @returns Promise
     */
    post(url, data) {
        return new Promise((resolve, reject) => {
            axios({
                    method: 'post',
                    url,
                    data: qs.stringify(data),
                    headers:setHeader(),
                })
                .then(res => {
                    resolve(res)
                })
                .catch(err => {
                    reject(err)
                });
        })
    },

    postFormData(url, data) {
        return new Promise((resolve, reject) => {
            axios({
                method: 'post',
                url,
                data:data,
                headers:
                {
                    "token": store.getState().LoginReducer.token,
                    "Content-Type": "multipart/form-data",
                },
            })
            .then(res => {
                resolve(res)
            })
            .catch(err => {
                reject(err)
            });
        })
    },

    get(url, data) {
        return new Promise((resolve, reject) => {
            axios({
                    method: 'get',
                    url,
                    params: data,
                    headers:setHeader(),
                })
                .then(res => {
                    resolve(res)
                })
                .catch(err => {
                    reject(err)
                })
        })
    },
     /**
     * @param {String} url
     * @param {Object} data
     * @returns Promise
     */
      put(url, data) {
        return new Promise((resolve, reject) => {
            axios({
                    method: 'put',
                    url,
                    data: qs.stringify(data),
                    headers:setHeader(),
                })
                .then(res => {
                    resolve(res)
                })
                .catch(err => {
                    reject(err)
                });
        })
    },
     /**
     * @param {String} url
     * @param {Object} data
     * @returns Promise
     */
      delete(url, data) {
        return new Promise((resolve, reject) => {
            axios({
                    method: 'delete',
                    url,
                    data: qs.stringify(data),
                    headers:setHeader(),
                })
                .then(res => {
                    resolve(res)
                })
                .catch(err => {
                    reject(err)
                });
        })
    },
};