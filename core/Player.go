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
  playerManager *PlayerManager
  name          string
  x, y, z       float64
}

// PlayerJSON ...
type PlayerJSON struct {
  Name    string
  X, Y, Z float64
}

// ToJSON ...
func (p *Player) ToJSON() PlayerJSON {
  return PlayerJSON{p.name, p.x, p.y, p.z}
}

// NewPlayer ...
func NewPlayer(playerManager *PlayerManager, name string) *Player {
  player := new(Player)
  player.playerManager = playerManager
  player.name = name
  return player
}

// FilePath ...
func (p *Player) FilePath() string {
  return path.Join(p.playerManager.worldPath, PlayerDir, p.name+".dat")
}

// X ...
func (p *Player) X() float64 {
  return p.x
}

// Y ...
func (p *Player) Y() float64 {
  return p.y
}

// Z ...
func (p *Player) Z() float64 {
  return p.z
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
  tree := nbt.NewNbtTree(re)
  positionList := tree.Root().Entries["Pos"].(*nbt.TagNodeList)
  positionX := positionList.Get(0).(*nbt.TagNodeDouble)
  positionY := positionList.Get(1).(*nbt.TagNodeDouble)
  positionZ := positionList.Get(2).(*nbt.TagNodeDouble)
  p.x = positionX.Data
  p.y = positionY.Data
  p.z = positionZ.Data
  return p
}
