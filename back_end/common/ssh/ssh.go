package ssh

import (
	//"fmt"

	"github.com/Juniper/go-netconf/netconf"
	"golang.org/x/crypto/ssh"
	//logger "alto_server/common/log"
)

var SSHSESSIONHANDLE netconf.Session

//var GLOBALNETCONFTRANSPORT netconf.Transport

// func init() {
// 	sshConfig := BuildConfig()

// 	sshSessionPtr, err := netconf.DialSSH("10.56.25.36:832", sshConfig)
// 	if err != nil {
// 		fmt.Println(err)
// 		//log.Fatal(err)
// 	}

// 	defer sshSessionPtr.Close()
// 	SSHSESSIONHANDLE.SessionID = sshSessionPtr.SessionID
// 	SSHSESSIONHANDLE.Transport = sshSessionPtr.Transport
// 	SSHSESSIONHANDLE.ServerCapabilities = sshSessionPtr.ServerCapabilities
// 	SSHSESSIONHANDLE.ErrOnWarning = sshSessionPtr.ErrOnWarning

// 	//GLOBALNETCONFTRANSPORT = sshSessionPtr.Transport
// 	fmt.Println("alto ssh init :", SSHSESSIONHANDLE.SessionID)
// 	logger.LoggerToErrorlog().Debug(SSHSESSIONHANDLE.SessionID)

// }

func GetSessionHandle() netconf.Session {
	return SSHSESSIONHANDLE
}

func BuildConfig() *ssh.ClientConfig {
	sshConfig := &ssh.ClientConfig{
		User:            "admin",
		Auth:            []ssh.AuthMethod{ssh.Password("Netconf$150")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return sshConfig
}
