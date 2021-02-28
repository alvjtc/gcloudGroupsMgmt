//Copyright 2021 Álvaro José Teijido Carpente
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/spf13/cobra"
	"grtool/internal/app/groups"
	"os"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	ValidArgs: groups.GetGroupValidArgs(),
	Run:       runList,
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringP("domain", "d", "", "Help message for domain")
	listCmd.MarkFlagRequired("domain")

	listCmd.Flags().StringP("output", "o", "", "Help message for domain")
}

func runList(cmd *cobra.Command, args []string) {
	domain, _ := cmd.Flags().GetString("domain")

	groupList, err := groups.GetAllGroups(Googler.GoogleDirectorySrv, domain)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	output, _ := cmd.Flags().GetString("output")

	var stream *os.File

	if output != "" {
		fmt.Printf("Listing all groups from domain %s to file %s:\n", domain, output)

		stream, err = os.Create(output)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		defer func() {
			if err := stream.Close(); err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
		}()
	} else {
		fmt.Printf("Listing all groups from domain %s:\n\n", domain)

		stream = os.Stdout
	}

	w := csv.NewWriter(stream)

	w.Write(groups.GetHeaders(args))

	for _, g := range groupList {
		values, err := g.ToSlice(args)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		w.Write(values)
	}

	w.Flush()
}
