import React from 'react';
import {useSelector,useDispatch} from 'react-redux'
import { useNavigate } from "react-router-dom";
// import { connect } from 'react-redux'
import AppBanner, { AppBannerLogo, AppBannerContent, AppBannerDivider } from '@nokia-csf-uxr/ccfk/AppBanner';
import HorizontalNavigation from '@nokia-csf-uxr/ccfk/HorizontalNavigation';
import { AppHeader } from '@nokia-csf-uxr/ccfk/App';
import '../../css/index.css'
import LanguageSelector from '../language'
import TabLink from '../tabLink'
import UserAccount from '../userAccount'
import OltOverview from '../oltOverview'
import MENU from '../../global/menu';
import { clearAllTabIndex } from '../../actions/global'

const TabHeader = (props) => {
  const tabIndex = useSelector((state) => state.GlobalReducer.tabIndex)
  const dispatch = useDispatch()
  const navigate = useNavigate();
  const goHome=()=>{
    navigate('/home',{ replace: false })
    dispatch(clearAllTabIndex())
  }

  return (
    <AppHeader>
      <AppBanner>
        <AppBannerLogo style={{width:"10%", cursor:"pointer" }} onClick={() => {goHome()}}/>
        <AppBannerContent  style={{width:"70%",justifyContent: 'space-between'}}>
          <HorizontalNavigation scroll style={{width:"100%"}}>
            {
              MENU.MAIN_TAB_INFO.map((info)=>{
                //console.log("info  =", info);
                return (<TabLink key={info.index} tab_id={info.index} label={info.label} selected={tabIndex === info.index} > </TabLink>)
              })
            }
          </HorizontalNavigation>
        </AppBannerContent>
        <AppBannerContent style={{width:"30%"}}>
        <AppBannerDivider style={{padding:"0rem",margin:"0rem"}}/>
        <OltOverview></OltOverview>
        <UserAccount></UserAccount>
        <LanguageSelector light></LanguageSelector>
      </AppBannerContent>
      </AppBanner>
      
    </AppHeader>
  );
};
export default TabHeader