package gcfg

import (
	logger "alto_server/common/log"
	"alto_server/conf"
	"fmt"
)

type globalCfg struct {
	PageSize       int  `ini:"PAGE_SIZE"`
	RPCLogOutput   bool `ini:"RPC_LOG_OUTPUT"`
	NetconfTimeout int  `ini:"NETCONF_TIME_OUT"`
}

func NewGlobalCfg() *globalCfg {
	a := new(globalCfg)
	a.PageSize = 10
	a.RPCLogOutput = false
	a.NetconfTimeout = 3
	return a
}

var instance *globalCfg

func InitGlobalCfg() *globalCfg {
	if instance == nil { // ensure init one times only
		instance = NewGlobalCfg()
		err := conf.Cfg.Section("app").MapTo(&instance)
		if err != nil {
			fmt.Println("load global config failed!  err", err)
		}
	}
	return instance
}

func GetGlobalCfg() *globalCfg {
	if instance == nil {
		instance = NewGlobalCfg()
		fmt.Println("get global list failed!  return default list")
		logger.SystemLogger.Error("get global list failed!  return default list")
		// panic("Feature list has not been initialized!")
	}
	return instance
}
