package system

import (
	"alto_server/common/db"
	"alto_server/common/feature"
	logger "alto_server/common/log"
	"alto_server/common/models"
	"alto_server/common/pkg/e"
	"alto_server/common/rpc"
	"alto_server/common/utils"
	. "alto_server/common/utils"
	"alto_server/constants"
	"bytes"
	"encoding/xml"
	"fmt"
	"html/template"
	"io"
	"strconv"
	"strings"

	"github.com/Juniper/go-netconf/netconf"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/ssh"
)

func getOlts() ([]models.OltInfoWithLTCard, error) {

	var olts []models.OltInfoWithLTCard
	// index := 1
	fmt.Printf("open success")
	db := db.DbOriginal

	rows, err := db.Query("SELECT * FROM olt")
	if err != nil {
		logger.SystemLogger.Debug("select sqlite file failed, err:", err)
		fmt.Println("select sqlite file failed, err:", err)
		return olts, err
	}
	for rows.Next() {
		var id int
		var ip string
		var userName string
		var password string
		var oltType string
		var LtNum int
		var lt1Info string
		var lt2Info string
		var lt3Info string
		var lt4Info string
		var lt5Info string
		var lt6Info string
		var lt7Info string
		var lt8Info string
		var lt9Info string
		var lt10Info string
		var lt11Info string
		var lt12Info string
		var lt13Info string
		var lt14Info string
		err = rows.Scan(&id, &ip, &userName, &password, &oltType, &LtNum, &lt1Info, &lt2Info, &lt3Info, &lt4Info, &lt5Info, &lt6Info, &lt7Info, &lt8Info, &lt9Info, &lt10Info, &lt11Info, &lt12Info, &lt13Info, &lt14Info)
		if err != nil {
			logger.SystemLogger.Debug("scan sqlite file failed, err:", err)
			logger.SystemLogger.Debug("scan sqlite file failed, err:", err)
			fmt.Println("scan sqlite file failed, err:", err)
			return olts, err
		}
		olt := models.OltInfoWithLTCard{}
		olt.IP = ip
		olt.Username = userName
		olt.Password = password
		olt.Status = "Connected"
		olt.OltType = oltType
		olt.OltLtNum = LtNum
		olt.Lt1Info = lt1Info
		olt.Lt2Info = lt2Info
		olt.Lt3Info = lt3Info
		olt.Lt4Info = lt4Info
		olt.Lt5Info = lt5Info
		olt.Lt6Info = lt6Info
		olt.Lt7Info = lt7Info
		olt.Lt8Info = lt8Info
		olt.Lt9Info = lt9Info
		olt.Lt10Info = lt10Info
		olt.Lt11Info = lt11Info
		olt.Lt12Info = lt12Info
		olt.Lt13Info = lt13Info
		olt.Lt14Info = lt14Info
		olts = append(olts, olt)
		// index++
	}
	return olts, nil
}
func todoSystemHandler(c *gin.Context) {
	RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "Success", "data": "2323232"})
}

// @BasePath /nokia-alto
// PingExample godoc
// @Summary get the olt list
// @Schemes
// @Description find all the connected olt device
// @Tags System
// @Accept mpfd
// @Produce json
// @Param token header string true "登录信息"
// @Success      200  {json}   {"message":"Success","status":200,"timestamp":"","data":""}
// @Failure      20001  {string}  {"Token鉴权失败"}
// @Router /system/oltList [get]
func getOLTListHandler(c *gin.Context) {
	olts, err := getOlts()
	if err != nil {
		AbortWithStatusJSON(c, e.SUCCESS, e.ERROR_FIND_OLTS_FAIL, err.Error())
	}
	var featurelist = feature.GetFeatureList()

	RES(c, e.SUCCESS, gin.H{
		"olt_list":    olts,
		"status":      e.SUCCESS,
		"featurelist": featurelist,
		"message":     "OLTs get success.",
	})
}

// @BasePath /nokia-alto
// PingExample godoc
// @Summary add new olt
// @Schemes
// @Description add a new olt to the controled list
// @Tags System
// @Accept mpfd
// @Produce json
// @Param token header string true "登录信息"
// @Param        IP   formData      string  true  "IP"
// @Param        Username   formData      string  true  "Username"
// @Param        Password   formData      string  true  "Password"
// @Success      200  {json}   {"message":"Success","status":200,"timestamp":"","data":""}
// @Failure      20001  {string}  {"Token鉴权失败"}
// @Router /system/oltList [put]
func addOLTListHandler(c *gin.Context) {
	var olt models.OltInfo
	if err := c.ShouldBind(&olt); err != nil {
		AbortWithStatusJSON(c, e.SUCCESS, e.INVALID_PARAMS, e.GetMessage(e.INVALID_PARAMS))
	}
	value := olt.IP + "/" + olt.Username + "/" + olt.Password + "/" + olt.OltType + "/" + strconv.Itoa(olt.OltLtNum)
	fmt.Printf("addOLTListHandler   new olt  value: %+v \n", value)
	// err := conf.Cfg.Section("olt").NewKey(key, value)

	db := db.DbOriginal
	stmt, err := db.Prepare("update olt set user_name=?,passwd=?,type=?,lt_num=? where ip=?")
	PanickErr(err)

	res, err := stmt.Exec(olt.Username, olt.Password, olt.OltType, strconv.Itoa(olt.OltLtNum), olt.IP)
	PanickErr(err)
	affect, err := res.RowsAffected()
	PanickErr(err)
	if affect == 0 {
		stmt, err := db.Prepare("INSERT INTO olt(ip,user_name, passwd, type,lt_num) values(?,?,?,?,?)")
		PanickErr(err)

		res, err := stmt.Exec(olt.IP, olt.Username, olt.Password, olt.OltType, strconv.Itoa(olt.OltLtNum))
		PanickErr(err)

		id, err := res.LastInsertId()
		PanickErr(err)
		RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "Success", "type": "insert", "data": id})
	} else {
		RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "Success", "type": "update", "data": affect})
	}
}

