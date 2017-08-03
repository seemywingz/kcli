package main

import(
  "os"
  "fmt"
  "flag"
)

// Global Flags
var force bool

func do() {
  doFlags := flag.NewFlagSet("do", flag.ExitOnError)
  doFlags.BoolVar(&force, "f", false, "force")
  doFlags.Parse(flag.Args()[1:])
  if doFlags.Parsed() {
    if len(doFlags.Args()) == 0 {
      fmt.Println("Must provide something to do!")
      os.Exit(1)
    }
	  fmt.Printf("Run Command: %q\n", doFlags.Args())
  }
}

func parseFlags() {
  // Global Flags
  flag.BoolVar(&force, "f", false, "force")
  flag.Parse()

  if len(flag.Args()) <= 0 {
    fmt.Println("Must Provide Command")
    os.Exit(1)
  }

  switch flag.Args()[0] {
  case "do":
    do()
  default:
  }
}

func main() {
  fmt.Println("**  KCLI  **")
  parseFlags()
}
