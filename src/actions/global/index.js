import { GlobalSets } from "../constants"
const changeTabIndex = data => ({type:GlobalSets.ChangeTabIndex,data:data})
const changeSubTabIndex = data => ({type:GlobalSets.ChangeSubTabIndex,data:data})
const clearAllTabIndex = () => ({type:GlobalSets.ClearAllTabIndex})
const updateOltInfor = (data) => ({type:GlobalSets.UpdateOltInfor,data:data})
const clearOltInfor = () => ({type:GlobalSets.ClearOltInfor})
export {
    changeTabIndex,
    changeSubTabIndex,
    clearAllTabIndex,
    updateOltInfor,
    clearOltInfor,
}