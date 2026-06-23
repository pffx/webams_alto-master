import React, { useEffect, useState } from 'react';
import DeleteIcon from '@nokia-csf-uxr/ccfk-assets/DeleteIcon';
import RefreshIcon from '@nokia-csf-uxr/ccfk-assets/RefreshIcon';
import ArrowBoldDownCircleOutlineIcon from '@nokia-csf-uxr/ccfk-assets/ArrowBoldDownCircleOutlineIcon';
import { getContext, getEasing, getDuration  } from '@nokia-csf-uxr/ccfk/common';
import { ListItemBasic, ListItemText, ListItemTextContainer, ListItemSubText, ListItemActions } from '@nokia-csf-uxr/ccfk/List';
import Tooltip from '@nokia-csf-uxr/ccfk/Tooltip';
import IconButton from '@nokia-csf-uxr/ccfk/IconButton';
import LinearProgressIndicatorDeterminate from '@nokia-csf-uxr/ccfk/LinearProgressIndicatorDeterminate';
import { useTransition, animated } from 'react-spring';
import { SPACING_16, SPACING_12, SPACING_48, SPACING_8 } from '@nokia-csf-uxr/freeform-design-tokens/tokens/spacing';
import { TRANSITION_ACCELERATE } from '@nokia-csf-uxr/freeform-design-tokens/tokens/transition';
import { SPEED_SLOW } from '@nokia-csf-uxr/freeform-design-tokens/tokens/speed';
import { TRANSPARENT } from '@nokia-csf-uxr/freeform-design-tokens/tokens/colors';
import UTILS from '../../../../global/utils';
const LISTITEM_TEXT_STYLES = {
  padding: '0 0 0.0625rem 0'
};
const LISTITEM_SECONDARY_STYLES = {
  padding: `0 ${SPACING_16} 0.0625rem ${SPACING_16}`,
}
const TOOLTIP_PROPS = {
  placement: 'bottom',
  trigger: ['hover'],
  fallbackPlacements: ['top'],
  modifiers: { offset: { offset: '[0, 12]' } },
  tooltip: ''
};
/** Example of how to build the list item for FileUpload */
const FileUploadListItem = (props) => {
  const {
    fileName,
    status,
    secondaryContent,
    disabled,
    progress,
    errorMessage,
    onFocus,
    onBlur,
    onPointerEnter,
    onPointerLeave,
    onRetry,
    onDelete,
    onDownload,
    allowDownLoad,
    index,
    ...otherProps
  } = props;
  const $theme = getContext(({ theme }) => theme.FILE_UPLOAD);
  const [hideProgress, setHideProgress] = useState(progress >= 100);
  const [fadeOut, setFadeOut] = useState(false);
  const [isFocused, setIsFocused] = useState(false);
  const [isHovered, setIsHovered] = useState(false);
  const [isDeleting, setIsDeleting] = useState(false);
  const handleFocus = (e) => {
    setIsFocused(true);
    onFocus && onFocus(e);
  };
  const handleBlur = (e) => {
    setIsFocused(false);
    onBlur && onBlur(e);
  };
  const handlePointerEnter = (e) => {
    setIsHovered(true);
    onPointerEnter && onPointerEnter(e);
  };
  const handlePointerLeave = (e) => {
    setIsHovered(false);
    onPointerLeave && onPointerLeave(e);
  };
  const resetHoverAndFocus = () => {
    setIsFocused(false);
    setIsHovered(false);
  }
  const handleRetry = (e) => {
    resetHoverAndFocus();
    onRetry && onRetry(e);
  };
  const handleDownload = (e) => {
    onDownload && onDownload(e);
  }
  const handleDelete = (e) => {
    setIsDeleting(true);
    /** wait for the end of animation, then call onDelete callback. */
    setTimeout(() => {
      resetHoverAndFocus();
      onDelete && onDelete(e);
    }, getDuration(SPEED_SLOW));
  };
  const getProgress = () => Math.round((progress / 100) * 100);
  useEffect(() => {
    if (progress >= 100) {
      setFadeOut(true);
      setTimeout(() => setHideProgress(true), 500);
    }
    if (fadeOut && progress < 100) {
      setFadeOut(false);
      setHideProgress(false);
    }
  }, [progress]);
  const LIST_ITEM_STYLE = {
    height: errorMessage ? '3.75rem' : SPACING_48,
    minHeight: errorMessage ? '3.75rem' : SPACING_48,
    position: 'relative',
    overflowY: 'hidden',
    overflowX: 'hidden',
    scrollbars: 'none',
    "::webkitScrollbar": { display: 'none' },
    msOverflowStyle: 'none',
    scrollbarWidth: 'none'
  };
  const LIST_ITEM_ERROR_TEXT_STYLES = {
    paddingBottom: '0.0625rem',
    color: $theme.error,
  };
  const actionButtonStyle = isFocused || isHovered || errorMessage ? { opacity: 1, width: 'auto', height: 'auto' } : { opacity: 0 };
  /** Animation when trying to delete this item by clicking the delete button. It requires react-spring version 8. If you don't need this animation. Feel free to delete this part of code. */
  const defaultAnimationConfig = {
    config: {
      duration: getDuration(SPEED_SLOW),
      easing: getEasing(TRANSITION_ACCELERATE)
    },
    from: { opacity: 1, height: LIST_ITEM_STYLE.height },
    enter: { opacity: 1, height: LIST_ITEM_STYLE.height},
    leave: { opacity: 0, height: '0rem', minHeight: '0rem' }
  };
  const transitions = useTransition(!isDeleting, null, defaultAnimationConfig);
  return transitions.map(({ item, key, props: transitionsStyles }) =>
  item && (
    <animated.div style={{ ...LIST_ITEM_STYLE, ...transitionsStyles }} key={key}>
      {/* This is a hack for ListIteBasic. When it is the first child, hasBottomBorder will add both top and bottom border. 
      As it will always be the first child of animated.div, add this empty div to avoid this behaviour */}
      {index!==0 && <div></div>}
      <ListItemBasic
        disabled={disabled}
        hasBottomBorder
        onPointerEnter={handlePointerEnter}
        onPointerLeave={handlePointerLeave}
        onFocus={handleFocus}
        onBlur={handleBlur}
        {...otherProps}
        style={LIST_ITEM_STYLE}
      >
        {/* Section for File name */}
        <ListItemTextContainer style={{ padding: 0 }}>
          <ListItemText style={LISTITEM_TEXT_STYLES} >{fileName}</ListItemText>
          {errorMessage && <ListItemSubText style={LIST_ITEM_ERROR_TEXT_STYLES} >{errorMessage}</ListItemSubText>}
        </ListItemTextContainer>
        {/* Section for secondary content like size or uploaded Date */}
        {!isFocused && !isHovered && !errorMessage && (
          <>
            {/* Transparent div added to give spacing for the secondaryContent so it will not be overlapped by filename as position:absolute is needed for proper RTL LTR */}
            <div style={{ color:`${TRANSPARENT}`, position: 'relative', overflow: 'unset', padding: 0, right: !UTILS.isRTL() ? '0rem' : 'auto', left: !UTILS.isRTL() ? 'auto' : '0rem', marginLeft: `${SPACING_8}`, marginRight: `${SPACING_8}` }}>{secondaryContent}</div>
            <ListItemTextContainer style={{ position: 'absolute', overflow: 'unset', padding: 0, right: !UTILS.isRTL() ? '0rem' : 'auto', left: !UTILS.isRTL() ? 'auto' : '0rem', marginLeft: `${SPACING_8}`, marginRight: `${SPACING_8}` }}>
              <ListItemSubText style={LISTITEM_SECONDARY_STYLES}>{secondaryContent}</ListItemSubText>
            </ListItemTextContainer>
          </>
        )}
        {/* Section for action buttons */}
        {(
          <ListItemActions style={{ overflow: 'visible', ...actionButtonStyle }}>
              {errorMessage ? (
                <IconButton aria-label='Retry' disabled={disabled} onClick={handleRetry}>
                  <Tooltip {...TOOLTIP_PROPS} tooltip="Retry">
                    <RefreshIcon />
                  </Tooltip>
                </IconButton>
              ): ( allowDownLoad &&
                <IconButton aria-label='Download' disabled={disabled} onClick={handleDownload}>
                  <Tooltip {...TOOLTIP_PROPS} tooltip="Download">
                    <ArrowBoldDownCircleOutlineIcon />
                  </Tooltip>
                </IconButton>
              )}
              <IconButton
                disabled={disabled}
                aria-label='Delete'
                onClick={handleDelete}
                rippleStyle={{ margin: UTILS.isRTL() ? `0 ${SPACING_12} 0 0` : `0 0 0 ${SPACING_12}`, zIndex: 0 }}
              >
                <Tooltip {...TOOLTIP_PROPS} tooltip="Delete">
                  <DeleteIcon />
                </Tooltip>
              </IconButton>
          
          </ListItemActions>
        )}
        {/* Section for progress indicator */}
        {progress != null && progress <= 100 && !hideProgress && (
          <LinearProgressIndicatorDeterminate
            style={{
              position: 'absolute',
              bottom: 0,
              left: 0,
              width: '100%'
            }}
            progress={getProgress()}
            fade={fadeOut}
          />
        )}
      </ListItemBasic>
    </animated.div>
  ));
};
export default FileUploadListItem;