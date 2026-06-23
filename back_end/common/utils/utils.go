package utils

import (
	"alto_server/common/pkg/e"
	"alto_server/conf"
	"alto_server/constants"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
	"time"
	"unicode"

	"github.com/gin-gonic/gin"
)

func ConvertToLowerCase(inputStr string) string {
	r := []rune(inputStr)
	for i, c := range r {
		r[i] = unicode.ToLower(c)
	}
	lowerCaseStr := string(r)
	fmt.Println(lowerCaseStr)
	return lowerCaseStr
}

// RES 返回信息自动根据code插入message
func RES(c *gin.Context, code int, obj gin.H) {
	if obj["message"] == "" {
		obj["message"] = e.GetMessage(code)
	}
	obj["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	c.JSON(code, obj)
}

// Abort next middleware.
func AbortWithStatusJSON(c *gin.Context, status int, error_num int, message string) {
	c.AbortWithStatusJSON(status, gin.H{
		"status":  error_num,
		"message": message,
	})
	// c.AbortWithStatusJSON(e.SUCCESS, gin.H{
	// 	"status":  e.ERROR_AUTH,
	// 	"message": e.GetMessage(e.ERROR_AUTH),
	// })
}

func PanickErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
func ArrayToString(arr []string) string {
	var result string
	for _, i := range arr { //遍历数组中所有元素追加成string
		result += i
	}
	return result
}

func IsDebugMode(mode string) bool {
	if mode == "debug" {
		return true
	} else {
		return false
	}
}

func IsReleaseMode(mode string) bool {
	if mode == "release" {
		return true
	} else {
		return false
	}
}
func IsTestMode(mode string) bool {
	if mode == "test" {
		return true
	} else {
		return false
	}
}

//	func GetFilePathHeader() string {
//		electron_mode, err := conf.Cfg.Section("").GetKey("ELECTRON_MODE")
//		if err != nil {
//			return "resources/"
//		}
//		if electron_mode.String() == "false" {
//			return "./"
//		} else {
//			return "resources/"
//		}
//	}
func GetReactAppFilePath() string {
	electron_mode, err := conf.Cfg.Section("").GetKey("ELECTRON_MODE")
	if err != nil {
		return "resources/build"
	}
	if electron_mode.String() == "false" {
		return "./app/build"
	} else {
		return "resources/build"
	}
}
func GetLogsFilePath(file string) string {
	electron_mode, err := conf.Cfg.Section("").GetKey("ELECTRON_MODE")
	if err != nil {
		return "logs/" + file
	}
	if electron_mode.String() == "false" {
		return "./logs/" + file
	} else {
		return "logs/" + file
	}
}
func GetDatabasePath() string {
	electron_mode, err := conf.Cfg.Section("").GetKey("ELECTRON_MODE")
	if err != nil {
		return "resources/db/alto.db"
	}
	if electron_mode.String() == "false" {
		return "./db/alto.db"
	} else {
		return "resources/db/alto.db"
	}
}
func GetConfIniPath() string {
	electron_mode, err := conf.Cfg.Section("").GetKey("ELECTRON_MODE")
	if err != nil {
		return "conf/app.ini"
	}
	if electron_mode.String() == "false" {
		return "./conf/app.ini"
	} else {
		return "conf/app.ini"
	}
}
func GetUploadPath() string {
	electron_mode, err := conf.Cfg.Section("").GetKey("ELECTRON_MODE")
	if err != nil {
		return "resources/upload/"
	}
	if electron_mode.String() == "false" {
		return "./upload/"
	} else {
		return "resources/upload/"
	}
}
func GetSoftwarePath() string {
	electron_mode, err := conf.Cfg.Section("").GetKey("ELECTRON_MODE")
	if err != nil {
		return "software/"
	}
	if electron_mode.String() == "false" {
		return "./software/"
	} else {
		return "software/"
	}
}
func GetTemplatePath(file string) string {
	electron_mode, err := conf.Cfg.Section("").GetKey("ELECTRON_MODE")
	if err != nil {
		return "resources/template/" + file
	}
	if electron_mode.String() == "false" {
		return "./template/" + file
	} else {
		return "resources/template/" + file
	}
}
func GetMigrationToolPrefixPath() string {
	electron_mode, err := conf.Cfg.Section("").GetKey("ELECTRON_MODE")
	if err != nil {
		return constants.ELECTRON_MIGRATION_TOOL_XSL_PREFIX
	}
	if electron_mode.String() == "false" {
		return constants.MIGRATION_TOOL_XSL_PREFIX
	} else {
		return constants.ELECTRON_MIGRATION_TOOL_XSL_PREFIX
	}
}
func MapStringStringToJsonString(input map[string]string) string {
	bytes, _ := json.Marshal(input)
	result := string(bytes)
	result = strings.Replace(result, "\\", "", -1)
	result = strings.Replace(result, "\"{", "{", -1)
	result = strings.Replace(result, "}\"", "}", -1)
	return result
}

func MapIntStringToJsonString(input map[int]string) string {
	bytes, _ := json.Marshal(input)
	result := string(bytes)
	result = strings.Replace(result, "\\", "", -1)
	result = strings.Replace(result, "\"{", "{", -1)
	result = strings.Replace(result, "}\"", "}", -1)
	return result
}
func JsonToString(input []byte) string {
	result := string(input)
	result = strings.Replace(result, "\\", "", -1)
	result = strings.Replace(result, "\"{", "{", -1)
	result = strings.Replace(result, "}\"", "}", -1)
	return result
}

func FindLocalIp(matchPattern string) string {
	var finalIp string
	cmd := exec.Command("cmd", "/c", "ipconfig")
	if out, err := cmd.StdoutPipe(); err != nil {
		fmt.Println(err)
	} else {
		defer out.Close()
		if err := cmd.Start(); err != nil {
			fmt.Println(err)
		}
		if opBytes, err := ioutil.ReadAll(out); err != nil {
			log.Fatal(err)
		} else {
			str := string(opBytes)
			var strs = strings.Split(str, "\r\n")
			if 0 != len(strs) {
				for _, value := range strs {

					vidx := strings.Index(value, "IPv4")
					//说明已经找到该ip
					if vidx != -1 {

						ip4lines := strings.Split(value, ":")
						if len(ip4lines) == 2 {
							ip4str := ip4lines[1]
							finalIp = strings.TrimSpace(ip4str)
							fmt.Printf(finalIp + "\r\n")
							if strings.Contains(ip4str, matchPattern) {
								fmt.Printf(finalIp + "\r\n")
								return finalIp
							}
						}
					}
				}
			}
		}
	}
	return ""
}

//	func GetNtPortFromIp(oltIp string) string {
//		var node models.Node
//		result := db.DbHandle.Where("olt_ip = ?", oltIp).First(&node)
//		if result.Error != nil {
//			return constants.NT_PORT
//		}
//		if strings.Contains(node.OltType, "DF") {
//			return constants.DF_PORT
//		} else {
//			return constants.NT_PORT
//		}
//	}
