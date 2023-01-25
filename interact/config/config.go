package config

import (
	"github.com/spf13/viper"
	"os"
)

func InitConfig() {
	// 获取当前工作路径
	workDir, _ := os.Getwd()
	// 配置文件的文件名
	viper.SetConfigName("config")
	// 配置文件的文件类型
	viper.SetConfigType("yml")
	// 配置文件所在路径(绝对路径) (拼接工作路径即可)
	viper.AddConfigPath(workDir + "/config")
	// 读取配置
	err := viper.ReadInConfig()
	if err != nil {
		return
	}

}
