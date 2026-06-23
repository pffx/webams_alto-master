package olt

import (
	"alto_server/common/db"
	logger "alto_server/common/log"
	"alto_server/common/models"
	"alto_server/common/pkg/e"
	"alto_server/common/rpc"
	"alto_server/common/utils"
	"alto_server/constants"
	"bytes"
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"github.com/go-xmlfmt/xmlfmt"

	"github.com/Juniper/go-netconf/netconf"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/ssh"
)

// backupHandler godoc
// @BasePath /nokia-alto
// @Summary backup olt configuration
// @schemes http https
// @Description backup the configuration from indicated olt
// @Tags Provisioning
// @Accept x-www-form-urlencoded
// @Produce json
// @Param token header string true "登录信息"
// @Param oltId  formData    string true "olt id"
// @Param dstPort  formData    string true "port of olt card"
// @Success      200  {json}   {success}
// @Failure      20001  {string}  {Token鉴权失败}
// @Router /v1/provisioning/olt/backup [post]
func backupHandler(c *gin.Context) {
	oltId := c.DefaultPostForm("oltId", "192.168.1.1")
	logger.SystemLogger.Debug("===start backupHandler")
	// sshConfig := &ssh.ClientConfig{
	// 	User:            "admin",
	// 	Auth:            []ssh.AuthMethod{ssh.Password("admin")},
	// 	HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	// }

	dstPort := c.DefaultPostForm("dstPort", "830") // OLT type
	fmt.Println(dstPort)
	logger.SystemLogger.Debug(dstPort)
	//dstIpAddr := c.DefaultPostForm("oltId", "192.168.1.1")
	dstIpAddr := oltId
	fmt.Println(dstIpAddr)
	logger.SystemLogger.Debug(dstIpAddr)
	rpcInfo := rpc.RPCInfor{
		IP:           dstIpAddr,
		Port:         dstPort,
		TemplatePath: "backup.tpl",
	}
	reply, err := rpc.RunStaticRPC(rpcInfo, "", logger.SystemLogger, false)

	// s, err := netconf.DialSSHTimeout(dstIpAddr+":"+dstPort, sshConfig, 3*time.Second)
	// if err != nil {
	// 	logger.LoggerToErrorlog().Debug("Dial SSH failed, err:", err)
	// 	fmt.Println("Dial SSH failed, err:", err)
	// 	utils.RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "set fail", "data": "err response"})
	// }
	// defer s.Close()
	// tmpl, err := template.ParseFiles(utils.GetTemplatePath("backup.tpl"))
	// if err != nil {
	// 	logger.LoggerToErrorlog().Debug("create template failed, err:", err)
	// 	fmt.Println("create template failed, err:", err)
	// 	utils.RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "set fail", "data": "err response"})
	// }
	// buf := new(bytes.Buffer)
	// tmpl.Execute(buf, "xx")
	// logger.SystemLogger.Debug(buf.String())
	// reply, err := s.Exec(netconf.RawMethod(buf.String()))

	if err != nil {
		fmt.Println(err)
		if reply == nil {
			utils.RES(c, e.SUCCESS, gin.H{"status": e.ERROR_NETCONF_EXEC_FAIL, "message": "exec netconf fail", "data": "err response"})
			return
		}
		v := models.RpcError{}
		err = xml.Unmarshal([]byte(reply.Data), &v)
		if err != nil {
			fmt.Println("Unmarshal  failed, err:", err)
			logger.SystemLogger.Debug("Unmarshal  failed, err:", err)
			utils.RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "set fail", "data": "err response"})
			return
		}
		fmt.Println(v.ErrorAppTag)
		fmt.Println(v.ErrorMessage)
		//fmt.Println(v.ErrorInfo)
		logger.SystemLogger.Debug(v.ErrorAppTag)
		logger.SystemLogger.Debug(v.ErrorMessage)
		utils.RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "set fail", "data": "err response"})
		return
	} else {
		now := time.Now().Unix()
		tmpDirName := utils.GetSoftwarePath()
		tmpFileName := strconv.FormatInt(now, 10) + "_backup.xml"
		fmt.Printf("Reply: %+v", tmpDirName+tmpFileName)
		newStr := strings.Replace(reply.Data, "<data>", "", 1)
		newStr = strings.Replace(newStr, "</data>", "", 1)
		newStr = xmlfmt.FormatXML(newStr, "", "  ")
		err2 := ioutil.WriteFile(tmpDirName+tmpFileName, []byte(newStr), 0666) //写入文件(字节数组)
		if err2 != nil {
			utils.RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "set fail", "data": "err response"})
			return
		}

		//fmt.Printf("Reply: %+v", reply.Data)
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type, newtoken, Transfer-Encoding, Content-Transfer-Encoding, Content-Disposition")
		c.Header("Cache-Control", "no-cache")
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Disposition", "attachment; filename="+tmpFileName)
		c.Header("Content-Transfer-Encoding", "binary")
		//c.Header("Transfer-Encoding", "chunked")
		c.Header("Content-Length", "-1")
		c.File(tmpDirName + tmpFileName)
	}
}

