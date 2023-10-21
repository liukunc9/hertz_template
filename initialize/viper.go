package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/liukunc9/hertz_template/config"
	"github.com/liukunc9/hertz_template/global"
	"github.com/spf13/viper"
)

func initViper() {
	v := viper.New()
	v.SetConfigFile(config.DefaultConfigFile)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.Config); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.Config); err != nil {
		fmt.Println(err)
	}

	global.Viper = v
}
