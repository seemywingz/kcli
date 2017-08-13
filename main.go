package main

import (
	"github.com/voxelbrain/goptions"
)

type optionsDef struct {
	Help    goptions.Help `goptions:"-h, --help, description='Show this help'"`
	Tail    bool          `goptions:"-t, --tail, description='Follow the logs'"`
	Verbose bool          `goptions:"-v, --verbose, description='Verbose'"`

	goptions.Verbs
	Status struct {
	} `goptions:"status"`

	Life struct {
	} `goptions:"life"`

	Mtg struct {
		Set     string `goptions:"-s, --set, description='TLA (three letter abreviation) for Set Name'"`
		Name    string `goptions:"-n, --name, description='Name of the card to lookup'"`
		SetName string `goptions:"--set-name, description='Full Set Name'"`
	} `goptions:"mtg"`

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
	case "mtg":
		Mtg()
	default:
		ListConfig()
	}

}
