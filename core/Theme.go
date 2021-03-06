package core

import (
  "encoding/xml"
  "io/ioutil"
  "log"
  "os"
)

// Query ...
type Query struct {
  Blocks []Block `xml:"Block"`
}

// Block ...
type Block struct {
  ID    byte   `xml:"id,attr"`
  Red   uint8  `xml:"red,attr"`
  Green uint8  `xml:"green,attr"`
  Blue  uint8  `xml:"blue,attr"`
  Alpha uint8  `xml:"alpha,attr"`
  Name  string `xml:"name,attr"`
}

// Theme ...
type Theme struct {
  mMap   map[byte]Block
  mTheme string
}

// NewTheme ...
func NewTheme(pTheme string) *Theme {
  theme := new(Theme)
  theme.mTheme = pTheme
  theme.LoadTheme()
  return theme
}

// Reload ...
func (t *Theme) Reload() {
  t.LoadTheme()
}

// LoadTheme ...
func (t *Theme) LoadTheme() {
  xmlFile, err := os.Open("public/themes/default/theme.xml")
  if err != nil {
    log.Fatalf("Cant load theme file", err)
  }
  defer xmlFile.Close()

  b, _ := ioutil.ReadAll(xmlFile)
  var q Query
  xml.Unmarshal(b, &q)

  t.mMap = make(map[byte]Block)
  for _, block := range q.Blocks {
    t.mMap[block.ID] = block
  }
}

// GetByID ...
func (t *Theme) GetByID(pID byte) Block {
  return t.mMap[pID]
}

// GetMap ...
func (t *Theme) GetMap() map[byte]Block {
  return t.mMap
}
