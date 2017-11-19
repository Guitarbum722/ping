package ping

import (
	"os/exec"
)

// Ping will ping host 3 times and returns the summary of the ping results.
func Ping(host string) (string, error) {
	output, err := exec.Command("ping", "-c 3", host).Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
