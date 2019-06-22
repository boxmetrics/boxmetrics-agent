package cli

import (
	"github.com/boxmetrics/boxmetrics-agent/internal/pkg/errors"
	"os/exec"
)

// StartServer start agent server
func StartServer() (int, error) {
	pid, err := readPID()
	if err != nil {
		return 0, err
	}

	if pid != -1 {
		return 0, errors.New("server already running")
	}

	server := exec.Command("bash", "-c", "./bin/server")
	err = server.Start()
	if err != nil {
		return 0, err
	}

	pid, err = writePID(server.Process.Pid)
	if err != nil {
		return 0, err
	}

	return pid, nil
}
