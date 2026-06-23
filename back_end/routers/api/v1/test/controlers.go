package test

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"strconv"
	"text/template"
	"time"

	//"log"
	"alto_server/common/pkg/e"
	. "alto_server/common/utils" //RES

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/ssh"

	alto_ssh "alto_server/common/ssh"

	"github.com/Juniper/go-netconf/netconf"
	_ "github.com/mattn/go-sqlite3"
)

type IpInfo struct {
	IpAddr       string
	NetMask      string
	DefaultRoute string
}
type rpcError struct {
	ErrorType     string `xml:"error-type"`
	ErrorTag      string `xml:"error-tag"`
	ErrorSeverity string `xml:"error-severity"`
	ErrorAppTag   string `xml:"error-app-tag"`
	ErrorMessage  string `xml:"error-message"`
	ErrorInfo     string `xml:"error-info"`
}

func uploadXmlHandler(c *gin.Context) {

	sshConfig := &ssh.ClientConfig{
		User:            "admin",
		Auth:            []ssh.AuthMethod{ssh.Password("Netconf$150")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	dstPort := c.PostForm("dstPort")
	fmt.Println(dstPort)
	ipAddr := c.PostForm("ipAddr")
	fmt.Println(ipAddr)
	netMask := c.PostForm("netMask")
	fmt.Println(netMask)
	gateway := c.PostForm("gataway")
	fmt.Println(gateway)
	dstIpAddr := c.PostForm("dstIpAddr")
	fmt.Println(dstIpAddr)
	file, err := c.FormFile("file")
	if err != nil {
		c.String(500, "上传图片出错")
	}
	ipInfo := IpInfo{
		IpAddr:       ipAddr,
		NetMask:      netMask,
		DefaultRoute: gateway,
	}
	s, err := netconf.DialSSH(dstIpAddr+":"+dstPort, sshConfig)
	now := time.Now().Unix()

	defer s.Close()
	tmpFileName := "./upload/" + strconv.FormatInt(now, 10) + file.Filename
	c.SaveUploadedFile(file, tmpFileName)
	buf := new(bytes.Buffer)
	tmpl, err := template.ParseFiles(tmpFileName)
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	tmpl.Execute(buf, ipInfo)
	reply, err := s.Exec(netconf.RawMethod(buf.String()))

	if err != nil {
		v := rpcError{}
		err = xml.Unmarshal([]byte(reply.Data), &v)
		fmt.Println(v.ErrorAppTag)
		fmt.Println(v.ErrorMessage)
		//fmt.Println(v.ErrorInfo)
		RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "set fail", "data": v.ErrorMessage})
	} else {
		fmt.Println("result ok")
		RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "Success", "data": reply})
	}
}

func testRpcHandler(c *gin.Context) {
	tmpl, err := template.ParseFiles(GetTemplatePath("setSysLogServer.tpl"))
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	buf := new(bytes.Buffer)
	//logger.LoggerToErrorlog().Debug("xxxx")
	//logger.LoggerToSyslog().Debug("xxxx")
	tmpl.Execute(buf, "xx")
	//fmt.Println("===parsed temp:", buf.String())
	reply, err := alto_ssh.SSHSESSIONHANDLE.Exec(netconf.RawMethod(buf.String()))
	//fmt.Printf("Reply: %+v", reply.Data)
	if err != nil {
		v := rpcError{}
		err = xml.Unmarshal([]byte(reply.Data), &v)
		fmt.Println(v.ErrorAppTag)
		fmt.Println(v.ErrorMessage)
		//fmt.Println(v.ErrorInfo)
		RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "set fail", "data": v.ErrorMessage})
	} else {
		fmt.Println("result ok")
		RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "Success", "data": reply})
	}
}
func someXMLHandler(c *gin.Context) {
	/*
		tmpl, err := template.ParseFiles("./template/firstCfgOobMf2.tpl")
		if err != nil {
			fmt.Println("create template failed, err:", err)
			return
		}
		logger.LoggerToErrorlog().Debug("xxxx")
		logger.LoggerToSyslog().Debug("xxxx")
		ipInfo := IpInfo{
			IpAddr:       "135.251.246.191",
			NetMask:      "255.255.255.0",
			DefaultRoute: "135.251.247.1",
		}
		tmpl.Execute(c.Writer, ipInfo)
	*/
}

// @BasePath /v1
// getBasicInfoHandler godoc
// @Summary test the get request
// @Schemes
// @Description test the get request
// @Tags V1,Test
// @Accept json
// @Produce json
// @Success      200  {string}   success
// @Router /v1/test/basicInfo [get]
func getBasicInfoHandler(c *gin.Context) {
	fmt.Println("come getBasicInfoHandler ")
	// token := c.Query("token")
	// fmt.Printf("token: %+v", token)
	//var tmpSession = netconf.NewSession(alto_ssh.SSHSESSIONHANDLE.Transport)
	buf := new(bytes.Buffer)
	tmpl, err := template.ParseFiles(GetTemplatePath("getRunningSystem.tpl"))
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	tmpl.Execute(buf, "xx")
	fmt.Println("getBasicInfoHandler sessionId:", alto_ssh.SSHSESSIONHANDLE.SessionID)
	//fmt.Println("getBasicInfoHandler sessionId:", tmpSession.SessionID)
	//reply, err := tmpSession.Exec(netconf.RawMethod(buf.String()))
	reply, err := alto_ssh.SSHSESSIONHANDLE.Exec(netconf.RawMethod(buf.String()))
	PanickErr(err)
	//fmt.Printf("Reply: %+v", reply)
	if err != nil {
		RES(c, e.SUCCESS, gin.H{
			"status":  e.ERROR_SSH_CREATE_FAIL,
			"message": err.Error(),
		})
		AbortWithStatusJSON(c, e.SUCCESS, e.ERROR_SSH_CREATE_FAIL, err.Error())
	} else {
		RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "Success", "data": reply})
	}

}

// @BasePath /v1
// setBasicInfoHandler godoc
// @Summary test the post request
// @Schemes
// @Description test the post request
// @Tags V1,Test
// @Accept json
// @Produce json
// @Success 200 {string} successful!
// @Router /v1/test/basicInfo [post]
func setBasicInfoHandler(c *gin.Context) {
	RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "Success", "data": "2323232"})
}
