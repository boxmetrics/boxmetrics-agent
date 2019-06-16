package main

import (
	"github.com/boxmetrics/boxmetrics-agent/internal/pkg/agent"
)

func main() {
	agent.InitConfig()

	agent.CreateServer()
}
