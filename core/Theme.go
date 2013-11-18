package core


import (
  "os"
  "io/ioutil"
  "encoding/xml"
)



type Query struct {
  Blocks []Block `xml:"Block"`
}
type Block struct {
  Id byte `xml:"id,attr"`
  Red uint8 `xml:"red,attr"`
  Green uint8 `xml:"green,attr"`
  Blue uint8 `xml:"blue,attr"`
  Alpha uint8 `xml:"alpha,attr"`
}


func LoadTheme(p_Theme string) map[byte]Block {
  s_Logger.Debug("Load xml theme")
  xmlFile, err := os.Open("./public/themes/default/theme.xml")
  if err != nil {
    s_Logger.Fatal("Cant load theme file")
  }
  defer xmlFile.Close()

  b, _ := ioutil.ReadAll(xmlFile)
  var q Query
  xml.Unmarshal(b, &q)

  m := make(map[byte]Block)

  for _, block := range q.Blocks {
    m[block.Id] = block
  }

  return m
}

