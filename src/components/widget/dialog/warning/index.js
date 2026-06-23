import React, { useRef } from 'react';
import PropTypes from 'prop-types';

import Dialog, { DialogTitle, DialogContent, DialogFooter } from '@nokia-csf-uxr/ccfk/Dialog';
import Button from '@nokia-csf-uxr/ccfk/Button';
import { AlertBar, AlertMessage,} from '@nokia-csf-uxr/ccfk/Alert';
//import { patternBackgroundStyle } from '@nokia-csf-uxr/ccfk/StorybookHelpers';

const DIALOG_HEIGHT = '12.5rem';
const DIALOG_WIDTH = '26.75rem';
const WARNING_STYLE = {
  top: `calc((100vh - ${DIALOG_HEIGHT}) / 2)`,
  height: DIALOG_HEIGHT,
  minHeight: DIALOG_HEIGHT,
  left: `calc((100vw - ${DIALOG_WIDTH}) / 2)`,
  right: `calc((100vw - ${DIALOG_WIDTH}) / 2)`,
  width: DIALOG_WIDTH,
};

const WarningDialog = (props) => {
  const {
    onRightClick,
    onLeftClick,
    title,
    body,
    rButton,
    lButton,
    open,
    style,
  } = props;
  const appElementRef = useRef(null);
  return (
    <>
      <div style={ style }
        ref={appElementRef}
      />
      <Dialog
        appElement={appElementRef}
        isOpen={open}
        isMaskLight={false}
        ariaHideApp={false}
        style={{ content: WARNING_STYLE }}
        onRequestClose={(event) => { console.log(event); }}
      >
        <AlertBar variant="WARNING" />
        <DialogTitle title={title} />
        <DialogContent
          isTopDividerVisible={false}
          isBottomDividerVisible={false}
        >
          <AlertMessage
            variant="WARNING"
            message={body}
          />
        </DialogContent>
        <DialogFooter>
          {lButton && <Button  onClick={() => { onLeftClick();}}>{lButton}</Button>}
          <Button
            autoFocus
            onClick={() => {
                onRightClick();
            }}>
            {rButton}
          </Button>
        </DialogFooter>
      </Dialog>
    </>
  );
}

WarningDialog.propTypes = {
  onRightClick:PropTypes.func,
  onLeftClick:PropTypes.func,
  title: PropTypes.string,
  body: PropTypes.string,
  rButton: PropTypes.string,
  lButton: PropTypes.string,
  open: PropTypes.bool.isRequired,
  style: PropTypes.object,
};

WarningDialog.defaultProps = {
  title: "Warning!",
  body:"",
  rButton: "Confirm",
  lButton: undefined,
  open:false,
  style:{},
}

export default WarningDialog;