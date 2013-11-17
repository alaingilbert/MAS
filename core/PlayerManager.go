package core


import (
  "os"
  "path"
  "strings"
)


const PLAYER_DIR = "players"


type PlayerManager struct {
  m_WorldPath string
}


func NewPlayerManager(p_WorldPath string) *PlayerManager {
  playerManager := PlayerManager{}
  playerManager.m_WorldPath = p_WorldPath
  return &playerManager
}


func (p *PlayerManager) GetPlayers() []*Player {
  var players []*Player
  return players
}


func (p *PlayerManager) GetPlayer(p_Name string) *Player {
  player := NewPlayer(p, p_Name).Player()
  return player
}


func (p *PlayerManager) PlayerNames() []string {
  playersDir, _ := os.Open(path.Join(p.m_WorldPath, PLAYER_DIR))
  defer playersDir.Close()
  var newFiles []string
  files, _ := playersDir.Readdirnames(0)
  for _, file := range files {
    if !strings.HasSuffix(file, "dat") {
      continue
    }
    file = strings.TrimRight(file, ".dat")
    newFiles = append(newFiles, file)
  }
  return newFiles
}
