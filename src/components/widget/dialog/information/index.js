import React, { useRef } from 'react';
import PropTypes from 'prop-types';

import Dialog, { DialogTitle, DialogContent, DialogFooter } from '@nokia-csf-uxr/ccfk/Dialog';
import Button, { ButtonText } from '@nokia-csf-uxr/ccfk/Button';
import Typography from '@nokia-csf-uxr/ccfk/Typography';



const DIALOG_HEIGHT = '12.5rem';
const DIALOG_WIDTH = '26.75rem';
const INFOR_STYLE = {
  top: `calc((100vh - ${DIALOG_HEIGHT}) / 2)`,
  height: DIALOG_HEIGHT,
  minHeight: DIALOG_HEIGHT,
  left: `calc((100vw - ${DIALOG_WIDTH}) / 2)`,
  right: `calc((100vw - ${DIALOG_WIDTH}) / 2)`,
  width: DIALOG_WIDTH,
};

const InformationDialog = (props) => {
  const {
    onClick,
    title,
    body,
    button,
    open,
    style,
  } = props;
  const appElementRef = useRef(undefined);
  return (
    <>
      <div style={ style }
        ref={appElementRef}
      />
      <Dialog
        appElement={appElementRef}
        isOpen={open}
        isMaskLight={false}
        style={{ content: INFOR_STYLE }}
        onRequestClose={(event) => { console.log(event); }}
      >
        <DialogTitle title={title} />
        <DialogContent
          isTopDividerVisible={false}
          isBottomDividerVisible={false}
          style={{ overflow: 'hidden' }}
        >
          <Typography variant="BODY">{body}</Typography>
        </DialogContent>
        <DialogFooter>
          <Button 
            autoFocus
            onClick={() => {
              onClick();
            }}>{button}</Button>
        </DialogFooter>
      </Dialog>
    </>
  );
}
  
InformationDialog.propTypes = {
  onClick:PropTypes.func,
  title: PropTypes.string,
  body: PropTypes.string,
  button: PropTypes.string,
  open: PropTypes.bool.isRequired,
  style: PropTypes.object,
};

InformationDialog.defaultProps = {
  title: "Information",
  body:"",
  button: "Confirm",
  open:false,
  style:{},
}
export default InformationDialog;