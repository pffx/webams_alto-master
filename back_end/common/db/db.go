package db

import (
	logger "alto_server/common/log"
	"alto_server/common/models"
	"alto_server/common/utils"
	"bytes"
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Juniper/go-netconf/netconf"
	"golang.org/x/crypto/ssh"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var SyncLock sync.WaitGroup
var ResetcLock sync.WaitGroup
var DbHandle *gorm.DB
var DbOriginal *sql.DB

func InitSqliteConn() {
	config := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		}}

	database, err := gorm.Open(sqlite.Open(utils.GetDatabasePath()), config)
	if err != nil {
		logger.SystemLogger.Debug("init sqlite fail, err:", err)
		fmt.Println("init sqlite fail, err:", err)
		return
	}
	if err := database.Exec("PRAGMA journal_mode=WAL;").Error; err != nil {
		logger.SystemLogger.Debug("fail to set wal mode, err:", err.Error())
		return
	}
	DbHandle = database
	DbOriginal, _ = DbHandle.DB()
}
func GetOltInfoWithIp(oltIp string) (models.OltInfoWithLTCard, error) {

	olt := models.OltInfoWithLTCard{}
	db := DbHandle

	oltTableInfo := models.Olt{}
	db.Where("ip=?", oltIp).Find(&oltTableInfo)

	olt.IP = oltTableInfo.Ip
	olt.Username = oltTableInfo.UserName
	olt.Password = oltTableInfo.Passwd
	olt.Status = "Connected"
	olt.OltType = oltTableInfo.Type
	olt.OltLtNum = oltTableInfo.LtNum
	olt.Lt1Info = oltTableInfo.Lt1
	olt.Lt2Info = oltTableInfo.Lt2
	olt.Lt3Info = oltTableInfo.Lt3
	olt.Lt4Info = oltTableInfo.Lt4
	olt.Lt5Info = oltTableInfo.Lt5
	olt.Lt6Info = oltTableInfo.Lt6
	olt.Lt7Info = oltTableInfo.Lt7
	olt.Lt8Info = oltTableInfo.Lt8
	olt.Lt9Info = oltTableInfo.Lt9
	olt.Lt10Info = oltTableInfo.Lt10
	olt.Lt11Info = oltTableInfo.Lt11
	olt.Lt12Info = oltTableInfo.Lt12
	olt.Lt13Info = oltTableInfo.Lt13
	olt.Lt14Info = oltTableInfo.Lt14

	return olt, nil
}
func GetFirstOltInfo() (models.OltInfoWithLTCard, error) {

	olt := models.OltInfoWithLTCard{}
	db := DbHandle

	oltTableInfo := models.Olt{}
	db.First(&oltTableInfo)

	olt.IP = oltTableInfo.Ip
	olt.Username = oltTableInfo.UserName
	olt.Password = oltTableInfo.Passwd
	olt.Status = "Connected"
	olt.OltType = oltTableInfo.Type
	olt.OltLtNum = oltTableInfo.LtNum
	olt.Lt1Info = oltTableInfo.Lt1
	olt.Lt2Info = oltTableInfo.Lt2
	olt.Lt3Info = oltTableInfo.Lt3
	olt.Lt4Info = oltTableInfo.Lt4
	olt.Lt5Info = oltTableInfo.Lt5
	olt.Lt6Info = oltTableInfo.Lt6
	olt.Lt7Info = oltTableInfo.Lt7
	olt.Lt8Info = oltTableInfo.Lt8
	olt.Lt9Info = oltTableInfo.Lt9
	olt.Lt10Info = oltTableInfo.Lt10
	olt.Lt11Info = oltTableInfo.Lt11
	olt.Lt12Info = oltTableInfo.Lt12
	olt.Lt13Info = oltTableInfo.Lt13
	olt.Lt14Info = oltTableInfo.Lt14

	return olt, nil
}
func UpdateLtInfoWithIp(oltIp string, ltIndex string, ltInfo string) (int, error) {

	fmt.Printf("UpdateLtInfoWithIp %+v %+v %+v\r\n", ltIndex, ltInfo, oltIp)
	db := DbHandle
	//oltTableInfo := models.Olt{}
	//db.Where("ip=?", oltIp).Find(&oltTableInfo)
	logger.OltLogger.Infof("UpdateLtInfoWithIp %s %s %s", ltIndex, ltInfo, oltIp)

	execResult := db.Exec("update olt set "+ltIndex+" = ? where ip=?;", ltInfo, oltIp)

	return int(execResult.RowsAffected), nil
}

