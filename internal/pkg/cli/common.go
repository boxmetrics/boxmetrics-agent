package cli

import (
	"io/ioutil"
	"strconv"
	"os"
)

func readPID() (int, error) {

	if _, err := os.Stat("boxagent.pid"); os.IsNotExist(err) {
		return -1, nil
	} else if err != nil {
		return 0, err
	}

	data, err := ioutil.ReadFile("boxagent.pid")
	if err != nil {
		return 0, err
	}

	pid, err := strconv.Atoi(string(data))
	if err != nil {
		return 0, err
	}

	return pid, nil
}

func writePID(pid int) (int, error) {
	file, err := os.OpenFile("boxagent.pid", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return 0, err
	}

	file.WriteString(strconv.Itoa(pid))
	file.Close()

	return pid, nil
}