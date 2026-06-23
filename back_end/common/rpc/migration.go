package rpc

import (
	logger "alto_server/common/log"
	"alto_server/common/models"
	"alto_server/common/utils"
	"alto_server/constants"
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func GetBackFile(rpcInfo RPCInfor, fileName string) error {

	reply, err := RunStaticFSRPC(rpcInfo, "", logger.SystemLogger, false)
	if err != nil {
		logger.SystemLogger.Error(err.Error())
		return err
	}
	v := FilterRPCReply{}
	err = xml.Unmarshal([]byte(reply.Data), &v)
	if err != nil || reply.Data == "<data></data>" {
		logger.SystemLogger.Error("no data")
		return err
	}
	tmpDirName := utils.GetSoftwarePath()

	// fmt.Printf("Reply: %+v", tmpDirName+tmpFileName)
	// newStr := strings.Replace(reply.Data, "<data>", "", 1)
	// newStr = strings.Replace(newStr, "</data>", "", 1)
	// newStr = xmlfmt.FormatXML(newStr, "", "  ")
	logger.SystemLogger.Debug("tmpFileName, ", fileName)
	//logger.SystemLogger.Debug("reply.Data, ", reply.Data)
	err = utils.ProcessRPCFile(reply.Data, tmpDirName+fileName, logger.SystemLogger)
	if err != nil {
		logger.SystemLogger.Error(err.Error())
		return err
	}
	return nil
}
func GetYmlFile(fileName string, neType string, targetVer string) error {
	logger.SystemLogger.Info(neType)
	ymlContent := fmt.Sprintf("ne_type: %s\nversion: %s\nhash:11000111\n", neType, targetVer)
	// Create the temporary YML file
	tmpYmlFile, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("Failed to create the temporary YML file: %w", err)
	}
	if _, err := tmpYmlFile.WriteString(ymlContent); err != nil {
		tmpYmlFile.Close()
		os.Remove(fileName)
		return fmt.Errorf("Failed to write content to the YML file: %w", err)
	}
	tmpYmlFile.Close()
	return nil
}
func GetRealDir(targetDir string, matchStr string) (string, error) {
	logger.SystemLogger.Infof("GetRealDir %s %s  ", targetDir, matchStr)
	entries, err := os.ReadDir(targetDir)
	if err != nil {
		fmt.Println("read fail:\n", err.Error())
		return "", err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			dirName := entry.Name()
			if strings.Contains(dirName, matchStr) {
				fmt.Println(dirName)
				return dirName, nil
			}
		}
	}
	return "", errors.New(targetDir + " not exist: " + matchStr)
}
func GetRealZipName(targetDir string, matchStr string) (string, error) {
	// 构建正则表达式：前缀-数字.数字-数字.zip
	// 注意：前缀可能包含特殊字符（如-），需先转义
	escapedPrefix := regexp.QuoteMeta(matchStr)
	pattern := fmt.Sprintf(`^%s-(\d+)\.(\d+)-(\d+)\.zip$`, escapedPrefix)
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", err
	}

	// 遍历目录下的所有文件
	entries, err := os.ReadDir(targetDir)
	if err != nil {
		return "", err
	}
	for _, entry := range entries {
		if !entry.IsDir() {
			filename := entry.Name()
			// 匹配正则
			if re.MatchString(filename) {
				return filename, nil
			}
		}
	}

	return "", errors.New(targetDir + " not exist: " + matchStr)

}

