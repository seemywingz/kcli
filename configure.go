package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	gtills "github.com/seemywingz/gtills"
)

const configDir = ".kcli"
const configFileName = "config"

var homeDir, configFile string

type jsonConfig struct {
	Fname string `json:"fname"`
	Lname string `json:"lname"`
	Email string `json:"email"`
	IP    string `json:"ip"`
}

var config jsonConfig

// SaveConfig : writes the current config to disk
func SaveConfig() {
	data, jsoEerr := json.Marshal(config)
	gtills.EoE(jsoEerr, "Error Parsing Json:")
	gtills.EoE(ioutil.WriteFile(configFile, data, 0644), "Error Saving Config File:")
}

// ListConfig : prints the current config
func ListConfig() {
	configJSON, err := json.MarshalIndent(config, "", "   ")
	gtills.EoE(err, "Error Parsing Json")
	fmt.Println("📖  Reading Config", configFile, "\n", string(configJSON))
}

// Configure : Gather User Informaton and save it to config file
func Configure() {

	switch {
	case options.Config.Name != "":
		names := strings.Split(options.Config.Name, " ")
		config.Fname = names[0]
		config.Lname = names[1]
		SaveConfig()
		ListConfig()
		return
	case options.Config.Email != "":
		config.Email = options.Config.Email
		SaveConfig()
		ListConfig()
		return
	case options.Config.List:
		ListConfig()
		return
	default:
		fmt.Println("📝  Writing", configFile)
		gtills.SetFromInput(&config.Fname, "\nFirst Name: ")
		gtills.SetFromInput(&config.Lname, " Last Name: ")
		gtills.SetFromInput(&config.Email, "     Email: ")
	}

	if gtills.Confirm("Save Configuratuon File?") {
		SaveConfig()
		fmt.Println("\n✨  Configuration File Saved Successfully")
		os.Exit(0)
	} else {
		fmt.Println("\n🚫  Configuration File Not Saved")
	}
}

// GetConfig : Check to see if there is a config file, if not create one
func GetConfig() {
	homeDir := gtills.GetHomeDir()
	if homeDir == "" {
		os.Exit(1)
	}
	configFile = filepath.Join(homeDir, configDir, configFileName)
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		fmt.Println("❗  CONFIG NOT FOUND")
		if gtills.Confirm("⚙  Want to Create one now?") {
			err := os.MkdirAll(filepath.Join(homeDir, configDir), os.ModePerm)
			gtills.EoE(err, "Error Creating Config Directory:")
			Configure()
		} else {
			fmt.Println("⏩  Skipping Configuration File Creation")
			os.Exit(10)
		}
	} else { // config exists
		jsonFile, err := ioutil.ReadFile(configFile)
		gtills.EoE(err, "Error Reading Config File:")
		json.Unmarshal(jsonFile, &config)
	}
}
