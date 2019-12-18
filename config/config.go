package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"hot/utils"
	"os"
)

func init() {
	viper.SetConfigName("config.yaml") //把json文件换成yaml文件，只需要配置文件名 (不带后缀)即可
	viper.AddConfigPath(".")           //添加配置文件所在的路径
	err := viper.ReadInConfig()
	if err != nil {
		utils.Error.Println(err)
		os.Exit(1)
	}
	viper.WatchConfig() //监听配置变化
	viper.OnConfigChange(func(e fsnotify.Event) {
		utils.Info.Println(e.Name)
	})
}
