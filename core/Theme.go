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
  m_Theme string
}


func NewTheme(p_Theme string) *Theme {
  theme := Theme{}
  theme.m_Theme = p_Theme
  return &theme
}


func (t *Theme) Reload() {
  t.LoadTheme()
}


func (t *Theme) LoadTheme() {
  xmlFile, err := os.Open("public/themes/default/theme.xml")
  if err != nil {
    s_Logger.Fatal("Cant load theme file", err)
  }
  defer xmlFile.Close()

  b, _ := ioutil.ReadAll(xmlFile)
  var q Query
  xml.Unmarshal(b, &q)

  t.m_Map = make(map[byte]Block)
  for _, block := range q.Blocks {
    t.m_Map[block.Id] = block
  }
}


func (t *Theme) GetById(p_Id byte) Block {
  return t.m_Map[p_Id]
}