// resetAllHandler godoc
// @BasePath /nokia-alto
// @Summary reset the configuration of all cards
// @schemes http https
// @Description reset the configuration of all cards
// @Tags Provisioning
// @Accept x-www-form-urlencoded
// @Produce json
// @Param token header string true "登录信息"
// @Param oltId  formData    string true "olt id"
// @Param OltPort  formData    string true "port of olt card"
// @Success      200  {json}   {success}
// @Failure      20001  {string}  {Token鉴权失败}
// @Router /v1/provisioning/olt/backup [post]

func resetReallyDo() {
	var dstPortArr = [3]string{"834", "833", "832"}
	for _, v := range dstPortArr {
		if v == "832" {
			time.Sleep(20 * time.Second)
		} else {
			time.Sleep(2 * time.Second)
		}
		fmt.Println(v)
		logger.SystemLogger.Debug(v)

		tmpl, err := template.ParseFiles(utils.GetTemplatePath("reset.tpl"))
		if err != nil {
			fmt.Println("create template failed, err:", err)
			logger.SystemLogger.Debug("get reset template failed, err:", err)
			continue
		}
		buf := new(bytes.Buffer)
		tmpl.Execute(buf, "xx")
		fmt.Printf("request: %+v", buf.String())
		logger.SystemLogger.Debug(buf.String())

		sshConfig := &ssh.ClientConfig{
			User:            "admin",
			Auth:            []ssh.AuthMethod{ssh.Password("admin")},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}

		s, err := netconf.DialSSHTimeout("192.168.1.1:"+v, sshConfig, 3*time.Second)
		if err != nil {
			fmt.Println("Dial SSH failed, err:", err)
			logger.SystemLogger.Debug("Dial SSH failed, err:", err)
			continue
		}

		// s, err := DialSShRPCTimeout("192.168.1.1:" + v)
		// if err != nil {
		// 	fmt.Println("Dial SSH failed, err:", err)
		// 	logger.LoggerToErrorlog().Debug("Dial SSH failed, err:", err)
		// 	continue
		// }
		defer s.Close()
		reply, err := s.Exec(netconf.RawMethod(buf.String()))
		if err != nil {
			v := models.RpcError{}
			err = xml.Unmarshal([]byte(reply.Data), &v)
			if err != nil {
				fmt.Println("Unmarshal  failed, err:", err)
				logger.SystemLogger.Debug("Unmarshal  failed, err:", err)
				continue
			}
			fmt.Println(v.ErrorAppTag)
			fmt.Println(v.ErrorMessage)
			fmt.Println(v.ErrorInfo)
			logger.SystemLogger.Debug(v.ErrorAppTag)
			logger.SystemLogger.Debug(v.ErrorMessage)
			logger.SystemLogger.Debug(v.ErrorInfo)
			continue
		} else {
			fmt.Println(reply.Data)
			logger.SystemLogger.Debug(reply.Data)
			if v == "832" {
				db.ResetLtInfoWhenActive()
			}

		}
	}
}
func resetAllHandler(c *gin.Context) {
	logger.SystemLogger.Debug("===start resetAllHandler")
	go resetReallyDo()
	utils.RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "Success", "data": "reply"})
}

// resetHandler godoc
// @BasePath /nokia-alto
// @Summary reset the oam configuration
// @schemes http https
// @Description reset the oam configuration for the indicated olt
// @Tags Provisioning
// @Accept x-www-form-urlencoded
// @Produce json
// @Param token header string true "登录信息"
// @Param oltId  formData    string true "olt id"
// @Param OltPort  formData    string true "port of olt card"
// @Success      200  {json}   {success}
// @Failure      20001  {string}  {Token鉴权失败}
// @Router /v1/provisioning/olt/backup [post]
func resetHandler(c *gin.Context) {
	logger.SystemLogger.Debug("===start resetHandler")
	var oltInfo models.OltParamInfo
	if err := c.ShouldBind(&oltInfo); err != nil {
		utils.AbortWithStatusJSON(c, e.SUCCESS, e.INVALID_PARAMS, err.Error())
	}
	fmt.Println(oltInfo)
	logger.SystemLogger.Debug(oltInfo)
	rpcInfo := rpc.RPCInfor{
		IP:           oltInfo.OltId,
		Port:         oltInfo.OltPort,
		TemplatePath: "reset.tpl",
	}
	reply, err := rpc.RunStaticRPC(rpcInfo, "", logger.SystemLogger, false)

	if err != nil {
		utils.RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "set fail", "data": err.Error()})
		return
	}
	fmt.Println(reply.Data)
	logger.SystemLogger.Debug(reply.Data)
	if strings.Compare(oltInfo.OltPort, constants.NT_PORT) == 0 {
		db.ResetLtInfoWhenActive()
	}
	utils.RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "Success", "data": reply})
}

