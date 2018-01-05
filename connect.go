package main

import (
	"time"
	"net"
)

func Connected(address string) bool {
	var timeout = (1*time.Second) / 10

	// So, we try to connect to the address within the time limit.
	// Here, if we can't make it in 0.1 second, it shall return the error.
	_, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return false
	}
	return true
}
