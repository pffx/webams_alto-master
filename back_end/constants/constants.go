package constants

const (
	//ont pon type start
	XGSPON  = "XGSPON"
	GPON    = "GPON"
	GSPON25 = "25GSPON"

	BASE_OLT_RELEASE = "2412"

	COMMA = ","

	HARDWARE_INFO_LIST_PATH = "./conf/onu_hardware_infos.xlsx"

	//Tennant start
	NOKIA_ORG_CODE = "A01"
	//Tennant end

	//This items corresponds to the items in the database
	BRIDGE_MODE = "FALSE"
	ROUTER_MODE = "TRUE"
	POE_MODE    = "TRUE"
	TYPE_C      = "TRUE"
	//ont pon type end

	//OLT capacity  part start
	MAX_LT_NUMBER = 14
	LT_PORT_START = 832 // invalaid port
	DEBUG_PORT    = "923"
	DF_DEBUG_PORT = "2222"
	COMMAND_PORT  = "940"

	DF_PORT   = "830"
	IHUB_PORT = "831"
	NT_PORT   = "832"
	LT1_PORT  = "833"
	LT2_PORT  = "834"
	LT3_PORT  = "835"
	LT4_PORT  = "836"
	LT5_PORT  = "837"
	LT6_PORT  = "838"
	LT7_PORT  = "839"
	LT8_PORT  = "840"
	LT9_PORT  = "841"
	LT10_PORT = "842"
	LT11_PORT = "843"
	LT12_PORT = "844"
	LT13_PORT = "845"
	LT14_PORT = "846"
	//OLT capacity part end

	DEBUG_PORT_USER = "root"
	DEBUG_PORT_PWD  = "2x2=4"

	//ONU deployment status in DB start
	IMPORTED                = "imported"
	DISCOVERED              = "discovered"
	DISCOVERED_AND_IMPORTED = "discovered_and_imported"
	//ONU deployment status in DB end
	//ploam status
	UNDETECTED     = "undetected"
	NOT_ACTIVATED  = "notActivated"
	LOSS_OF_SIGNAL = "lossOfSignal"
	DYING_GASP     = "dyingGasp"
	ONLINE         = "online"
	//onu scan
	LT_BOARD_IS_PLANNED = 1

	//template path start
	TEMPLATE_OLT          = "olt/"
	TEMPLATE_ONU          = "onu/"
	TEMPLATE_SERVICE      = "service/"
	TEMPLATE_SERVICE_PREF = "./template/service/"
	TEMPLATE_HGW          = "hgw/"
	TEMPLATE_HGW_PREF     = "./template/hgw/"
	TEMPLATE_CFG          = "cfg/"
	TEMPLATE_CFG_PREF     = "./template/cfg/"

	TEMPLATE_OLT_FS = "template/olt/"
	TEMPLATE_ONU_FS = "template/onu/"
	TEMPLATE_FS     = "template/"
	//template path end

	//log module defination start
	LOG_MOD_ONT    = "ONT"
	LOG_MOD_OLT    = "OLT"
	LOG_MOD_SYSTEM = "SYSTEM"
	//log module defination end

	//define chain
	CHAIN_DIR                        = "/mnt/persistent/ONT/chain_index"
	CHAIN_DB_NODES_DIR               = "/mnt/persistent/ONT/chain_index/nodes"
	CHAIN_OLT_BACKUP_FILE_DIR        = "/mnt/persistent/ONT/chain_index/olt"
	CHAIN_OLT_USER_DATA_DIR          = "/mnt/persistent/ONT/chain_index/userData"
	CHAIN_KOGEN_MASTER_KEY_DIR       = "/mnt/persistent/ONT/chain_index/masterKey"
	CHAIN_KOGEN_MASTER_KEY_FILE_PATH = "/mnt/persistent/ONT/chain_index/masterKey/masterKey"
	CHAIN_KOGEN_MASTER_KEY_FILE_NAME = "masterKey"
	//ONT config file path
	ONT_CONFIG_FILE_DIR = "/mnt/persistent/ONT"

	ONU_STATE_WAIT_FOR_ONLINE = "bbf-xpon-onu-types:onu-present-and-no-v-ani-known-and-in-o5"
	ONU_STATE_ONLINED         = "bbf-xpon-onu-types:onu-present-and-on-intended-channel-termination"

	HISTORY_ORDER_STATUS_USED    = "deployed"
	HISTORY_ORDER_STATUS_DELETED = "deleted"

	OLT_SOFTWARE_DOWNLOAD = "application_software"
	ONU_SOFTWARE_DOWNLOAD = "ont_software"

	TYPEB_STATUS_PROTECTED     = "type-b-protected"
	TYPEB_STATUS_NOT_PROTECTED = "type-b-not-protected-channel-pair-in-service"
	TYPEB_STATUS_NOT_DEPLOYED  = "type-b-not-deployed"

	TYPEB_STATUS_ACTIVE           = "active"
	TYPEB_STATUS_AVAIL_NOT_ACTIVE = "not-active-but-available"
	TYPEB_STATUS_NOT_AVAIL        = "not-available"

	CVC_TYPE_HYBRID = "1:1"
	CVC_TYPE_VOICE  = "N:1"
	CVC_U2U_ENABLE  = "true"
	CVC_U2U_DISABLE = "false"

	ALARM_MAJOR  = "major"
	ALARM_MINOR  = "minor"
	ALARM_CLEARD = "cleared"

	SFP_NOKIA_GPON     = "NOKIA_GPON"
	SFP_NOKIA_NGPON    = "NOKIA_NGPON"
	SFP_NOKIA_MPM      = "NOKIA_MPM"
	SFP_NOKIA_25GSPON  = "NOKIA_25GSPON"
	SFP_NOKIA_ETHERNET = "NOKIA_ETHERNET"

	DISTANCE_CHECKER = 15

	ONU_ID_MIX_NUMBER     = 0
	ONU_ID_MIN_NUMBER_STR = "0"
	ONU_ID_MAX_NUMBER     = 127
	ONU_ID_MAX_NUMBER_STR = "127"

	KOGEN_OPID_DEFAULT = "NKBB"
	KOGEN_OPID_JAPAN   = "JPDX"

	SERVICE_NOT_DEPLOYED = "notDeployed"
	SERVICE_DEPLOYED     = "deployed"

	MASTER_KEY_NAME               = "master_key"
	SERIVICES_DB_BACKUP_PREFIX    = "services_"
	ONUS_DB_BACKUP_PREFIX         = "onus_"
	CUSTOMERS_DB_BACKUP_PREFIX    = "customers_"
	NODES_DB_BACKUP_PREFIX        = "nodes_"
	USERS_DB_BACKUP_PREFIX        = "users_"
	ORGS_DB_BACKUP_PREFIX         = "orgs_"
	TYPEBS_DB_BACKUP_PREFIX       = "typebs_"
	DB_LOCAL_STORE_DECOMPRESS_DIR = "tmp/tmpbackup/"
	DB_LOCAL_STORE_DIR            = "tmp/"

	OLT_DB_ONLINE_STATUS           = "ONLINE"
	OLT_DB_OFFLINE_STATUS          = "OFFLINE"
	MIGRATION_TOOL_EXTRATED_PREFIX = "migrationTools/xsl/extracted/"
	MIGRATION_TOOL_XSL_PREFIX      = "migrationTools/xsl/"

	MIGRATION_TOOL_EXTRATED                 = "extracted/"
	ELECTRON_MIGRATION_TOOL_EXTRATED_PREFIX = "resources/migrationTools/xsl/extracted/"
	ELECTRON_MIGRATION_TOOL_XSL_PREFIX      = "resources/migrationTools/xsl/"
)
