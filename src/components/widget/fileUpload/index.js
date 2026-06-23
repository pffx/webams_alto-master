import React, { useRef, useState } from 'react';
import { isFirefox } from 'react-device-detect';
import Hyperlink from '@nokia-csf-uxr/ccfk/Hyperlink'
import Typography from '@nokia-csf-uxr/ccfk/Typography'
import Label from '@nokia-csf-uxr/ccfk/Label';
import FileUpload, { FileUploadSection, FileUploadIcon, FileUploadList,FileUploadLabelContainer } from '@nokia-csf-uxr/ccfk/FileUpload';
import InlineFeedbackMessage from '@nokia-csf-uxr/ccfk/InlineFeedbackMessage';
import { TextInputLabelContent } from '@nokia-csf-uxr/ccfk/TextInput';
import { getEasing, getDuration } from '@nokia-csf-uxr/ccfk/common/parse-tokens';
import { NEUTRAL_GREY_300 } from '@nokia-csf-uxr/freeform-design-tokens/tokens/colors';
import { SPACING_4 } from '@nokia-csf-uxr/freeform-design-tokens/tokens/spacing';
import { SPEED_SLOW } from '@nokia-csf-uxr/freeform-design-tokens/tokens/speed';
import { TRANSITION_ACCELERATE } from '@nokia-csf-uxr/freeform-design-tokens/tokens/transition';
import FileUploadListItem from './fileUploadListItem';
import PropTypes from 'prop-types';
import UTILS from '../../../global/utils';
const INLINEFEEDBACK_STYLE = {
  border: 'none',
  margin: SPACING_4,
  position: 'absolute',
  top: 0,
  width: 'calc(100% - 0.5rem)',
  zIndex: 3
};
const ROW_STYLE= {
  display: "flex",
  justifyContent: "flex-start",
  alignItems: "center",
  flexDirection: "row",
  flexWrap: "nowrap",
}
const INLINEFEEDBACK_ANIMATION = {
  from: { opacity: 0, ...INLINEFEEDBACK_STYLE },
  enter: { opacity: 1, ...INLINEFEEDBACK_STYLE },
  leave: { opacity: 0, ...INLINEFEEDBACK_STYLE },
  config: {
    duration: getDuration(SPEED_SLOW),
    easing: getEasing(TRANSITION_ACCELERATE)
  }
};


