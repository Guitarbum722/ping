package ping

import (
	"os/exec"
	"strings"
)

// Ping will ping host 3 times and returns the summary of the ping results.
func Ping(host string) (string, error) {
	output, err := exec.Command("ping", host, "-c 3").Output()
	if err != nil {
		return "", err
	}
	lines := strings.Split(string(output), "\n")
	return lines[len(lines)-2], nil
}
