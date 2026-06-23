package conf

import (
	"fmt"
	"os"

	"github.com/go-ini/ini"
)

var (
	// Cfg .
	Cfg     *ini.File
	Feature *ini.File
)

// 加载配置文件
func init() {
	var err error
	// Cfg, err = ini.Load("conf/app.ini")
	Cfg, err = ini.LoadSources(ini.LoadOptions{
		AllowPythonMultilineValues: true,
		IgnoreInlineComment:        true,
	}, "conf/app.ini")
	if err != nil {
		fmt.Println("loade config err")
		os.Exit(1)
	}
	Feature, err = ini.Load("conf/feature.ini")
	if err != nil {
		fmt.Println("loade feature err")
		os.Exit(1)
	}
}
