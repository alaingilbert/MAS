package core


import (
  "os"
)


// World represent a minecraft world.
type World struct {
  m_Path string
  m_RegionManager *RegionManager
  m_PlayerManager *PlayerManager
}


// NewWorld instantiate a world object.
// It returns a pointer to a world.
func NewWorld(p_Path string) *World {
  world := &World{}
  world.m_Path = p_Path
  world.m_RegionManager = NewRegionManager(world.m_Path)
  world.m_PlayerManager = NewPlayerManager(world.m_Path)
  return world
}


func (w *World) PathValid() bool {
  _, err := os.Stat(w.m_Path)
  return err == nil
}


// Path get the path to the world directory.
// It returns a path to the world directory.
func (w *World) Path() string {
  return w.m_Path
}


// RegionManager get the region manager.
// It returns a pointer to the region manager.
func (w *World) RegionManager() *RegionManager {
  return w.m_RegionManager
}


func (w *World) PlayerManager() *PlayerManager {
  return w.m_PlayerManager
}
