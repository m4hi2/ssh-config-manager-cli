package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	dump()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func dump(fileLocation string) {
	data, err := os.ReadFile(fileLocation)
	check(err)
	fmt.Print(string(data))
	fmt.Printf("\n")
}

func writeConfig(fileLocation string, configString string) {
	f, err := os.OpenFile(fileLocation, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	defer f.Close()
	_, error := f.WriteString(configString)
	check(error)
}

func generateHostString(host string,
	hostName string,
	port string,
	user string) string {
	return fmt.Sprintf("Host %s\n\t HostName %s\n\t Port %s\n\t User %s\n", host, hostName, port, user)
}

// extraction of username and hostname and port
// user@10.0.0.1 -p 22 -> user, 10.0.0.1, 22
func extractParams(sshLogin string) (string, string, string) {
	port := "22"
	userAndHost := "none"
	user := "none"
	host := "none"
	if strings.Contains(sshLogin, "-p") {
		hostAndPort := strings.Split(sshLogin, "-p")
		userAndHost = strings.TrimSpace(hostAndPort[0])
		port = strings.TrimSpace(hostAndPort[1])
	} else {
		userAndHost = sshLogin
	}

	if strings.Contains(userAndHost, "@") {
		userAndHostName := strings.Split(userAndHost, "@")
		user = strings.TrimSpace(userAndHostName[0])
		host = strings.TrimSpace(userAndHostName[1])
	}

	return user, host, port

}
