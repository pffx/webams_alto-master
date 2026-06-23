package olt

import (
	"alto_server/common/db"
	logger "alto_server/common/log"
	"alto_server/common/models"
	"alto_server/common/pkg/e"
	"alto_server/common/rpc"
	"alto_server/common/utils"
	"alto_server/constants"
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/elliotchance/sshtunnel"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/ssh"
)

type CurrentState struct {
	XMLName   xml.Name `xml:"current-state"`
	State     string   `xml:"state" json:"State"`
	Timestamp string   `xml:"timestamp" json:"Timestamp"`
}
type LastDownloadState struct {
	XMLName      xml.Name `xml:"last-download-state"`
	State        string   `xml:"state" json:"State"`
	Timestamp    string   `xml:"timestamp" json:"Timestamp"`
	SoftwareName string   `xml:"software-name" json:"SoftwareName"`
}

type Download struct {
	XMLName           xml.Name          `xml:"download"`
	CurrentState      CurrentState      `json:"CurrentState"`
	LastDownloadState LastDownloadState `json:"LastDownloadState"`
}
type ConfigDownload struct {
	XMLName           xml.Name          `xml:"config-download"`
	CurrentState      CurrentState      `json:"CurrentState"`
	LastDownloadState LastDownloadState `json:"LastDownloadState"`
}
type Software2 struct {
	XMLName        xml.Name       `xml:"software"`
	Name           string         `xml:"name" json:"Name"`
	Download       Download       `json:"Download"`
	Revisions      Revisions      `json:"Revisions"`
	ConfigDownload ConfigDownload `json:"ConfigDownload"`
}

type Software1 struct {
	XMLName  xml.Name  `xml:"software"`
	Software Software2 `json:"Software"`
}

type Component struct {
	XMLName  xml.Name  `xml:"component"`
	Name     string    `xml:"name" json:"Name"`
	Software Software1 `json:"Software"`
}
type HardwareState struct {
	XMLName   xml.Name  `xml:"hardware-state"`
	Component Component `json:"Component"`
}

type SoftwareVersionData struct {
	XMLName       xml.Name      `xml:"data"`
	HardwareState HardwareState `json:"HardwareState"`
	Ip            string
	HostName      string
	Port          string
}

type Revision struct {
	XMLName           xml.Name `xml:"revision"`
	Name              string   `xml:"name" json:""`
	DownloadTimestamp string   `xml:"download-timestamp" json:""`
	Version           string   `xml:"version" json:""`
	IsValid           string   `xml:"is-valid" json:""`
	IsCommitted       string   `xml:"is-committed" json:""`
	IsActive          string   `xml:"is-active" json:""`
}
type Revisions struct {
	XMLName  xml.Name   `xml:"revisions"`
	Revision []Revision `xml:"revision" json:""`
}

func returnOltsSoftware(c *gin.Context) (string, error) {
	logger.SystemLogger.Debug("===start returnOltsSoftware")
	var result map[int]string
	var index = 0
	result = make(map[int]string)
	db, err := sql.Open("sqlite3", utils.GetDatabasePath())
	if err != nil {
		fmt.Println("open db fail:", err)
	}
	defer db.Close()
	utils.PanickErr(err)
	rows, err := db.Query("SELECT * FROM olt")
	utils.PanickErr(err)
	for rows.Next() {
		var id int
		var name string
		var ip string
		var port int
		var userName string
		var password string
		var oltType string
		var LtNum int
		err = rows.Scan(&id, &name, &ip, &port, &userName, &password, &oltType, &LtNum)
		fmt.Printf("returnOltsSoftware    ip: %+v", ip)
		utils.PanickErr(err)
		result[index], err = returnOltSoftware(c, ip)
		if err != nil {
			fmt.Printf("returnOltsSoftware    with err: ")
			return "", err
		}
		index++
	}
	bytes, _ := json.Marshal(result)
	stringData := string(bytes)
	return stringData, nil
}

