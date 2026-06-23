import React from 'react';
import PropTypes from 'prop-types';
import { useTranslation } from 'react-i18next';
import { toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import Typography from '@nokia-csf-uxr/ccfk/Typography';
import ExpansionPanelsContainer, {
  ExpansionPanel,
  ExpansionPanelTitle,
  ExpansionPanelHeader,
  ExpansionPanelBody,
  ExpansionPanelButton
} from '@nokia-csf-uxr/ccfk/ExpansionPanels';

import TextInput, { TextInputLabelContent } from '@nokia-csf-uxr/ccfk/TextInput';
import Label, {LabelHelpIcon} from '@nokia-csf-uxr/ccfk/Label';
import Tooltip from '@nokia-csf-uxr/ccfk/Tooltip';
import HelpCircleIcon from '@nokia-csf-uxr/ccfk-assets/HelpCircleIcon';
import UploadIcon from '@nokia-csf-uxr/ccfk-assets/UploadIcon';
import DownloadIcon from '@nokia-csf-uxr/ccfk-assets/DownloadIcon';
import StopCircleOutlineIcon from '@nokia-csf-uxr/ccfk-assets/StopCircleOutlineIcon';
import StopIcon from '@nokia-csf-uxr/ccfk-assets/StopIcon';
import Button, { ButtonText, ButtonIcon } from '@nokia-csf-uxr/ccfk/Button';
import ButtonsRow from '@nokia-csf-uxr/ccfk/ButtonsRow';
import HorizontalDivider from '@nokia-csf-uxr/ccfk/HorizontalDivider'
// import Select from "../../../../components/widget/select"
import ConfirmDialog from '../../../../components/widget/dialog/confirm';
import GLOBAL, {SPACING, TOAST_CONF, COLOR } from "../../../../global"
import utils from "../../../../global/utils"
import { KEYDOWN, CLICK, ENTER_KEY, SPACE_KEY } from '../../../../global/keybaord';
import AXIOS from '../../../../axios';
import {API_SoftwareAction,API_SoftwareMigrationUpload,API_SoftwareMigration } from "../../../../global/API"


const PADDING = `${SPACING.SPACING_16} ${SPACING.SPACING_24}`;

const ExpansionPanels = (props) => {
  const {
    data,
    softwareServer,
    onChange,
    style,
  } = props;
  const { t } = useTranslation();
  const { disablePanel, selectPanel, ...otherProps } = props;
  const [expanded, setExpanded] = React.useState(undefined);
  const [actionPanel, setActionPanel] = React.useState(undefined);
  const [softwarePath, setSoftwarePath] = React.useState("");
  const [targetVersion, setTargetVersion] = React.useState("");
  const [softwareName, setSoftwareName] = React.useState("");
  const [confirmMessage, setConfirmMessage] = React.useState("");
  const [showConfirm, setShowConfirm] = React.useState(false);
  const [confirmType, setConfirmType] = React.useState("");
  const headerRef = React.useRef(null);
  const SUBTEXT_COLOR = "#616161"
  const DISABLED_SUBTEXT_COLOR = "#FFFF00"

  const isExpanded = id=>{
    // console.log("isExpanded   id = ",id)
    // console.log("isExpanded   expanded = ",expanded)
    if(expanded === id){
      return true
    }else{
      return false
    }
  }

  const handleExpansion = id => (e) => {
    const newId = expanded === id ? undefined : id;
    // console.log("handleExpansion   id = ",id)
    // console.log("handleExpansion   expanded = ",expanded)
    if (e.type === KEYDOWN) {
      if (
        e.target.getAttribute('data-test') === 'header' && // check if keydown from header
        (e.key === SPACE_KEY || e.key === ENTER_KEY)
      ) {
        setExpanded(newId);
      }
    }
    if (e.type === CLICK) {
      setExpanded(newId);
    }
  };
  const sendUploading = (panel)=>{
    onChange(panel,softwarePath,softwareName)
    
    let uploadData = {
      oltId: panel.ip,
      dstPort:panel.port,
      action: "download",
      url: utils.slashCompatibly(softwarePath) + "/" + softwareName,
      name:softwareName
    }
    AXIOS
      .put(API_SoftwareAction,uploadData)
      .then((res) => {
        console.log(res)
        if(res.data.status === GLOBAL.ERROR_NUM.Success){
          toast.success("Software upload success!",TOAST_CONF)
        }else if (res.data.status === GLOBAL.ERROR_NUM.ActionFailed){
          toast.error("Software upload Faild, please check and retry!",TOAST_CONF)
        }
      })
      .catch((err) => {
    });

  }

  const sendMigrate = (panel)=>{
    // onChange(panel,softwarePath,softwareName)
    
    let migrationData = {
      oltId: panel.ip,
      dstPort:panel.port,
      target:targetVersion
    }
    AXIOS
      .put(API_SoftwareMigration,migrationData)
      .then((res) => {
        console.log(res)
        if(res.data.status === GLOBAL.ERROR_NUM.Success){
          toast.success("Configuarion migration success!",TOAST_CONF)
        }else if (res.data.status === GLOBAL.ERROR_NUM.ActionFailed){
          toast.error("Configuration migration Faild, please check and retry!",TOAST_CONF)
        }
      })
      .catch((err) => {
    });

  }
  const sendActiveWithMigration = (panel)=>{
    let name = getActiveSoftwareName(panel)
    
    let data = {
      oltId: panel.ip,
      dstPort:panel.port,
      action: "active_migrate",
      name: name,
    }
    AXIOS
      .put(API_SoftwareAction,data)
      .then((res) => {
        console.log(res)
        if(res.data.status === GLOBAL.ERROR_NUM.Success){
          toast.success("Software active success!",TOAST_CONF)
        }else if (res.data.status === GLOBAL.ERROR_NUM.ActionFailed){
          toast.error("Software active Faild, please check and retry!",TOAST_CONF)
        }
      })
      .catch((err) => {
    });

  }
  const sendMigrateUpload = (panel)=>{
    // onChange(panel,softwarePath,softwareName)
    
    let uploadData = {
      oltId: panel.ip,
      dstPort:panel.port,
    }
    AXIOS
      .put(API_SoftwareMigrationUpload,uploadData)
      .then((res) => {
        console.log(res)
        if(res.data.status === GLOBAL.ERROR_NUM.Success){
          toast.success("Migration file upload success!",TOAST_CONF)
        }else if (res.data.status === GLOBAL.ERROR_NUM.ActionFailed){
          toast.error("Migration file upload Faild, please check and retry!",TOAST_CONF)
        }
      })
      .catch((err) => {
    });

  }
  const checkSoftwarePathAndName = ()=>{
    // console.log("checkSoftwarePath   software = ",softwareServer)
    // console.log("checkSoftwarePath   softwareName = ",softwareName)
    // console.log("checkSoftwarePath   softwarePath = ",softwarePath)
    let pathSets = utils.slashCompatibly(softwarePath).split('/')
    // console.log("checkSoftwarePath   pathSets = ",pathSets)
    if(pathSets.length <= 1){
      toast.error("Please input the full path of the firmware for OLT!",TOAST_CONF)
      return false
    }
    let obj = softwareServer
    let found = false
    for (var index=0; index < pathSets.length; index++){
      let tmpKey = pathSets[index]
      if(index === pathSets.length-1){
        if(typeof(obj[tmpKey]) === "object"){
          Object.keys(obj[tmpKey]).map(key => {
            if(softwareName === key){
              found = true
            }
          })
          if(!found){
            toast.error("Wrong software name!",TOAST_CONF)
            break
          }
        }else{
          found = false
          toast.error("Wrong software path! Please check and correct the File Path",TOAST_CONF)
          break
        }
      }else{
        if(typeof(obj[tmpKey]) === "object"){
          obj = obj[tmpKey]
        }else{
          found = false
          toast.error("Wrong software path!",TOAST_CONF)
          break
        }
      }

    }
    if(!found){
      return false
    }
    return true
  }
  const getActiveSoftwareName = (panel)=>{
    let name = ""
    let version1 = panel.version1
    let version2 = panel.version2
    if(version1.active && version1.commit && !version2.active && !version2.commit && version2.valid){
      name = version2.name
    }
    if(version2.active && version2.commit && !version1.active && !version1.commit && version1.valid){
      name = version1.name
    }
    if(!version1.active && version1.commit && version2.active && !version2.commit && version2.valid){
      name = version1.name
    }
    if(!version2.active && version2.commit && version1.active && !version1.commit && version1.valid){
      name = version2.name
    }
    return name
  }
  const getCommitSoftwareName = (panel)=>{
    let name = ""
    let version1 = panel.version1
    let version2 = panel.version2
    if(!version1.active && version1.commit && version2.active && !version2.commit && version2.valid){
      name = version2.name
    }
    if(!version2.active && version2.commit && version1.active && !version1.commit && version1.valid){
      name = version1.name
    }
    return name
  }
  const handleUploading = (panel)=>{
    console.log("handleUploading     panel = ",panel)
    if(softwarePath === "" || softwareName === ""){
      toast.error("Please input the path and name of the firmware for OLT!",TOAST_CONF)
      return
    }
    // if(!checkSoftwarePathAndName()){
    //   return
    // }
    setConfirmMessage(`Firmware ${softwarePath}/${softwareName} will be downloaded by OLT!`)
    setConfirmType("download")
    setShowConfirm(true)
    setActionPanel(panel)
  }

  const handleActiveWithMigrate= (panel)=>{
    console.log("handleActiveWithMigrate     panel = ",panel)
    let name = getActiveSoftwareName(panel)
    if(name === ""){
      toast.error("No need to active",TOAST_CONF)
      return
    }

    setConfirmMessage(`Frimware ${name} will be actived with migration!`)
    setConfirmType("active_migrate")
    setShowConfirm(true)
    setActionPanel(panel)
  }

  const handleActive = (panel)=>{
    console.log("handleActive     panel = ",panel)
    let name = getActiveSoftwareName(panel)
    if(name === ""){
      toast.error("No need to active",TOAST_CONF)
      return
    }

    setConfirmMessage(`Firmware ${name} will be actived!`)
    setConfirmType("active")
    setShowConfirm(true)
    setActionPanel(panel)
  }
  const handleCommit = (panel)=>{
    console.log("handleCommit     panel = ",panel)
    let name = getCommitSoftwareName(panel)
    if(name === ""){
      toast.error("No need to commit",TOAST_CONF)
      return
    }
    
    setConfirmMessage(`Firmware ${name} will be commited!`)
    setConfirmType("commit")
    setShowConfirm(true)
    setActionPanel(panel)
  }

  const handleMigrate = (panel)=>{
    console.log("handleMigrate     panel = ",panel)
    if(targetVersion === ""){
      toast.error("Please input the target version!",TOAST_CONF)
      return
    }

    setConfirmMessage(`You will generate a new migration file to ${targetVersion} !`)
    setConfirmType("migrate")
    setShowConfirm(true)
    setActionPanel(panel)
  }

  const handleUploadMigration = (panel)=>{
    console.log("handleUploadMigration     panel = ",panel)
    // if(targetVersion === ""){
    //   toast.error("No need to migrate",TOAST_CONF)
    //   return
    // }

    setConfirmMessage(`You will upload the new migration file to the OLT !`)
    setConfirmType("migrate_upload")
    setShowConfirm(true)
    setActionPanel(panel)
  }
  const sendActive = (panel)=>{
    let name = getActiveSoftwareName(panel)
    
    let data = {
      oltId: panel.ip,
      dstPort:panel.port,
      action: "active",
      name: name,
    }
    AXIOS
      .put(API_SoftwareAction,data)
      .then((res) => {
        // console.log(res)
        if(res.data.status === GLOBAL.ERROR_NUM.Success){
          toast.success("Software active success!",TOAST_CONF)
        }else if (res.data.status === GLOBAL.ERROR_NUM.ActionFailed){
          toast.error("Software active Faild, please check and retry!",TOAST_CONF)
        }
      })
      .catch((err) => {
    });
  }
  const sendCommit = (panel)=>{
    let name = getCommitSoftwareName(panel)
    
    let data = {
      oltId: panel.ip,
      dstPort:panel.port,
      action: "commit",
      name: name,
    }
    AXIOS
      .put(API_SoftwareAction,data)
      .then((res) => {
        console.log(res)
        if(res.data.status === GLOBAL.ERROR_NUM.Success){
          toast.success("Software commit success!",TOAST_CONF)
        }else if (res.data.status === GLOBAL.ERROR_NUM.ActionFailed){
          toast.error("Software commit Faild, please check and retry!",TOAST_CONF)
        }
      })
      .catch((err) => {
    });
  }
  const onConfirm = (data, type)=>{
    // console.log("onConfirm  data = ",data)
    // console.log("onConfirm  type = ",type)
    switch (type){
      case "active":
        sendActive(data)
        break;
      case "commit":
        sendCommit(data)
        break;
      case "download":
        sendUploading(data)
        break;
      case "migrate":
        sendMigrate(data)
        break;
      case "migrate_upload":
        sendMigrateUpload(data)
        break;
      case "active_migrate":
        sendActiveWithMigration(data)
        break;
      default:
          console.log("onConfirm  no matched type = ",type)
    }
    setActionPanel(undefined);
    setConfirmMessage("");
    setShowConfirm(false);
    setConfirmType("");
  }
  const renderVersionTitle = (panel)=>{
    return(
      <div className='row'>
        <Typography style={{fontWeight:"700"}}> {"Version"} </Typography>
        <Typography style={{fontWeight:"700"}}> {"Release"} </Typography>
        <Typography style={{fontWeight:"700"}}> {"Valid"} </Typography>
        <Typography style={{fontWeight:"700"}}> {"Active"} </Typography>
        <Typography style={{fontWeight:"700"}}> {"Committed"} </Typography>
        <Typography style={{fontWeight:"700"}}> {"Download Timestamp"} </Typography>
        {/* {panel.download_status === "in-progress" ?null:<Typography style={{fontWeight:"700"}}> {"Download Status"} </Typography>} */}
        <Typography style={{fontWeight:"700"}}> {"Download Status"} </Typography>
      </div>
    )
  }
  const renderVersion1 = (panel)=>{
    // console.log("renderVersion1   panel = ",panel)
    return(
      panel ===""|| panel ===undefined || panel.version1.name ===""
      ?<></>
      :<div className='row'>
        <Typography> {panel.version1.name} </Typography>
        <Typography> {panel.version1.release} </Typography>
        <Typography> {panel.version1.valid ?"Yes":"No"} </Typography>
        <Typography> {panel.version1.active ?"Yes":"No"} </Typography>
        <Typography> {panel.version1.commit ?"Yes":"No"} </Typography>
        <Typography> {panel.version1.timestamp} </Typography>
        <Typography> {panel.download_result} </Typography>
      </div>
    )
  }
  const renderVersion2 = (panel)=>{
    return(
      panel ===""|| panel ===undefined || panel.version2.name ===""
      ?<></>
      : <div className='row'>
        <Typography> {panel.version2.name} </Typography>
        <Typography> {panel.version2.release} </Typography>
        <Typography> {panel.version2.valid ?"Yes":"No"} </Typography>
        <Typography> {panel.version2.active ?"Yes":"No"} </Typography>
        <Typography> {panel.version2.commit ?"Yes":"No"} </Typography>
        <Typography> {panel.version2.timestamp} </Typography>
        <Typography> {panel.download_result} </Typography>
      </div>
    )
  }
  // const renderExpandedHeader = ()=>{
  //   return(
  //     <Typography style={{ color: (disablePanel) ? DISABLED_SUBTEXT_COLOR : SUBTEXT_COLOR }} variant="BODY">
  //      OLT Version Operation
  //      </Typography>
  //   )
  // }
  // const renderSoftwareStatus = (panel)=>{
  //   return(
  //     <Typography style={{ width: '5%' }}>
  //       {
  //       isExpanded(panel.id)
  //       ?""
  //       :panel.download_status === "in-progress"?<DownloadIcon></DownloadIcon>:""
  //       }
  //     </Typography>
  //   )
  // }
  const renderMigrationUploadingStatus = (panel)=>{
  console.log("renderMigrationUploadingStatus panel = ",panel)
    if (panel.config_download_status === "config-preparing"){
      return (
        <div className='row3'>
          <DownloadIcon ></DownloadIcon>
          <Typography variant="CAPTION"  style={{marginLeft:"1.5rem"}}>
            Migration file donwloading, please wait!
          </Typography>
        </div>
        )
    }else if (panel.config_download_status === "failed"){
      return(
        <Typography variant="CAPTION" style={{marginLeft:"1.5rem"}}>
          Migration file download failure reason: {panel.config_download_result}
        </Typography>
      )
    }else if (panel.config_download_status === "idle"){
      return(
        <div className='row3'>
          <StopCircleOutlineIcon ></StopCircleOutlineIcon>
          <Typography variant="CAPTION"  style={{marginLeft:"1.5rem"}}>
            Migration file download process is not running! The last process is : {panel.config_download_result}
          </Typography>
        </div>
      )
    }else{
      return(
        <></>
      )
    }
  }

 const renderSoftwareStatus = (panel)=>{
  // console.log("renderSoftwareStatus panel = ",panel)
    if (panel.download_status === "in-progress"){
      return (
        <div className='row3'>
          <DownloadIcon ></DownloadIcon>
          <Typography variant="CAPTION"  style={{marginLeft:"1.5rem"}}>
            Software donwloading, please wait!
          </Typography>
        </div>
        )
    }else if (panel.download_status === "failed"){
      return(
        <Typography variant="CAPTION" style={{marginLeft:"1.5rem"}}>
          Current download failure reason: {panel.download_result}
        </Typography>
      )
    }else if (panel.download_status === "idle"){
      return(
        <div className='row3'>
          <StopCircleOutlineIcon ></StopCircleOutlineIcon>
          <Typography variant="CAPTION"  style={{marginLeft:"1.5rem"}}>
            Current download process is not running! The last process is : {panel.download_result}
          </Typography>
        </div>
      )
    }else{
      return(
        <></>
      )
    }
  }
  return (
    <ExpansionPanelsContainer>
      {data.map(panel => (
        <ExpansionPanel
          disabled={disablePanel}
          selected={selectPanel}
          expanded={isExpanded(panel.id)}
          style={{marginBottom:"0.5rem"}}
          key={panel.id}
          id={panel.id}
          {...otherProps}
        >
          <ExpansionPanelHeader
            data-test="header"
            ref={headerRef}
            role="button"
            aria-expanded={expanded === panel.id}
            style={{ justifyContent: 'space-between',minHeight:"9rem" }}
            onKeyDown={handleExpansion(panel.id)}
            onClick={handleExpansion(panel.id)}
          >
            <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center',minHeight:"9rem", width: '95%' }}>
              <ExpansionPanelTitle style={{width:"6rem",color:COLOR.Nokia_Blue,fontWeight:"700"}}>{panel.ip + "("+ panel.hostname+")"}</ExpansionPanelTitle>
              <div style={{ width: '85%' }}>
                <div style={{ width: '100%', marginBottom:"0.3rem"}}>
                    {renderVersionTitle(panel)}
                </div>
                <div style={{ width: '100%',marginBottom:"1rem" }}>
                    {/*isExpanded(panel.id) ? renderExpandedHeader() : */renderVersion1(panel)}
                </div>
                    {isExpanded(panel.CardName)&&<HorizontalDivider style={{background:"#EFEFEF"}}/>}
                <div style={{ width: '100%' }}>
                    {/*isExpanded(panel.id) ? '' : */renderVersion2(panel)}
                </div>
              </div>
              {/* {renderSoftwareStatus(panel)} */}

            </div>
            <ExpansionPanelButton
              onClick={handleExpansion(panel.id)}
              style={{ width:"5%" }}
              iconButtonProps={{
                'aria-hidden': true,
              }}
            />
          </ExpansionPanelHeader>
          <ExpansionPanelBody style={{ marginTop: '1rem',marginBottom: '1rem' }}
          >
            <div style={{ marginLeft: "1rem"}}>
              {renderMigrationUploadingStatus(panel)}
            </div>
            <div style={{ display: 'flex', marginLeft: '0.5rem', padding: PADDING,justifyContent: 'space-between' }} className="row">
              <div className='row' style={{ marginRight: '0.9375rem', overflow: 'hidden' }}>
                <div style={{ width: '29.375rem',marginRight: '0.9375rem'}}>
                  <Label >
                    <TextInputLabelContent required>{"Target Migration Version"}</TextInputLabelContent>
                    <Tooltip placement="top" trigger={["hover", "focus" ]} fallbackPlacements={["bottom", "right" ]} modifiers={{offset: {offset: "[0, 8]"}}}
                      tooltip='Input the target version of the software to be used by OLT.'>
                      <LabelHelpIcon>
                          <HelpCircleIcon />
                      </LabelHelpIcon>
                    </Tooltip>
                  </Label>
                  <TextInput placeholder="25.12" id={panel.active_version} value={targetVersion} onChange={(event) => {setTargetVersion(event.target.value)}} aria-label={panel.active_version} variant="underlined" />
                </div>
              </div>
              <ButtonsRow>
              {/* <Button>{t('button.cancel')}</Button> */}
              <Button 
                onClick={() => {
                  handleMigrate(panel);
                }}
                >{t('button.migrate')}
              </Button>
              <Button 
                onClick={() => {
                  handleUploadMigration(panel);
                }}
                >{t('button.upload')}
              </Button>
            </ButtonsRow>
            </div>
            <HorizontalDivider/>
            <div style={{ marginLeft: "1rem"}}>
              {renderSoftwareStatus(panel)}
            </div>
            <div style={{ display: 'flex', marginLeft: '0.5rem', padding: PADDING,justifyContent: 'space-between' }} className="row">
              <div className='row' style={{ marginRight: '0.9375rem', overflow: 'hidden' }}>
                <div style={{ width: '29.375rem',marginRight: '0.9375rem'}}>
                  <Label >
                    <TextInputLabelContent required>{"OLT Software File Path"}</TextInputLabelContent>
                    <Tooltip placement="top" trigger={["hover", "focus" ]} fallbackPlacements={["bottom", "right" ]} modifiers={{offset: {offset: "[0, 8]"}}}
                      tooltip='Input the full path of the software separated by "/". '>
                      <LabelHelpIcon>
                          <HelpCircleIcon />
                      </LabelHelpIcon>
                    </Tooltip>
                  </Label>
                  <TextInput placeholder="eg:lightspan_2203.038/L6GQAG" id={panel.active_version} value={softwarePath} onChange={(event) => {setSoftwarePath(event.target.value)}} aria-label={panel.active_version} variant="underlined" />
                </div>
                <div >
                  <Label >
                    <TextInputLabelContent required >{"OLT Software File Name"}</TextInputLabelContent>
                  </Label>
                  <TextInput placeholder="eg:L6GQAG2203.038" id={panel.active_version} value={softwareName} onChange={(event) => {setSoftwareName(event.target.value)}} aria-label={panel.active_version} variant="underlined" />
                </div>
              </div>
              <Button onClick={() => {handleUploading(panel);}}
                >
                <ButtonIcon>
                  <UploadIcon />
                </ButtonIcon>
                <ButtonText aria-label="Add active_version">{t('button.upload')}</ButtonText>
              </Button>
            </div>
            <HorizontalDivider style={{ marginLeft:  "1rem", marginRight: "1rem" }}/>
            <ButtonsRow>
              {/* <Button>{t('button.cancel')}</Button> */}
              <Button 
                onClick={() => {
                  handleActiveWithMigrate(panel);
                }}
                >{t('button.active_with_migrate')}
              </Button>
              <Button 
                onClick={() => {
                  handleActive(panel);
                }}
                >{t('button.active')}
              </Button>
              <Button 
                onClick={() => {
                  handleCommit(panel);
                }}
                >{t('button.commit')}
              </Button>
            </ButtonsRow>
          </ExpansionPanelBody>
        </ExpansionPanel>
      ))}
      <ConfirmDialog
        open={showConfirm}
        body={confirmMessage}
        data={actionPanel}
        type={confirmType}
        onCancel={
          ()=>{
            setActionPanel(undefined);
            setConfirmMessage("");
            setShowConfirm(false);
            setConfirmType("");
          }
        }
        onConfirm={onConfirm}
        >
        </ConfirmDialog>
    </ExpansionPanelsContainer>
  );
};
ExpansionPanels.propTypes = {
    data:PropTypes.array.isRequired,
    softwareServer:PropTypes.object.isRequired,
    onChange:PropTypes.func.isRequired,
    style:PropTypes.object,
};

ExpansionPanels.defaultProps = {
    data:[],
    softwareServer:{},
    style:{ },
}


export default ExpansionPanels;