import UTILS from "../utils"
const MAIN_TAB_INFO = [
  {index: 'home', label: 'tab.home'},
  // {index: 'service', label: 'tab.service'},
  // {index: 'ports', label: 'tab.ports'},
  // {index: 'onts', label: 'tab.ont'},
  // {index: 'system', label: 'tab.system'},
  {index: 'maintenance', label: 'tab.maintenance'},
  // {index: 'alarm_log', label: 'tab.alarm_log'},
]
  
// const MAINTENANCE_TAB_INFO = [
//   {index: 'backup_restore', label: 'tab.backup_restore'},
//   {index: 'reset', label: 'tab.reset'},
//   {index: 's_management', label: 'tab.smanagement'},
//   {index: 'deploy', label: 'tab.deploy'},
//   // {index: 'zero_touch', label: 'tab.zero_touch'},
//   {index: 'load_config', label: 'tab.load_config'},
//   // {index: 'command', label: 'tab.command'},
//   // {index: 'management', label: 'tab.management'},
// ]

const MAINTENANCE_TAB_INFO =(featurelist)=>{
    let defs = [
      {index: 'load_config', label: 'tab.load_config'},
      {index: 'backup_restore', label: 'tab.backup_restore'},
      {index: 's_management', label: 'tab.smanagement'},
      {index: 'reset', label: 'tab.reset'},
    ]
   
    if(UTILS.hasFeature(featurelist,"WEBGUI_DEPLOYMENT")){
        defs.push({index: 'deploy', label: 'tab.deploy'})
    }
    // defs.push({index: 'load_config', label: 'tab.load_config'})
    return defs
}

const SYSTEM_TAB_INFO = [
  {index: 'management', label: 'Management'},
  {index: 'hardware', label: 'HardWare'},
  {index: 's_management', label: 'Soft Management'},
]

const ALARMLOG_TAB_INFO = [
  {index: 'alarm', label: 'tab.alarm'},
  {index: 'log', label: 'tab.log'},
]
export default {
  MAIN_TAB_INFO,
  MAINTENANCE_TAB_INFO,
  SYSTEM_TAB_INFO,
  ALARMLOG_TAB_INFO,
}