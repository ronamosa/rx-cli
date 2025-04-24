/*
Copyright © 2022 RON AMOSA <ron@cloudbuilder.io>
*/
package main

import "rx/cmd"

// Version information set by build process using ldflags
var (
	Version   = "dev"
	BuildTime = "unknown"
)

func main() {
	cmd.SetVersionInfo(Version, BuildTime)
	cmd.Execute()
}
