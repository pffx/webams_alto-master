package rpc

// Some common functions unrelated to specific business operations.
//

import (
	"alto_server/common/gcfg"
	logger "alto_server/common/log"
	ncssh "alto_server/common/netconf2/transport/ssh"
	"alto_server/common/templateFS"
	"alto_server/common/utils"

	// "alto_server/daemon/netconfSessionPool"
	"context"
	"errors"
	"strings"

	// "alto_server/common/templateFS"
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"text/template"

	netconf "alto_server/common/netconf2"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
)

//***********************************************************************************
//Common function for rpc operation start *******************************************
//***********************************************************************************

// Execute a tpl file.
// the tpl file path is defined in RPCInfor
// the input param will be used to replace the {{.XX}} in tpl file
// func RunStaticRPCWrapper(rpcInfor RPCInfor, param any, modLogger *logrus.Logger, logRecord bool) (*netconf.RPCReply, error) {
// 	_, err := os.Stat(rpcInfor.TemplatePath)
// 	if err == nil {
// 		return RunStaticRPCWithRealPath(rpcInfor, param, modLogger, logRecord)
// 	} else {
// 		return RunStaticFSRPC(rpcInfor, param, modLogger, logRecord)
// 	}
// }

func RunStaticRPCWithRealPath(rpcInfor RPCInfor, param any, modLogger *logrus.Logger, logRecord bool) (*netconf.RPCReply, error) {
	var gcfg = gcfg.GetGlobalCfg()
	tmpl, err := template.ParseFiles(rpcInfor.TemplatePath)
	if err != nil {
		modLogger.Debug("create template failed, path is :", rpcInfor.TemplatePath)
		modLogger.Debug("create template failed, err:", err.Error())
		fmt.Println("RunStaticRPCWithRealPath create template file failed, err:", err)
		return nil, err
	}
	buf := new(bytes.Buffer)
	tmpl.Execute(buf, param)
	if gcfg.RPCLogOutput {
		fmt.Println("****************************************")
		fmt.Println(buf.String())
		fmt.Println("****************************************")
	}
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	session, err := StartSessionCreate(rpcInfor.IP, rpcInfor.Port)
	if err != nil {
		modLogger.Debug("Dial SSH failed, err:", err)
		fmt.Println("Dial SSH failed, err:", err)
		return nil, err
	}
	defer session.Close(ctx)

	reply, err := session.Do(ctx, (buf.String()))
	if err != nil {
		fmt.Println("exec template failed, err:", err)

		modLogger.Error(fmt.Sprint("RPC IP:", rpcInfor.IP, " Port:", rpcInfor.Port, " template:", rpcInfor.TemplatePath))
		modLogger.Error("request:", buf.String())
		modLogger.Error("exec template failed, err:", err.Error())

		if reply != nil && gcfg.RPCLogOutput {
			fmt.Println(reply.Data)
		}
		return nil, err
	} else {
		if logRecord {
			//need record the failure case no need check the record flag
			modLogger.Info(fmt.Sprint("RPC IP:", rpcInfor.IP, " Port:", rpcInfor.Port, " template:", rpcInfor.TemplatePath))
			modLogger.Info("request:", buf.String())
			modLogger.Info("response:", reply.Data)
		}
		if reply != nil && gcfg.RPCLogOutput {
			fmt.Println(reply.Data)
		}
	}
	if err := reply.Err(netconf.SevError); err != nil {
		return reply, err
	}
	return reply, nil
}
func RunStaticRPC(rpcInfor RPCInfor, param any, modLogger *logrus.Logger, logRecord bool) (*netconf.RPCReply, error) {
	var gcfg = gcfg.GetGlobalCfg()
	tmpl, err := template.ParseFiles(utils.GetTemplatePath(rpcInfor.TemplatePath))
	if err != nil {
		modLogger.Debug("create template failed, path is :", rpcInfor.TemplatePath)
		modLogger.Debug("create template failed, err:", err.Error())
		fmt.Println("RunStaticRPC create template file failed, err:", err)
		return nil, err
	}
	buf := new(bytes.Buffer)
	tmpl.Execute(buf, param)
	if gcfg.RPCLogOutput {
		fmt.Println("****************************************")
		fmt.Println(buf.String())
		fmt.Println("****************************************")
	}
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	session, err := StartSessionCreate(rpcInfor.IP, rpcInfor.Port)
	if err != nil {
		modLogger.Debug("Dial SSH failed, err:", err)
		fmt.Println("Dial SSH failed, err:", err)
		return nil, err
	}
	defer session.Close(ctx)
	reply, err := session.Do(ctx, (buf.String()))
	if err != nil {
		fmt.Println("exec template failed, err:", err)

		modLogger.Error(fmt.Sprint("RPC IP:", rpcInfor.IP, " Port:", rpcInfor.Port, " template:", rpcInfor.TemplatePath))
		modLogger.Error("request:", buf.String())
		modLogger.Error("exec template failed, err:", err.Error())

		if reply != nil && gcfg.RPCLogOutput {
			fmt.Println(reply.Data)
		}
		return nil, err
	} else {
		if logRecord {
			//need record the failure case no need check the record flag
			modLogger.Info(fmt.Sprint("RPC IP:", rpcInfor.IP, " Port:", rpcInfor.Port, " template:", rpcInfor.TemplatePath))
			modLogger.Info("request:", buf.String())
			modLogger.Info("response:", reply.Data)
		}
		if reply != nil && gcfg.RPCLogOutput {
			fmt.Println(reply.Data)
		}
	}
	if err := reply.Err(netconf.SevError); err != nil {
		return reply, err
	}
	return reply, nil
}
func RunStaticRPCWithExtParams(rpcInfor RPCInfor, param any, modLogger *logrus.Logger, logRecord bool, extParams string) (*netconf.RPCReply, error) {
	var gcfg = gcfg.GetGlobalCfg()
	tmpl, err := template.ParseFiles(utils.GetTemplatePath(rpcInfor.TemplatePath))
	if err != nil {
		modLogger.Debug("create template failed, path is :", rpcInfor.TemplatePath)
		modLogger.Debug("create template failed, err:", err.Error())
		fmt.Println("RunStaticRPC create template file failed, err:", err)
		return nil, err
	}
	buf := new(bytes.Buffer)
	tmpl.Execute(buf, param)

	//for extParams
	if strings.Compare(extParams, "") != 0 {
		var extParamsMap map[string]interface{}

		// 解析JSON字符串到map
		err = json.Unmarshal([]byte(extParams), &extParamsMap)
		if err != nil {
			modLogger.Error("parse ext param fail:", err.Error())
		} else {
			tmpl.Execute(buf, extParamsMap)
		}
	}
	if gcfg.RPCLogOutput {
		fmt.Println("****************************************")
		fmt.Println(buf.String())
		fmt.Println("****************************************")
	}

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	session, err := StartSessionCreate(rpcInfor.IP, rpcInfor.Port)
	if err != nil {
		modLogger.Debug("Dial SSH failed, err:", err)
		fmt.Println("Dial SSH failed, err:", err)
		return nil, err
	}
	defer session.Close(ctx)
	reply, err := session.Do(ctx, (buf.String()))
	if err != nil {
		fmt.Println("exec template failed, err:", err)

		modLogger.Error(fmt.Sprint("RPC IP:", rpcInfor.IP, " Port:", rpcInfor.Port, " template:", rpcInfor.TemplatePath))
		modLogger.Error("request:", buf.String())
		modLogger.Error("exec template failed, err:", err.Error())

		if reply != nil && gcfg.RPCLogOutput {
			fmt.Println(reply.Data)
		}
		return nil, err
	} else {
		if logRecord {
			//need record the failure case no need check the record flag
			modLogger.Info(fmt.Sprint("RPC IP:", rpcInfor.IP, " Port:", rpcInfor.Port, " template:", rpcInfor.TemplatePath))
			modLogger.Info("request:", buf.String())
			modLogger.Info("response:", reply.Data)
		}
		if reply != nil && gcfg.RPCLogOutput {
			fmt.Println(reply.Data)
		}
	}
	if err := reply.Err(netconf.SevError); err != nil {
		return reply, err
	}
	return reply, nil
}

