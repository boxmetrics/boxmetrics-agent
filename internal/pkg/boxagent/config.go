package boxagent

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Config is the global configuration varaible
var Config = viper.New()

// SetConfig set configuration
func SetConfig() {
	Config.SetConfigName("boxagent")
	Config.SetEnvPrefix("boxagent")
	Config.AutomaticEnv()

	Config.AddConfigPath("/etc/boxmetrics/")
	Config.AddConfigPath("$HOME/.boxmetrics/")
	Config.AddConfigPath("./configs/")
	Config.AddConfigPath(".")
	setDefault()

	err := Config.ReadInConfig()
	if err != nil {
		fmt.Println("No configuration file loaded")
	} else {
		Config.WatchConfig()
		Config.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("Config file changed:", e.Name)
		})
	}
}

func setDefault() {
	Config.SetDefault("test", "fefaef")
}
