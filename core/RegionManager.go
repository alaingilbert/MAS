package core


import (
  "path"
  "log"
  "os"
  "strconv"
  "strings"
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
// p_RegionX coordinate of the region on the X axis.
// p_RegionZ coordinate of the region on the Z axis.
// It returns a pointer to a region.
func (r *RegionManager) GetRegion(p_RegionX, p_RegionZ int) *Region {
  return NewRegion(r, p_RegionX, p_RegionZ)
}


// GetRegionFromXYZ get a specific region from a global world coordinate.
// p_X x world coordinate.
// p_Y y world coordinate.
// p_Z z world coordinate.
// It returns a pointer to a region.
func (r *RegionManager) GetRegionFromXYZ(p_X, p_Y, p_Z int) *Region {
  return NewRegionFromXYZ(r, p_X, p_Y, p_Z)
}


func (r *RegionManager) RegionFileNames() []string {
  tilesDirectory, err := os.Open(path.Join(r.m_RegionPath, REGION_DIR))
  if err != nil {
    log.Fatal(err)
  }
  defer tilesDirectory.Close()
  files, err := tilesDirectory.Readdirnames(0)
  var newFiles []string
  for _, file := range files {
    if !strings.HasSuffix(file, "mca") {
      continue
    }
    newFiles = append(newFiles, file)
  }
  return newFiles
}


func (r *RegionManager) RegionsCoordinates() [][2]int {
  files := r.RegionFileNames()
  result := make([][2]int, len(files))
  for index, fileName := range files {
    splits := strings.SplitN(fileName, ".", 4)
    regionX, _ := strconv.Atoi(splits[1])
    regionZ, _ := strconv.Atoi(splits[2])
    result[index][0] = regionX
    result[index][1] = regionZ
  }
  return result
}


// RegionPath get the region folder path.
// It returns the region folder path.
func (r *RegionManager) RegionPath() string {
  return r.m_RegionPath
}
