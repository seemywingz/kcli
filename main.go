package main

import(
  "fmt"
  "github.com/voxelbrain/goptions"
)

// Global Flags
var force bool

func main() {
  options := struct {
		Force bool          `goptions:"-f, --force, description='Fuce - Rho - Dah'"`
		Help  goptions.Help `goptions:"-h, --help, description='Show this help'"`

		goptions.Verbs
		Do struct {
			Action  string `goptions:"-a, --name, obligatory, description='Name of the entity to be deleted'"`
			Force bool   `goptions:"-f, --force, description='Force removal'"`
		} `goptions:"do"`

	}{ // Default values go here
		// Force: true,
	}
	goptions.ParseAndFail(&options)

  fmt.Println(options)


}