// restoreHandler godoc
// @BasePath /nokia-alto
// @Summary restore olt configuration
// @schemes http https
// @Description restore the configuration to indicated olt
// @Tags Provisioning
// @Accept x-www-form-urlencoded
// @Produce json
// @Param file formData file true "file"
// @Param dstPort formData string true "port of olt card"
// @Param oltId formData string true "olt id"
// @Success      200  {json}   {success}
// @Failure      20001  {string}  {Token鉴权失败}
// @Router /v1/provisioning/olt/restore [post]
func restoreHandler(c *gin.Context) {
	oltId := c.DefaultPostForm("oltId", "192.168.1.1")
	logger.SystemLogger.Debug("===start restoreHandler")
	dstPort := c.DefaultPostForm("dstPort", "830") // OLT type
	fmt.Println(dstPort)
	logger.SystemLogger.Debug(dstPort)
	dstIpAddr := oltId
	fmt.Printf("oltId: %+v", oltId)

	file, err := c.FormFile("file")
	if err != nil {
		utils.AbortWithStatusJSON(c, e.SUCCESS, e.INVALID_PARAMS, err.Error())
		return
	}
	fmt.Println(file.Filename)
	now := time.Now().Unix()
	tmpFileName := utils.GetUploadPath() + strconv.FormatInt(now, 10) + file.Filename
	c.SaveUploadedFile(file, tmpFileName)

	tmpl, err := template.ParseFiles(tmpFileName)
	if err != nil {
		fmt.Println("create template failed, err:", err)
		logger.SystemLogger.Debug("create template failed, err:", err)
		utils.AbortWithStatusJSON(c, e.SUCCESS, e.ERROR_SSH_CREATE_FAIL, err.Error())
		return
	}
	buf := new(bytes.Buffer)
	tmpl.Execute(buf, "xx")

	netconfPrefix := `<edit-config><target><running/></target><config>`
	netconfSuffix := `</config></edit-config>`
	if strings.Compare(dstPort, constants.IHUB_PORT) == 0 {
		netconfPrefix = `<edit-config><target><candidate/></target><config>`
	}
	newStr := strings.Replace(buf.String(), "<data>", "", 1)
	newStr = strings.Replace(newStr, "</data>", "", 1)
	newStr = netconfPrefix + newStr + netconfSuffix
	//fmt.Println(newStr)
	logger.SystemLogger.Debug(newStr)
	rpcInfo := rpc.RPCInfor{
		IP:           dstIpAddr,
		Port:         dstPort,
		TemplatePath: tmpFileName,
	}
	reply, err := rpc.RunStringRPC(rpcInfo, newStr)

	if err != nil {
		utils.RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "set fail", "data": "err response"})
		return
	}
	fmt.Println(reply.Data)
	logger.SystemLogger.Debug(reply.Data)
	//for ihub need commit the candidate data
	if strings.Compare(dstPort, constants.IHUB_PORT) == 0 {
		commitRPC := rpc.NewCommit()
		_, err := rpc.RunGeneratedRPC(*commitRPC, rpcInfo, logger.SystemLogger, true)
		if err != nil {
			logger.SystemLogger.Error("commitCandidata failed:", err.Error())
			utils.RES(c, e.ERROR_NETCONF_EXEC_FAIL, gin.H{})
			return
		}
	}

	utils.RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": reply, "data": "Sent rpc success!"})

}
func doPingGW(c *gin.Context) {
	logger.SystemLogger.Debug("===start testPingOltGw")

	sshConfig := &ssh.ClientConfig{
		User:            "admin",
		Auth:            []ssh.AuthMethod{ssh.Password("admin")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	var oltInfo models.OltParamInfo
	dstIP := c.DefaultPostForm("dstIP", "")
	if err := c.ShouldBind(&oltInfo); err != nil {
		utils.AbortWithStatusJSON(c, e.SUCCESS, e.INVALID_PARAMS, err.Error())
	}
	// dstPort := c.DefaultPostForm("dstPort", "830")
	fmt.Println(oltInfo)
	fmt.Println(dstIP)
	logger.SystemLogger.Debug(oltInfo)
	activeTmpl, err := template.ParseFiles(utils.GetTemplatePath("enableLemiDebug.tpl"))
	if err != nil {
		fmt.Println("create template failed, err:", err)
		logger.SystemLogger.Debug("create template failed, err:", err)
		utils.RES(c, e.SUCCESS, gin.H{"status": e.ERROR_TEMPLATE_CREATE_FAIL, "message": "set fail", "data": "err response"})
	}
	buf := new(bytes.Buffer)
	activeTmpl.Execute(buf, "xxx")
	fmt.Printf("Reply: %+v", buf.String())

	s, err := netconf.DialSSHTimeout(oltInfo.OltId+":"+oltInfo.OltPort, sshConfig, 3*time.Second)
	if err != nil {
		fmt.Println("Dial SSH failed, err:", err)
		logger.SystemLogger.Debug("Dial SSH failed, err:", err)
		utils.AbortWithStatusJSON(c, e.SUCCESS, e.ERROR_SSH_CREATE_FAIL, err.Error())
	}
	defer s.Close()
	reply, err := s.Exec(netconf.RawMethod(buf.String()))
	if err != nil {
		v := models.RpcError{}
		err = xml.Unmarshal([]byte(reply.Data), &v)
		if err != nil {
			fmt.Println("Unmarshal  failed, err:", err)
			logger.SystemLogger.Debug("Unmarshal  failed, err:", err)
			utils.RES(c, e.SUCCESS, gin.H{"status": e.ERROR_NETCONF_EXEC_FAIL, "message": "enableLemiDebug fail", "data": "err response"})
		}
		fmt.Println(v.ErrorAppTag)
		fmt.Println(v.ErrorMessage)
		fmt.Println(v.ErrorInfo)
		logger.SystemLogger.Debug(v.ErrorAppTag)
		logger.SystemLogger.Debug(v.ErrorMessage)
		logger.SystemLogger.Debug(v.ErrorInfo)
		utils.RES(c, e.SUCCESS, gin.H{"status": e.ERROR_NETCONF_EXEC_FAIL, "message": "enableLemiDebug fail", "data": v.ErrorMessage})
	}
	logger.SystemLogger.Debug(reply.Data)

	// cli := models.Cli{
	// 	Addr: "169.254.0.1:2222",
	// 	User: "root",
	// 	Pwd:  "2x2=4",
	// }
	cli := models.Cli{
		Addr: "169.254.0.1:2222",
		User: constants.DEBUG_PORT_USER,
		Pwd:  constants.DEBUG_PORT_PWD,
	}
	fmt.Println("start connect")
	client, err := cli.Connect()
	if err != nil {
		fmt.Println("connect to cli failed, err:", err)
		logger.SystemLogger.Debug("ssh connect to 169.254.0.1:2222 failed, err:", err)
		return
	}
	fmt.Println("connect succeed")
	defer client.Client.Close()
	attempts := c.DefaultPostForm("attempts", "5")
	fmt.Println("ping " + dstIP + " -c" + attempts + " -w 1 -A >/tmp/pingResult")

	res1, err := client.Run("ping " + dstIP + " -c" + attempts + " -w 1 -A >/tmp/pingResult")
	if err != nil {
		fmt.Println("exec fail: ", err)
		logger.SystemLogger.Debug("exec fail: ", err)
	}
	// fmt.Println(routeInfo)
	fmt.Println(res1)

	res2, err := client.Run("cat /tmp/pingResult")
	if err != nil {
		fmt.Println("rm -rf /mnt/persistent/confd-cdb/* fail: ", err)
		logger.SystemLogger.Debug("rm -rf /mnt/persistent/confd-cdb/* fail: ", err)
		utils.AbortWithStatusJSON(c, e.SUCCESS, e.ERROR_SSH_CLI_EXEC_FAIL, err.Error())
	} else {
		fmt.Println(res2)
		logger.SystemLogger.Debug(res2)
		utils.RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "Success", "data": res2})
	}
}
func doPingAlto(c *gin.Context) {
	utils.RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "Success", "data": "22222"})
}
func doTraceRouteGW(c *gin.Context) {
	logger.SystemLogger.Debug("===start testPingOltGw")

	sshConfig := &ssh.ClientConfig{
		User:            "admin",
		Auth:            []ssh.AuthMethod{ssh.Password("admin")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	var oltInfo models.OltParamInfo
	dstIP := c.DefaultPostForm("dstIP", "")
	if err := c.ShouldBind(&oltInfo); err != nil {
		utils.AbortWithStatusJSON(c, e.SUCCESS, e.INVALID_PARAMS, err.Error())
	}
	// dstPort := c.DefaultPostForm("dstPort", "830")
	fmt.Println(oltInfo)
	fmt.Println(dstIP)
	logger.SystemLogger.Debug(oltInfo)
	activeTmpl, err := template.ParseFiles(utils.GetTemplatePath("enableLemiDebug.tpl"))
	if err != nil {
		fmt.Println("create template failed, err:", err)
		logger.SystemLogger.Debug("create template failed, err:", err)
		utils.RES(c, e.SUCCESS, gin.H{"status": e.ERROR_TEMPLATE_CREATE_FAIL, "message": "set fail", "data": "err response"})
	}
	buf := new(bytes.Buffer)
	activeTmpl.Execute(buf, "xxx")
	fmt.Printf("Reply: %+v", buf.String())

	s, err := netconf.DialSSHTimeout(oltInfo.OltId+":"+oltInfo.OltPort, sshConfig, 3*time.Second)
	if err != nil {
		fmt.Println("Dial SSH failed, err:", err)
		logger.SystemLogger.Debug("Dial SSH failed, err:", err)
		utils.AbortWithStatusJSON(c, e.SUCCESS, e.ERROR_SSH_CREATE_FAIL, err.Error())
	}
	defer s.Close()
	reply, err := s.Exec(netconf.RawMethod(buf.String()))
	if err != nil {
		v := models.RpcError{}
		err = xml.Unmarshal([]byte(reply.Data), &v)
		if err != nil {
			fmt.Println("Unmarshal  failed, err:", err)
			logger.SystemLogger.Debug("Unmarshal  failed, err:", err)
			utils.RES(c, e.SUCCESS, gin.H{"status": e.ERROR_NETCONF_EXEC_FAIL, "message": "enableLemiDebug fail", "data": "err response"})
		}
		fmt.Println(v.ErrorAppTag)
		fmt.Println(v.ErrorMessage)
		fmt.Println(v.ErrorInfo)
		logger.SystemLogger.Debug(v.ErrorAppTag)
		logger.SystemLogger.Debug(v.ErrorMessage)
		logger.SystemLogger.Debug(v.ErrorInfo)
		utils.RES(c, e.SUCCESS, gin.H{"status": e.ERROR_NETCONF_EXEC_FAIL, "message": "enableLemiDebug fail", "data": v.ErrorMessage})
	}
	logger.SystemLogger.Debug(reply.Data)

	cli := models.Cli{
		Addr: "169.254.0.1:2222",
		User: "root",
		Pwd:  "2x2=4",
	}
	fmt.Println("start connect")
	client, err := cli.Connect()
	if err != nil {
		fmt.Println("connect to cli failed, err:", err)
		logger.SystemLogger.Debug("ssh connect to 169.254.0.1:2222 failed, err:", err)
		return
	}
	fmt.Println("connect succeed")
	defer client.Client.Close()
	// attempts := c.DefaultPostForm("attempts", "5")
	fmt.Println("traceroute " + dstIP + " -w 1 >/tmp/pingResult")

	res1, err := client.Run("traceroute " + dstIP + " -w 1  >/tmp/pingResult")
	if err != nil {
		fmt.Println("exec fail: ", err)
		logger.SystemLogger.Debug("exec fail: ", err)
	}
	// fmt.Println(routeInfo)
	fmt.Println(res1)

	res2, err := client.Run("cat /tmp/pingResult")
	if err != nil {
		fmt.Println("rm -rf /mnt/persistent/confd-cdb/* fail: ", err)
		logger.SystemLogger.Debug("rm -rf /mnt/persistent/confd-cdb/* fail: ", err)
		utils.AbortWithStatusJSON(c, e.SUCCESS, e.ERROR_SSH_CLI_EXEC_FAIL, err.Error())
	} else {
		fmt.Println(res2)
		logger.SystemLogger.Debug(res2)
		utils.RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "Success", "data": res2})
	}
}
func doTraceRouteAlto(c *gin.Context) {
	utils.RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "Success", "data": "22222"})
}

