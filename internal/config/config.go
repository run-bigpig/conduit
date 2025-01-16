package config

import (
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/run-bigpig/conduit/internal/constants"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func Init() {
	createConfigFile(constants.ConfigPath)
	viper.SetConfigFile(constants.ConfigPath) // 指定配置文件路径
	viper.AddConfigPath(".")                  // 还可以在工作目录中查找配置
	err := viper.ReadInConfig()               // 查找并读取配置文件
	if err != nil {                           // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func Write(key string, value any, writeFile bool) {
	viper.Set(key, value)
	if !writeFile {
		return
	}
	if err := viper.WriteConfig(); err != nil {
		log.Errorf("write config error: %s", err)
	}
}

func createConfigFile(filePath string) {
	//判断文件是否存在
	_, err := os.Stat(filePath)
	//判断是否是文件不存在
	if os.IsNotExist(err) {
		//创建文件夹
		err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
		if err != nil {
			log.Fatalf("create config file error: %s", err)
		}
		//创建文件
		file, err := os.Create(filePath)
		if err != nil {
			log.Errorf("create config file error: %s", err)
		}
		file.Close()
	}
}

func Read(key string) any {
	return viper.Get(key)
}

func ReadString(key string) string {
	return viper.GetString(key)
}

func ReadInt(key string) int {
	return viper.GetInt(key)
}