const uploadTimeComparator = (a, b) => {
  if ( a.lastModified < b.lastModified ){
    return -1;
  }
  if ( a.lastModified > b.lastModified ){
    return 1;
  }
  return 0;
};
const FileUploadComponent = (props) => {
  const [isDragActive, setIsDragActive] = useState(false);
  const [acceptedFiles, setAcceptedFiles] = useState([]);
  const [rejectedFiles, setRejectedFiles] = useState([]);
  const [showGeneralErrorMessage, setShowGeneralErrorMessage] = useState(false);
  const generalErrorMessage = useRef();
  const acceptedFilesRef = useRef(acceptedFiles);
  const rejectedFilesRef = useRef([]);
  const openDialog = useRef();
  const allFiles = acceptedFiles.concat(rejectedFiles).sort(uploadTimeComparator);
 
  const error = isDragActive && acceptedFiles.length >= props.maxNum;
  const MAX_WIDTH = props.verticalLayout ? "32rem" :"48rem"

  const updateProgress = (file) => {
    let percentage = 0;
    let delay = 0
    for (percentage = 0, delay = 0; percentage <= 100; percentage += 2, delay += 50) {
      const p = percentage;
      setTimeout(() => {
        if (file != null && file.status === 'uploading') {
          //console.log("file uploading   file=",file)
          file.progress = p;
          if( p === 100 ) {
            // set error randomly to files
            // if (Math.floor(Math.random() * 10) % 2 === 1) {
            //   file.errorMessage = 'Unknown error while file was uploading.';
            // }
            file.status = 'complete';
          }
          setAcceptedFiles([...acceptedFilesRef.current]);
        }
      }, delay);
    }
  };
  /** simulate the upload processing */
  const simulateUploading = () => {
    acceptedFilesRef.current.forEach(({ file }) => {
      //console.log("simulateUploading   file=",file)
      // if (file.status === 'pending') {
      //   file.status = 'uploading';
      //   file.progress = 0;
      //   setAcceptedFiles([...acceptedFilesRef.current]);
      //   setTimeout(() => { updateProgress(file); }, Math.floor(Math.random() * 2000));
      // }
      file.status = 'complete';
    });
  };
  const acceptedFilesCallback = (data) => {
    //console.log('acceptedFiles data: ', data);
    if(acceptedFilesRef.current.length>=props.maxNum){
      console.log("The number of files has reached the maximum!")
      return
    }
    const newData = data.map(file => {
      // if uploaded file name exists, set error message to inline notification.
      if (acceptedFilesRef.current.findIndex(({ file: acceptedFile }) => acceptedFile.name === file.name) !== -1) {
        generalErrorMessage.current = 'Some files have been already uploaded.';
        setShowGeneralErrorMessage(true);
      } else {
        file.status = 'pending';
        return { file };
      }
    }).filter((element) =>  element != null);
    if (newData.length > 0) {
      acceptedFilesRef.current = newData.concat(acceptedFilesRef.current);
      setAcceptedFiles(acceptedFilesRef.current);
      simulateUploading();
      props.onFilesChange(acceptedFilesRef.current)
    }
  };
  const fileRejections = (data) => {
    //console.log('fileRejections: ', data);
    if (data[0] && data[0].errors) {
      generalErrorMessage.current = data[0].errors[0].message;
      setShowGeneralErrorMessage(true);
    }
  };
  const handleFeedbackClose = () => {
    setShowGeneralErrorMessage(false);
  };
  const handleDelete = (deleteFile) => {
    //console.log("handleDelete    ")
    if (deleteFile.error) {
      const files = [...rejectedFiles];
      const indexToDelete = files.findIndex(({ file }) => file.name === deleteFile.name);
      files.splice(indexToDelete, 1);
      rejectedFilesRef.current = files;
      setRejectedFiles(files);
    } else {
      const files = [...acceptedFiles];
      const indexToDelete = files.findIndex(({ file }) => file.name === deleteFile.name);
      files.splice(indexToDelete, 1);
      acceptedFilesRef.current = files;
      setAcceptedFiles(acceptedFilesRef.current);
    };
  };
  const handleRetry = (retryFile) => {
    if (retryFile.errorMessage) {
      acceptedFilesRef.current.forEach(({ file }) => {
        if (file.name === retryFile.name) {
          file.status = 'pending';
          file.errorMessage = undefined;
          file.progress = 0;
        }
      });
      setAcceptedFiles([...acceptedFilesRef.current]);
      simulateUploading();
    }
  }
  const FilenameLabel = (
    <Label htmlFor="id-1" verticalLayout style={{ width: !isFirefox ? 'fit-content' : '-moz-fit-content' }}>
      <TextInputLabelContent>
        FileName
      </TextInputLabelContent>
    </Label>
  );

  const UploadedLabel = (
    <Label htmlFor="id-1" verticalLayout style={{ width: !isFirefox ? 'fit-content' : '-moz-fit-content' }}>
      <TextInputLabelContent>
        Uploaded
      </TextInputLabelContent>
    </Label>
  );
  const DragAndDropTextBlock = (
    <>
    Drag and drop files here, or&nbsp;
      <Hyperlink
        aria-label='Browse'
        href="" style={{ cursor:'pointer' }}
        onClick={(e) => { e.preventDefault(); openDialog.current && openDialog.current(); }}>
          browse
      </Hyperlink>
    </>
  );
  const getMaxFileTitle=()=>{
    if(props.maxNum === 1){
      return "Maximum " + props.maxNum + " file"
    }else{
      return "Maximum " + props.maxNum + " files"
    }
  }
  return (
    <div style={{ width: MAX_WIDTH, marginBottom: '0.125rem' }}>
      {/* <Label verticalLayout> */}
      <div style={props.verticalLayout? {}:ROW_STYLE}>
      {/* <Label style={{width:"20%"}}>{props.title}</Label> */}
      <Label verticalLayout style={{ width: !isFirefox ? 'fit-content' : '-moz-fit-content' }}>
        <TextInputLabelContent>
        {props.title}
        </TextInputLabelContent>
      </Label>
      <FileUpload
        dragStatus={(status) => { setIsDragActive(status.isDragActive); }}
        acceptedFiles={acceptedFilesCallback}
        fileRejections={fileRejections}
        open={(open) => { openDialog.current = open}}
        maxFiles={props.maxNum}
        error={error}
        {...props}
      >
        <InlineFeedbackMessage
          variant="error"
          show={showGeneralErrorMessage}
          onClose={handleFeedbackClose}
          animation={INLINEFEEDBACK_ANIMATION}
          closeButton
        >
          {generalErrorMessage.current}
        </InlineFeedbackMessage>
        {allFiles.length > 0 && (
          <>
          {!isDragActive && (<FileUploadLabelContainer>
            {!UTILS.isRTL() && FilenameLabel}
            {UploadedLabel}
            {UTILS.isRTL() && FilenameLabel}
          </FileUploadLabelContainer>)}
           {!isDragActive && (<FileUploadList>
              {allFiles.map(({ file }, index) =>
                <FileUploadListItem
                  // allowDownLoad
                  id={`${file.name}-${index}`}
                  key={`${file.name}-${index}`}
                  index={index}
                  fileName={file.name}
                  errorMessage={file.errorMessage}
                  secondaryContent={UTILS.sizeFormatter(file.size)}
                  progress={file.progress}
                  status={file.status}
                  onDelete={() => handleDelete(file)}
                  // onDownload={() => console.log(`Download file: ${file.name}.`)}
                  onRetry={() => file.errorMessage && handleRetry(file)}
                />
              )}
            </FileUploadList>)}
          </>
        )}
        <FileUploadSection>
          { (acceptedFiles.length === 0 || isDragActive || error) && <FileUploadIcon/> }
          { !isDragActive && (
            <>
              <Typography variant="BODY">
                {acceptedFiles.length === props.maxNum ? 
                'Maximum number of files reached' : DragAndDropTextBlock
                }</Typography>
              {acceptedFiles.length < props.maxNum && <Typography variant="CAPTION" style={{ color: NEUTRAL_GREY_300 }}>{getMaxFileTitle()}</Typography>}
            </>
          )}
          { isDragActive && <Typography variant="TITLE_16" style={{ fontSize: '0.875rem' }}>{error ? 'Maximum files reached' : 'Drop to add files(s)'}</Typography> }
        </FileUploadSection>
      </FileUpload>
      </div>
    </div >)
}
FileUploadComponent.propTypes = {
  maxNum: PropTypes.number.isRequired,
  onFilesChange: PropTypes.func.isRequired,
  title: PropTypes.string,
  verticalLayout:PropTypes.bool
};
FileUploadComponent.defaultProps = {
  maxNum: 1,
  onFilesChange: undefined,
  title: 'File',
  verticalLayout: true,
}
export default FileUploadComponent;