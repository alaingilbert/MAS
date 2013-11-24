package core

import (
  "os"
)

// World represent a minecraft world.
type World struct {
  mPath          string
  mRegionManager *RegionManager
  mPlayerManager *PlayerManager
}

// NewWorld instantiate a world object.
// It returns a pointer to a world.
func NewWorld(pPath string) *World {
  world := &World{}
  world.mPath = pPath
  world.mRegionManager = NewRegionManager(world.mPath)
  world.mPlayerManager = NewPlayerManager(world.mPath)
  return world
}

// PathValid ...
func (w *World) PathValid() bool {
  _, err := os.Stat(w.mPath)
  return err == nil
}

// Path get the path to the world directory.
// It returns a path to the world directory.
func (w *World) Path() string {
  return w.mPath
}

// RegionManager get the region manager.
// It returns a pointer to the region manager.
func (w *World) RegionManager() *RegionManager {
  return w.mRegionManager
}

// PlayerManager ...
func (w *World) PlayerManager() *PlayerManager {
  return w.mPlayerManager
}
