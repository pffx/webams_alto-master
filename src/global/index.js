import { 
  POPPY_RED_300,
  PUMPKIN_ORANGE_300,
  CANARY_YELLOW_300,
  NEUTRAL_GREY_150,
  NOKIA_BLUE_300,
  APPLE_GREEN_300,
  SEA_GREEN_300,
  NOKIA_BLUE_500,
} from '@nokia-csf-uxr/freeform-design-tokens/tokens/colors';
import { toast } from 'react-toastify';
import Tokens from '@nokia-csf-uxr/freeform-design-tokens';
// import ic_warning_kpi from '@nokia-csf-uxr/ccfk/node_modules/@nokia-csf-uxr/ccfk-assets/icons/kpi_status_warning.svg';
// import ic_minor_kpi from '@nokia-csf-uxr/ccfk/node_modules/@nokia-csf-uxr/ccfk-assets/icons/kpi_status_minor.svg';
// import ic_major_kpi from '@nokia-csf-uxr/ccfk/node_modules/@nokia-csf-uxr/ccfk-assets/icons/kpi_status_major.svg';
// import ic_cleared_kpi from '@nokia-csf-uxr/ccfk/node_modules/@nokia-csf-uxr/ccfk-assets/icons/kpi_status_cleared.svg';
// import ic_status_up from '@nokia-csf-uxr/ccfk/node_modules/@nokia-csf-uxr/ccfk-assets/icons/arrow_line_up_circle.svg';
// import ic_status_down from '@nokia-csf-uxr/ccfk/node_modules/@nokia-csf-uxr/ccfk-assets/icons/arrow_line_down_circle.svg';
// import ic_critical_kpi from '@nokia-csf-uxr/ccfk/node_modules/@nokia-csf-uxr/ccfk-assets/icons/kpi_status_critical.svg';
import ic_warning_kpi from '@nokia-csf-uxr/ccfk-assets/icons/kpi_status_warning.svg';
import ic_minor_kpi from '@nokia-csf-uxr/ccfk-assets/icons/kpi_status_minor.svg';
import ic_major_kpi from '@nokia-csf-uxr/ccfk-assets/icons/kpi_status_major.svg';
import ic_cleared_kpi from '@nokia-csf-uxr/ccfk-assets/icons/kpi_status_cleared.svg';
import ic_status_up from '@nokia-csf-uxr/ccfk-assets/icons/arrow_line_up_circle.svg';
import ic_status_down from '@nokia-csf-uxr/ccfk-assets/icons/arrow_line_down_circle.svg';
import ic_critical_kpi from '@nokia-csf-uxr/ccfk-assets/icons/kpi_status_critical.svg';
// import kpi_status_indeterminate from '@nokia-csf-uxr/ccfk-assets/icons/kpi_status_indeterminate.svg';

// import { initializeConnect } from 'react-redux/es/components/connect';


// const localIp = window.location.host.split(':')[0];

// export const SocketURI = process.env.NODE_ENV === "production"? 'ws://'+localIp+':5500': 'ws://127.0.0.1:5500';
// export const SocketURI = process.env.NODE_ENV === "production"? 'ws://'+localIp+':5500': 'ws://127.0.0.1:5600/ws';


export const INITIALIAL_OLT_IP = "192.168.1.1"

export const COLOR={
  Background:"#eeeeee",
  Critical: POPPY_RED_300,
  Major: PUMPKIN_ORANGE_300,
  Minor: CANARY_YELLOW_300,
  Warning: NOKIA_BLUE_300,
  Indeterminate: NEUTRAL_GREY_150,
  Up:APPLE_GREEN_300,
  Down: POPPY_RED_300,
  Nokia_Blue: NOKIA_BLUE_500,
}

