package agent

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"net"
)

// Config is the global configuration varaible
var Config = viper.New()

// InitConfig set configuration
func InitConfig() {
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
	Config.SetDefault("protocol", "http")
	Config.SetDefault("host", getMainIP())
	Config.SetDefault("http_port", 4455)
	Config.SetDefault("https_port", 5544)
	Config.SetDefault("ssl_crt", "certificates/server.crt")
	Config.SetDefault("ssl_key", "certificates/server.key")
	Config.SetDefault("jwt_auth", false)
}

func getMainIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		Log.WithField("error", err).Error("unable to get local interfaces")
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return "localhost"
}
