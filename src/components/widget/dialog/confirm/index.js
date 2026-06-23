import React, { useRef } from 'react';
import PropTypes from 'prop-types';

import Dialog, { DialogTitle, DialogContent, DialogFooter } from '@nokia-csf-uxr/ccfk/Dialog';
import Button from '@nokia-csf-uxr/ccfk/Button';
import { AlertBar, AlertMessage,} from '@nokia-csf-uxr/ccfk/Alert';

const DIALOG_HEIGHT = '12.5rem';
const DIALOG_WIDTH = '26.75rem';
const CONFIRM_STYLE = {
  top: `calc((100vh - ${DIALOG_HEIGHT}) / 2)`,
  height: DIALOG_HEIGHT,
  minHeight: DIALOG_HEIGHT,
  left: `calc((100vw - ${DIALOG_WIDTH}) / 2)`,
  right: `calc((100vw - ${DIALOG_WIDTH}) / 2)`,
  width: DIALOG_WIDTH,
};

const ConfirmationDialog = (props) => {
  const {
    onConfirm,
    onCancel,
    title,
    body,
    open,
    style,
    data,
    type,
  } = props;
  const appElementRef = useRef(null);
  return(
    <>
      <div style={ style }
        ref={appElementRef}
      />
      <Dialog
        appElement={appElementRef}
        isOpen={open}
        ariaHideApp={false}
        isMaskLight={false}
        style={{ content: CONFIRM_STYLE }}
        onRequestClose={(event) => { console.log(event); }}
      >
        <AlertBar variant="CONFIRM" />
        <DialogTitle title={title} />
        <DialogContent
          isTopDividerVisible={false}
          isBottomDividerVisible={false}
          style={{ padding: '4px 21px 0 24px', overflow: 'hidden' }}
        >
          <AlertMessage 
            message={body}
            variant="CONFIRM"
          />
        </DialogContent>
        <DialogFooter>
          <Button onClick={() => {
              onCancel();
            }}
          >Cancel</Button>
          <Button
            autoFocus
            onClick={() => {
              onConfirm(data,type);
            }}>Confirm</Button>
        </DialogFooter>
      </Dialog>
    </>
  );
}

ConfirmationDialog.propTypes = {
  onConfirm:PropTypes.func,
  onCancel:PropTypes.func,
  title: PropTypes.string,
  body: PropTypes.string,
  open: PropTypes.bool.isRequired,
  data: PropTypes.object,
  style: PropTypes.object,
  type:PropTypes.string,
};

ConfirmationDialog.defaultProps = {
  title: "Are you Confirmed?",
  body:"",
  open:false,
  style:{},
  data:{},
  type:"",
}
export default ConfirmationDialog;