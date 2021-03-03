/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package cmd

import (
	"fmt"

	"github.com/hyperledger-labs/fabric-private-chaincode/client_sdk/go/sample/helloworld/pkg"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(storeCmd)
}

var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "store assets",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		client := pkg.NewClient(config)
		res := client.Call("storeAsset", args...)
		fmt.Println("> " + res)
	},
}
