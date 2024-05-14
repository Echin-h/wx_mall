package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"wxapp/pkg/fs"
)

var config *Conf

func LoadConfig() {
	workDir, _ := os.Getwd()
	if !fs.FileExists(workDir + "/config/config.yaml") {
		panic(fmt.Errorf("config fs not found"))
	}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config fs: %s \n", err))
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("Config Unmalshal error: %s \n", err))
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		err := viper.Unmarshal(&config)
		if err != nil {
			panic(fmt.Errorf("Config Unmalshal error: %s \n", err))
		}
	})
	viper.WatchConfig()
}

func GetConfig() *Conf {
	return config
}
