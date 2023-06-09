/*
Copyright © 2023 OpenFGA

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
package stores

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/openfga/fga-cli/lib/cmd-utils"
	"github.com/spf13/cobra"
)

// getCmd represents the get command.
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get store",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		clientConfig := cmdutils.GetClientConfig(cmd)
		fgaClient, err := clientConfig.GetFgaClient()
		if err != nil {
			fmt.Printf("Failed to initialize FGA Client due to %v", err)
			os.Exit(1)
		}
		store, err := fgaClient.GetStore(context.Background()).Execute()
		if err != nil {
			fmt.Printf("Failed to get store %v due to %v", clientConfig.StoreID, err)
			os.Exit(1)
		}

		storeJSON, err := json.Marshal(store)
		if err != nil {
			fmt.Printf("Failed to get store due to %v", err)
			os.Exit(1)
		}
		fmt.Print(string(storeJSON))
	},
}

func init() {
	getCmd.Flags().String("store-id", "", "Store ID")
	err := getCmd.MarkFlagRequired("store-id")
	if err != nil { //nolint:wsl
		fmt.Print(err)
		os.Exit(1)
	}
}