func GenerateMigrationFiles(card string, oltInfo models.Olt, fromV string, targetV string) (error, string, string) {
	oltId := oltInfo.Ip
	oltType := oltInfo.Type
	oltSeries := utils.GetOltSeriesByType(oltType)
	oltType = strings.ToUpper(oltType)
	oltSeries = strings.ToUpper(oltSeries)
	logger.SystemLogger.Infof("GenerateMigrationFiles %s %s %s %s ", card, oltId, fromV, targetV)
	port, err := utils.GetOLTPortIndexByCard2(card)
	if err != nil {
		logger.SystemLogger.Error(err.Error())
		return err, "", ""
	}
	rpcInfo := RPCInfor{
		IP:           oltId,
		Port:         port,
		TemplatePath: constants.TEMPLATE_FS + "backup.tpl",
	}
	if strings.Compare(card, "ihub") == 0 {
		rpcInfo.TemplatePath = constants.TEMPLATE_FS + "backupIhub.tpl"
	}
	//file in the zip is not allowded to use .
	oltIdForName := strings.ReplaceAll(oltId, ".", "_")
	tmpDirName := utils.GetSoftwarePath()
	oldXmlName := oltIdForName + "_" + port + "_old.xml"
	newYmlName := oltIdForName + "_" + port + "_new.yml"
	newXmlName := oltIdForName + "_" + port + "_new.xml"
	newXmlPath := tmpDirName + newXmlName
	oldXmlPath := tmpDirName + oldXmlName
	newYmlPath := tmpDirName + newYmlName
	err = GetBackFile(rpcInfo, oldXmlName)
	if err != nil {
		logger.SystemLogger.Debug("GetBackFile err = ", err)
		return err, "", ""
	}
	realDirName, err := GenerateMigrationDirName(oltInfo, card)
	if err != nil {
		logger.SystemLogger.Debug("GetBackFile err = ", err)
		return err, "", ""
	}
	fmt.Println(realDirName)
	err = utils.ExecuteXSLTransformations(oldXmlPath, newXmlPath, fromV, targetV, logger.SystemLogger, realDirName)
	if err != nil {
		fmt.Println(err.Error())
		logger.SystemLogger.Debug("ExecuteXSLTransformations err = ", err)
		return err, "", ""
	}
	neType := ""

	if strings.Compare(port, "831") == 0 {
		ntType, err := GetBoardType(oltId, "nt")
		if err != nil {
			logger.SystemLogger.Debug("GetBoardType err = ", err)
			return err, "", ""
		}
		neType = "LS-" + oltSeries + "-IHUB-" + ntType //"LS-MF-IHUB-LMNT-B"
	} else if strings.Compare(port, "832") == 0 {
		ntType, err := GetBoardType(oltId, "nt")
		if err != nil {
			logger.SystemLogger.Debug("GetBoardType err = ", err)
			return err, "", ""
		}
		neType = "LS-" + oltSeries + "-" + ntType //"LS-MF-LMNT-B"
	} else if strings.Compare(port, "830") == 0 {
		neType = "LS-DF-CFXR-E"
	} else {
		ltType, err := GetBoardType(oltId, card)
		if err != nil {
			logger.SystemLogger.Debug("GetBoardType err = ", err)
			return err, "", ""
		}
		neType = "LS-" + oltSeries + "-" + ltType //"LS-MF-LWLT-C"
	}
	//generate yml
	err = GetYmlFile(newYmlPath, neType, targetV)
	if err != nil {
		logger.SystemLogger.Debug("GetYmlFile err = ", err)
		return err, "", ""
	}
	return nil, newXmlPath, newYmlPath
}
func GetBoardType(oltId string, boardType string) (string, error) {
	templatePath := ""
	ltIndex := models.LtIndex{
		LtIndex: "1",
	}
	if strings.Compare(boardType, "Chassis") == 0 {
		templatePath = constants.TEMPLATE_FS + "getChassisType.tpl"
	} else if strings.Compare(boardType, "nt") == 0 {
		templatePath = constants.TEMPLATE_FS + "getNtType.tpl"
	} else if strings.Contains(boardType, "lt") {
		ltIndex.LtIndex = boardType[2:]
		templatePath = constants.TEMPLATE_FS + "getLtType.tpl"
	}
	rpcInfo := RPCInfor{
		IP:           oltId,
		Port:         "832",
		TemplatePath: templatePath,
	}
	reply, err := RunStaticFSRPC(rpcInfo, ltIndex, logger.SystemLogger, true)
	if err != nil {
		logger.SystemLogger.Error(err.Error())
		return "", err
	}
	v := FilterRPCReply{}
	err = xml.Unmarshal([]byte(reply.Data), &v)
	if err != nil || reply.Data == "<data></data>" {
		logger.SystemLogger.Error("no data")
		return "", err
	}
	tmpType := v.HardwareState.Component[0].ModelName
	//fmt.Println(tmpType)
	return tmpType, nil
}
func GenerateMigrationDirName(oltInfo models.Olt, card string) (string, error) {
	dirName := ""
	oltId := oltInfo.Ip
	oltType := oltInfo.Type
	oltSeries := utils.GetOltSeriesByType(oltType)
	oltType = strings.ToLower(oltType)
	oltSeries = strings.ToLower(oltSeries)
	zipPath := utils.GetMigrationToolPrefixPath()
	extractPath := utils.GetMigrationToolPrefixPath() + constants.MIGRATION_TOOL_EXTRATED
	if strings.Compare(card, "nt") == 0 {
		chassisType, err := GetBoardType(oltId, "Chassis")
		if err != nil {
			logger.SystemLogger.Debug("GetBoardType err = ", err)
			return "", err
		}
		ntType, err := GetBoardType(oltId, "nt")
		if err != nil {
			logger.SystemLogger.Debug("GetBoardType err = ", err)
			return "", err
		}
		//device-extension-ls-mf-ihub-lmnt-b-mf2-lmxr-b-25.9-459
		dirName = "device-extension-ls-" + oltSeries + "-" + strings.ToLower(ntType) + "-" + oltType + "-" + strings.ToLower(chassisType)
		zipName, err := GetRealZipName(zipPath, dirName)
		if err != nil {
			//find, if not find, use second name: device-extension-ls-mf-lmnt-b-25.9-459
			dirName = "device-extension-ls-" + oltSeries + "-" + strings.ToLower(ntType)
			zipName, err = GetRealZipName(zipPath, dirName)
			if err != nil {
				logger.SystemLogger.Debug("GetRealZipName err = ", err)
				return "", err
			}
			targetName := zipName[:len(zipName)-4]
			//check if dir exist, no action, or unzip
			err := utils.UncompressFiles(zipPath+zipName, extractPath+targetName)
			if err != nil {
				return "", err
			} else {
				return extractPath + targetName, nil
			}
		} else {
			targetName := zipName[:len(zipName)-4]
			//check if dir exist, no action, or unzip
			err := utils.UncompressFiles(zipPath+zipName, extractPath+targetName)
			if err != nil {
				return "", err
			} else {
				return extractPath + targetName, nil
			}
		}

		//then unzip and return
	} else if strings.Contains(card, "lt") {
		ltType, err := GetBoardType(oltId, card)
		if err != nil {
			logger.SystemLogger.Debug("GetBoardType err = ", err)
			return "", err
		}
		//device-extension-ls-mf-lwlt-c-25.9-459
		dirName = "device-extension-ls-" + oltSeries + "-" + strings.ToLower(ltType)
		zipName, err := GetRealZipName(zipPath, dirName)
		if err != nil {
			return "", err
		} else {
			targetName := zipName[:len(zipName)-4]
			//check if dir exist, no action, or unzip
			err := utils.UncompressFiles(zipPath+zipName, extractPath+targetName)
			if err != nil {
				return "", err
			} else {
				return extractPath + targetName, nil
			}
		}
	} else if strings.Compare(card, "ihub") == 0 {
		chassisType, err := GetBoardType(oltId, "Chassis")
		if err != nil {
			logger.SystemLogger.Debug("GetBoardType err = ", err)
			return "", err
		}
		ntType, err := GetBoardType(oltId, "nt")
		if err != nil {
			logger.SystemLogger.Debug("GetBoardType err = ", err)
			return "", err
		}
		//device-extension-ls-mf-ihub-lmnt-b-mf2-lmxr-b-25.9-459
		dirName = "device-extension-ls-" + oltSeries + "-" + "ihub-" + strings.ToLower(ntType) + "-" + oltType + "-" + strings.ToLower(chassisType)

		zipName, err := GetRealZipName(zipPath, dirName)
		if err != nil {
			//device-extension-ls-mf-ihub-lmnt-b-25.9-459
			dirName = "device-extension-ls-" + oltSeries + "-" + "ihub-" + strings.ToLower(ntType)
			zipName, err = GetRealZipName(zipPath, dirName)
			if err != nil {
				logger.SystemLogger.Debug("GetRealZipName err = ", err)
				return "", err
			}
			targetName := zipName[:len(zipName)-4]
			//check if dir exist, no action, or unzip
			err := utils.UncompressFiles(zipPath+zipName, extractPath+targetName)
			if err != nil {
				return "", err
			} else {
				return extractPath + targetName, nil
			}
		} else {
			targetName := zipName[:len(zipName)-4]
			//check if dir exist, no action, or unzip
			err := utils.UncompressFiles(zipPath+zipName, extractPath+targetName)
			if err != nil {
				return "", err
			} else {
				return extractPath + targetName, nil
			}
		}
	}
	logger.SystemLogger.Debug(dirName)
	return dirName, nil
}
func MigrateRunningConfiguration(oltInfo models.Olt, card string, fromV string, targetV string) error {
	oltId := oltInfo.Ip
	//oltType := oltInfo.Type
	//oltSeries := utils.GetOltSeriesByType(oltType)
	port, err := utils.GetOLTPortIndexByCard2(card)
	if err != nil {
		logger.SystemLogger.Error(err.Error())
		return err
	}
	logger.SystemLogger.Errorf("MigrateRunningConfiguration %s %s %s %s %s", oltId, card, port, fromV, targetV)
	tmpDirName := utils.GetSoftwarePath()
	err, xmlPath, ymlPath := GenerateMigrationFiles(card, oltInfo, fromV, targetV)
	if err != nil {
		logger.SystemLogger.Debug("GenerateMigrationFiles err = ", err)
		return err
	}

	//zip
	zipName := oltId + "_" + card + "_migration.zip"
	zipPath := tmpDirName + zipName
	inputArray := []string{xmlPath, ymlPath}
	if strings.Compare(card, "nt") == 0 {
		//for nt, need also get ihub part, then zip with nt file together
		port = "831"
		card = "ihub"
		err, xmlPath, ymlPath := GenerateMigrationFiles(card, oltInfo, fromV, targetV)
		if err != nil {
			logger.SystemLogger.Debug("GenerateMigrationFiles err = ", err)
			return err
		}
		inputArray = append(inputArray, xmlPath)
		inputArray = append(inputArray, ymlPath)
	}
	err = utils.GenerateZip(inputArray, zipPath)
	if err != nil {
		logger.SystemLogger.Debug("GenerateConfigurationPackages err = ", err)
		return err
	}

	return nil

}
func DownloadMigrationConfig(oltId string, port string, softName string, zipName string) {
	var serverLocalIp string = utils.FindLocalIp("192.168.1.")
	var localPort string = "5600"
	var prefixStr string = "http://" + serverLocalIp + ":" + localPort + "/software/" + zipName

	softwareDownloadInfo := models.SoftwareDownloadInfo{
		SoftwareName: softName,
		SoftwareUrl:  prefixStr,
	}
	fmt.Println(prefixStr)
	rpcInfo := RPCInfor{
		IP:           oltId,
		Port:         port,
		TemplatePath: constants.TEMPLATE_FS + "downloadMigration.tpl",
	}
	_, err := RunStaticFSRPC(rpcInfo, softwareDownloadInfo, logger.SystemLogger, true)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func GetDownloadMigrationConfig(oltId string, port string) {

	rpcInfo := RPCInfor{
		IP:           oltId,
		Port:         port,
		TemplatePath: constants.TEMPLATE_FS + "getMigrationStatus.tpl",
	}
	reply, err := RunStaticFSRPC(rpcInfo, "", logger.SystemLogger, true)
	if err != nil {
		fmt.Println(err.Error())
	}
	v := FilterRPCReply{}
	err = xml.Unmarshal([]byte(reply.Data), &v)
	if err != nil || reply.Data == "<data></data>" {
		logger.SystemLogger.Error("no data")
	}
	fmt.Println("CurrentState", v.HardwareState.Component[0].SoftwaresOlt.SoftwareOlt[0].ConfigDownload.CurrentState.State)
	fmt.Println("LastDownloadState:", v.HardwareState.Component[0].SoftwaresOlt.SoftwareOlt[0].ConfigDownload.LastDownloadState.State)
	fmt.Println("LastDownloadState:", v.HardwareState.Component[0].SoftwaresOlt.SoftwareOlt[0].ConfigDownload.LastDownloadState.SoftwareName)
}

func GetCurPlatformVersion(oltId string, port string) (string, error) {
	version := ""
	rpcInfo := RPCInfor{
		IP:           oltId,
		Port:         port,
		TemplatePath: constants.TEMPLATE_FS + "getSoftwareStatusWithPlatform.tpl",
	}

	reply, err := RunStaticFSRPC(rpcInfo, "", logger.SystemLogger, false)
	if err != nil {
		return "", err
	}
	// fmt.Println("reply  = ", reply.Data)
	v, err := UnmarshalReply(reply)
	// v := FilterRPCReply{}
	// err = xml.Unmarshal([]byte(reply.Data), &v)
	if err != nil {
		fmt.Println("no data")
		return "", err
	}
	version = v.SystemState.Platform.SoftwareRelease.Name
	fmt.Println("GetCurPlatformVersion  = ", version)
	return version, nil
}
func GetCurStandbyVersionName(oltId string, port string) (string, error) {
	version := ""
	rpcInfo := RPCInfor{
		IP:           oltId,
		Port:         port,
		TemplatePath: constants.TEMPLATE_FS + "getSoftwareStatusWithPlatform.tpl",
	}

	reply, err := RunStaticFSRPC(rpcInfo, "", logger.SystemLogger, false)
	if err != nil {
		return "", err
	}
	// fmt.Println("reply  = ", reply.Data)
	v, err := UnmarshalReply(reply)
	// v := FilterRPCReply{}
	// err = xml.Unmarshal([]byte(reply.Data), &v)
	if err != nil {
		fmt.Println("no data")
		return "", err
	}
	if strings.Compare(v.HardwareState.Component[0].SoftwaresOlt.SoftwareOlt[0].Revisions.Revision[0].IsActive, "true") == 0 {
		version = v.HardwareState.Component[0].SoftwaresOlt.SoftwareOlt[0].Revisions.Revision[1].Name
	} else {
		version = v.HardwareState.Component[0].SoftwaresOlt.SoftwareOlt[0].Revisions.Revision[0].Name
	}
	fmt.Println("GetCurStandbyVersion  = ", version)
	return version, nil
}
