package main

import (
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/boxmetrics/boxmetrics-agent/internal/pkg/boxagent"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	boxagent.InitConfig()

	http.HandleFunc("/ws/v1", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

		if err != nil {
			boxagent.Log.WithField("error", err).Error("websocket error")
		}

		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			// Log message

			logfields := logrus.Fields{"remote": conn.RemoteAddr(), "text": string(msg)}
			boxagent.Log.WithFields(logfields).Info("receive")

			// Write message back to browser
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/websocket.html")
	})

	var httpErr error
	protocol := boxagent.Config.GetString("protocol")
	host := boxagent.Config.GetString("host")
	port := boxagent.Config.GetInt(strings.Join([]string{protocol, "_port"}, ""))
	addr := strings.Join([]string{host, ":", strconv.Itoa(port)}, "")
	url := strings.Join([]string{protocol, "://", addr}, "")

	logfields := logrus.Fields{"host": host, "port": port, "url": url}
	boxagent.Log.WithFields(logfields).Info("server started")

	if protocol == "https" {
		crt := boxagent.Config.GetString("ssl_crt")
		key := boxagent.Config.GetString("ssl_key")
		if _, err := os.Stat(crt); err != nil {
			boxagent.Log.WithField("error", err).Fatal("could not find certificate file")
		}
		if _, err := os.Stat(key); err != nil {
			boxagent.Log.WithField("error", err).Fatal("could not find key file")
		}
		httpErr = http.ListenAndServeTLS(addr, crt, key, nil)

	} else {
		httpErr = http.ListenAndServe(addr, nil)
	}

	if httpErr != nil {
		boxagent.Log.WithField("error", httpErr).Fatal("listener fatal error")
	}

}
