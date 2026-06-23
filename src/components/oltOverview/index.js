import React from 'react';
import { useTranslation } from 'react-i18next';
import {useSelector,useDispatch} from 'react-redux'

// import {AppContent, AppContentWrapper, AppBody } from '@nokia-csf-uxr/ccfk/App';
import OverlayPanel from '@nokia-csf-uxr/ccfk/OverlayPanel';
import { OverlayPanelHeader } from '@nokia-csf-uxr/ccfk/OverlayPanel';
import ToggleButton from '@nokia-csf-uxr/ccfk/ToggleButton';
import InfoOutlineIcon from '@nokia-csf-uxr/ccfk-assets/DeviceIcon';
import { AppToolbarButtonRow } from '@nokia-csf-uxr/ccfk/AppToolbar';
import Tooltip from '@nokia-csf-uxr/ccfk/Tooltip';
import HorizontalDivider from '@nokia-csf-uxr/ccfk/HorizontalDivider';
// import FloatingPanel from '@nokia-csf-uxr/ccfk/FloatingPanel/FloatingPanel';
// import FloatingPanelHeader from '@nokia-csf-uxr/ccfk/FloatingPanel/FloatingPanelHeader';
// import Card from '@nokia-csf-uxr/ccfk/Card';

// import Label from '@nokia-csf-uxr/ccfk/Label';
import Typography from '@nokia-csf-uxr/ccfk/Typography';

import utils from '../../global/utils';

const TOOLTIP_PROPS = {
    placement: 'left',
    trigger: ['hover', 'focus'],
    modifiers: { offset: { offset: '0, 10' } },
    tooltip: 'Overview of Selected OLT',
  };
// const theme = getContext(({ theme }) => theme.FLOATING_PANEL.TYPOGRAPHY);

const OltOverview = (props)=>{
  const { t } = useTranslation();
  const { variant, ...otherProps } = props;

  const dispatch = useDispatch()
  const {oltInfor} = useSelector((state) => state.GlobalReducer)
  const toggleRef = React.useRef(null); 

  const [showOverview, setshowOverview] = React.useState(false);

//   const handleClosePopup = () => {
//     setshowOverview(false)
//   };
    const renderOLTType=()=>{
        if(oltInfor.type ==="DF16"){
            return oltInfor.type
        }else{
            let num = oltInfor.ltNum ==0 ? "No ": oltInfor.ltNum
            return oltInfor.type + " ( "+num + "  LT)"
        }
    }

  const renderOverview=()=>{
    // console.log("renderOverview  oltInfor = ",oltInfor)
    return(
        <Typography style={{ margin: "1rem" }}>
            {
            oltInfor.ip === "" && oltInfor.hostname === ""
            ?utils.renderNullSelectedOlt()
            :<table>
                <tbody id="systemSummary" >
                {/* <tr>
                    <td>Hostname</td>
                    <td>{oltInfor.hostname}</td>
                </tr> */}
                <tr>
                    <td>IP Address</td>
                    <td>{oltInfor.ip}</td>
                </tr>
                <tr>
                    <td>OLT Type</td>
                    <td>{renderOLTType()}</td>
                </tr>
                {/* <tr>
                    <td>Active Software	</td>
                    <td>{oltInfor.software}</td>
                </tr> */}
                </tbody>
            </table>
            }
            
        </Typography>
    )
  }

  return (
    <div>
        <AppToolbarButtonRow style={{ padding: 0}}>
            <Tooltip {...TOOLTIP_PROPS}>
                <ToggleButton
                    ref={toggleRef}
                    active={showOverview}
                    onClick={() => {
                        setshowOverview(!showOverview);
                    }}
                    >
                    <InfoOutlineIcon  size={'2rem'}/>
                </ToggleButton>
            </Tooltip>
        </AppToolbarButtonRow>

        <OverlayPanel
            visible={showOverview}
            style={{ marginTop:"3rem",minWidth: '20rem',minHeight:"20rem" }}
        >
            <OverlayPanelHeader 
            title={"OLT Overview"}
            // buttons={renderButtons()}
            />
            <HorizontalDivider />

            <div style={{ overflow: 'auto' }}>
              {oltInfor ? renderOverview() :<>Loading...</>}
            </div>
        </OverlayPanel>
    </div>
    );

}
export default OltOverview