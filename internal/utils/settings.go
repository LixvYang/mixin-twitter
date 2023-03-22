package utils

import (
	"fmt"
	"os"
	"sync"

	"github.com/spf13/viper"
)

var InitSetting sync.Once

type Config struct {
	Server   Server
	Database Database
	Mixin    Mixin
}

type Server struct {
	DomainName string
	IP         string
	AppMode    string
	HttpPort   string
}

type Database struct {
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
}

type Mixin struct {
	Pin        string
	ClientId   string
	SessionId  string
	PinToken   string
	PrivateKey string
	AppSecret  string
}

var Conf Config

func Init() {
	viper.AddConfigPath(".")
	RootDir, _ := os.Getwd()
	viper.SetConfigFile(fmt.Sprintf("%s/configs/env.yaml", RootDir))
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("Fatal to unmarshal config file: %s \n", err))
	}
}
