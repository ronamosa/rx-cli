/*
Copyright © 2022 RAMOSA <ron@cloudbuilder.io>
*/
package cmd

import (
	"embed"
	"fmt"
	"net"
	"os"
	"rx/cmd/util"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

//go:embed templates/*
var tmplFS embed.FS

// notesCmd represents the notes command
var notesCmd = &cobra.Command{
	Use:   "notes <target-name> <target-ip>",
	Short: "Create a structured markdown notes template",
	Long: `Create a new markdown file with a comprehensive template for 
penetration testing notes, including common commands and checklists.

This will:
1. Create a directory with the specified target name
2. Generate a markdown file with a template inside that directory
3. Fill in the IP address and target name in appropriate locations

Examples:
  rx create notes MyTarget 10.10.10.10
  rx create notes HackTheBox 172.16.10.2

Parameters:
  <target-name>  Name of the target/challenge (used for directory and file naming)
  <target-ip>    IP address of the target machine`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		// assign vars
		projectName := strings.ToUpper(args[0])
		targetIP := args[1]

		// parse the ip address
		addr := net.ParseIP(targetIP)

		// check if ip is valid
		if addr == nil {
			fmt.Println("Invalid address:", targetIP)
			os.Exit(1)
		}

		_, err := createNotes(projectName, targetIP)
		if err != nil {
			fmt.Println("Create Failed: ", err)
		} else {
			fmt.Println("Create Successful.")
		}
	},
}

func createNotes(name string, ipaddress string) (bool, error) {

	// check if folder already exists
	fileExists, err := util.FileExists(name)
	if err != nil {
		return false, err
	}
	if fileExists {
		fmt.Printf("Folder %v already exists\n", name)
		return false, err
	} else {
		// create directory
		_, err := util.CreateFolder(strings.ToUpper(name))
		if err != nil {
			//fmt.Println(err)
			return false, err
		}

		// create target struct
		target := Target{
			Name:      name,
			IPAddress: ipaddress,
			Port:      "",
		}

		notes_tmpl := template.Must(template.ParseFS(tmplFS, "templates/notes.tmpl"))

		fmt.Println("Creating new notes file:", strings.ToUpper(name)+"/"+"notes-"+name+".md...")

		file, err := os.Create(strings.ToUpper(name) + "/" + "notes-" + name + ".md")
		if err != nil {
			fmt.Println(err)
			return false, nil
		} else {
			notes_tmpl.Execute(file, target)
			return true, nil
		}
	}
}

func init() {
	createCmd.AddCommand(notesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//notesCmd.PersistentFlags().String("foo", "", "A help for foo")
	//notesCmd.PersistentFlags().String("name", "", "name of the project, box, room.")
	//notesCmd.PersistentFlags().String("ip", "", "ip address of the target, for the room.")
	notesCmd.PersistentFlags().StringP("add", "a", "", "add a section to the notes")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//notesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Marking Flags Required
	//notesCmd.MarkPersistentFlagRequired("name")
	//notesCmd.MarkPersistentFlagRequired("ip")
}
