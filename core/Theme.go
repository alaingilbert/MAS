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


type Theme struct {
  m_Map map[byte]Block
}


func LoadTheme(p_Theme string) *Theme {
  xmlFile, err := os.Open("./public/themes/default/theme.xml")
  if err != nil {
    s_Logger.Fatal("Cant load theme file")
  }
  defer xmlFile.Close()

  b, _ := ioutil.ReadAll(xmlFile)
  var q Query
  xml.Unmarshal(b, &q)

  theme := Theme{}
  theme.m_Map = make(map[byte]Block)
  for _, block := range q.Blocks {
    theme.m_Map[block.Id] = block
  }

  return &theme
}


func (t *Theme) GetById(p_Id byte) Block {
  return t.m_Map[p_Id]
}