func pingHandler(c *gin.Context) {
	mode := c.DefaultPostForm("mode", "ping")
	source := c.DefaultPostForm("source", "alto")
	if mode == "ping" {
		if source == "alto" {
			doPingAlto(c)
		} else {
			doPingGW(c)
		}
	} else if mode == "traceroute" {
		if source == "alto" {
			doTraceRouteAlto(c)
		} else {
			doTraceRouteGW(c)
		}
	} else {
		utils.AbortWithStatusJSON(c, e.SUCCESS, e.INVALID_PARAMS, "Invalid mode, only support ping and trace route")
	}

}

// @BasePath /nokia-alto
// PingExample godoc
// @Summary upload the config yang file
// @Schemes
// @Description upload the config yang file
// @Tags Provisioning
// @Accept mpfd
// @Produce json
// @Param token header string true "登录信息"
// @Param file formData file true "文件"
// @Param dstPort formData string true " tcurrent we have DF(830) and MF(832)"
// @Param ipAddr formData string true "IP address"
// @Param netMask formData string true "Net mask"
// @Param gataway formData string true "Gateway"
// @Param dstIpAddr formData string true "Destination IP address"
// @Success      200  {json}   {"message":"Success","status":200,"timestamp":"","data":""}
// @Failure      20001  {string}  {"Token鉴权失败"}
// @Router /provisioning/configFile [post]
func uploadConfigXmlHandler(c *gin.Context) {
	logger.SystemLogger.Debug("===start uploadConfigXmlHandler")
	// sshConfig := &ssh.ClientConfig{
	// 	User:            "admin",
	// 	Auth:            []ssh.AuthMethod{ssh.Password("admin")},
	// 	HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	// }

	dstPort := c.DefaultPostForm("dstPort", "830") // OLT type
	fmt.Println(dstPort)
	logger.SystemLogger.Debug(dstPort)
	ipAddr := c.DefaultPostForm("ipAddr", "10.10.10.10") // template
	fmt.Println(ipAddr)
	logger.SystemLogger.Debug(ipAddr)
	netMask := c.DefaultPostForm("netMask", "255.255.255.0") // template
	fmt.Println(netMask)
	logger.SystemLogger.Debug(netMask)
	gateway := c.DefaultPostForm("gateway", "10.10.10.1") // template
	fmt.Println(gateway)
	logger.SystemLogger.Debug(gateway)
	dstIpAddr := c.PostForm("oltId")
	fmt.Println(dstIpAddr)
	logger.SystemLogger.Debug(dstIpAddr)
	file, err := c.FormFile("file")
	fmt.Println(file.Filename)
	if err != nil {
		utils.RES(c, e.SUCCESS, gin.H{
			"status":  e.INVALID_PARAMS,
			"message": err.Error(),
		})
		return
	}
	ipInfo := models.IpInfo{
		IpAddr:  ipAddr,
		NetMask: netMask,
		GateWay: gateway,
	}

	rpcInfo := rpc.RPCInfor{
		IP:   dstIpAddr,
		Port: dstPort,
		// TemplatePath: "activeSoftware.tpl",
	}

	now := time.Now().Unix()

	tmpFileName := utils.GetUploadPath() + strconv.FormatInt(now, 10) + file.Filename
	c.SaveUploadedFile(file, tmpFileName)
	buf := new(bytes.Buffer)
	tmpl, err := template.ParseFiles(tmpFileName)
	if err != nil {
		fmt.Println("create template failed, err:", err)
		logger.SystemLogger.Debug("create template failed, err:", err)
		utils.AbortWithStatusJSON(c, e.SUCCESS, e.ERROR_SSH_CREATE_FAIL, err.Error())
	}
	tmpl.Execute(buf, ipInfo)
	reply, err := rpc.RunStringRPC(rpcInfo, buf.String())

	// reply, err := s.Exec(netconf.RawMethod(buf.String()))
	logger.SystemLogger.Debug(buf.String())
	fmt.Println(buf.String())
	if err != nil {
		logger.SystemLogger.Debug("upload config xml RunStringRPC failed, err:", err)
		utils.RES(c, e.SUCCESS, gin.H{"status": e.ERROR_NETCONF_EXEC_FAIL, "message": "load initial cfg fail", "data": err.Error()})
		return
	}
	fmt.Println("result ok")
	logger.SystemLogger.Debug(reply.Data)
	utils.RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": reply, "data": "Sent rpc success!"})
}

