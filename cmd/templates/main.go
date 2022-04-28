package main

import (
	"os"
	"text/template"
)

type Target struct {
	Name      string
	IPAddress string
}

func main() {

	name := Target{"HTBName", "192.168.1.1"}
	template, _ := template.ParseFiles("notes.tmpl")
	template.Execute(os.Stdout, name)
}