func returnOltSoftware(c *gin.Context, oltIp string) (string, error) {
	//logger.SystemLogger.Debug("===start returnOltSoftware")
	// sshConfig := &ssh.ClientConfig{
	// 	User:            "admin",
	// 	Auth:            []ssh.AuthMethod{ssh.Password("admin")},
	// 	HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	// }

	dstPort := c.DefaultQuery("dstPort", "830") // OLT type
	//logger.SystemLogger.Debug(dstPort)
	dstIpAddr := oltIp
	fmt.Println(dstIpAddr)
	//logger.SystemLogger.Debug(dstIpAddr)

	// s, err := netconf.DialSSHTimeout(dstIpAddr+":"+dstPort, sshConfig, 3*time.Second)
	// if err != nil {
	// 	logger.LoggerToSyslog().Debug("Dial SSH failed, err:", err)
	// 	logger.LoggerToErrorlog().Debug("Dial SSH failed, err:", err)
	// 	fmt.Println("Dial SSH failed, err:", err)
	// 	return "", err
	// }
	// defer s.Close()

	if dstPort == constants.NT_PORT {
		rpc.EnableAllDebug(dstIpAddr, dstPort, "admin", "admin")
		// activeTmpl, err := template.ParseFiles(GetTemplatePath("enableLemiDebug.tpl"))
		// if err == nil {
		// 	buf := new(bytes.Buffer)
		// 	activeTmpl.Execute(buf, "xx")
		// 	s.Exec(netconf.RawMethod(buf.String()))
		// } else {
		// 	logger.LoggerToSyslog().Debug("enable lemi to NT failed, err:", err)
		// 	logger.LoggerToErrorlog().Debug("enable lemi to NT failed, err:", err)
		// }
	}

	rpcInfo := rpc.RPCInfor{
		IP:           dstIpAddr,
		Port:         dstPort,
		TemplatePath: "getSoftwareStatus.tpl",
	}
	reply, err := rpc.RunStaticRPC(rpcInfo, "", logger.SystemLogger, false)

	// tmpl, err := template.ParseFiles(GetTemplatePath("getSoftwareStatus.tpl"))
	// if err != nil {
	// 	logger.LoggerToSyslog().Debug("create getSoftwareStatus failed, err:", err)
	// 	logger.LoggerToErrorlog().Debug("create getSoftwareStatus failed, err:", err)
	// 	fmt.Println("create getSoftwareStatus failed, err:", err)
	// 	return "", err
	// }
	// buf := new(bytes.Buffer)
	// tmpl.Execute(buf, "xx")
	// logger.LoggerToSyslog().Debug(buf.String())
	// reply, err := s.Exec(netconf.RawMethod(buf.String()))

	if err != nil {
		fmt.Println(err)
		logger.SystemLogger.Debug("exec getSoftwareStatus failed, err:", err)
		if reply == nil {
			utils.RES(c, e.SUCCESS, gin.H{"status": e.ERROR_NETCONF_EXEC_FAIL, "message": "exec netconf fail", "data": "err response"})
			return "", err
		}
		v := models.RpcError{}
		err = xml.Unmarshal([]byte(reply.Data), &v)
		if err != nil {
			fmt.Println("Unmarshal  failed, err:", err)
			logger.SystemLogger.Debug("Unmarshal  failed, err:", err)
			utils.RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "get fail", "data": "err response"})
			return "", err
		}
		fmt.Println(v.ErrorAppTag)
		fmt.Println(v.ErrorMessage)
		//fmt.Println(v.ErrorInfo)
		logger.SystemLogger.Debug(v.ErrorAppTag)
		logger.SystemLogger.Debug(v.ErrorMessage)
		return "", err
	} else {
		//fmt.Printf("Reply: %+v", reply.Data)
		//logger.SystemLogger.Debug(reply.Data)
		v := SoftwareVersionData{}
		err := xml.Unmarshal([]byte(reply.Data), &v)
		if err != nil {
			fmt.Printf("error: %v", err)
			logger.SystemLogger.Error("xml parse fail:", err)
			utils.RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "get fail", "data": "err response"})
			return "", err
		}
		v.Ip = dstIpAddr
		// v.HostName = "DF-16"
		v.Port = dstPort
		jsonData, err := json.Marshal(v)
		//fmt.Printf("software info: %v", jsonData)
		if err != nil {
			fmt.Printf("error: %v", err)
			logger.SystemLogger.Error("json parse fail:", err)
			return "", err
		}
		return string(jsonData), nil
	}
}