// @BasePath /nokia-alto
// PingExample godoc
// @Summary undeploy the MF2 webgui
// @Schemes
// @Description undeploy the MF2 webgui
// @Tags Provisioning
// @Accept mpfd
// @Produce json
// @Param token header string true "登录信息"
// @Param dstPort formData string true " tcurrent we only have MF(832)"
// @Success      200  {json}   {"message":"Success","status":200,"timestamp":"","data":""}
// @Failure      20001  {string}  {"Token鉴权失败"}
// @Router /provisioning/mf_gui [post]
func webguiUndeployHandler(c *gin.Context) {

	softwareUrl := "rm -f /root/*.tar && export BOARD_PASS=2x2=4 && /isam/scripts/webgui_manager undeploy all"

	fmt.Println(softwareUrl)
	logger.SystemLogger.Debug(softwareUrl)
	remoteSshIp := "169.254.0.1:2222"
	cli := models.Cli{
		Addr: remoteSshIp,
		User: "root",
		Pwd:  "2x2=4",
	}
	client, err := cli.Connect()
	if err != nil {
		fmt.Println("ssh connect to "+remoteSshIp+" failed, err:", err)
		logger.SystemLogger.Debug("ssh connect to "+remoteSshIp+" failed, err:", err)
		return
	}
	defer client.Client.Close()
	logger.SystemLogger.Debug("need undeploy")
	res, err := client.Run(softwareUrl)

	if err != nil {
		fmt.Println("softwareUrl fail: softwareUrl fail: ", err)
		logger.SystemLogger.Debug(err)
	} else {
		fmt.Println(res)
		logger.SystemLogger.Debug(res)
	}
	res2, err := client.Run("/isam/scripts/webgui_manager query all")

	if err != nil {
		fmt.Println("/isam/scripts/webgui_manager query all fail: ", err)
	} else {
		fmt.Println(res2)
		logger.SystemLogger.Debug(res2)
	}
	var returnStr string = ""
	strArr := strings.Split(res2, "\n")
	for i := 0; i < len(strArr); i++ {
		if strings.Contains(strArr[i], "active nt") || strings.Contains(strArr[i], "lt-1") || strings.Contains(strArr[i], "lt-2") {
			returnStr = returnStr + strings.TrimSpace(strArr[i]) + ","

		}
	}
	logger.SystemLogger.Debug(returnStr)

	utils.RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "Success", "data": returnStr})

}

