package sshconfigmanagerintenal

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/m4hi2/ssh-config-manager-cli/pkg/parser"
	"github.com/manifoldco/promptui"
)

func backUpFile(filePath string) {
	originalFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Can not find ssh config file\n Error: %s", err)
		log.Printf("Creating empty file as backup file...")
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

type PromptContent struct {
	ErrorMsg string
	Label    string
}

func PromptGetInput(pc PromptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.ErrorMsg)
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     pc.Label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}