func RunStaticFSRPC(rpcInfor RPCInfor, param any, modLogger *logrus.Logger, logRecord bool) (*netconf.RPCReply, error) {
	// tmpl, err := template.ParseFiles(UTILS.GetTemplatePath(rpcInfor.TemplatePath))
	var gcfg = gcfg.GetGlobalCfg()

	tmpl, err := template.ParseFS(templateFS.TemplateFS, rpcInfor.TemplatePath)
	if err != nil {
		modLogger.Debug("create template failed, file path is :", rpcInfor.TemplatePath)
		modLogger.Debug("create template failed, err:", err.Error())
		fmt.Println("RunStaticFSRPC create template file failed, err:", err)
		return nil, err
	}
	buf := new(bytes.Buffer)
	tmpl.Execute(buf, param)
	//logger.SystemLogger.Debug(buf.String())
	// fmt.Println("run tmpl, tmpl:", buf.String())
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	session, err := StartSessionCreate(rpcInfor.IP, rpcInfor.Port)
	if err != nil {
		modLogger.Debug("Dial SSH failed, err:", err)
		fmt.Println("Dial SSH failed, err:", err)
		return nil, err
	}
	defer session.Close(ctx)
	reply, err := session.Do(ctx, (buf.String()))
	if err != nil {
		fmt.Println("exec static FS template failed, err:", err)

		modLogger.Error(fmt.Sprint("RPC IP:", rpcInfor.IP, " Port:", rpcInfor.Port, " template:", rpcInfor.TemplatePath))
		modLogger.Error("request:", buf.String())
		modLogger.Error("exec template failed, err:", err.Error())

		return nil, err
	} else {
		if logRecord {
			modLogger.Info(fmt.Sprint("RPC IP:", rpcInfor.IP, " Port:", rpcInfor.Port, " template:", rpcInfor.TemplatePath))
			modLogger.Info("request:", buf.String())
			modLogger.Info("response:", reply.Data)
		}
		if reply != nil && gcfg.RPCLogOutput {
			fmt.Println(reply.Data)
		}
	}
	if err := reply.Err(netconf.SevError); err != nil {
		return reply, err
	}
	return reply, nil
}

