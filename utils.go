package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"strconv"
	"strings"
)

// GetJenkinsBuildNumber : get request to find last build number
func GetJenkinsBuildNumber(jenkinsURL string) int64 {
	res := SendRequest("GET", jenkinsURL+"/lastBuild/buildNumber", "")
	buildNum, err := strconv.ParseInt(string(res), 10, 0)
	EoE(err, "Error Parsing Build Number")
	return buildNum
}

// SendRequest : send http request to provided url
func SendRequest(method, url, body string) []byte {
	reader := strings.NewReader(body)

	req, err := http.NewRequest(method, url, reader)
	EoE(err, "Error Formatting HTTP Request")

	client := http.Client{}
	res, err := client.Do(req)
	EoE(err, "Error Getting HTTP Response")

	resData, err := ioutil.ReadAll(res.Body)
	EoE(err, "Error Parsing Response")
	return resData
}

// EoE : exit with error code 1 and print if err is notnull
func EoE(err error, msg string) {
	if err != nil {
		fmt.Println("❌  "+msg, err)
		os.Exit(1)
	}
}

// GetHomeDir : returns a full path to user's home dorectory
func GetHomeDir() string {
	usr, err := user.Current()
	EoE(err, "Failed to get Home Directory")
	if usr.HomeDir != "" {
		return usr.HomeDir
	}
	// Maybe it's cross compilation without cgo support. (darwin, unix)
	return os.Getenv("HOME")
}

// Confirm : return confirmation based on user input
func Confirm(q string) bool {
	a := GetInput(q + " (Y/n) ")
	var res bool
	switch a {
	case "":
		fallthrough
	case "y":
		fallthrough
	case "Y":
		res = true
	case "n":
	case "N":
		res = false
	default:
		return Confirm(q)
	}
	return res
}

// GetInput : return string of user input
func GetInput(q string) string {
	if q != "" {
		fmt.Print(q)
	}
	reader := bufio.NewReader(os.Stdin)
	ans, _ := reader.ReadString('\n')
	return strings.TrimRight(ans, "\n")
}

// SetFromInput : set value of provided var to the value of user input
func SetFromInput(a *string, q string) {
	*a = strings.TrimRight(GetInput(q), "\n")
}