func activeOltSoftware(c *gin.Context) {
	logger.SystemLogger.Debug("===start activeOltSoftware")
	// sshConfig := &ssh.ClientConfig{
	// 	User:            "admin",
	// 	Auth:            []ssh.AuthMethod{ssh.Password("admin")},
	// 	HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	// }

	dstPort := c.DefaultPostForm("dstPort", constants.DF_PORT)
	fmt.Println(dstPort)
	logger.SystemLogger.Debug(dstPort)
	dstIpAddr := c.DefaultPostForm("oltId", "192.168.1.1")
	fmt.Println(dstIpAddr)
	logger.SystemLogger.Debug(dstIpAddr)
	softwareName := c.DefaultPostForm("name", "")
	fmt.Println(softwareName)
	logger.SystemLogger.Debug(softwareName)
	softwareDownloadName := models.SoftwareDownloadName{
		SoftwareName: softwareName,
	}
	rpc.EnableAllDebug(dstIpAddr, dstPort, "admin", "admin")
	//finish enable lemi, then start clean db.
	remoteSshIp := "169.254.0.1:2222"
	if strings.Compare(dstPort, constants.DF_PORT) != 0 && strings.Compare(dstPort, constants.IHUB_PORT) != 0 && strings.Compare(dstPort, constants.NT_PORT) != 0 {
		numDstPort, err := strconv.Atoi(dstPort)
		if err != nil {
			logger.SystemLogger.Debug("convert dstport to number fail:" + dstPort)
			utils.RES(c, e.SUCCESS, gin.H{"status": e.ERROR_NETCONF_EXEC_FAIL, "message": "convert fail", "data": dstPort})
			return
		}
		tunnelIpLastPart := numDstPort - 832 + 2
		logger.SystemLogger.Debug("tunnel to 169.254.1." + strconv.Itoa(tunnelIpLastPart))
		tunnel := sshtunnel.NewSSHTunnel(
			"root@169.254.0.1:2222",
			ssh.Password("2x2=4"),
			"169.254.1."+strconv.Itoa(tunnelIpLastPart)+":2222",
			"0",
		)
		go tunnel.Start()
		time.Sleep(1000 * time.Millisecond)
		remoteSshIp = "127.0.0.1:" + strconv.Itoa(tunnel.Local.Port)
	}

	cli := models.Cli{
		Addr: remoteSshIp,
		User: constants.DEBUG_PORT_USER,
		Pwd:  constants.DEBUG_PORT_PWD,
	}
	client, err := cli.Connect()
	if err != nil {
		fmt.Println("clean db. ssh connect to "+remoteSshIp+" failed, err:", err)
		logger.SystemLogger.Debug("clean db, ssh connect to "+remoteSshIp+" failed, err:", err)
		utils.RES(c, e.SUCCESS, gin.H{"status": e.ERROR_SSH_CLI_EXEC_FAIL, "message": "ssh fail", "data": err.Error()})
		return
	}
	logger.SystemLogger.Debug(remoteSshIp)
	defer client.Client.Close()
	//DF
	res, err := client.Run("rm -rf /isam/slot_default/fast_db/*")
	if err != nil {
		fmt.Println("rm -rf /isam/slot_default/fast_db/* fail: ", err)
		logger.SystemLogger.Debug("rm -rf /isam/slot_default/fast_db/* fail: ", err)
	}
	fmt.Println(res)
	logger.SystemLogger.Debug("rm -rf /isam/slot_default/fast_db/*")
	logger.SystemLogger.Debug(res)
	res1, err := client.Run("rm -rf /mnt/persistent/confd-cdb/*")
	if err != nil {
		fmt.Println("rm -rf /mnt/persistent/confd-cdb/* fail: ", err)
		logger.SystemLogger.Debug("rm -rf /mnt/persistent/confd-cdb/* fail: ", err)
	}
	fmt.Println(res1)
	logger.SystemLogger.Debug("rm -rf /mnt/persistent/confd-cdb/*")
	logger.SystemLogger.Debug(res1)
	res2, err := client.Run("rm -f /mnt/persistent/ihub/config/*_config.cfg")
	if err != nil {
		fmt.Println("rm -f /mnt/persistent/ihub/config/*_config.cfg fail: ", err)
		logger.SystemLogger.Debug("rm -f /mnt/persistent/ihub/config/*_config.cfg fail: ", err)
	}
	fmt.Println(res2)
	logger.SystemLogger.Debug("rm -f /mnt/persistent/ihub/config/*_config.cfg")
	logger.SystemLogger.Debug(res2)
	res3, err := client.Run("rm -rf /mnt/nand-dbase/confd-cdb/*.*")
	if err != nil {
		fmt.Println("rm -rf /mnt/nand-dbase/confd-cdb/*.* fail: ", err)
		logger.SystemLogger.Debug("rm -rf /mnt/nand-dbase/confd-cdb/*.* fail: ", err)
	}
	fmt.Println(res3)
	logger.SystemLogger.Debug("rm -rf /mnt/nand-dbase/confd-cdb/*.*")
	logger.SystemLogger.Debug(res3)
	res4, err := client.Run("rm -rf /mnt/nand/backup/*")
	if err != nil {
		fmt.Println("rm -rf /mnt/nand/backup/* fail: ", err)
		logger.SystemLogger.Debug("rm -rf /mnt/nand/backup/* fail: ", err)
	}
	fmt.Println(res4)
	logger.SystemLogger.Debug("rm -rf /mnt/nand/backup/*")
	logger.SystemLogger.Debug(res4)
	res7, err := client.Run("rm -rf /mnt/emmc-persistent/Dbase/*")
	if err != nil {
		fmt.Println("rm -rf /mnt/emmc-persistent/Dbase/* fail: ", err)
		logger.SystemLogger.Debug("rm -rf /mnt/emmc-persistent/Dbase/* fail: ", err)
	}
	fmt.Println(res7)
	logger.SystemLogger.Debug("rm -rf /mnt/emmc-persistent/Dbase/*")
	logger.SystemLogger.Debug(res7)
	res5, err := client.Run("sync")
	if err != nil {
		fmt.Println("rm -rf /mnt/nand/backup/* fail: ", err)
		logger.SystemLogger.Debug("rm -rf /mnt/nand/backup/* fail: ", err)
	}
	fmt.Println(res5)
	logger.SystemLogger.Debug("sync")
	logger.SystemLogger.Debug(res5)
	res6, err := client.Run("ifconfig |grep 169.254.")
	if err != nil {
		fmt.Println("ifconfig |grep 169.254.1 ", err)
		logger.SystemLogger.Debug("ifconfig |grep 169.254.1 ", err)
	}
	fmt.Println(res6)
	logger.SystemLogger.Debug("ifconfig |grep 169.254.1")
	logger.SystemLogger.Debug(res6)
	fmt.Printf("cleandb done \r\n")
	logger.SystemLogger.Debug("cleandb done")

	rpcInfo := rpc.RPCInfor{
		IP:           dstIpAddr,
		Port:         dstPort,
		TemplatePath: "activeSoftware.tpl",
	}
	activeReply, err := rpc.RunStaticRPC(rpcInfo, softwareDownloadName, logger.SystemLogger, false)

	if err != nil {
		v := models.RpcError{}
		logger.SystemLogger.Debug(activeReply.Data)
		err = xml.Unmarshal([]byte(activeReply.Data), &v)
		if err != nil {
			fmt.Println("Unmarshal  failed, err:", err)
			logger.SystemLogger.Debug("Unmarshal  failed, err:", err)
			utils.RES(c, e.SUCCESS, gin.H{"status": e.ERROR_NETCONF_EXEC_FAIL, "message": "set fail", "data": "err response"})
			return
		}
		fmt.Println(v.ErrorAppTag)
		fmt.Println(v.ErrorMessage)
		fmt.Println(v.ErrorInfo)
		logger.SystemLogger.Debug(v.ErrorAppTag)
		logger.SystemLogger.Debug(v.ErrorMessage)
		logger.SystemLogger.Debug(v.ErrorInfo)
		utils.RES(c, e.SUCCESS, gin.H{"status": e.ERROR_NETCONF_EXEC_FAIL, "message": "set fail", "data": v.ErrorMessage})
		return
	} else {
		fmt.Println(activeReply.Data)
		logger.SystemLogger.Debug(activeReply.Data)
		if dstPort == constants.NT_PORT {
			db.ResetLtInfoWhenActive()
		}
		utils.RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "Success", "data": activeReply})
	}
}
func downloadOltSoftware(c *gin.Context) {
	logger.SystemLogger.Debug("===start downloadOltSoftware")
	// sshConfig := &ssh.ClientConfig{
	// 	User:            "admin",
	// 	Auth:            []ssh.AuthMethod{ssh.Password("admin")},
	// 	HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	// }

	dstPort := c.DefaultPostForm("dstPort", "830")
	fmt.Println(dstPort)
	logger.SystemLogger.Debug(dstPort)
	dstIpAddr := c.DefaultPostForm("oltId", "192.168.1.1")
	fmt.Println(dstIpAddr)
	logger.SystemLogger.Debug(dstIpAddr)
	softwareUrl := c.DefaultPostForm("url", "")
	fmt.Println(softwareUrl)
	//firstly enable lemi and add iptables
	if strings.Compare(dstPort, "830") == 0 || strings.Compare(dstPort, "831") == 0 || strings.Compare(dstPort, "832") == 0 {
		fmt.Println("no need add iptables")
	} else {
		err := rpc.EnableAllDebug(dstIpAddr, "832", "admin", "admin")

		// s2, err := netconf.DialSSHTimeout(dstIpAddr+":832", sshConfig, 3*time.Second)
		// if err != nil {
		// 	fmt.Println("Dial SSH failed, err:", err)
		// 	logger.LoggerToSyslog().Debug("Dial SSH failed, err:", err)
		// 	RES(c, e.SUCCESS, gin.H{"status": e.ERROR_NETCONF_EXEC_FAIL, "message": "dail fail", "data": "dail err response"})
		// }
		// defer s2.Close()
		// logger.LoggerToSyslog().Debug("lt card download, need firstly  add iptables rules")
		// activeTmpl, err := template.ParseFiles(GetTemplatePath("enableLemiDebug.tpl"))
		// if err != nil {
		// 	fmt.Println("create template failed, err:", err)
		// 	logger.LoggerToSyslog().Debug("create template failed, err:", err)
		// }
		// buf := new(bytes.Buffer)
		// activeTmpl.Execute(buf, "xx")
		// fmt.Printf("enable request: %+v", buf.String())
		// logger.LoggerToSyslog().Debug("enable request:", buf.String())

		// reply, err := s2.Exec(netconf.RawMethod(buf.String()))
		if err != nil {

			// v := models.RpcError{}
			// err = xml.Unmarshal([]byte(reply.Data), &v)
			// if err != nil {
			// 	fmt.Println("active olt Unmarshal  failed, err:", err)
			// 	logger.LoggerToSyslog().Debug("active olt Unmarshal  failed, err:", err)
			// } else {
			// 	logger.LoggerToSyslog().Debug(reply.Data)
			// 	logger.LoggerToSyslog().Debug(v.ErrorAppTag)
			// 	logger.LoggerToSyslog().Debug(v.ErrorMessage)
			// 	logger.LoggerToSyslog().Debug(v.ErrorInfo)
			logger.SystemLogger.Debug("active olt, no need stop when enable lemi fail", err.Error())
			// }
		} else {
			logger.SystemLogger.Debug("enable lemi debug success, sshing")
			// logger.LoggerToSyslog().Debug(reply.Data)
			remoteSshIp := "169.254.0.1:2222"
			cli := models.Cli{
				Addr: remoteSshIp,
				User: constants.DEBUG_PORT_USER,
				Pwd:  constants.DEBUG_PORT_PWD,
			}
			client, err := cli.Connect()
			if err != nil {
				fmt.Println("ssh connect to "+remoteSshIp+" failed, err:", err)
				logger.SystemLogger.Debug("ssh connect to "+remoteSshIp+" failed, err:", err)
				utils.RES(c, e.SUCCESS, gin.H{"status": e.ERROR_SSH_CLI_EXEC_FAIL, "message": "ssh fail", "data": err.Error()})
				return
			}
			defer client.Client.Close()
			logger.SystemLogger.Debug("need delete old rules and add new rules")
			res, err := client.Run("iptables -t nat -D POSTROUTING -s 169.254.1.1/24  -j SNAT --to-source 192.168.1.1")
			fmt.Println(res)
			logger.SystemLogger.Debug(res)
			if err != nil {
				fmt.Println("iptables -t nat -D POSTROUTING -s 169.254.1.1/24  -j SNAT --to-source 192.168.1.1 fail: ", err)
				logger.SystemLogger.Debug("iptables -t nat -D POSTROUTING -s 169.254.1.1/24  -j SNAT --to-source 192.168.1.1 fail: ", err)
				//AbortWithStatusJSON(c, e.SUCCESS, e.ERROR_SSH_CLI_EXEC_FAIL, err.Error())
			}
			res1, err := client.Run("iptables -t nat -I POSTROUTING -s 169.254.1.1/24  -j SNAT --to-source 192.168.1.1")
			fmt.Println(res1)
			logger.SystemLogger.Debug(res1)
			if err != nil {
				fmt.Println("iptables -t nat -I POSTROUTING -s 169.254.1.1/24  -j SNAT --to-source 192.168.1.1 fail: ", err)
				logger.SystemLogger.Debug("failed added iptables -t nat -I POSTROUTING -s 169.254.1.1/24  -j SNAT --to-source 192.168.1.1 fail: ", err.Error())
				//AbortWithStatusJSON(c, e.SUCCESS, e.ERROR_SSH_CLI_EXEC_FAIL, err.Error())
				utils.RES(c, e.SUCCESS, gin.H{"status": e.ERROR_SSH_CLI_EXEC_FAIL, "message": "ssh fail", "data": err.Error()})
				return
			}
			logger.SystemLogger.Debug("===downloadOltSoftware add iptables done")
		}
	}

	//http://192.168.1.64:5600/software/lightspan_2209.422/L6GQAG/L6GQAG2209.422
	//path should be lightspan_2209.422/L6GQAG/L6GQAG2209.422
	//we should add prefix http://192.168.1.64:5600/software/
	var serverLocalIp string = utils.FindLocalIp("192.168.1.")
	var localPort string = "5600"
	var prefixStr string = "http://" + serverLocalIp + ":" + localPort + "/software/"
	softwareUrl = prefixStr + softwareUrl
	fmt.Println(softwareUrl)
	logger.SystemLogger.Debug(softwareUrl)
	softwareName := c.DefaultPostForm("name", "")
	fmt.Println(softwareName)
	logger.SystemLogger.Debug(softwareName)
	softwareDownloadInfo := models.SoftwareDownloadInfo{
		SoftwareName: softwareName,
		SoftwareUrl:  softwareUrl,
	}

	rpcInfo := rpc.RPCInfor{
		IP:           dstIpAddr,
		Port:         dstPort,
		TemplatePath: "downloadSoftware.tpl",
	}
	reply, err := rpc.RunStaticRPC(rpcInfo, softwareDownloadInfo, logger.SystemLogger, false)

	// tmpl, err := template.ParseFiles(GetTemplatePath("downloadSoftware.tpl"))
	// if err != nil {
	// 	fmt.Println("create template failed, err:", err)
	// 	logger.LoggerToSyslog().Debug("get downloadSoftware template failed, err:", err)
	// 	RES(c, e.SUCCESS, gin.H{"status": e.ERROR_TEMPLATE_CREATE_FAIL, "message": "create downloadSoftware fail", "data": err.Error()})
	// }
	// buf := new(bytes.Buffer)
	// tmpl.Execute(buf, softwareDownloadInfo)
	// fmt.Printf("Reply: %+v", buf.String())
	// logger.LoggerToSyslog().Debug("downloadSoftware request:", buf.String())

	// s, err := netconf.DialSSHTimeout(dstIpAddr+":"+dstPort, sshConfig, 3*time.Second)
	// if err != nil {
	// 	fmt.Println("Dial SSH failed, err:", err)
	// 	logger.LoggerToSyslog().Debug("Dial SSH failed, err:", err)
	// 	RES(c, e.SUCCESS, gin.H{"status": e.ERROR_SSH_CREATE_FAIL, "message": "create ssh to  downloadSoftware fail", "data": err.Error()})
	// }
	// defer s.Close()
	// reply, err := s.Exec(netconf.RawMethod(buf.String()))
	if err != nil {
		logger.SystemLogger.Debug(reply.Data)
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
		fmt.Println(v.ErrorInfo)
		logger.SystemLogger.Debug(v.ErrorAppTag)
		logger.SystemLogger.Debug(v.ErrorMessage)
		logger.SystemLogger.Debug(v.ErrorInfo)
		utils.RES(c, e.SUCCESS, gin.H{"status": e.ERROR_ACTION_FAILED, "message": "set fail", "data": v.ErrorMessage})
		return
	} else {
		fmt.Println(reply.Data)
		logger.SystemLogger.Debug(reply.Data)
		utils.RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "Success", "data": reply})
	}
}
func commitOltSoftware(c *gin.Context) {
	logger.SystemLogger.Debug("===start commitOltSoftware")
	// sshConfig := &ssh.ClientConfig{
	// 	User:            "admin",
	// 	Auth:            []ssh.AuthMethod{ssh.Password("admin")},
	// 	HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	// }

	dstPort := c.DefaultPostForm("dstPort", "830")
	fmt.Println(dstPort)
	logger.SystemLogger.Debug(dstPort)
	dstIpAddr := c.DefaultPostForm("oltId", "192.168.1.1")
	fmt.Println(dstIpAddr)
	logger.SystemLogger.Debug(dstIpAddr)
	softwareName := c.DefaultPostForm("name", "")
	fmt.Println(softwareName)
	logger.SystemLogger.Debug(softwareName)
	softwareDownloadName := models.SoftwareDownloadName{
		SoftwareName: softwareName,
	}
	// s, err := netconf.DialSSHTimeout(dstIpAddr+":"+dstPort, sshConfig, 3*time.Second)
	// if err != nil {
	// 	fmt.Println("Dial SSH failed, err:", err)
	// 	logger.LoggerToSyslog().Debug("Dial SSH failed, err:", err)
	// 	RES(c, e.SUCCESS, gin.H{"status": e.ERROR_SSH_CREATE_FAIL, "message": "create ssh to  downloadSoftware fail", "data": err.Error()})
	// }
	// defer s.Close()
	if strings.Compare(dstPort, constants.DF_PORT) != 0 && strings.Compare(dstPort, constants.IHUB_PORT) != 0 && strings.Compare(dstPort, constants.NT_PORT) != 0 {
		logger.SystemLogger.Debug("lt card commit, need firstly del added iptables rules")
		err := rpc.EnableAllDebug(dstIpAddr, dstPort, "admin", "admin")

		// activeTmpl, err := template.ParseFiles(GetTemplatePath("enableLemiDebug.tpl"))
		if err != nil {
			// 	fmt.Println("create template failed, err:", err)
			// 	logger.LoggerToSyslog().Debug("create template failed, err:", err)
			// }
			// buf := new(bytes.Buffer)
			// activeTmpl.Execute(buf, softwareDownloadName)
			// fmt.Printf("enable request: %+v", buf.String())
			// logger.LoggerToSyslog().Debug("enable request:", buf.String())

			// reply, err := s.Exec(netconf.RawMethod(buf.String()))
			// if err != nil {
			// 	logger.LoggerToSyslog().Debug(reply.Data)
			// 	v := models.RpcError{}
			// 	err = xml.Unmarshal([]byte(reply.Data), &v)
			// 	if err != nil {
			// 		fmt.Println("active olt Unmarshal  failed, err:", err)
			// 		logger.LoggerToSyslog().Debug("active olt Unmarshal  failed, err:", err)
			// 	} else {
			// 		logger.LoggerToSyslog().Debug(v.ErrorAppTag)
			// 		logger.LoggerToSyslog().Debug(v.ErrorMessage)
			// 		logger.LoggerToSyslog().Debug(v.ErrorInfo)
			logger.SystemLogger.Debug("active olt, no need stop when enable lemi fail")
			// 	}
		} else {
			logger.SystemLogger.Debug("enable lemi debug success, sshing")
			// logger.LoggerToSyslog().Debug(reply.Data)
			remoteSshIp := "169.254.0.1:2222"
			cli := models.Cli{
				Addr: remoteSshIp,
				User: constants.DEBUG_PORT_USER,
				Pwd:  constants.DEBUG_PORT_PWD,
			}
			client, err := cli.Connect()
			if err != nil {
				fmt.Println("ssh connect to "+remoteSshIp+" failed, err:", err)
				logger.SystemLogger.Debug("ssh connect to "+remoteSshIp+" failed, err:", err)
			} else {
				defer client.Client.Close()
				logger.SystemLogger.Debug("iptables -t nat -D POSTROUTING -s 169.254.1.1/24  -j SNAT --to-source 192.168.1.1")
				res, err := client.Run("iptables -t nat -D POSTROUTING -s 169.254.1.1/24  -j SNAT --to-source 192.168.1.1")
				if err != nil {
					fmt.Println("iptables -t nat -D POSTROUTING -s 169.254.1.1/24  -j SNAT --to-source 192.168.1.1 fail: ", err)
					logger.SystemLogger.Debug("iptables -t nat -D POSTROUTING -s 169.254.1.1/24  -j SNAT --to-source 192.168.1.1 fail: ", err)
				}
				fmt.Println(res)
				logger.SystemLogger.Debug(res)
			}
		}
	}

	rpcInfo := rpc.RPCInfor{
		IP:           dstIpAddr,
		Port:         dstPort,
		TemplatePath: "commitSoftware.tpl",
	}
	reply, err := rpc.RunStaticRPC(rpcInfo, softwareDownloadName, logger.SystemLogger, false)

	// tmpl, err := template.ParseFiles(GetTemplatePath("commitSoftware.tpl"))
	// if err != nil {
	// 	fmt.Println("create template failed, err:", err)
	// 	logger.LoggerToSyslog().Debug("get commitSoftware template failed, err:", err)
	// 	RES(c, e.SUCCESS, gin.H{"status": e.ERROR_TEMPLATE_CREATE_FAIL, "message": "create fail", "data": err.Error()})
	// }
	// buf := new(bytes.Buffer)
	// tmpl.Execute(buf, softwareDownloadName)
	// fmt.Printf("commit request: %+v", buf.String())
	// logger.LoggerToSyslog().Debug("commit request: ", buf.String())

	// reply, err := s.Exec(netconf.RawMethod(buf.String()))
	if err != nil {
		logger.SystemLogger.Debug(reply.Data)
		v := models.RpcError{}
		err = xml.Unmarshal([]byte(reply.Data), &v)
		if err != nil {
			fmt.Println("Unmarshal  failed, err:", err)
			logger.SystemLogger.Debug("Unmarshal  failed, err:", err)
			utils.RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "set fail", "data": "err response"})
		}
		fmt.Println(v.ErrorAppTag)
		fmt.Println(v.ErrorMessage)
		fmt.Println(v.ErrorInfo)
		logger.SystemLogger.Debug(v.ErrorAppTag)
		logger.SystemLogger.Debug(v.ErrorMessage)
		logger.SystemLogger.Debug(v.ErrorInfo)
		utils.RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "set fail", "data": v.ErrorMessage})
		return
	} else {
		fmt.Println(reply.Data)
		logger.SystemLogger.Debug(reply.Data)
		utils.RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "Success", "data": reply})
	}
}