//Generate a tpl, and then execute it
//the input rpc is a struct.

func RunGeneratedRPC(rpc any, rpcInfor RPCInfor, modLogger *logrus.Logger, logRecord bool) (*netconf.RPCReply, error) {
	var gcfg = gcfg.GetGlobalCfg()
	output, err := xml.MarshalIndent(rpc, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	tmpl, err := template.New("xx").Parse(string(output))
	if err != nil {
		modLogger.Error("create tpl failed, err:", err)
		fmt.Println("create tpl failed, err:", err)
		return nil, err
	}
	buf := new(bytes.Buffer)
	tmpl.Execute(buf, "xx")
	if gcfg.RPCLogOutput {
		fmt.Println("****************************************")
		fmt.Println(buf.String())
		fmt.Println("****************************************")
	}

	//show the detail rpc log in the screen
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	session, err := StartSessionCreate(rpcInfor.IP, rpcInfor.Port)
	if err != nil {
		modLogger.Debug("Dial SSH failed, err:", err)
		fmt.Println("Dial SSH failed, err:", err)
		return nil, err
	}
	defer session.Close(ctx)
	reply, err := session.Do(ctx, (buf.String()))
	if err != nil {
		fmt.Println("exec generated template failed, err:", err)

		modLogger.Error(fmt.Sprint("RPC IP:", rpcInfor.IP, " Port:", rpcInfor.Port, " template:", rpcInfor.TemplatePath))
		modLogger.Error("request:", buf.String())
		modLogger.Error("exec template failed, err:", err.Error())

		return nil, err
	} else {
		if logRecord {
			modLogger.Info(fmt.Sprint("RPC IP:", rpcInfor.IP, " Port:", rpcInfor.Port))
			modLogger.Info("request:", buf.String())
			modLogger.Info("response:", reply.Data)
		}
		if reply != nil && gcfg.RPCLogOutput {
			fmt.Println(reply.Data)
		}
	}
	if err := reply.Err(netconf.SevError); err != nil {
		return reply, err
	}
	return reply, nil
}

func RunStringRPC(rpcInfor RPCInfor, param string) (*netconf.RPCReply, error) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	session, err := StartSessionCreate(rpcInfor.IP, rpcInfor.Port)
	if err != nil {
		logger.SystemLogger.Debug("Dial SSH failed, err:", err)
		fmt.Println("Dial SSH failed, err:", err)
		return nil, err
	}
	defer session.Close(ctx)
	reply, err := session.Do(ctx, param)
	if err != nil {
		logger.SystemLogger.Debug("exec string template failed, err:", err)
		fmt.Println("exec string template failed, err:", err)
		return nil, err
	}
	logger.SystemLogger.Debug("exec template success :", reply.Data)
	return reply, nil
}
func StartSessionCreate(oltIp string, realPortNum string) (*netconf.Session, error) {
	sshConfig := &ssh.ClientConfig{
		User:            "admin",
		Auth:            []ssh.AuthMethod{ssh.Password("admin")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Config: ssh.Config{
			Ciphers: []string{"aes256-ctr"},
		},
	}

	sshConfig.SetDefaults()
	ctx := context.Background()

	transport, err := ncssh.Dial(ctx, "tcp", oltIp+":"+realPortNum, sshConfig)
	if err != nil {
		fmt.Println(oltIp + ":" + realPortNum + ":admin:admin StartSessionCreate Dial:" + err.Error())
		return nil, err
	}

	session, err := netconf.Open(transport)
	if err != nil {
		fmt.Println("StartSessionCreate Open:" + err.Error())
		return nil, err
	}

	return session, nil
}

//***********************************************************************************
//Common function for rpc operation end *********************************************
//***********************************************************************************

func UnmarshalReply(reply *netconf.RPCReply) (FilterRPCReply, error) {
	rpcReply := FilterRPCReply{}
	err := xml.Unmarshal([]byte(reply.Data), &rpcReply)
	if err != nil {
		return rpcReply, err
	}
	if reply.Data == "<data></data>" || reply.Data == "<data/>" {
		return rpcReply, errors.New("rpc reply is null")
	}
	return rpcReply, nil
}

//***********************************************************************************
//Common function by rpc request end *********************************************
//***********************************************************************************

func EnableAllDebug(dstIpAddr string, dstPort string, username string, password string) error {
	rpcInfo := RPCInfor{
		Account:      username,
		Password:     password,
		IP:           dstIpAddr,
		Port:         dstPort,
		TemplatePath: "enableLemiDebug.tpl",
	}

	_, err := RunStaticRPC(rpcInfo, "", logger.SystemLogger, false)

	if err != nil {
		fmt.Println("run template for enable debug port failed, err:", err)
		return err
	}
	return nil
	//logger.OltLogger.Info("Enable Debug port for OLT : " + dstIpAddr)
}
