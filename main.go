package main

import (
	"fmt"
	"os"
)

func main() {
	dump()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func dump() {
	data, err := os.ReadFile("/home/m4hi2/ssh-config-manager-cli/config")
	check(err)
	fmt.Print(string(data))
}

func generateHostString(host string,
	hostName string,
	port string,
	user string) string {
	return fmt.Sprintf("Host %s\n\t HostName %s\n\t Port %s\n\t User %s\n", host, hostName, port, user)
}
}