func GetSuffixName(filePath string) string {
	return path.Ext(filePath)
}
func GetAllFile(pathname string) []string {
	rd, _ := ioutil.ReadDir(pathname)
	var filePath []string
	// fmt.Printf("rd: %+v", rd)
	for _, fi := range rd {
		if fi.IsDir() {
			subdirectory := GetAllFile(pathname + fi.Name() + "\\")

			// filePath = append(filePath, pathname + "/" + fi.Name())
			filePath = append(filePath, utils.ArrayToString(subdirectory))
		} else if fi.Name() == ".gitignore" {
			//ignore
		} else {
			filePath = append(filePath, pathname+"/"+fi.Name())
		}
	}
	return filePath
}

func GetAllFileRoot(pathname string) string {
	rd, _ := ioutil.ReadDir(pathname)
	var files map[string]string
	var index = 0
	files = make(map[string]string)
	for _, fi := range rd {
		if fi.IsDir() {
			files[fi.Name()] = GetAllFileRoot(pathname + fi.Name() + "/")
		} else if fi.Name() == ".gitignore" {
			//ignore
		} else {
			files[fi.Name()] = fi.Name()
		}
		index++
	}
	stringData := utils.MapStringStringToJsonString(files)

	// fmt.Println("##############################################################  ")
	// fmt.Println("stringData " + stringData)
	return stringData
}