export const TableColorMap = {
  Cleared: SEA_GREEN_300,
  Critical: POPPY_RED_300,
  Down: POPPY_RED_300,
  Major: PUMPKIN_ORANGE_300,
  Minor: CANARY_YELLOW_300,
  UP: APPLE_GREEN_300,
  Warning: NOKIA_BLUE_300
}
export const TableIconMap = {
  Cleared: ic_cleared_kpi,
  Critical: ic_critical_kpi,
  Down: ic_status_down,
  Major: ic_major_kpi,
  Minor: ic_minor_kpi,
  UP: ic_status_up,
  Warning: ic_warning_kpi
}
export const SPACING = {
  SPACING_0 : Tokens.SPACING.SPACING_0,//"0rem",
  SPACING_2 : Tokens.SPACING.SPACING_2,//"0.125rem",
  SPACING_4 : Tokens.SPACING.SPACING_4,//"0.25rem",
  SPACING_8 : Tokens.SPACING.SPACING_8,//"0.5rem",
  SPACING_12 : Tokens.SPACING.SPACING_12,//"0.75rem",
  SPACING_16 : Tokens.SPACING.SPACING_16,//"1rem",
  SPACING_24 : Tokens.SPACING.SPACING_24,//"1.5rem",
  SPACING_32 : Tokens.SPACING.SPACING_32,//"2rem",
  SPACING_48 : Tokens.SPACING.SPACING_48,// "3rem",
  SPACING_64 : Tokens.SPACING.SPACING_64,//"4rem",
}

export const TOAST_CONF = {
  position: toast.POSITION.TOP_CENTER,
  autoClose: 2000,
  draggable: true,
}
export const ERROR_NUM = {
  Success: 200,
  ActionFailed: 10002,
}

export const LANGUAGES = [
  {key:'en',label:"English"},
  {key:'cn',label:"简体中文"},
];

export const CARD_PORT = {
  "nt":832,
  "lt":931,
  "shelf":933
}
export const OLT_TYPE_PORT = {
  "df":830,
  "ihub":831,
  "mf_lt0":832,//named NT card
  "mf_lt1":833,
  "mf_lt2":834,
  "mf_lt3":835,
  "mf_lt4":836,
  "mf_lt5":837,
  "mf_lt6":838,
  "mf_lt7":839,
  "mf_lt8":840,
  "mf_lt9":841,
  "mf_lt10":842,
  "mf_lt11":843,
  "mf_lt12":844,
  "mf_lt13":845,
  "mf_lt14":846,
  // "mf_lt15":847,
  // "mf_lt16":848,
}
export const OLT_PORT_NAME = {
  "830":"DF2",
  "832":"NT",//named NT card
  "833":"LT1",
  "834":"LT2",
  "835":"LT3",
  "836":"LT4",
  "837":"LT5",
  "838":"LT6",
  "839":"LT7",
  "840":"LT8",
  "841":"LT9",
  "842":"LT10",
  "843":"LT11",
  "844":"LT12",
  "845":"LT13",
  "846":"LT14",
  // "847":"LT15",
  // "848":"LT16",
}
export const CARD = [
  {label: 'Card', isHeader: true },
  {value:"nt",label:"NT"},
  {value:"lt",label:"LT"},
  {value:"shelf",label:"Shelf"},
]
export const OLT_TYPE = [
  {label: 'Type', isHeader: true },
  {value:"DF16",label:"DF16"},
  {value:"MF2",label:"MF2"},
  {value:"MF14",label:"MF14"},
]
export const OLT_LT_NUM = [
  {label: 'Number', isHeader: true },
  {value:1,label:"1"},
  {value:2,label:"2"},
  {value:14,label:"14"},
]

export const ALLOWED_NE = [
  "LS-DF-CFXR-E",
  "LS-DF-CFXR-H",
  "LS-FX-FGLT-D",
  "LS-FX-FWLT-C",
  "LS-MF-LWLT-C",
  "LS-MF-LMNT-B",
]
export default {
  // SocketURI,
  LANGUAGES,
  COLOR,
  TableColorMap,
  TableIconMap,
  CARD_PORT,
  OLT_TYPE_PORT,
  OLT_PORT_NAME,
  OLT_TYPE,
  OLT_LT_NUM,
  ERROR_NUM,
  TOAST_CONF,
  SPACING,
  CARD,
}