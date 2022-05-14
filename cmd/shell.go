/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

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

		shellType := args[0]
		LHOST := "192.168.1.2"
		LPORT := "443"

		switch {
		case shellType == "php":
			createPhpSh(LHOST, LPORT)
		case shellType == "py":
			createPySh(LHOST, LPORT)
		case shellType == "bash":
			createBashSh(LHOST, LPORT)
		case shellType == "bin":
			createBinSh(LHOST, LPORT)
		default:
			fmt.Println("Did you select an available shell?")
		}
	},
}

func createPhpSh(ipaddress string, port string) (bool, error) {
	fmt.Printf("PHP shell for %v:%s", ipaddress, port)
	return true, nil
}

func createPySh(ipaddress string, port string) (bool, error) {
	fmt.Printf("Python shell for %v:%s", ipaddress, port)
	return true, nil
}

func createBashSh(ipaddress string, port string) (bool, error) {
	fmt.Printf("Bash shell for %v:%s ", ipaddress, port)
	return true, nil
}

func createBinSh(ipaddress string, port string) (bool, error) {
	fmt.Printf("Bin shell for %v:%s", ipaddress, port)
	return true, nil
}

func init() {
	createCmd.AddCommand(shellCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// shellCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// shellCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