func getAvailableSoftwareList(c *gin.Context) {
	result := GetAllFileRoot(utils.GetSoftwarePath())
	// fmt.Println("##############################################################")
	utils.RES(c, e.SUCCESS, gin.H{
		"software_list": result,
		"status":        e.SUCCESS,
		"message":       "Services list  get success.",
	})
}

func getOltSoftwareList(c *gin.Context) {
	// oltId := "10.106.132.53"
	oltId := c.DefaultQuery("oltId", "")
	var oltStr string = ""
	var err error = nil
	if oltId == "" {
		oltStr, err = returnOltsSoftware(c)
	} else {
		oltStr, err = returnOltSoftware(c, oltId)
	}
	if err != nil {
		fmt.Printf("getOltSoftwareList after with error")
		utils.RES(c, e.SUCCESS, gin.H{"status": e.ERROR, "message": "get fail", "data": err})
	}
	fmt.Printf("##############################################################  \n")
	fmt.Printf("getOltSoftwareList    oltStr: %+v \n", oltStr)
	fmt.Printf("##############################################################  \n")
	utils.RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "Success", "olt_software_info": oltStr})
}

func handleOltSoftware(c *gin.Context) {
	oltId := c.DefaultPostForm("oltId", "")
	action := c.DefaultPostForm("action", "")
	path := c.DefaultPostForm("url", "")
	name := c.DefaultPostForm("name", "")
	if oltId == "" {
		utils.AbortWithStatusJSON(c, e.SUCCESS, e.INVALID_PARAMS, "Please specific an olt")
	} else {
		switch action {
		case "active":
			if name == "" {
				utils.AbortWithStatusJSON(c, e.SUCCESS, e.INVALID_PARAMS, "Please specific software name")
			} else {
				activeOltSoftware(c)
			}
		case "commit":
			if name == "" {
				utils.AbortWithStatusJSON(c, e.SUCCESS, e.INVALID_PARAMS, "Please specific software name")
			} else {
				commitOltSoftware(c)
			}
		case "download":
			if path == "" || name == "" {
				utils.AbortWithStatusJSON(c, e.SUCCESS, e.INVALID_PARAMS, "Please specific an olt software")
			} else {
				downloadOltSoftware(c)
			}
		case "active_migrate":
			if name == "" {
				utils.AbortWithStatusJSON(c, e.SUCCESS, e.INVALID_PARAMS, "Please specific software name")
			} else {
				activeOltSoftwareWithMigrate(c)
			}
		default:
			utils.AbortWithStatusJSON(c, e.SUCCESS, e.INVALID_PARAMS, "Please use the right operation")
		}
	}
}

