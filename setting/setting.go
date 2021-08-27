package setting

import (
	"gopkg.in/ini.v1"
	"os"
)

var Conf = new(AppConfig)

// AppConfig 应用程序配置
type AppConfig struct {
	Release      bool `ini:"release"`
	Port         int  `ini:"port"`
	*MySQLConfig `ini:"mysql"`
}

// MySQLConfig 数据库配置
type MySQLConfig struct {
	User     string `ini:"user"`
	Password string `ini:"password"`
	DB       string `ini:"db"`
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
}

func Init(path string) error {

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	return ini.MapTo(Conf, file)
}
