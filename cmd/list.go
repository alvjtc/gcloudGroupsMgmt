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
	"bufio"
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
	Run: runList,
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
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
	if output != "" {
		fmt.Printf("Listing all groups from domain %s to file %s:\n", domain, output)

		fo, err := os.Create(output)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		defer func() {
			if err := fo.Close(); err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
		}()

		w := bufio.NewWriter(fo)

		// CSV Header
		if _, err := fmt.Fprintln(w, "group_name"); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		for _, g := range groupList {
			if _, err := fmt.Fprintf(w, "%s\n", g.Email); err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
		}
		w.Flush()
	} else {
		fmt.Printf("Listing all groups from domain %s:\n\n", domain)

		for _, g := range groupList {
			fmt.Println(g.Email)
		}
	}
}