func handleOltConfigMigration(c *gin.Context) {
	// var node models.Node
	//targetVersion := c.DefaultPostForm("targetVersion", "25.6")
	// if err := c.ShouldBind(&node); err != nil {
	// 	utils.RES(c, e.INVALID_PARAMS, gin.H{})
	// 	// logger.SystemLogger.Debug("input param error", err.Error())
	// 	return
	// }
	// result := db.DbHandle.Where("olt_ip = ? ", node.OltIp).First(&node)
	// if result.Error != nil {
	// 	fmt.Println("can not find that OLT when read software")
	// 	utils.RES(c, e.ERROR_FIND_OLT_FAIL, gin.H{})
	// 	return
	// }

	oltId := c.DefaultPostForm("oltId", "192.168.1.1")
	// oltId := "10.106.132.53"
	target := c.DefaultPostForm("target", "25.9")
	dstPort := c.DefaultPostForm("dstPort", constants.DF_PORT)
	fmt.Println("handleOltConfigMigration target: ", target)
	card, err := utils.GetOLTCardByPort(dstPort)
	if err != nil {
		logger.SystemLogger.Error(err.Error())
		utils.RES(c, e.ERROR_ACTION_FAILED, gin.H{"status": e.SUCCESS, "message": "Success"})
		return
	}
	logger.SystemLogger.Infof("%s %s %s %s ", oltId, target, dstPort, card)
	curVersion, err := rpc.GetCurPlatformVersion(oltId, dstPort)
	if err != nil {
		logger.SystemLogger.Error(err.Error())
		utils.RES(c, e.ERROR_ACTION_FAILED, gin.H{"status": e.SUCCESS, "message": "Success"})
		return
	}
	oltInfo := models.Olt{}
	db.DbHandle.Where("ip=?", oltId).Find(&oltInfo)
	//go rpc.MigrateRunningConfiguration(oltId, "lt1", "25.3", targetVersion)
	rpc.MigrateRunningConfiguration(oltInfo, card, curVersion, target)
	utils.RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "Success"})
}

