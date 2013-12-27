package core

import (
  "log"
  "os"
  "path"
  "strings"
)

// PlayerDir ...
const PlayerDir = "players"

// PlayerManager ...
type PlayerManager struct {
  mWorldPath string
}

// NewPlayerManager ...
func NewPlayerManager(pWorldPath string) *PlayerManager {
  playerManager := new(PlayerManager)
  playerManager.mWorldPath = pWorldPath
  return playerManager
}

// GetPlayers ...
func (p *PlayerManager) GetPlayers() []*Player {
  var players []*Player
  for _, name := range p.PlayerNames() {
    players = append(players, p.GetPlayer(name))
  }
  return players
}

// GetPlayer ...
func (p *PlayerManager) GetPlayer(pName string) *Player {
  player := NewPlayer(p, pName).Player()
  return player
}

// PlayerNames ...
func (p *PlayerManager) PlayerNames() []string {
  playersDir, err := os.Open(path.Join(p.mWorldPath, PlayerDir))
  if err != nil {
    log.Panic(err)
  }
  defer playersDir.Close()
  var newFiles []string
  files, _ := playersDir.Readdirnames(0)
  for _, file := range files {
    if !strings.HasSuffix(file, "dat") {
      continue
    }
    file = strings.TrimSuffix(file, ".dat")
    newFiles = append(newFiles, file)
  }
  return newFiles
}
