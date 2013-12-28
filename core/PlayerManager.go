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
  worldPath string
}

// NewPlayerManager ...
func NewPlayerManager(worldPath string) *PlayerManager {
  playerManager := new(PlayerManager)
  playerManager.worldPath = worldPath
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
func (p *PlayerManager) GetPlayer(name string) *Player {
  player := NewPlayer(p, name).Player()
  return player
}

// PlayerNames ...
func (p *PlayerManager) PlayerNames() []string {
  playersDir, err := os.Open(path.Join(p.worldPath, PlayerDir))
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