// @BasePath /nokia-alto
// PingExample godoc
// @Summary deploy the MF2 webgui
// @Schemes
// @Description deploy the MF2 webgui
// @Tags Provisioning
// @Accept mpfd
// @Produce json
// @Param token header string true "登录信息"
// @Param dstPort formData string true " tcurrent we only have MF(832)"
// @Success      200  {json}   {"message":"Success","status":200,"timestamp":"","data":""}
// @Failure      20001  {string}  {"Token鉴权失败"}
// @Router /provisioning/mf_gui [post]
func webguiDeployHandler(c *gin.Context) {

	// s, err := DialSShRPCTimeout("192.168.1.1:832")
	// if err != nil {
	// 	fmt.Println("Dial SSH failed, err:", err)
	// 	logger.SystemLogger.Debug("Dial SSH failed, err:", err)
	// 	RES(c, e.SUCCESS, gin.H{"status": e.ERROR_NETCONF_EXEC_FAIL, "message": "dail fail", "data": "dail err response"})
	// }
	sshConfig := &ssh.ClientConfig{
		User:            "admin",
		Auth:            []ssh.AuthMethod{ssh.Password("admin")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	s, err := netconf.DialSSHTimeout("192.168.1.1:832", sshConfig, 3*time.Second)
	if err != nil {
		fmt.Println("Dial SSH failed, err:", err)
		logger.SystemLogger.Debug("Dial SSH failed, err:", err)
		utils.RES(c, e.SUCCESS, gin.H{"status": e.ERROR_NETCONF_EXEC_FAIL, "message": "dail fail", "data": "dail err response"})
	}

	defer s.Close()
	// Get Chassis type
	chassisTmpl, err := template.ParseFiles(utils.GetTemplatePath("getChassisInfo.tpl"))
	if err != nil {
		fmt.Println("create chassis template failed, err:", err)
		logger.SystemLogger.Debug("create chassis template failed, err:", err)
		utils.RES(c, e.SUCCESS, gin.H{"status": e.ERROR_NETCONF_EXEC_FAIL, "message": "chassis tpl fail", "data": "tpl err"})
	}
	chassisBuf := new(bytes.Buffer)
	chassisTmpl.Execute(chassisBuf, "xx")
	chassisReply, err := s.Exec(netconf.RawMethod(chassisBuf.String()))
	if err != nil {
		fmt.Println("get chassis info fail:", err)
		logger.SystemLogger.Debug("get chassis info fail:", err)
	}
	type ChassisInfo struct {
		HardwareState struct {
			Component struct {
				ModelName string `xml:"model-name"`
			} `xml:"component"`
		} `xml:"hardware-state"`
	}
	var chassisInfo ChassisInfo
	xml.Unmarshal([]byte(chassisReply.Data), &chassisInfo)
	chassisType := chassisInfo.HardwareState.Component.ModelName
	fmt.Println("ChassisType:", chassisType)
	logger.SystemLogger.Debug("ChassisType: " + chassisType)
	// Get NT type
	tmpl, err := template.ParseFiles(utils.GetTemplatePath("getNtInfo.tpl"))
	if err != nil {
		fmt.Println("create nt template failed, err:", err)
		logger.SystemLogger.Debug("create nt template failed, err:", err)
		utils.RES(c, e.SUCCESS, gin.H{"status": e.ERROR_NETCONF_EXEC_FAIL, "message": "dail fail", "data": "dail err response"})
	}
	buf := new(bytes.Buffer)
	tmpl.Execute(buf, "xx")
	fmt.Printf("request: %+v\r\n", buf.String())
	logger.SystemLogger.Debug("request: ", buf.String())
	reply, err := s.Exec(netconf.RawMethod(buf.String()))
	if err != nil {
		fmt.Println("get nt info fail:", err)
		logger.SystemLogger.Debug("get nt info fail:", err)
	} else {
		fmt.Printf("response: %+v\r\n", reply.Data)
		logger.SystemLogger.Debug("response: ", reply.Data)
		if reply.Data == "<data></data>" {
			fmt.Printf("nta is null now")
			logger.SystemLogger.Debug("nta is null now")
		} else {
			v := models.LtBoardInfoData{}
			err := xml.Unmarshal([]byte(reply.Data), &v)
			if err != nil {
				fmt.Printf("error: %v\r\n", err)
				logger.SystemLogger.Debug("parse error:", err)
				utils.RES(c, e.SUCCESS, gin.H{"status": e.ERROR_NETCONF_EXEC_FAIL, "message": "parse fail", "data": "dail err response"})
			}
			fmt.Printf("NTa modelName: %+v\r\n", v.HardwareState.Component.ModelName)
			boardType := v.HardwareState.Component.ModelName
			fmt.Println("BoardType:", boardType)
			logger.SystemLogger.Debug("BoardType: " + boardType)

			installPrefix := chassisType + "_" + boardType

			readerInfos, err := ioutil.ReadDir("software/")
			if err != nil {
				fmt.Println(err)
				logger.SystemLogger.Debug("read dir error:", err)
				utils.RES(c, e.SUCCESS, gin.H{"status": e.ERROR_NETCONF_EXEC_FAIL, "message": "parse fail", "data": "dail err response"})
			}
			var webguiInstallName string
			for _, info := range readerInfos {
				if !info.IsDir() && strings.HasPrefix(info.Name(), installPrefix) {
					webguiInstallName = info.Name()
					break
				}
			}
			if webguiInstallName == "" {
				utils.RES(c, e.SUCCESS, gin.H{"status": e.ERROR_NETCONF_EXEC_FAIL, "message": "cannot find the installation package", "data": ""})
				return
			}

			var serverLocalIp string = utils.FindLocalIp("192.168.1.")
			var localPort string = "5600"
			var prefixStr string = "wget http://" + serverLocalIp + ":" + localPort + "/software/"

			newInstallName := strings.TrimPrefix(webguiInstallName, chassisType+"_")

			softwareUrl := prefixStr + webguiInstallName +
				" && mv /root/" + webguiInstallName + " /root/" + newInstallName +
				" && ln -s /root/" + newInstallName + " /root/" + boardType + "_WEBGUI_SOFTWARE.tar" +
				" && export BOARD_PASS=2x2=4 && /isam/scripts/webgui_manager deploy all"

			fmt.Println(softwareUrl)
			logger.SystemLogger.Debug(softwareUrl)
			remoteSshIp := "169.254.0.1:2222"
			cli := models.Cli{
				Addr: remoteSshIp,
				User: "root",
				Pwd:  "2x2=4",
			}
			client, err := cli.Connect()
			if err != nil {
				fmt.Println("ssh connect to "+remoteSshIp+" failed, err:", err)
				logger.SystemLogger.Debug("ssh connect to "+remoteSshIp+" failed, err:", err)
			}
			defer client.Client.Close()
			logger.SystemLogger.Debug("need delete old rules and add new rules")
			res, err := client.Run(softwareUrl)
			fmt.Println(res)
			logger.SystemLogger.Debug(res)
			if err != nil {
				fmt.Println("softwareUrl fail: softwareUrl fail: ", err)
			}
			res2, err := client.Run("/isam/scripts/webgui_manager query all")
			fmt.Println(res2)
			logger.SystemLogger.Debug(res2)
			if err != nil {
				fmt.Println("/isam/scripts/webgui_manager query all fail: ", err)
			}
			var returnStr string = ""
			strArr := strings.Split(res2, "\n")
			for i := 0; i < len(strArr); i++ {
				if strings.Contains(strArr[i], "active nt") || strings.Contains(strArr[i], "lt-1") || strings.Contains(strArr[i], "lt-2") {
					returnStr = returnStr + strings.TrimSpace(strArr[i]) + ","

				}
			}
			logger.SystemLogger.Debug(returnStr)

			utils.RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "Success", "data": returnStr})
			return
		}
	}

	utils.RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "Success", "data": "2323232"})
}

