package core


import (
  "bytes"
  "compress/gzip"
  "fmt"
  "mas/nbt"
  "io/ioutil"
  "path"
  "strings"
)


type Player struct {
  m_PlayerManager *PlayerManager
  m_Name string
  m_X, m_Y, m_Z float64
}


type PlayerJson struct {
  Name string
  X, Y, Z float64
}


func (p *Player) ToJson() PlayerJson {
  return PlayerJson{p.m_Name, p.m_X, p.m_Y, p.m_Z}
}


func NewPlayer(p_PlayerManager *PlayerManager, p_Name string) *Player {
  player := Player{}
  player.m_PlayerManager = p_PlayerManager
  player.m_Name = p_Name
  return &player
}


func (p *Player) FilePath() string {
  return path.Join(p.m_PlayerManager.m_WorldPath, PLAYER_DIR, p.m_Name + ".dat")
}


func (p *Player) X() float64 {
  return p.m_X
}


func (p *Player) Y() float64 {
  return p.m_Y
}


func (p *Player) Z() float64 {
  return p.m_Z
}


func (p *Player) Player() *Player {
  filePath := p.FilePath()
  file, err := ioutil.ReadFile(filePath)
  if err != nil {
    s_Logger.Debug(err)
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
  p.m_X = positionX.Data()
  p.m_Y = positionY.Data()
  p.m_Z = positionZ.Data()
  return p
}
