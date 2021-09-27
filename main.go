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