func connectOltHandler(c *gin.Context) {
	// sshConfig := &ssh.ClientConfig{
	// 	User:            "admin",
	// 	Auth:            []ssh.AuthMethod{ssh.Password("Netconf$150")},
	// 	HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	// }
	// s, err := netconf.DialSSH(dstIpAddr+":"+dstPort, sshConfig)
	// defer
	//now := time.Now().Unix()
	RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "Success", "data": "2323232"})
}

func ntGWHandler(c *gin.Context) {
	logger.SystemLogger.Debug("===start getPingGW")

	sshConfig := &ssh.ClientConfig{
		User:            "admin",
		Auth:            []ssh.AuthMethod{ssh.Password("admin")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	var oltInfo models.OltParamInfo
	if err := c.ShouldBind(&oltInfo); err != nil {
		AbortWithStatusJSON(c, e.SUCCESS, e.INVALID_PARAMS, err.Error())
	}
	// dstPort := c.DefaultPostForm("dstPort", "830")
	fmt.Println(oltInfo)
	logger.SystemLogger.Debug(oltInfo)
	activeTmpl, err := template.ParseFiles(GetTemplatePath("enableLemiDebug.tpl"))
	if err != nil {
		fmt.Println("create template failed, err:", err)
		logger.SystemLogger.Debug("create template failed, err:", err)
		RES(c, e.SUCCESS, gin.H{"status": e.ERROR_TEMPLATE_CREATE_FAIL, "message": "set fail", "data": "err response"})
	}
	buf := new(bytes.Buffer)
	activeTmpl.Execute(buf, "xxx")
	fmt.Printf("Reply: %+v", buf.String())

	s, err := netconf.DialSSH(oltInfo.OltId+":"+oltInfo.OltPort, sshConfig)
	if err != nil {
		fmt.Println("Dial SSH failed, err:", err)
		logger.SystemLogger.Debug("Dial SSH failed, err:", err)
		AbortWithStatusJSON(c, e.SUCCESS, e.ERROR_SSH_CREATE_FAIL, err.Error())
	}
	defer s.Close()
	reply, err := s.Exec(netconf.RawMethod(buf.String()))
	if err != nil {
		v := models.RpcError{}
		err = xml.Unmarshal([]byte(reply.Data), &v)
		if err != nil {
			fmt.Println("Unmarshal  failed, err:", err)
			logger.SystemLogger.Debug("Unmarshal  failed, err:", err)
			RES(c, e.SUCCESS, gin.H{"status": e.ERROR_NETCONF_EXEC_FAIL, "message": "enableLemiDebug fail", "data": "err response"})
		}
		fmt.Println(v.ErrorAppTag)
		fmt.Println(v.ErrorMessage)
		fmt.Println(v.ErrorInfo)
		logger.SystemLogger.Debug(v.ErrorAppTag)
		logger.SystemLogger.Debug(v.ErrorMessage)
		logger.SystemLogger.Debug(v.ErrorInfo)
		RES(c, e.SUCCESS, gin.H{"status": e.ERROR_NETCONF_EXEC_FAIL, "message": "enableLemiDebug fail", "data": v.ErrorMessage})
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
	res, err := client.Run("ip route|grep default|awk '{print $3}'")
	if err != nil {
		fmt.Println("rm -rf /isam/slot_default/fast_db/* fail: ", err)
		logger.SystemLogger.Debug("rm -rf /isam/slot_default/fast_db/* fail: ", err)
		AbortWithStatusJSON(c, e.SUCCESS, e.ERROR_SSH_CLI_EXEC_FAIL, err.Error())
	}
	routeInfo := strings.Replace(res, "\n", "", -1)
	routeInfo = strings.Replace(routeInfo, " ", "", -1)
	RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "Success", "data": routeInfo})
}

func rpcPushHandler(c *gin.Context) {
	dstPort := c.DefaultPostForm("dstPort", constants.DF_PORT)
	dstIpAddr := c.DefaultPostForm("oltId", "192.168.1.1")
	file, err := c.FormFile("file")

	if err != nil {
		utils.RES(c, e.INVALID_PARAMS, gin.H{})
		return
	}

	fileObj, err := file.Open()
	if err != nil {
		utils.RES(c, e.ERROR_NO_FILE, gin.H{"message": "Failed to open the file"})
		return
	}
	defer fileObj.Close()

	contentBytes, err := io.ReadAll(fileObj)
	if err != nil {
		utils.RES(c, e.ERROR_NO_FILE, gin.H{"message": "Failed to read the file content"})
		return
	}
	rInfor := rpc.RPCInfor{
		IP:   dstIpAddr,
		Port: dstPort,
	}
	reply, err := rpc.RunStringRPC(rInfor, string(contentBytes))

	if err != nil {
		logger.SystemLogger.Info("rpcPushHandler   failed OLT ip is: ", dstIpAddr)
		utils.RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": err.Error()})
		return
	}
	logger.SystemLogger.Info("rpcPushHandler   success OLT ip is: ", dstIpAddr)
	utils.RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": reply.Data})
}
