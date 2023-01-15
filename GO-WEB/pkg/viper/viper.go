package viper

import (
	"github.com/robot007num/go/go-web/model/config"
	"github.com/spf13/viper"
)

func InitViper() {
	allConfig := config.AllConfig{}

	viper := viper.New()
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic("Viper readConfig err: " + err.Error())
	}

	if err := viper.Unmarshal(&allConfig); err != nil {
		panic("Viper Unmarshal err: " + err.Error())
	}

	//fmt.Printf("%+v\n", allConfig)
	allConfig.InitAllConfig()
}
