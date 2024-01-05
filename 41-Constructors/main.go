package main

import (
	"fmt"
	"strings"
)

type ipAddress struct {
	ip    string
	parts []string
}

func (ip ipAddress) network() string {
	return strings.Join(ip.parts[0:3], ".")
}

func newIp(ip string) (ipAddress, error) {
	parts := strings.Split(ip, ".")

	if len(parts) < 4 {
		return ipAddress{}, fmt.Errorf("Invalid IP address")
	}

	return ipAddress{ip, parts}, nil
}

func main() {
	ip, err := newIp("192.168.1.1")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(ip.network())
	}
}
