/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-github/v32/github"
	"github.com/inconshreveable/go-update"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update binary from github",
	Long:  `Update binary from github`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		if list {

			client := github.NewClient(nil)

			ctx := context.Background()
			releases, _, err := client.Repositories.ListReleases(ctx, "mennink", "cmdtest", nil)

			if err != nil {
				fmt.Println(err)
			}
			for _, r := range releases {
				fmt.Printf("release : %s\n", r.GetTagName())
			}
		}

	},
}

var list bool

func doUpdate(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = update.Apply(resp.Body, update.Options{})
	if err != nil {
		return err

	}
	return nil
}

func init() {
	updateCmd.PersistentFlags().BoolVarP(&list, "list", "l", false, "List versions")
	rootCmd.AddCommand(updateCmd)

}
