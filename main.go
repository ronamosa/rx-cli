/*
Copyright © 2022 RON AMOSA <ron@cloudbuilder.io>
*/
package main

import "rx/cmd"

// Version information set by build process
var (
	Version   = "dev"
	BuildTime = "unknown"
)

func main() {
	cmd.Execute()
}
