import React, { useState, useRef } from 'react';
import PropTypes from 'prop-types';
// import { SPACING_24 } from '@nokia-csf-uxr/freeform-design-tokens/tokens/spacing';
import { CAPTION } from '@nokia-csf-uxr/freeform-design-tokens/tokens/typography';

import Dialog from '@nokia-csf-uxr/ccfk/Dialog';
import Button from '@nokia-csf-uxr/ccfk/Button';
import TextArea from '@nokia-csf-uxr/ccfk/TextArea';

import {
  AlertBar,
  AlertMessage,
  AlertDetails
} from '@nokia-csf-uxr/ccfk/Alert';

import {
  DialogTitle,
  DialogContent,
} from '@nokia-csf-uxr/ccfk/Dialog';
import GLOBAL, {SPACING } from "../../../../global"

const EXPANDED_HEIGHT_SHORT = '15.9rem';
const EXPANDED_HEIGHT_LONG = '19.1rem';
const COLLAPSED_HEIGHT = '11rem';
const DIALOG_WIDTH = '26.75rem';

const ERROR_STYLE = (expanded, longMessage) => {
  const resolvedHeight = expanded ? longMessage ? EXPANDED_HEIGHT_LONG : EXPANDED_HEIGHT_SHORT : COLLAPSED_HEIGHT;
  return ({
    height: resolvedHeight,
    top: `calc((100vh - ${resolvedHeight}) / 2)`,
    left: `calc((100vw - ${DIALOG_WIDTH}) / 2)`,
    right: `calc((100vw - ${DIALOG_WIDTH}) / 2)`,
    width: DIALOG_WIDTH,
  });
};

const EXPANDED_MESSAGE_NODE = (msg)=>{
  return(
    <TextArea 
      readOnly
      value={msg}
      style={{ height: '100%' }}
      textareaProps={{ style: { padding: `${SPACING.SPACING_24}`, backgroundColor: 'transparent', ...CAPTION} }}
    />
  )}

const ErrorDialog = (props) => {
  const {
    onRightClick,
    onLeftClick,
    title,
    body,
    extralBody,
    button,
    open,
    longMessage,
    style,
  } = props;
  const [expanded, setExpanded] = useState(false);
  const handleExpansion = () => {
    setExpanded(!expanded);
  };
  const appElementRef = useRef(null);
  return (
    <>
      <div style={style}
        ref={appElementRef}
      />
      <Dialog
        isOpen={open}
        appElement={appElementRef}
        ariaHideApp={false}
        isMaskLight={false}
        style={{ content: { ...ERROR_STYLE(expanded, longMessage) } }}
        onRequestClose={(event) => { console.log(event); }}
      >
        <AlertBar variant="ERROR" />
        <DialogTitle style={{ overflow: (expanded && longMessage) ? 'unset' : undefined }} title={title}  />
        <DialogContent style={{ flex: (expanded && longMessage) ? 'unset' : undefined }}>
          <AlertMessage 
            message={body}
            variant="ERROR"
          />
        </DialogContent>
        <AlertDetails
          expanded={expanded}
          expansionButtonProps={{
            onExpansionChange: handleExpansion,
          }}
          expandedMessage={longMessage ? EXPANDED_MESSAGE_NODE(extralBody) : extralBody}
          buttons={[<Button onClick={() => { onLeftClick();}} key="button-1">CANCEL</Button>, <Button autoFocus onClick={() => { onRightClick();}} key="button-2">{button}</Button>]}
        />
      </Dialog>
    </>
  );
};
ErrorDialog.propTypes = {
  onRightClick:PropTypes.func,
  onLeftClick:PropTypes.func,
  longMessage: PropTypes.bool,
  title: PropTypes.string,
  body: PropTypes.string,
  extralBody: PropTypes.string,
  button: PropTypes.string,
  open: PropTypes.bool.isRequired,
  style: PropTypes.object,
};

ErrorDialog.defaultProps = {
  title: "Error",
  extralBody:"Nothing!",
  longMessage:true,
  body:"",
  button: "OK",
  open:false,
  style:{},
}
export default ErrorDialog;