func UpdateLtNumWithIp(oltIp string, ltNum int) (int, error) {

	fmt.Printf("UpdateLtNumWithIp %+v %+v\r\n", ltNum, oltIp)
	logger.OltLogger.Infof("UpdateLtNumWithIp %d %s", ltNum, oltIp)
	db := DbHandle
	execResult := db.Exec("update olt set lt_num = ? where ip=?;", ltNum, oltIp)

	return int(execResult.RowsAffected), nil
}
func AddLtNum(oltIp string) {
	oltInfo, err := GetOltInfoWithIp(oltIp)
	if err == nil {
		newLtNum := oltInfo.OltLtNum + 1
		UpdateLtNumWithIp(oltIp, newLtNum)
	} else {
		logger.SystemLogger.Debug("GetOltInfoWithIp fail, err:", err)
	}
}
func SetLtNumToZero(oltIp string) {
	UpdateLtNumWithIp(oltIp, 0)
}
func PlanLTPeriodically() {

	oltInfoWithLt, err := GetFirstOltInfo()
	if err != nil {
		fmt.Println("GetOltInfoWithIp failed, err:", err)
		return
	}
	var ltCnt = 0
	switch oltInfoWithLt.OltType {
	case "MF2":
		ltCnt = 2
	case "MF14":
		ltCnt = 14
	default:
		return
	}
	var ltArr = [14]string{oltInfoWithLt.Lt1Info, oltInfoWithLt.Lt2Info, oltInfoWithLt.Lt3Info, oltInfoWithLt.Lt4Info, oltInfoWithLt.Lt5Info, oltInfoWithLt.Lt6Info, oltInfoWithLt.Lt7Info, oltInfoWithLt.Lt8Info, oltInfoWithLt.Lt9Info, oltInfoWithLt.Lt10Info, oltInfoWithLt.Lt11Info, oltInfoWithLt.Lt12Info, oltInfoWithLt.Lt13Info, oltInfoWithLt.Lt14Info}
	sshConfig := &ssh.ClientConfig{
		User:            "admin",
		Auth:            []ssh.AuthMethod{ssh.Password("admin")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	s, err := netconf.DialSSHTimeout(oltInfoWithLt.IP+":832", sshConfig, 3*time.Second)
	if err != nil {
		fmt.Println("Dial SSH failed, err:", err)
		return
	}
	defer s.Close()
	sIhub, err := netconf.DialSSHTimeout(oltInfoWithLt.IP+":831", sshConfig, 3*time.Second)
	if err != nil {
		fmt.Println("Dial SSH failed, err:", err)
		return
	}
	defer sIhub.Close()
	var i = 0
	needUpdateDb := false
	for i = 1; i <= ltCnt; i++ {
		fmt.Println("lt" + strconv.Itoa(i) + ":" + ltArr[i-1] + "\r\n")
		ltArr[i-1] = strings.Replace(ltArr[i-1], "'", "\"", -1)
		var lTInfoStruct models.LTInfoStruct
		err := json.Unmarshal([]byte(ltArr[i-1]), &lTInfoStruct)
		if err != nil {
			continue
		} else {
			fmt.Println(lTInfoStruct.Planned)
		}
		ltIndex := models.LtIndex{
			LtIndex: strconv.Itoa(i),
		}
		tmpl, err := template.ParseFiles(utils.GetTemplatePath("getLtInfo.tpl"))
		if err != nil {
			fmt.Println("create template failed, err:", err)
			continue
		}
		buf := new(bytes.Buffer)
		tmpl.Execute(buf, ltIndex)
		fmt.Printf("request: %+v\r\n", buf.String())
		reply, err := s.Exec(netconf.RawMethod(buf.String()))
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Reply: %+v\r\n", reply.Data)
		if reply.Data == "<data></data>" {
			//fmt.Printf("lt" + strconv.Itoa(i) + " is null now no need plan")
			continue
		}
		v := models.LtBoardInfoData{}
		err = xml.Unmarshal([]byte(reply.Data), &v)
		if err != nil {
			fmt.Printf("error: %v\r\n", err)
			return
		}
		fmt.Printf("lt2 modelName: %+v\r\n", v.HardwareState.Component.ModelName)
		planLtInfo := models.PlanLtInfo{
			LtIndex:     strconv.Itoa(i),
			LtIndexAdd8: strconv.Itoa(i + 8 - 1),
			ModelName:   v.HardwareState.Component.ModelName,
		}
		tmpl, err = template.ParseFiles(utils.GetTemplatePath("planLTC.tpl"))
		if err != nil {
			fmt.Println("create template failed, err:", err)
			return
		}
		buf = new(bytes.Buffer)
		tmpl.Execute(buf, planLtInfo)
		fmt.Printf("plan request: %+v\r\n", buf.String())
		logger.OltLogger.Debug("plan request: %+v\r\n", buf.String())
		reply, err = s.Exec(netconf.RawMethod(buf.String()))
		if err != nil {
			fmt.Println(err)
			logger.OltLogger.Debug(err.Error())
			continue
		}
		logger.OltLogger.Debug("plan Reply: %+v\r\n", reply.Data)
		if reply.Data != "<ok/>" {
			fmt.Printf(" plan lt fail ")
			logger.OltLogger.Debug("plan lt fail")
		}

		planLtInfo.ModelName = utils.ConvertToLowerCase(v.HardwareState.Component.ModelName)
		tmpl, err = template.ParseFiles(utils.GetTemplatePath("planIhub.tpl"))
		if err != nil {
			fmt.Println("create template failed, err:", err)
			logger.OltLogger.Debug(err.Error())
			return
		}
		buf = new(bytes.Buffer)
		tmpl.Execute(buf, planLtInfo)
		fmt.Printf("plan ihub request: %+v\r\n", buf.String())
		logger.OltLogger.Debug("plan ihub request: %+v\r\n", buf.String())
		reply, err = sIhub.Exec(netconf.RawMethod(buf.String()))
		if err != nil {
			fmt.Println(err)
			logger.OltLogger.Debug(err.Error())
			continue
		}
		logger.OltLogger.Debug("plan ihub Reply: %+v\r\n", reply.Data)

		tmpl, err = template.ParseFiles(utils.GetTemplatePath("commit.tpl"))
		if err != nil {
			fmt.Println("create template failed, err:", err)
			logger.OltLogger.Debug(err.Error())
			return
		}
		buf = new(bytes.Buffer)
		tmpl.Execute(buf, planLtInfo)
		fmt.Printf("plan ihub commit request: %+v\r\n", buf.String())
		logger.OltLogger.Debug("plan ihub commit request: %+v\r\n", buf.String())
		reply, err = sIhub.Exec(netconf.RawMethod(buf.String()))
		if err != nil {
			fmt.Println(err)
			logger.OltLogger.Debug(err.Error())
			continue
		}
		logger.OltLogger.Debug("plan ihub commit Reply: %+v\r\n", reply.Data)
		needUpdateDb = true
	}
	if needUpdateDb {
		GetLtNum()
	}
}

func ResetLtInfoWhenActive() {
	var i = 0
	oltInfoWithLt, err := GetFirstOltInfo()
	if err != nil {
		fmt.Println("GetOltInfoWithIp failed, err:", err)
		return
	}
	for i = 1; i < 15; i++ {
		UpdateLtInfoWithIp(oltInfoWithLt.IP, "lt"+strconv.Itoa(i), "{'planned':'0'}")
	}
	SetLtNumToZero(oltInfoWithLt.IP)
}

func GetLtNum() {
	sshConfig := &ssh.ClientConfig{
		User:            "admin",
		Auth:            []ssh.AuthMethod{ssh.Password("admin")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	fmt.Printf("GetLtNum: \r\n")
	oltInfoWithLt, err := GetFirstOltInfo()
	if err != nil {
		fmt.Println("GetOltInfoWithIp failed, err:", err)
		return
	}
	var ltCnt = 0
	if oltInfoWithLt.OltType == "MF2" {
		ltCnt = 2
	} else if oltInfoWithLt.OltType == "MF14" {
		ltCnt = 14
	} else {
		return
	}
	var ltArr = [14]string{oltInfoWithLt.Lt1Info, oltInfoWithLt.Lt2Info, oltInfoWithLt.Lt3Info, oltInfoWithLt.Lt4Info, oltInfoWithLt.Lt5Info, oltInfoWithLt.Lt6Info, oltInfoWithLt.Lt7Info, oltInfoWithLt.Lt8Info, oltInfoWithLt.Lt9Info, oltInfoWithLt.Lt10Info, oltInfoWithLt.Lt11Info, oltInfoWithLt.Lt12Info, oltInfoWithLt.Lt13Info, oltInfoWithLt.Lt14Info}
	oltIp := oltInfoWithLt.IP
	s, err := netconf.DialSSHTimeout(oltIp+":832", sshConfig, 3*time.Second)
	if err != nil {
		fmt.Println("Dial SSH failed, err:", err)
		logger.OltLogger.Debug("dial ssh fail:", err.Error())
		return
	}
	defer s.Close()
	var i = 0
	totalLtNum := 0
	for i = 1; i <= ltCnt; i++ {
		fmt.Println("lt" + strconv.Itoa(i) + ":" + ltArr[i-1] + "\r\n")
		ltArr[i-1] = strings.Replace(ltArr[i-1], "'", "\"", -1)
		var lTInfoStruct models.LTInfoStruct
		err := json.Unmarshal([]byte(ltArr[i-1]), &lTInfoStruct)
		if err != nil {
			continue
		} else {
			fmt.Println(lTInfoStruct.Planned)
			//if lTInfoStruct.Planned == "1" {
			//      continue
			//}
		}
		ltIndex := models.LtIndex{
			LtIndex: strconv.Itoa(i),
		}
		tmpl, err := template.ParseFiles(utils.GetTemplatePath("getLtNum.tpl"))
		if err != nil {
			fmt.Println("create template failed, err:", err)
			return
		}
		buf := new(bytes.Buffer)

		tmpl.Execute(buf, ltIndex)
		fmt.Printf("request: %+v\r\n", buf.String())
		reply, err := s.Exec(netconf.RawMethod(buf.String()))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Reply: %+v\r\n", reply.Data)
			if reply.Data == "<data></data>" {
				fmt.Printf("lt" + strconv.Itoa(i) + " is null now no need plan")
			} else {
				totalLtNum = totalLtNum + 1
				UpdateLtInfoWithIp(oltIp, "lt"+strconv.Itoa(i), "{'planned':'1'}")
			}
		}

	}
	fmt.Printf("totalLtNum: %+v\r\n", totalLtNum)
	UpdateLtNumWithIp(oltIp, totalLtNum)
}

func TaskPeriod() {
	PlanLTPeriodically()
	// GetLtNum()
}

func TaskPeriod2() {
	GetLtNum()
}
