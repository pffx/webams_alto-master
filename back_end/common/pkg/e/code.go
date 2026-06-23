package e

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400
	//********************************************************
	//Common Error
	ERROR_EXIST_TAG     = 10001
	ERROR_ACTION_FAILED = 10002
	//********************************************************
	//Error for authentication
	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
	ERROR_PERMISSION               = 20005
	//********************************************************
	//Error for SSH and template
	ERROR_SSH_CREATE_FAIL      = 30001
	ERROR_TEMPLATE_CREATE_FAIL = 30002
	ERROR_SSH_CLI_EXEC_FAIL    = 30003
	//********************************************************
	ERROR_FIND_OLTS_FAIL = 40001
	ERROR_NO_FILE        = 40002
	//Error for netconf
	ERROR_NETCONF_EXEC_FAIL = 50001
)
