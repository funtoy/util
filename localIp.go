package util

import (
	"net"
)

func GetLocalIP() string {
	list, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, addr := range list {
		if IPNet, ok := addr.(*net.IPNet); ok && !IPNet.IP.IsLoopback() {
			if IPNet.IP.To4() != nil {
				return IPNet.IP.String()
			}
		}
	}
	return ""
}
