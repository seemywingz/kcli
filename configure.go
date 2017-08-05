package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

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

func saveConfig() {
	data, jsoEerr := json.Marshal(config)
	gt.E(jsoEerr, "Error Parsing Json:")
	gt.E(ioutil.WriteFile(configFile, data, 0644), "Error Saving Config File:")
}

// Configure : Gather User Informaton and save it to config file
func Configure() {

	switch {
	case options.Config.Name != "":
		names := strings.Split(options.Config.Name, " ")
		config.Fname = names[0]
		config.Lname = names[1]
		saveConfig()
		return
	case options.Config.Email != "":
		config.Email = options.Config.Email
		saveConfig()
		return
	case options.Config.List:
		configData, err := ioutil.ReadFile(configFile)
		gt.E(err, "Error Reading Config File:")
		fmt.Println(string(configData))
		return
	default:
		fmt.Println("📝  Writing", configFile)
		gt.SetFromInput(&config.Fname, "\nFirst Name: ")
		gt.SetFromInput(&config.Lname, " Last Name: ")
		gt.SetFromInput(&config.Email, "     Email: ")
	}

	save := gt.Confirm("Save Configuratuon File?")
	if save == true {
		saveConfig()
		fmt.Println("\n✨  Configuration File Saved Successfully")
	} else {
		fmt.Println("\n🚫  Configuration File Not Saved")
	}
}

// GetConfig : Check to see if there is a config file, if not create one
func GetConfig() {
	homeDir := gt.GetHomeDir()
	if homeDir == "" {
		os.Exit(1)
	}
	configFile = filepath.Join(homeDir, configDir, configFileName)
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		fmt.Println("❗  CONFIG NOT FOUND")
		ans := gt.Confirm("⚙  Want to Create one now?")
		if ans {
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
