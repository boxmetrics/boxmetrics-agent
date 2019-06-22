package cli

import (
	"os/exec"
)

// StartServer start agent server
func StartServer() (int, error) {
	server := exec.Command("bash", "-c", "./bin/server")
	err := server.Start()
	if err != nil {
		return 0, err
	}

	pid, err := writePID(server.Process.Pid)
	if err != nil {
		return 0, err
	}

	return pid, nil
}
