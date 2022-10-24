package gobalConfig

import (
	"github.com/spf13/viper"
	"log"
)

var (
	ServerPort string
	FrontMode  bool
	GinMode    string
)

func init() {
	log.Println("正在应用配置文件...")
	viper.SetConfigName("conf")
	viper.SetConfigType("yml")
	viper.AddConfigPath("conf")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicln("viper load fail ...")
		return
	}

	ServerPort = viper.GetString("server.port")
	GinMode = viper.GetString("server.ginMode")
	FrontMode = viper.GetBool("config.frontMode")
}
