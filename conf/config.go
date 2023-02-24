package conf

import (
	"fmt"
	"github.com/go-jwt/model"
	"github.com/spf13/viper"
)

var (
	user   string
	pass   string
	ip     string
	port   string
	dbname string
)

func ConfigInit() {
	//使用viper读取配置文件
	viper.SetConfigName("config")  // name of config file (without extension)
	viper.SetConfigType("yaml")    // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./conf/") // path to look for the config file in
	err := viper.ReadInConfig()    // Find and read the config file
	if err != nil {                // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	//初始化MySQL
	ReadMySqlConf()
	dsn := user + ":" + pass + "@tcp(" + ip + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	model.MySqlInit(dsn)
}

func ReadMySqlConf() {
	user = viper.GetString("MySql.user")
	pass = viper.GetString("MySql.pass")
	ip = viper.GetString("MySql.ip")
	port = viper.GetString("MySql.port")
	dbname = viper.GetString("MySql.dbname")
}
