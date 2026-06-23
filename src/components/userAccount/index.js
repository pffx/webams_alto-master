import React from 'react';
import { useTranslation } from 'react-i18next';
import UserAccountSummary, { UserAccountSummaryButton, UserAccountSummaryPopup, UserAccountSummaryUsername, UserAccountSummaryDescription, UserAccountSummaryList, UserAccountSummaryHeader, UserAccountSummaryFooter } from '@nokia-csf-uxr/ccfk/UserAccountSummary';
import Button from '@nokia-csf-uxr/ccfk/Button';
import { ListItemBasic, ListItemText } from '@nokia-csf-uxr/ccfk/List';
import ButtonsRow from '@nokia-csf-uxr/ccfk/ButtonsRow';
import { AlertBar } from '@nokia-csf-uxr/ccfk/Alert';
import NokiaLogo from '@nokia-csf-uxr/ccfk/NokiaLogo';
import Typography from '@nokia-csf-uxr/ccfk/Typography';
import Dialog, { DialogContent } from '@nokia-csf-uxr/ccfk/Dialog';
import { getContext } from '@nokia-csf-uxr/ccfk/common';
import TOKENS from '@nokia-csf-uxr/freeform-design-tokens';
import Hyperlink from '@nokia-csf-uxr/ccfk/Hyperlink';
import Avatar from '@nokia-csf-uxr/ccfk/Avatar';

import {useSelector,useDispatch} from 'react-redux'
import { logoutAction } from '../../actions/login'
import { clearAllTabIndex } from '../../actions/global'
import { ARROW_DOWN, ENTER_KEY } from '../../global/keybaord';

//import Redirect from '../../router/redirect'
const TOOLTIP_PROPS = {
    placement: 'bottom',
    //trigger: ['hover', 'focus'],
    trigger: [''],
    modifiers: { offset: { offset: '0, 10' } },
    tooltip: 'User Info',
};

