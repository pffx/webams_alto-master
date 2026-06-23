package feature

import (
	"alto_server/common/injection"
	logger "alto_server/common/log"
	"alto_server/conf"
	"fmt"
	"reflect"
)

type featurelist struct {
	OPID string // default OPID is NKBB

	WEBGUI_DEPLOYMENT bool
}

func NewFeaturelist() *featurelist {
	a := new(featurelist)
	// a.WEBGUI_DEPLOYMENT = false
	value := reflect.ValueOf(a)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		if field.CanSet() {
			switch field.Kind() {
			case reflect.Int:
				field.SetInt(0)
			case reflect.String:
				field.SetString("false")
			case reflect.Bool:
				//set all bool params to false as default value
				field.SetBool(false)
			}
		}
	}
	a.OPID = injection.OPID
	return a
}

var instance *featurelist

func InitFeatureList() *featurelist {
	if instance == nil { // ensure init one times only
		instance = NewFeaturelist()
		err := conf.Feature.Section("featurelist").MapTo(&instance)
		if err != nil {
			fmt.Println("load feature config failed!  err", err)
		}
	}
	return instance
}

func GetFeatureList() *featurelist {
	if instance == nil {
		instance = NewFeaturelist()
		fmt.Println("get feature list failed!  return default list")
		logger.SystemLogger.Error("get feature list failed!  return default list")
		// panic("Feature list has not been initialized!")
	}
	return instance
}
