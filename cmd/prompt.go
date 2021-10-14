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

	"github.com/spf13/cobra"
)

// promptCmd represents the prompt command
var promptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "Provides a prompt to the user for adding host config interactively",
	Long:  `Provides a prompt to the user for adding host config interactively`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("prompt called")
	},
}

func init() {
	addCmd.AddCommand(promptCmd)
}