// @BasePath /nokia-alto
// PingExample godoc
// @Summary  show result
// @Schemes
// @Description  show result
// @Tags Provisioning
// @Accept mpfd
// @Produce json
// @Param token header string true "登录信息"
// @Success      200  {json}   {"message":"Success","status":200,"timestamp":"","data":""}
// @Failure      20001  {string}  {"Token鉴权失败"}
// @Router /provisioning/mf_gui [get]
func getWebguiDeploidResultHandler(c *gin.Context) {
	remoteSshIp := "169.254.0.1:2222"
	// cli := models.Cli{
	// 	Addr: remoteSshIp,
	// 	User: "root",
	// 	Pwd:  "2x2=4",
	// }
	cli := models.Cli{
		Addr: remoteSshIp,
		User: constants.DEBUG_PORT_USER,
		Pwd:  constants.DEBUG_PORT_PWD,
	}
	client, err := cli.Connect()
	if err != nil {
		fmt.Println("ssh connect to "+remoteSshIp+" failed, err:", err)
		logger.SystemLogger.Debug("ssh connect to "+remoteSshIp+" failed, err:", err)
		utils.RES(c, e.SUCCESS, gin.H{"status": e.ERROR_NETCONF_EXEC_FAIL, "message": "dail fail", "data": "dail err response"})
		return
	}
	defer client.Client.Close()
	res2, err := client.Run("export BOARD_PASS=2x2=4 && /isam/scripts/webgui_manager query all")
	//fmt.Println(res2)
	logger.SystemLogger.Debug(res2)
	if err != nil {
		fmt.Println("/isam/scripts/webgui_manager query all fail: ", err)
	}
	var returnStr string = ""
	strArr := strings.Split(res2, "\n")
	for i := 0; i < len(strArr); i++ {
		if strings.Contains(strArr[i], "active nt") || strings.Contains(strArr[i], "lt-1") || strings.Contains(strArr[i], "lt-2") {
			returnStr = returnStr + strings.TrimSpace(strArr[i]) + ","

		}
	}
	logger.SystemLogger.Debug(returnStr)

	utils.RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "Success", "data": returnStr})
}
