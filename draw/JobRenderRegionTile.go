package draw


import (
  "fmt"
  "mas/core"
)


type JobRenderRegionTile struct {
  m_RegionX, m_RegionZ int
  m_World *core.World
  m_Theme map[byte]core.Block
}


func NewJobRenderRegionTile(p_RegionX, p_RegionZ int, p_World *core.World, p_Theme map[byte]core.Block) JobRenderRegionTile {
  job := JobRenderRegionTile{p_RegionX, p_RegionZ, p_World, p_Theme}
  return job
}


func (j JobRenderRegionTile) Do() {
  regionX := j.m_RegionX
  regionZ := j.m_RegionZ
  region := j.m_World.RegionManager().GetRegion(regionX, regionZ)
  img := RenderRegionTile(region, j.m_Theme)
  Save(fmt.Sprintf("tiles/r.%d.%d.png", regionX, regionZ), img)
}
