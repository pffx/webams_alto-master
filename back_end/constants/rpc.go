package constants

const (

	//rpc action type
	OPERATION_DELETE  = "remove"
	OPERATION_DELETE2 = "delete"
	OPERATION_MERGE   = "merge"

	//defination of ont name in rpc start
	UNI_LAN  = "UNI_LAN"
	UNI_TEL  = "UNI_TEL"
	UNI_VOIP = "UNI_VOIP"
	UNI_10G  = "UNI_10G"
	NNI_VEIP = "NNI_VEIP"
	UNI_25G  = "UNI_25G"

	SLASH_UNI_LAN  = "/UNI_LAN"
	SLASH_UNI_TEL  = "/UNI_TEL"
	SLASH_UNI_VOIP = "/UNI_VOIP"
	SLASH_UNI_10G  = "/UNI_10G"
	SLASH_UNI_25G  = "/UNI_25G"
	SLASH_NNI_VEIP = "/NNI_VEIP"
	BAR_ONT        = "_ONT"
	//defination of ont name in rpc end

	//rpc name space defination start
	XMLNS_VERSION = "urn:ietf:params:xml:ns:netconf:base:1.0"
	//rpc name space defination end

	//defination of specific paramters in Restful API  start
	RESTFUL_ONU_DETAIL_RSSI            = "rssi"
	RESTFUL_ONU_DETAIL_SOFTWARE        = "software"
	RESTFUL_ONU_DETAIL_PORT_MACLIST    = "port_maclist"
	RESTFUL_ONU_DETAIL_802DOT1x_STATUS = "8021x"
	RESTFUL_ONU_DETAIL_OPERAT_STATUS   = "operational_status"
	RESTFUL_ONU_DETAIL_PORTS_STATUS    = "ports_status"
	RESTFUL_ONU_DETAIL_IPV4_ADDR       = "ipv4_addr"
	RESTFUL_ONU_DETAIL_IPV6_ADDR       = "ipv6_addr"
	//defination of specific paramters in Restful API  end

	//defination of OLT specific paramters in Restful API  start
	RESTFUL_OLT_DETAIL_OPERAT_STATUS         = "operational_status"
	RESTFUL_OLT_DETAIL_CPU                   = "cpu"
	RESTFUL_OLT_DETAIL_8021x_AUTH_SERVER     = "auth_server"
	RESTFUL_OLT_DETAIL_8021x_ACCOUNT_SERVER  = "account_server"
	RESTFUL_OLT_DETAIL_8021x_AUTH_SERVER_CFG = "auth_server_cfg"
	RESTFUL_OLT_DETAIL_NT                    = "nt"
	RESTFUL_OLT_DETAIL_HARDWARE              = "hw"
	RESTFUL_OLT_DETAIL_INTERFACE             = "interface"
	RESTFUL_OLT_DETAIL_LAG                   = "lag"
	//defination of OLT specific paramters in Restful API  end
)