const UserAccount = (props)=>{
  // const { elevationIndex, dark, variant, ...otherProps } = props;
  const { t } = useTranslation();
  const { variant, ...otherProps } = props;
  const {userInfo} = useSelector((state) => state.LoginReducer)
//   console.log("userInfo = ",userInfo)

  const dispatch = useDispatch()
  const THEME = getContext(({ theme }) => theme.USER_ACCOUNT_SUMMARY.ABOUT_PRODUCT);
  const AVATAR_PROPS = {
    // disabled,
    // hoverEffect,
    // badged,
    // shape,
    // onClick: !disabled && hoverEffect ? (e) => { console.log('Avatar onClick', e); } : undefined,
    // onKeyDown: !disabled && hoverEffect ? (e) => { console.log('Avatar onKeyDown', e); } : undefined,
    role: 'img',
    ...otherProps
  };

  const popupRef = React.useRef(null);
  const listRef = React.useRef(null);
  const [showUserAccountSummary, setShowUserAccountSummary] = React.useState(false);
  const [version, getVersion] =React.useState("2.1.0")
  const [showDialog, setShowDialog] = React.useState(false);
  const handleClosePopup = () => {
    setShowUserAccountSummary(false)
  };
  const handleCloseDialog = () => {
    setShowDialog(false);
  };
  const handleLogout = ()=>{
    dispatch(logoutAction())
    dispatch(clearAllTabIndex())
  }

  return (
    <UserAccountSummary
        onEscFocusOut={handleClosePopup}
        onClickOutside={handleClosePopup}
        {...otherProps}
    >
        <UserAccountSummaryButton
        buttonProps={{
            onClick: () => {
            setShowUserAccountSummary(!showUserAccountSummary);
            }
        }}
        tooltipProps={TOOLTIP_PROPS}
        variant={variant}
        >
        {"Settings"}
        </UserAccountSummaryButton>
        <UserAccountSummaryPopup
        // elevationProps={{ elevationIndex, dark }}
        ref={popupRef}
        open={showUserAccountSummary}
        onKeyDown={(e) => {
            // for accessibility, if user presses down arrow, focus goes to the List
            if (e.key === ARROW_DOWN && !listRef.current.contains(e.target)) {
            e.preventDefault();
            listRef.current && listRef.current.focus({ preventScroll: true });
            }
        }}
        >
        <UserAccountSummaryHeader>
            {/* remove the account display part for Alto */}
            {/* <div className='row'>
                <Avatar {...AVATAR_PROPS} size="xlarge">
                  {userInfo.account ? userInfo.account.substring(0, 3) : ""}
                </Avatar>
            <div style={{marginLeft:"3rem"}}>
                <UserAccountSummaryUsername>{userInfo.account}</UserAccountSummaryUsername>
                <UserAccountSummaryDescription>
                    {(userInfo.role ? userInfo.role : "") + " - " + (userInfo.department ? userInfo.department : "")}
                </UserAccountSummaryDescription>
            </div>
            </div> */}
            
        </UserAccountSummaryHeader>
        <UserAccountSummaryList isOverflowNecessary={false} ref={listRef}>
            <ListItemBasic
            onClick={() => {
                setShowDialog(true);
            }}
            onKeyDown={(event) => {
                if (event.key === ' ' || event.key === ENTER_KEY) {
                setShowDialog(true);
                }
            }}
            >
            <ListItemText>About Alto</ListItemText>
            </ListItemBasic>
        </UserAccountSummaryList>
        {/* <UserAccountSummaryFooter>
            <ButtonsRow>
            <Button onClick={handleLogout} >{t('button.logout')}</Button>
            </ButtonsRow>
        </UserAccountSummaryFooter> */}
        </UserAccountSummaryPopup>
        {
        <Dialog
            ariaHideApp={false}
            isOpen={showDialog}
            style={{
            content: {
                left: 'calc((100vw - 21.875rem) / 2)',
                top: 'calc((100vh - 15.3125rem) / 2)',
                width: '21.875rem',
                height: '15.3125rem',
                transform: 'initial'
            }
            }}
            onRequestClose={handleCloseDialog}
        >
            <AlertBar />
            <DialogContent isTopDividerVisible isBottomDividerVisible style={{ overflow: 'hidden', padding: 0, border: 0 }}>
                <div style={{paddingRight:"1.5rem", paddingLeft:"1.0rem",paddingTop:"1.0rem"}}>
                    <div className="accout-image-wrapper">
                        <NokiaLogo color={THEME.logo} iconProps={{ size: { height: '1.5rem' } }} />
                    </div>
                    <Typography variant="TITLE_24"
                        style={{
                        fontFamily: TOKENS.FONT_FAMILY.NOKIA_PURE_HEADLINE_LIGHT,
                        marginBottom: TOKENS.SPACING.SPACING_8,
                        }}>
                        {"Alto"}
                    </Typography>
                    <Typography variant="CAPTION" style={{
                        marginBottom: TOKENS.SPACING.SPACING_8,
                    }}>
                        {"Release " + version}
                    </Typography>
                    <Typography variant="CAPTION" style={{
                        marginBottom: TOKENS.SPACING.SPACING_8,
                    }}>
                        {"APAC RBC NPI Team"}
                    </Typography>
                    <Hyperlink
                        href="http://nokia.com"
                        target="_blank"
                        children="Terms and Conditions"
                        style={{
                            ...TOKENS.TYPOGRAPHY.MICRO
                        }}
                    />
                </div>
                <div style={{ paddingLeft:"0.5rem"}}>
                    <ButtonsRow>
                        <Typography variant="CAPTION" style={{ color: THEME.copyright }} >
                            &copy; {new Date().getFullYear()} Nokia
                        </Typography>
                        <div className="buttonWrapper">
                            <Button onClick={handleCloseDialog} aria-label="Close" children="CLOSE" />
                        </div>
                    </ButtonsRow>
                </div>
            </DialogContent>
        </Dialog>
        }
    </UserAccountSummary>
    );

}
export default UserAccount