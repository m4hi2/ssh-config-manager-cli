package sshconfigmanagerintenal

import (
	"fmt"
	"log"
	"os"
)

func backUpFile(filePath string) {
	originalFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Can not find ssh config file\n Error: %s", err)
	}

	backUp, _ := os.Create(fmt.Sprintf("%s.bak", filePath))
	backUp.Write(originalFile)
}
