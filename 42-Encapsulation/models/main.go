package models

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

func NewIp(ip string) (ipAddress, error) {
	parts := strings.Split(ip, ".")

	if len(parts) < 4 {
		return ipAddress{}, fmt.Errorf("Invalid IP address")
	}

	return ipAddress{ip, parts}, nil
}

func (ip ipAddress) GetParts() []string {
	return ip.parts
}
