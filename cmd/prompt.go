/*
Copyright © 2021 Mahir Labib Chowdhury <optical.mahir@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/m4hi2/ssh-config-manager-cli/pkg/parser"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// promptCmd represents the prompt command
var promptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "Provides a prompt to the user for adding host config interactively",
	Long:  `Provides a prompt to the user for adding host config interactively`,
	Run: func(cmd *cobra.Command, args []string) {
		addWithPrompt()
	},
}

func init() {
	addCmd.AddCommand(promptCmd)
}

func addWithPrompt() {
	hostPrompt := promptContent{
		errorMsg: "Enter a friendly host name",
		label:    "Host (Enter a name that you can remember for this server):",
	}
	host := promptGetInput(hostPrompt)

	hostNamePrompt := promptContent{
		errorMsg: "Enter the server address",
		label:    "HostName (Enter either the IP or the domain of the server):",
	}
	hostName := promptGetInput(hostNamePrompt)

	portPrompt := promptContent{
		errorMsg: "Enter the ssh port nubmer of the server",
		label:    "Port (Enter the ssh port. If not sure, enter 22):",
	}
	port := promptGetInput(portPrompt)

	userPrompt := promptContent{
		errorMsg: "Enter the login user name",
		label:    "User (Enter the username to be logged in as):",
	}
	user := promptGetInput(userPrompt)

	newConfig := parser.Config{
		Host:     host,
		HostName: hostName,
		Port:     port,
		User:     user,
	}

	fmt.Println(newConfig)
}

type promptContent struct {
	errorMsg string
	label    string
}

func promptGetInput(pc promptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.errorMsg)
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
		Label:     pc.label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)

	return result
}
