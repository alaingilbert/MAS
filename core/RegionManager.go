package core

import (
  "log"
  "os"
  "path"
  "strconv"
  "strings"
)

// REGION_DIR name of the regions directory.
const RegionDir = "region"

// RegionManager is used to manage the regions.
type RegionManager struct {
  mRegionPath string
}

// NewRegionManager instantiate a new region manager.
// It returns a pointer to a region manager.
func NewRegionManager(pRegionPath string) *RegionManager {
  regionManager := &RegionManager{}
  regionManager.mRegionPath = pRegionPath
  return regionManager
}

// GetRegion get a specific region.
// pRegionX coordinate of the region on the X axis.
// pRegionZ coordinate of the region on the Z axis.
// It returns a pointer to a region.
func (r *RegionManager) GetRegion(pRegionX, pRegionZ int) *Region {
  return NewRegion(r, pRegionX, pRegionZ)
}

// GetRegionFromXYZ get a specific region from a global world coordinate.
// pX x world coordinate.
// pY y world coordinate.
// pZ z world coordinate.
// It returns a pointer to a region.
func (r *RegionManager) GetRegionFromXYZ(pX, pY, pZ int) *Region {
  return NewRegionFromXYZ(r, pX, pY, pZ)
}

// RegionFileNames ...
func (r *RegionManager) RegionFileNames() []string {
  tilesDirectory, err := os.Open(path.Join(r.mRegionPath, RegionDir))
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

// RegionsCoordinates ...
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
  return r.mRegionPath
}
