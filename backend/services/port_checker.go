package services

import (
	"fmt"
	"net"
	"time"
)

// CheckPort checks if a port is listening on localhost
func CheckPort(port int) bool {
	timeout := 1 * time.Second
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("localhost:%d", port), timeout)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

// CheckPorts checks if at least one of the given ports is listening
// Returns true if any port is online, false if all are offline
func CheckPorts(ports []int) bool {
	for _, port := range ports {
		if CheckPort(port) {
			return true
		}
	}
	return false
}
