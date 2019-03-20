// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"fmt"
	"github.com/autom8ter/engine/plugin"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "enginectl",
	Long: fmt.Sprintf(`
----------------------------------------------------------------------------
8888888888                d8b                        888   888
888                       Y8P                        888   888
888                                                  888   888
8888888   88888b.  .d88b. 88888888b.  .d88b.  .d8888b888888888
888       888 "88bd88P"88b888888 "88bd8P  Y8bd88P"   888   888
888       888  888888  888888888  88888888888888     888   888
888       888  888Y88b 888888888  888Y8b.    Y88b.   Y88b. 888
8888888888888  888 "Y88888888888  888 "Y8888  "Y8888P "Y888888
                       888                                    
                  Y8b d88P                                    
                   "Y88P"
----------------------------------------------------------------------------
Download:
go get github.com/autom8ter/engine/enginectl
----------------------------------------------------------------------------
Expected Plugin Path:
$HOME/.plugins
----------------------------------------------------------------------------
Expected Plugin Export Name:
Plugin
----------------------------------------------------------------------------
Plugins Files Found:
%s
`, plugin.Files()),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

}