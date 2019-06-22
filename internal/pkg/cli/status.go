package cli

import (

)

// CheckStatus return if a server is running
func CheckStatus() (bool, error) {
	pid, err := readPID()
	if err != nil {
		return false, err
	}

	switch pid {
	case -1:
		return false, nil
	default:
		return true, nil
	}
}