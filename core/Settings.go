package core


import (
  "encoding/xml"
  "io/ioutil"
  "os"
)


type Settings struct {
  Theme string
  WorldPath string
  NbtVersion string
  WebServer WebServer
}

type WebServer struct {
  Host string
  Port int
}


func LoadSettings() (*Settings, error) {
  _, err := os.Stat("settings.xml")
  if err != nil {
    CreateSettingsFile()
  }

  settingsFile, err := ioutil.ReadFile("settings.xml")
  if err != nil {
    return nil, err
  }
  var settings Settings
  xml.Unmarshal(settingsFile, &settings)
  return &settings, nil
}


func CreateSettingsFile() {
  file, _ := os.Create("settings.xml")
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
}
