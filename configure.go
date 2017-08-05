package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	gt "github.com/seemywingz/gtills"
)

const configDir = ".myApp"
const configFileName = "config"

var homeDir, configFile string

type jsonConfig struct {
	Fname string `json:"fname"`
	Lname string `json:"lname"`
	Email string `json:"email"`
}

var config jsonConfig

// Configure : Gather User Informaton and save it to config file
func Configure() {

	fmt.Println("📝  ", configFile)
	gt.SetFromInput(&config.Fname, "\nFirst Name: ")
	gt.SetFromInput(&config.Lname, " Last Name: ")
	gt.SetFromInput(&config.Email, "     Email: ")

	data, jsoEerr := json.Marshal(config)
	gt.E(jsoEerr, "Error Parsing Json:")

	err := ioutil.WriteFile(configFile, data, 0644)
	gt.E(err, "Error Saving Config File:")
	fmt.Println("\n✨  Configuration File Saved Successfully")
}

// GetConfig : Check to see if there is a config file, if not create one
func GetConfig() {
	homeDir := gt.GetHomeDir()
	if homeDir == "" {
		os.Exit(1)
	}
	configFile = filepath.Join(homeDir, configDir, configFileName)
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		fmt.Println("❗ CONFIG NOT FOUND")
		ans := gt.GetInput("⚙  Want to Create one now? (Y/n): ")
		if ans == "y" || ans == "" {
			err := os.MkdirAll(filepath.Join(homeDir, configDir), os.ModePerm)
			gt.E(err, "Error Creating Config Directory:")
			Configure()
		} else {
			fmt.Println("⏩  Skipping Configuration File Creation")
			os.Exit(10)
		}
	} else { // config exists
		jsonFile, err := ioutil.ReadFile(configFile)
		gt.E(err, "Error Reading Config File:")
		json.Unmarshal(jsonFile, &config)
	}
}
