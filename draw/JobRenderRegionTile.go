package draw


import (
  "fmt"
  "mas/core"
)


// JobRenderRegionTile job that will render a region tile.
type JobRenderRegionTile struct {
  m_RegionX, m_RegionZ int
  m_World *core.World
  m_Theme map[byte]core.Block
}


// NewJobRenderRegionTile will instantiate a job to render a region tile.
// p_RegionX x axis coordinate of the region.
// p_RegionZ z axis coordinate of the region.
// p_World pointer to a minecraft world.
// p_Theme color theme to be used.
// It returns a JobRenderRegionTile.
func NewJobRenderRegionTile(p_RegionX, p_RegionZ int, p_World *core.World, p_Theme map[byte]core.Block) JobRenderRegionTile {
  job := JobRenderRegionTile{p_RegionX, p_RegionZ, p_World, p_Theme}
  return job
}


// Do is the function to be executed by the worker.
func (j JobRenderRegionTile) Do() {
  regionX := j.m_RegionX
  regionZ := j.m_RegionZ
  region := j.m_World.RegionManager().GetRegion(regionX, regionZ)
  img := RenderRegionTile(region, j.m_Theme)
  region.Dispose()
  Save(fmt.Sprintf("tiles/r.%d.%d.png", regionX, regionZ), img)
}
