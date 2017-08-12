package main

import (
	"github.com/voxelbrain/goptions"
)

type optionsDef struct {
	Help goptions.Help `goptions:"-h, --help, description='Show this help'"`
	Tail goptions.Help `goptions:"-t, --tail, description='Follow the logs'"`

	goptions.Verbs
	Status struct {
	} `goptions:"status"`

	Life struct {
	} `goptions:"life"`

	Config struct {
		Name  string `goptions:"-n, --name, description='Set Name in config file'"`
		Email string `goptions:"-e, --email, description='Set Email in config file'"`
		List  bool   `goptions:"-l, --list, description='List the current config'"`
	} `goptions:"config"`
}

var options optionsDef

func main() {
	goptions.ParseAndFail(&options)
	GetConfig()

	switch options.Verbs {
	case "config":
		Configure()
	case "life":
		GameOfLife()
	default:
		ListConfig()
	}

}
