package main

import (
	"github.com/boxmetrics/boxmetrics-agent/internal/pkg/agent"
)

// Start boxagent websocket
func Start() {
	agent.InitConfig()
	agent.CreateServer()
}

func main() {
	Start()
}
