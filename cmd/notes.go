/*
Copyright © 2022 RAMOSA <ron@cloudbuilder.io>

*/
package cmd

import (
	"errors"
	"fmt"
	"net"
	"os"
	"rx/cmd/util"
	"text/template"

	"github.com/spf13/cobra"
)

type Target struct {
	Name      string
	IPAddress string
}

// notesCmd represents the notes command
var notesCmd = &cobra.Command{
	Use:   "notes",
	Short: "create new notes markdown with templates.",
	Long: `Create a new markdown notes file for HTB or THM
and quickly get into hacking, For example:

rx notes add <box> <ip-address>`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			fmt.Println("len=", len(args))
			return errors.New("requires 2 args: <name> <ip-address>")
		}
		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		// assign vars
		projectName := args[0]
		targetIP := args[1]
		// parse the ip address
		addr := net.ParseIP(targetIP)

		// check if ip is valid
		if addr == nil {
			fmt.Println("Invalid address:", targetIP)
			os.Exit(1)
		}

		// check if the folder we want to create already exists
		fileExists, err := util.FileExists(projectName)
		if err != nil {
			fmt.Println(err)
		}
		// if it doesn't exist, create it.
		if fileExists {
			fmt.Println("File exists:", projectName)
		} else {
			fmt.Println("Creating directory:", projectName)

			// create directory
			_, err := util.CreateFolder(projectName)
			if err != nil {
				fmt.Println(err)
			}

			// create notes markdown
			target := Target{projectName, targetIP}
			template, _ := template.ParseFiles("notes.tmpl")
			template.Execute(os.Stdout, target)
		}
	},
}

func init() {
	createCmd.AddCommand(notesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	notesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	notesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
