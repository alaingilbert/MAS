package draw

import (
  "fmt"
  "mas/core"
)

// JobRenderRegionTile job that will render a region tile.
type JobRenderRegionTile struct {
  mRegionX, mRegionZ int
  mWorld              *core.World
  mTheme              *core.Theme
}

// NewJobRenderRegionTile will instantiate a job to render a region tile.
// pRegionX x axis coordinate of the region.
// pRegionZ z axis coordinate of the region.
// pWorld pointer to a minecraft world.
// pTheme color theme to be used.
// It returns a JobRenderRegionTile.
func NewJobRenderRegionTile(pRegionX, pRegionZ int, pWorld *core.World, pTheme *core.Theme) JobRenderRegionTile {
  job := JobRenderRegionTile{pRegionX, pRegionZ, pWorld, pTheme}
  return job
}

// Do is the function to be executed by the worker.
func (j JobRenderRegionTile) Do() {
  regionX := j.mRegionX
  regionZ := j.mRegionZ
  region := j.mWorld.RegionManager().GetRegion(regionX, regionZ)
  img := RenderRegionTile(region, j.mTheme)
  Save("", fmt.Sprintf("tiles/r.%d.%d.png", regionX, regionZ), img)
  sLogger.Debug("End drawing", regionX, regionZ)
}
