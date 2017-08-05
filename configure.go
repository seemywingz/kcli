package main

import(
  "os"
  "fmt"
  "io/ioutil"
  "encoding/json"
)

import gt "github.com/seemywingz/gtills"

var homeDir,configFile string

const configFileName = "/.myApp/config"

type jsonConfig struct{
  Fname string `json:"fname"`
  Lname string `json:"lname"`
  Email string `json:"email"`
}
var config jsonConfig

func Configure() {

  fmt.Println("📝  ",configFile)
  gt.GetInput(&config.Fname,"First Name:")
  gt.GetInput(&config.Lname,"Last Name:")
  gt.GetInput(&config.Email,"Email:")

  data, err := json.Marshal(config)
  if err != nil {
    fmt.Println("❌  Error converting json",err)
    os.Exit(1)
  }

  err = ioutil.WriteFile(configFile, data, 0644)
  if err != nil {
    fmt.Println("❌  Error Saving Config File",err)
    os.Exit(2)
  }
}

func GetConfig() {
  homeDir := gt.GetHomeDir()
  if homeDir == "" {
    os.Exit(1)
  }
  configFile = homeDir + configFileName
  if _, err := os.Stat(configFile); os.IsNotExist(err) {
    fmt.Println("❗ CONFIG NOT FOUND")
    var ans string
    gt.GetInput(&ans,"⚙  Want to Create one now? (Y/n)")
    if ans == "y" || ans == ""{
      Configure()
    }else{
      fmt.Println("⏩  Skipping Configuration File Creation")
      os.Exit(10)
    }
  }else{ // config exists
    jsonFile, err := ioutil.ReadFile(configFile)
    if err != nil {
      fmt.Println("❌  Error Reading Config File",err)
      os.Exit(2)
    }
    json.Unmarshal(jsonFile, &config)
  }
}
