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
	"fmt"
	"os"

	"github.com/m4hi2/ssh-config-manager-cli/pkg/parser"
	"github.com/m4hi2/ssh-config-manager-cli/pkg/sshconfigmanagerintenal"
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
	hostPrompt := sshconfigmanagerintenal.PromptContent{
		ErrorMsg: "Enter a friendly host name",
		Label:    "Host (Enter a name that you can remember for this server):",
	}
	host := sshconfigmanagerintenal.PromptGetInput(hostPrompt)

	hostNamePrompt := sshconfigmanagerintenal.PromptContent{
		ErrorMsg: "Enter the server address",
		Label:    "HostName (Enter either the IP or the domain of the server):",
	}
	hostName := sshconfigmanagerintenal.PromptGetInput(hostNamePrompt)

	portPrompt := sshconfigmanagerintenal.PromptContent{
		ErrorMsg: "Enter the ssh port nubmer of the server",
		Label:    "Port (Enter the ssh port. If not sure, enter 22):",
	}
	port := sshconfigmanagerintenal.PromptGetInput(portPrompt)

	userPrompt := sshconfigmanagerintenal.PromptContent{
		ErrorMsg: "Enter the login user name",
		Label:    "User (Enter the username to be logged in as):",
	}
	user := sshconfigmanagerintenal.PromptGetInput(userPrompt)

	newConfig := parser.Config{
		Host:     host,
		HostName: hostName,
		Port:     port,
		User:     user,
	}
	homeDir, _ := os.UserHomeDir()
	sshConfigFilePath := fmt.Sprintf("%s/.ssh/config", homeDir)
	sshconfigmanagerintenal.Add(newConfig, sshConfigFilePath)
}
