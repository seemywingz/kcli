package main

import(
  "os"
  "fmt"
  "bytes"
  "strings"
  "os/exec"
  "github.com/voxelbrain/goptions"
)

// Global Flags
var force bool

func concat(s1, s2 string) string{
  var buffer bytes.Buffer
  buffer.WriteString(s1)
  buffer.WriteString(" ")
  buffer.WriteString(s2)
  return buffer.String()
}

func main() {
  options := struct {
		Help  goptions.Help `goptions:"-h, --help, description='Show this help'"`
    Remainder goptions.Remainder

		goptions.Verbs
		Run struct {
			Sudo bool   `goptions:"-s, --sudo, description='Run as Superuser'"`
		} `goptions:"run"`

    Say struct {
			Phrase string `goptions:"-p, --phrase, description='Will attempt to say <phrase> outloud'"`
		} `goptions:"say"`

	}{ // Default values go here
		// Force: true,
	}

  err := goptions.Parse(&options)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  switch options.Verbs {
  case "run":
    cmd := options.Remainder[0]
    args := strings.Join(options.Remainder[1:], " ")
    fmt.Println("Running:",cmd)
    out, err := exec.Command(cmd,args).Output()
    if err != nil {
      fmt.Println("⛔  Error when running command:", cmd, err)
      os.Exit(1)
    }
    fmt.Println(out)
  case "say":
    phrase := concat(options.Say.Phrase, strings.Join(options.Remainder, " "))
    fmt.Println("Saying Phrase:", phrase)
    cmd := "say "+phrase
    exec.Command("sh","-c",cmd).Output()
  default:
    goptions.PrintHelp()
  }

}
