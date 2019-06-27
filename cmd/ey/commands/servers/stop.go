package servers

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	//"github.com/ess/ey/pkg/ey"
	"github.com/ess/ey/pkg/ey/eycore"
)

var stopCmd = &cobra.Command{
	Use: "stop <Server ID>",

	Short: "Stop a server",

	Long: `Stop a server

Given a server's Amazon ID, stop the server. If it is already stopped, this
is basically a no-op. If it is running, then we attempt to stop the server.`,

	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("Usage: ey servers stop <Server ID>")
		}

		token = viper.GetString("token")

		if len(token) == 0 {
			return fmt.Errorf("This operation requires Engine Yard API authentication.")
		}

		return nil
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		eycore.Setup(api, token)
		servers := eycore.NewServerService()
		serverID := args[0]

		server, err := servers.Get(serverID)
		if err != nil {
			return fmt.Errorf("could not find server with id %s", serverID)
		}

		if server.State == "running" {
			startErr := servers.Stop(server)
			if startErr != nil {
				return fmt.Errorf("could not stop server with id %s", serverID)
			}
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(stopCmd)
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
