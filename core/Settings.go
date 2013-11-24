package core

import (
  "encoding/xml"
  "fmt"
  "io/ioutil"
  "os"
)

// Settings ...
type Settings struct {
  Theme      string
  WorldPath  string
  NbtVersion string
  WebServer  WebServer
}

// WebServer ...
type WebServer struct {
  Host string
  Port int
}

// LoadSettings ...
func LoadSettings() (*Settings, error) {
  _, err := os.Stat("settings.xml")
  if err != nil {
    err := CreateSettingsFile()
    if err != nil {
      fmt.Println(err)
      return nil, err
    }
  }

  settingsFile, err := ioutil.ReadFile("settings.xml")
  if err != nil {
    return nil, err
  }
  var settings Settings
  xml.Unmarshal(settingsFile, &settings)
  return &settings, nil
}

// CreateSettingsFile ...
func CreateSettingsFile() error {
  file, err := os.Create("settings.xml")
  if err != nil {
    return err
  }
  defer file.Close()
  content := `<?xml version="1.0" encoding="UTF-8" ?>
<Settings>
  <Theme>default</Theme>
  <WorldPath>/Path/To/world</WorldPath>
  <NbtVersion>Anvil</NbtVersion>

  <WebServer>
    <Host>127.0.0.1</Host>
    <Port>8000</Port>
  </WebServer>
</Settings>`
  file.Write([]byte(content))
  return nil
}
