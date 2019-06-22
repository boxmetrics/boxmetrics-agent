package cli

import (
	"os"
	"github.com/boxmetrics/boxmetrics-agent/internal/pkg/errors"
)

// StopServer stop agent server
func StopServer() (bool, error) {
	pid, err := readPID()
	if err != nil {
		return false, err
	}

	if pid == -1 {
		return false, errors.New("server is not running")
	}

	if pid != 0 {
		proc, err := os.FindProcess(pid)
		if err != nil {
			return false, err
		}

		err = proc.Kill()
		if err != nil {
			return false, err
		}

		err = os.Remove("boxagent.pid")
		if err != nil {
			return false, err
		}
	}

	return true, nil
}