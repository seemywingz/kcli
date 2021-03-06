package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	gt "github.com/seemywingz/gtils"
)

const configDir = ".kcli"
const configFileName = "config"

var homeDir, configFile string

type jsonConfig struct {
	Fname    string `json:"fname"`
	Lname    string `json:"lname"`
	Email    string `json:"email"`
	Hostname string `json:"hostname"`
	IP       string `json:"ip"`
}

var config jsonConfig

// SaveConfig : writes the current config to disk
func SaveConfig() {
	data, jsoEerr := json.Marshal(config)
	gt.EoE("Error Parsing Json:", jsoEerr)
	gt.EoE("Error Saving Config File:", ioutil.WriteFile(configFile, data, 0644))
}

// ListConfig : prints the current config
func ListConfig() {
	println("")
	println("📖  Reading Config\n")
	println("First Name:📓 ", config.Fname)
	println(" Last Name:📓 ", config.Lname)
	println("     Email:📧 ", config.Email)
	println("  Hostname:🌐 ", config.IP)
	println("        IP:🌐 ", config.IP)
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
		print("\n")
		println("📝  Writing ", configFile)
		println("❗  Fields are Required\n")
		gt.SetFromInput(&config.Fname, "First Name:❗  ")
		gt.SetFromInput(&config.Lname, " Last Name:❗  ")
		gt.SetFromInput(&config.Email, "     Email:📧  ")
	}

	if gt.Confirm("Save Configuratuon File?") {
		SaveConfig()
		println("\n✨  Configuration File Saved Successfully")
		os.Exit(0)
	} else {
		println("\n🚫  Configuration File Not Saved")
	}
}

// GetConfig : Check to see if there is a config file, if not create one
func GetConfig() {
	homeDir = gt.GetHomeDir()
	if homeDir == "" {
		os.Exit(1)
	}
	configFile = filepath.Join(homeDir, configDir, configFileName)
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		println("❗  CONFIG NOT FOUND")
		if gt.Confirm("⚙  Want to Create one now?") {
			err := os.MkdirAll(filepath.Join(homeDir, configDir), os.ModePerm)
			gt.EoE("Error Creating Config Directory:", err)
			Configure()
		} else {
			println("⏩  Skipping Configuration File Creation")
			os.Exit(10)
		}
	} else { // config exists
		jsonFile, err := ioutil.ReadFile(configFile)
		gt.EoE("Error Reading Config File:", err)
		json.Unmarshal(jsonFile, &config)
		config.IP = gt.GetIP()
		name, err := os.Hostname()
		gt.EoE("Error Getting Hostname", err)
		config.Hostname = name
	}
}
