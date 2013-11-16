package core


import (
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


func (p *PlayerManager) GetPlayer(p_Name string) *Player {
  return NewPlayer()
}


func (p *PlayerManager) PlayerNames() []string {
  var players []string
  return players
}
