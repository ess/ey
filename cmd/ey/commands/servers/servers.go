package servers

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCmd = &cobra.Command{
	Use: "servers",

	Short: "Server-related commands",

	Long: `Server-related commands

This group provides access to server-related commands like "start a server"
and "stop a server."

Please see the command list below for more information.`,
}

func init() {
	api = viper.GetString("api")
	if len(api) == 0 {
		api = "https://api.engineyard.com"
	}

	token = viper.GetString("token")
}

var api string
var token string

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
