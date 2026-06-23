package main

import (
	. "alto_server/common/db"
	"alto_server/common/feature"
	logger "alto_server/common/log"
	"alto_server/conf"
	"alto_server/routers"
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

// func runSwagInit() {
// 	//swag init
// 	cmd := exec.Command("swag", "init")
// 	fmt.Println("Cmd", cmd.Args)
// 	var out bytes.Buffer
// 	cmd.Stdout = &out
// 	cmd.Stderr = os.Stderr
// 	err := cmd.Start()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(out.String())
// }

/*
func loadInitialOltList() {
	fmt.Println("loadInitialOltList")
	index := 1
	db := DbHandle
	for {
		key := "OLT" + strconv.Itoa(index)
		if conf.Cfg.Section("olt").HasKey(key) {
			info, _ := conf.Cfg.Section("olt").GetKey(key)
			ss := strings.Split(info.String(), "/")
			fmt.Printf("oltInfo: %+v \n", ss)
			olt := models.OltInfo{}
			olt.IP = ss[0]
			olt.Username = ss[1]
			olt.Password = ss[2]

			var dbIdx int
			var oltName string
			var oltIp string
			var oltPort int
			var userName string
			var password string
			var oltType string
			var ntList string
			var ltList string
			rows, err := db.Query("select * from olt")
			PanickErr(err)


			for rows.Next() {
				if hasRecord == false {
					rows.Scan(&dbIdx, &oltName, &oltIp, &oltPort, &userName, &password, &oltType, &ntList, &ltList)
				}
				if oltIp == olt.IP {
					fmt.Println("find the olt")
					hasRecord = true
				}
			}
			if hasRecord {
				fmt.Println("hasRecord " + strconv.Itoa(dbIdx) + " " + password + " " + oltIp)
				stmt, err := db.Prepare("update olt set user_name=?,passwd=? where id= ?")
				PanickErr(err)

				res, err := stmt.Exec(olt.Username, olt.Password, dbIdx)
				PanickErr(err)
				fmt.Println(res)
			}
			index++
		} else {
			break
		}
	}

}
*/
// @title Alto Restful API
// @version 1.0 版本
// @description Alto Restful description
// @BasePath /nokia-alto
// @query.collection.format multi
// @license.name Apache 2.0
// @contact.name   NPI Support
// @contact.url    http://localhost:5600/help/index.html#/
// @contact.email  support@nokia.com
// @tag.name Auth
// @tag.description User authorisation
// @tag.name User
// @tag.description User information management
// @tag.name System
// @tag.description System API
// @tag.name Test
// @tag.description Test API
// @tag.name V1
// @tag.description  API of V1 version
func main() {
	logfile, err := os.Create("logs/gin_console.log")
	if err != nil {
		fmt.Println("Could not create log file")
	}
	logger.InitLogger()
	//gin.SetMode(gin.DebugMode)
	gin.DefaultWriter = io.MultiWriter(logfile)
	runMode, err := conf.Cfg.Section("").GetKey("RUN_MODE")
	// if err == nil {
	// 	gin.SetMode(runMode.String())
	// 	if !IsReleaseMode(runMode.String()) {
	// 		runSwagInit()
	// 	}
	// } else {
	if err != nil {
		fmt.Println("load config failed!")
	}
	fmt.Printf("Alto is running in %+v mode now! \n", runMode)
	feature.InitFeatureList()
	InitSqliteConn()
	// port, err := conf.Cfg.Section("server").GetKey("HTTP_PORT")

	// if err != nil {
	// 	fmt.Println("load config failed!")
	// }
	// need check OLT info in ini file to update to db
	//loadInitialOltList()
	c := cron.New()
	c.AddFunc("0 */1 * * *", TaskPeriod)
	c.AddFunc("0 */10 * * *", TaskPeriod2)
	c.Start()
	r := routers.InitRouter()
	// fmt.Printf("Alto is running on port %+v now! \n", port)
	// r.Run(":" + port.String())
	r.Run(":5600")

}
