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
	initLogger()

	if err != nil {
		Log.Warn("no configuration file found")
	} else {
		Log.Info("configuration file loaded")
		Config.WatchConfig()
		Config.OnConfigChange(func(e fsnotify.Event) {
			Log.Info("configuration file changed:", e.Name)
		})
	}
}

func setDefault() {
	consoleLog := Logger{Type: "console", Format: "text", Level: "debug"}
	Config.SetDefault("loggers", []Logger{consoleLog})
	Config.SetDefault("protocol", "https")
	Config.SetDefault("host", "localhost")
	Config.SetDefault("http_port", 8080)
	Config.SetDefault("https_port", 9090)
	Config.SetDefault("ssl_crt", "certificates/server.crt")
	Config.SetDefault("ssl_key", "certificates/server.key")
}
