package main

import (
	"github.com/voxelbrain/goptions"
)

func main() {
	options := struct {
		Help goptions.Help `goptions:"-h, --help, description='Show this help'"`
		Tail goptions.Help `goptions:"-t, --tail, description='Follow the logs'"`

		goptions.Verbs
		// Request struct {
		// 	Name bool   `goptions:"-n, --name, description='Name of Instance'"`
		// } `goptions:"request"`

		Status struct {
		} `goptions:"status"`

		Config struct {
			Name string `goptions:"-n, --name, description='Set Name in config file'"`
		} `goptions:"config"`
	}{ // Default values go here
	// Force: true,
	}

	goptions.ParseAndFail(&options)
	GetConfig()

	switch options.Verbs {
	case "status":
		Status()
	case "config":
		Configure()
	default:
		// goptions.PrintHelp()
	}

}
