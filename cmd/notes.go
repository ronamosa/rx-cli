/*
Copyright © 2022 RAMOSA <ron@cloudbuilder.io>

*/
package cmd

import (
	"fmt"
	"net"
	"os"
	"rx/cmd/util"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

type Target struct {
	Name      string
	IPAddress string
}

// notesCmd represents the notes command
var notesCmd = &cobra.Command{
	Use:   "notes <name> <ip-address>",
	Short: "create new notes markdown with templates.",
	Long: `

	██████╗ ██╗  ██╗██╗  ██╗ █████╗  ██████╗██╗  ██╗
	██╔══██╗╚██╗██╔╝██║  ██║██╔══██╗██╔════╝██║ ██╔╝
	██████╔╝ ╚███╔╝ ███████║███████║██║     █████╔╝ 
	██╔══██╗ ██╔██╗ ██╔══██║██╔══██║██║     ██╔═██╗ 
	██║  ██║██╔╝ ██╗██║  ██║██║  ██║╚██████╗██║  ██╗
	╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝	
	`,
	Args: cobra.MinimumNArgs(2),
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
			fmt.Println("Creating directory:", strings.ToUpper(projectName))

			// create directory
			_, err := util.CreateFolder(strings.ToUpper(projectName))
			if err != nil {
				fmt.Println(err)
			}

			// create notes markdown
			target := Target{projectName, targetIP}
			template := template.Must(template.ParseFiles("/home/rxnamxsa/.config/rx/templates/notes.tmpl"))

			// write to file
			file, _ := os.Create(strings.ToUpper(projectName) + "/" + "notes-" + projectName + ".md")
			template.Execute(file, target)
		}
	},
}

func init() {
	createCmd.AddCommand(notesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//notesCmd.PersistentFlags().String("foo", "", "A help for foo")
	//notesCmd.PersistentFlags().String("name", "", "name of the project, box, room.")
	//notesCmd.PersistentFlags().String("ip", "", "ip address of the target, for the room.")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//notesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Marking Flags Required
	//notesCmd.MarkPersistentFlagRequired("name")
	//notesCmd.MarkPersistentFlagRequired("ip")
}
