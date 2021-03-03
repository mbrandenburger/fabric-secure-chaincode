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
	rootCmd.AddCommand(retrieveCmd)
}

var retrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "retrieve assets",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := pkg.NewClient(config)
		res := client.Call("retrieveAsset", args...)
		fmt.Println("> " + res)
	},
}
