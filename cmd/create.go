/*
Copyright © 2022 RON AMOSA <ron@cloudbuilder.io>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// declare target struct here
type Target struct {
	Name      string
	IPAddress string
	Port      string
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	TraverseChildren: true,
	Use:              "create",
	Short:            "Create notes or shells for your penetration testing",
	Long: `Create resources for your penetration testing workflow:

- Notes: Generate a markdown template with common pentest commands and checklist
- Shells: Generate various reverse shell payloads for different languages

Examples:
  rx create notes MyTarget 10.10.10.10       # Create note template
  rx create shell php --LHOST 10.0.0.1 --LPORT 4444  # PHP reverse shell
`,
	Args: cobra.MinimumNArgs(1),
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
