//const API_URI = process.env.NODE_ENV === "production"? 'http://'+localIp+':5600': 'http://127.0.0.1:5600';
const localIp = window.location.host.split(':')[0];
export const API_URI ='http://'+localIp+':5600';
export const API_ALTOV1 ="/nokia-alto/v1";
export const API_ALTO = "/nokia-alto";

// export const SocketURI = process.env.NODE_ENV === "production"? 'ws://'+localIp+':5500/ws': 'ws://127.0.0.1:5600/ws';
export const SocketURI =  'ws://'+localIp+':5600/ws';
export const API_GROUP = {
  USER:"/user",
  AUTH:"/auth",
  SOFTWARE:"/software",
  SYSTEM:"/system",
  TEST:"/test",
};
//Auth
export const API_Login="nokia-alto/auth/login"
export const API_Logout="nokia-alto/auth/logout"
export const API_CheckToken="nokia-alto/auth/token"
//System
export const API_OltList="nokia-alto/v1/system/oltList"
export const API_NTGATEWAY="nokia-alto/v1/system/nt_gw"
//Software
export const API_OltSoftwareInfo="nokia-alto/v1/software/olt/software"
export const API_ServerSoftware="nokia-alto/v1/software/olt/server_software"
export const API_SoftwareAction="nokia-alto/v1/software/olt/software"
export const API_SoftwareMigrationUpload="nokia-alto/v1/software/olt/software/migration_upload"
export const API_SoftwareMigration="nokia-alto/v1/software/olt/software/migration"
//Provisioning
export const API_ConfigFile="nokia-alto/v1/provisioning/olt/initial_file"
export const API_ProvisionService="nokia-alto/v1/provisioning/service/services"
export const API_ProvisionOltBackup="nokia-alto/v1/provisioning/olt/backup"
export const API_ProvisionOltRestore="nokia-alto/v1/provisioning/olt/restore"
export const API_ProvisionOltReset="nokia-alto/v1/provisioning/olt/reset"
export const API_ProvisionOltResetAll="nokia-alto/v1/provisioning/olt/reset_all"
export const API_ProvisionOltPing="nokia-alto/v1/provisioning/olt/ping"
export const API_ProvisionOltGuiDeploy="nokia-alto/v1/provisioning/olt/mf_gui"
export const API_ProvisionOltGuiUndeploy="nokia-alto/v1/provisioning/olt/mf_gui_un"
//WebSocket
export const WS_CMD=SocketURI+"/cmd"

//Test
export const API_BasicInfo="nokia-alto/v1/test/basicInfo"
