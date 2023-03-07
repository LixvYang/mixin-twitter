package utils

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

var InitSetting sync.Once

func Init() {
	viper.AddConfigPath(".")
	viper.SetConfigFile("env.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Println(viper.Get("database.DbName"))
}
