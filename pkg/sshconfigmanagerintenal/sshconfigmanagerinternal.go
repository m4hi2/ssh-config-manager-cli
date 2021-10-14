package sshconfigmanagerintenal

import (
	"fmt"
	"log"
	"os"

	"github.com/m4hi2/ssh-config-manager-cli/pkg/parser"
)

func backUpFile(filePath string) {
	originalFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Can not find ssh config file\n Error: %s", err)
	}

	backUp, _ := os.Create(fmt.Sprintf("%s.bak", filePath))
	backUp.Write(originalFile)
}

func Add(config parser.Config, filePath string) {
	backUpFile(filePath)
	fileContent, _ := os.ReadFile(filePath)
	parsedConfig := parser.Parse(string(fileContent))
	parsedConfig[config.Host] = &config
	newComposedConfig := parser.Compose(parsedConfig)
	os.WriteFile(filePath, []byte(newComposedConfig), 0644)

}
