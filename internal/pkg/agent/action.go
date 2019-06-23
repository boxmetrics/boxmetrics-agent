package agent

import (
	"github.com/boxmetrics/boxmetrics-agent/internal/pkg/errors"
	"github.com/boxmetrics/boxmetrics-agent/internal/pkg/info"
	"os"
	"os/exec"
	"strconv"
)

func dispatchEvent(e event) (interface{}, error) {
	switch e.Type {
	case Info:
		Log.Debug("info")
		return dispatchInfo(e)
	case Script:
		return dispatchScript(e)
	case Command:
		return dispatchCommand(e)
	default:
		return nil, errors.New("Event not support")
	}
}

func dispatchInfo(e event) (interface{}, error) {
	switch e.Value {
	case "memory":
		return info.Memory(e.Format)
	case "cpu":
		Log.Debug("cpu")
		return info.CPU(e.Format)
	case "cpuinfo":
		Log.Debug("cpuinfo")
		return info.CPUinfo()
	case "disks":
		Log.Debug("disks")
		return info.Disks(e.Format)
	case "containers":
		Log.Debug("containers")
		return info.Containers(e.Format)
	case "containersid":
		Log.Debug("containersid")
		return info.ContainersID()
	case "host":
		Log.Debug("host")
		return info.Host(e.Format)
	case "users":
		Log.Debug("users")
		return info.Users()
	case "network":
		Log.Debug("network")
		return info.Network(e.Format)
	case "connections":
		Log.Debug("connections")
		return info.Connections()
	case "processes":
		Log.Debug("processes")
		return info.Processes(e.Format)
	case "process":
		Log.Debug("process")
		return info.Process(int32(e.Options.Pid), e.Format)
	case "general":
		Log.Debug("general")
		return info.General(e.Format)
	default:
		return nil, errors.New("Info not support")
	}
}

func dispatchScript(e event) (interface{}, error) {
	switch e.Value {
	case "adduser":
		cmd := exec.Command("sudo", "scripts/add_user.sh")
		cmd.Args = append(cmd.Args, e.Options.Args...)

		r, err := cmd.Output()

		if err != nil {
			return nil, errors.Convert(err)
		}

		return string(r), nil
	case "killprocess":
		pid := e.Options.Pid
		proc, err := os.FindProcess(pid)
		if err != nil {
			return nil, errors.Convert(err)
		}

		err = proc.Kill()
		if err != nil {
			return nil, errors.Convert(err)
		}

		return "Process " + strconv.Itoa(pid) + " has stopped", nil
	default:
		return nil, errors.New("Script not support")
	}
}

func dispatchCommand(e event) (interface{}, error) {
	cmd := exec.Command(e.Value)
	cmd.Args = append(cmd.Args, e.Options.Args...)
	cmd.Env = e.Options.Env
	cmd.Dir = e.Options.Dir

	r, err := cmd.CombinedOutput()

	if err != nil {
		return nil, errors.Convert(err)
	}

	return string(r), nil
}
