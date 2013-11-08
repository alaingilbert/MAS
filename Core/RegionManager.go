package Core


type RegionManager struct {
  m_RegionPath string
}


func NewRegionManager(p_RegionPath string) *RegionManager {
  regionManager := &RegionManager{}
  regionManager.m_RegionPath = p_RegionPath
  return regionManager
}


func (r *RegionManager) GetRegion(p_X, p_Z int) *Region {
  return NewRegion(r, p_X, p_Z)
}


func (r *RegionManager) RegionPath() string {
  return r.m_RegionPath
}
