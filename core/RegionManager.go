package core


import (
  "path"
  "log"
  "os"
)


// REGION_DIR name of the regions directory.
const REGION_DIR = "region"


// RegionManager is used to manage the regions.
type RegionManager struct {
  m_RegionPath string
}


// NewRegionManager instantiate a new region manager.
// It returns a pointer to a region manager.
func NewRegionManager(p_RegionPath string) *RegionManager {
  regionManager := &RegionManager{}
  regionManager.m_RegionPath = p_RegionPath
  return regionManager
}


// GetRegion get a specific region.
// p_X coordinate of the region on the X axis.
// p_Z coordinate of the region on the Z axis.
// It returns a pointer to a region.
func (r *RegionManager) GetRegion(p_X, p_Z int) *Region {
  return NewRegion(r, p_X, p_Z)
}


func (r *RegionManager) RegionFileNames() []string {
  tilesDirectory, err := os.Open(path.Join(r.m_RegionPath, REGION_DIR))
  if err != nil {
    log.Fatal(err)
  }
  defer tilesDirectory.Close()
  files, err := tilesDirectory.Readdirnames(0)
  return files
}


// RegionPath get the region folder path.
// It returns the region folder path.
func (r *RegionManager) RegionPath() string {
  return r.m_RegionPath
}
