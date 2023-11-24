package main

import "fmt"

type server struct {
	host string
	user string
	port int
}

func execCommand(server *server, execCommand *string) {
	server.user = "user1"
	newCommand := "apt install php8.1"
	*execCommand = newCommand
}

func main() {
	provisionCmd := "apt install php8.2"

	server := server{
		host: "127.0.0.1",
		user: "root",
		port: 33,
	}

	execCommand(&server, &provisionCmd)

	fmt.Println(server, provisionCmd)
}
