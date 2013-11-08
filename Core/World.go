package Core


type World struct {
  m_Path string
  m_RegionManager *RegionManager
}


func NewWorld(p_Path string) *World {
  world := &World{}
  world.m_Path = p_Path
  world.m_RegionManager = NewRegionManager(world.m_Path)
  return world
}


func (w *World) Path() string {
  return w.m_Path
}


func (w *World) RegionManager() *RegionManager {
  return w.m_RegionManager
}
