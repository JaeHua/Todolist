package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	AppConf *AppConfig
	SqlConn string
}

type AppConfig struct {
	Name string
	Port string
	Mode string
}

func InitConf() *Config {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("conf")
	viper.SetConfigName("app")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("找不到配置文件..")
		} else {
			fmt.Println("配置文件出错..")
		}
	}

	return &Config{
		AppConf: &AppConfig{
			Name: viper.GetString("app.name"),
			Port: viper.GetString("app.port"),
			Mode: viper.GetString("app.mode"),
		},
		SqlConn: viper.GetString("mysql.conn"),
	}
}
