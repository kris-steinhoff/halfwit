// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// ForcePush is a flag to control whether the initia git push is done with --force
var ForcePush bool

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:                   "init [-f] git_repo_url",
	DisableFlagsInUseLine: true,
	Short:                 "Initialize a new directory for managed dotfiles",
	Long: `Create an empty directory for managed dofiles, populate it with README and
LICENSE files, and push to the specified repository URL.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		createDir(getDir())
	},
}

func createDir(dir string) {
	err := os.Mkdir(dir, 755)
	if err != nil {
		panic(err)
	}
}

func init() {
	initCmd.Flags().BoolVarP(&ForcePush, "force-push", "f", false, "Use a --force push to the git remote")
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
