package main

import (
	"fmt"
	"phptogo/42-Encapsulation/models"
)

func main() {
	ip, err := models.NewIp("192.168.1.1")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(ip.GetParts())
	}
}
