package core

import (
  "bytes"
  "compress/gzip"
  "fmt"
  "io/ioutil"
  "mas/nbt"
  "path"
  "strings"
)

// Player ...
type Player struct {
  mPlayerManager *PlayerManager
  mName          string
  mX, mY, mZ     float64
}

// PlayerJSON ...
type PlayerJSON struct {
  Name    string
  X, Y, Z float64
}

// ToJSON ...
func (p *Player) ToJSON() PlayerJSON {
  return PlayerJSON{p.mName, p.mX, p.mY, p.mZ}
}

// NewPlayer ...
func NewPlayer(pPlayerManager *PlayerManager, pName string) *Player {
  player := Player{}
  player.mPlayerManager = pPlayerManager
  player.mName = pName
  return &player
}

// FilePath ...
func (p *Player) FilePath() string {
  return path.Join(p.mPlayerManager.mWorldPath, PlayerDir, p.mName+".dat")
}

// X ...
func (p *Player) X() float64 {
  return p.mX
}

// Y ...
func (p *Player) Y() float64 {
  return p.mY
}

// Z ...
func (p *Player) Z() float64 {
  return p.mZ
}

// Player ...
func (p *Player) Player() *Player {
  filePath := p.FilePath()
  file, err := ioutil.ReadFile(filePath)
  if err != nil {
    return nil
  }
  reader, err := gzip.NewReader(bytes.NewReader(file))
  if err != nil {
    fmt.Println(err)
  }
  defer reader.Close()
  buf := new(bytes.Buffer)
  buf.ReadFrom(reader)
  s := buf.String()
  re := strings.NewReader(s)
  tree := nbt.NewNbtTree()
  tree.Init(re)
  positionList := tree.Root().Entries["Pos"].(nbt.TagNodeList)
  positionX := positionList.Get(0).(nbt.TagNodeDouble)
  positionY := positionList.Get(1).(nbt.TagNodeDouble)
  positionZ := positionList.Get(2).(nbt.TagNodeDouble)
  p.mX = positionX.Data()
  p.mY = positionY.Data()
  p.mZ = positionZ.Data()
  return p
}
