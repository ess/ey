package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ess/ey/cmd/ey/commands/servers"
)

var RootCmd = &cobra.Command{
	Use:   "ey",
	Short: "A CLI for Engine Yard",
	Long:  `A CLI for Engine Yard`,
}

// Execute attempts to run the root command and returns an error. If root
// returned cleanly, nothing is done and nil is returned. Otherwise, the
// root error is printed to the terminal and is then returned.
func Execute() error {
	err := RootCmd.Execute()

	if err != nil {
		fmt.Println(err)
	}

	return err
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.AddCommand(servers.RootCmd)
}

func initConfig() {
	viper.SetEnvPrefix("ey")
	viper.AutomaticEnv()
}

// Copyright © 2019 Engine Yard, Inc.
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
