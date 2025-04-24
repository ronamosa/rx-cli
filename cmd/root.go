/*
Copyright © 2022 RON AMOSA <ron@cloudbuilder.io>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// Version information
const (
	Version   = "0.0.2" // Set this to match release tags
	BuildTime = "2025-04-24"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rx",
	Short: "A CLI helper tool for penetration testing and CTF challenges",
	Long: `rx-cli - A hacker's helper tool for CTF challenges like HackTheBox and TryHackMe

Automates common setup tasks such as creating note templates, 
generating reverse shells, and organizing your hacking workflow.

Examples:
  rx create notes MyTarget 10.10.10.10
  rx create shell php --LHOST 10.10.14.20 --LPORT 443
`,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	Version: Version,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// Set custom version template
	rootCmd.SetVersionTemplate(`Version:    {{.Version}}
Build Date: ` + BuildTime + `
`)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rx.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
