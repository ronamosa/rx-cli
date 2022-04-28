/*
Copyright © 2022 RAMOSA <ron@cloudbuilder.io>

*/
package cmd

import (
	"fmt"
	"net"
	"os"

	"github.com/spf13/cobra"
)

// notesCmd represents the notes command
var notesCmd = &cobra.Command{
	Use:   "notes",
	Short: "create new notes markdown with templates.",
	Long: `Create a new markdown notes file for HTB or THM
and quickly get into hacking, For example:

rx notes add <box> <ip-address>`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("notes called")
		if len(args) != 2 {
			fmt.Println("Error: I need two args")
		} else {
			projectName := args[0]
			targetIP := args[1]

			addr := net.ParseIP(targetIP)
			if addr == nil {
				fmt.Println("Invalid address")
				os.Exit(1)
			}

			fmt.Println("Project Name=", projectName)
			fmt.Println("targetIp=", targetIP)
		}
	},
}

func init() {
	createCmd.AddCommand(notesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// notesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// notesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
