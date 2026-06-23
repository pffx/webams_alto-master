import moment from 'moment';
import { getContext } from '@nokia-csf-uxr/ccfk/common';
import Typography from '@nokia-csf-uxr/ccfk/Typography';
import {OLT_TYPE_PORT} from '../../global';



const hasFeature = (obj, key) =>{
  return obj.hasOwnProperty(key) && obj[key] === true;
}
const tools = {
  dateFormatter (date) {
    return moment(date).format('L')
  } ,
  isTestDataUsed () {
    // console.log("isTestDataUsed  env=",process.env)
    if(process.env.REACT_APP_TEST_DATA_MODE === "true"){
      return true
    }
    return false
  } ,
  /**
   * @param {String} size 
   * @param {Object} data
   * @returns Promise
  */
  sizeFormatter(size){
    if(size < 1024){
      return `${(size)} B`
    }else if(size < (1024*1024) && size >= 1024 ){
      return `${(size / 1024).toFixed(1)} KB`
    }else if (size < (1024*1024*1024) && size >= (1024*1024) ){
      return `${(size / (1024 * 1024)).toFixed(1)} MB`
    }else if (size < (1024*1024*1024*1024) && size >= (1024*1024*1024) ){
      return `${(size / (1024 * 1024 * 1024)).toFixed(1)} GB`
    }
    return `${(size / (1024 * 1024 * 1024)).toFixed(1)} GB`
  },

  isRTL(){
    //return getContext(({ RTL }) => RTL);
    return false
  },

  isValidIpAddress(address){
    if (address === "" || address === undefined){
      return false
    }
    let ipParts = address.split('/');
    let num;
    if (ipParts.length > 2) return false;
    if (ipParts.length == 2) {
        num = parseInt(ipParts[1]);
        if (num <= 0 || num > 32)
            return false;
    }

    if (ipParts[0] == '0.0.0.0' || ipParts[0] == '255.255.255.255')
        return false;

    let addrParts = ipParts[0].split('.');
    if ( addrParts.length != 4 ) {
        return false;
    }

    if ('0' == addrParts[0]) {
        return false;
    }

    for (i = 0; i < 4; i++) {
        let reg = /^[0-9]*$/;
        if (!reg.test(addrParts[i])) {
            return false;
        }
        if (isNaN(addrParts[i]) || addrParts[i] =="")
            return false;
        num = parseInt(addrParts[i]);
        if(isNaN(num)||/\s/.test(addrParts[i]))
            return false;
        if ( num < 0 || num > 255 )
            return false;
    }

    let addr1 = parseInt(addrParts[0]);
    let addr2 = parseInt(addrParts[1]);
    let addr3 = parseInt(addrParts[2]);
    let addr4 = parseInt(addrParts[3]);
    if (addr1 >= 1 && addr1 <= 223 && !ipParts[1]) {
        if (addr1 == 127) {
            return false;
        }
        if (addr4 == 255 || addr4 == 0) {
            return false;
        }
    }
    if (addr1 >= 1 && addr1 <= 127) {  //A
        if (addr1 == 127) {
            return false;
        }
        if ((addr2 == 255 && addr3 == 255 && addr4 == 255)
            || (addr2 == 0 && addr3 == 0 && addr4 == 0)) {
            return false;
        }
    }
    else if (addr1 >= 128 && addr1 <= 191) {  //B
        if ((addr3 == 255 && addr4 == 255)
            ||(addr3 == 0 && addr4 == 0)) {
            return false;
        }
    }
    else if (addr1 >= 192 && addr1 <= 223) {  //C
        if (addr4 == 255 || addr4 == 0) {
            return false;
        }
    }
    else { //D or E
        return false;
    }

    if (ipParts.length == 2) {
        let mask_num = 32 - parseInt(ipParts[1]);
        let ip = (addr1 << 24)|(addr2 << 16)|(addr3 << 8)|(addr4);
        let count = 1;
        for (var i = 0; i < mask_num; i++) {
            count = count * 2;
        }
        count--;

        let host_ip = (ip & count).toString(2);
        let str_mask = (count).toString(2);
        if (host_ip == "0") { //host self = 0
            return false;
        }
        else if (host_ip.length == str_mask.length){ //broadcast address
            if (host_ip.indexOf("01") < 0 && host_ip.indexOf("10") < 0) {
                return false;
            }
        }
    }
    return true;
  },

  getFileNameFromContentDisposition(contentDisposition){
    if (!contentDisposition) return null;
    const match = contentDisposition.match(/filename="?([^"]+)"?/);
    return match ? match[1] : null;
  },

  isEmptyString(str){
    if(str === "" || str === undefined){
      return true
    }else{
      return false
    }
  },
  isOltSelected(oltInfo){
    if(oltInfo.ip !== ""){
      return true
    }else {
      return false
    }
  },
  renderNullSelectedOlt(){
    return(
      <Typography variant="TITLE_16">Please select one OLT in Home page first!</Typography>
    )
  },
  //if it is DF, card will be not used
  getOltPort(oltInfor,card){
    if(oltInfor.type.startsWith("DF")){
      return OLT_TYPE_PORT["df"]
    }else{
      let index = "mf_lt" + card
      return OLT_TYPE_PORT[index]
    }
  },
  getOltPort2(oltInfor,card){
    if(oltInfor.type.startsWith("DF")){
      return OLT_TYPE_PORT["df"]
    }else{
      // let index = "mf_lt" + card
      return OLT_TYPE_PORT[card]
    }
  },
  isNum(num){
    var reg=/^[0-9]+.?[0-9]*$/; //判断字符串是否为数字 ，判断正整数用/^[1-9]+[0-9]*]*$/
    if(!reg.test(num)){
        return false;
    }else{
        return true;
    }
  },
  generateCardList(data){
    console.log("generateCardList   data = ",data)
    if(data.type ==undefined){
      return[]
    }
    var list = [
      {label: 'Card', isHeader: true },
      {value:"mf_lt0",label:"NT"},
    ]
    if(data.type.startsWith("DF")){
      return [
          {label: 'Card', isHeader: true },
          {value:"df",label:"DF"},
        ]
    }else if(data.type.startsWith("MF")){
      data.ltCardStatus.map(function(val,index){
        if(val == 1){
          let val = "mf_lt"+(index+1)
          let label = "LT"+(index+1)
          let tmp = {value:val,label:label}
          list.push(tmp)
        }
      })
      let val = "ihub"
      let label = "Ihub"
      let tmp = {value:val,label:label}
      list.push(tmp)
      return list
    }else{
      return[]
    }
  },
  generateLTCardPlanned(data){
    // console.log("generateLTCardPlanned   data = ",data)
    let list = [0,0,0,0,0,0,0,0,0,0,0,0,0,0]
    for(let i = 0;i<14;i++){
      let ltIndex = "Lt"+ (i+1) +"Info"
      if(data[ltIndex] == "{}"){
        continue
      }
      const searchRegExp = /'/g
      let tmp = data[ltIndex].replace(searchRegExp,'"')
      list[i] = JSON.parse(tmp).planned === "1"?1:0
    }
    // console.log("generateLTCardPlanned   list = ",list)
    return list
  },
  slashCompatibly(url){
    // console.log("slashCompatibly   url = ",url)
    const searchRegExp = /\\/g
    let tmp = ""
    tmp = url.replace(searchRegExp,'/')
    // console.log("slashCompatibly   tmp = ",tmp)
    return tmp
  },
  getSystem() {
		var sUserAgent = navigator.userAgent;
		var isWin = (navigator.platform == "Win32") || (navigator.platform == "Windows");
		var isMac = (navigator.platform == "Mac68K") || (navigator.platform == "MacPPC") || (navigator.platform == "Macintosh") || (navigator.platform == "MacIntel");
		if (isMac) return "Mac";
		var isUnix = (navigator.platform == "X11") && !isWin && !isMac;
		if (isUnix) return "Unix";
		var isLinux = (String(navigator.platform).indexOf("Linux") > -1);
		if (isLinux) return "Linux";
		if (isWin) {
			var isWin2K = sUserAgent.indexOf("Windows NT 5.0") > -1 || sUserAgent.indexOf("Windows 2000") > -1;
			if (isWin2K) return "Win2000";
			var isWinXP = sUserAgent.indexOf("Windows NT 5.1") > -1 || sUserAgent.indexOf("Windows XP") > -1;
			if (isWinXP) return "WinXP";
			var isWin2003 = sUserAgent.indexOf("Windows NT 5.2") > -1 || sUserAgent.indexOf("Windows 2003") > -1;
			if (isWin2003) return "Win2003";
			var isWinVista = sUserAgent.indexOf("Windows NT 6.0") > -1 || sUserAgent.indexOf("Windows Vista") > -1;
			if (isWinVista) return "WinVista";
			var isWin7 = sUserAgent.indexOf("Windows NT 6.1") > -1 || sUserAgent.indexOf("Windows 7") > -1;
			if (isWin7) return "Win7";
			var isWin10 = sUserAgent.indexOf("Windows NT 10") > -1 || sUserAgent.indexOf("Windows 10") > -1;
			if (isWin10) return "Win10";
		}
		return "other";
	},

  hasFeature,
}

export default tools