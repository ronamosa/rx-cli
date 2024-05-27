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
	"path/filepath"

	"github.com/spf13/cobra"
)

//go:embed templates/*
var tmplFS embed.FS

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
		_, err := util.CreateFolder("NOTE_OPD/"+strings.ToUpper(name))
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

		parentDir := "NOTE_OPD/"
		notes_tmpl := template.Must(template.ParseFS(tmplFS, "templates/notes.tmpl"))

		fmt.Println("Creating new notes file:", parentDir+strings.ToUpper(name)+"/"+"notes-"+name+".md...")

		// Defining the File Creation Part and storeing the file 
		// 
		subDir := strings.ToUpper(name) + "/"
		fileName := "notes-" + name + ".md"
		filePath := filepath.Join(parentDir, subDir, fileName)
	    file, err := os.Create(filePath)


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