func handleOltConfigMigrationUploading(c *gin.Context) {

	oltId := c.DefaultPostForm("oltId", "192.168.1.1")
	// oltId := "10.106.132.53"
	dstPort := c.DefaultPostForm("dstPort", constants.DF_PORT)
	card, err := utils.GetOLTCardByPort(dstPort)
	if err != nil {
		logger.SystemLogger.Error(err.Error())
		utils.RES(c, e.ERROR_ACTION_FAILED, gin.H{"status": e.SUCCESS, "message": "Success"})
		return
	}
	targetName, err := rpc.GetCurStandbyVersionName(oltId, dstPort)
	if err != nil {
		logger.SystemLogger.Error(err.Error())
		utils.RES(c, e.ERROR_ACTION_FAILED, gin.H{"status": e.SUCCESS, "message": "Success"})
		return
	}
	zipName := oltId + "_" + card + "_migration.zip"
	logger.SystemLogger.Infof("handleOltConfigMigrationUploading  %s %s %s %s %s", oltId, dstPort, card, targetName, zipName)

	rpc.DownloadMigrationConfig(oltId, dstPort, targetName, zipName)

	utils.RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "Success"})
}

func activeOltSoftwareWithMigrate(c *gin.Context) {
	dstPort := c.DefaultPostForm("dstPort", constants.DF_PORT)
	dstIpAddr := c.DefaultPostForm("oltId", "192.168.1.1")
	softwareName := c.DefaultPostForm("name", "")
	softwareDownloadName := models.SoftwareDownloadName{
		SoftwareName: softwareName,
	}
	logger.SystemLogger.Infof("activeOltSoftwareWithMigrate %s %s %s", dstIpAddr, dstPort, softwareName)

	rpcInfo := rpc.RPCInfor{
		IP:           dstIpAddr,
		Port:         dstPort,
		TemplatePath: "activeSoftware.tpl",
	}
	_, err := rpc.RunStaticRPC(rpcInfo, softwareDownloadName, logger.SystemLogger, true)

	if err != nil {
		logger.SystemLogger.Error(err.Error())
		utils.RES(c, e.ERROR_ACTION_FAILED, gin.H{"status": e.SUCCESS, "message": "Success"})
		return
	} else {
		utils.RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "Success"})
	}

}
