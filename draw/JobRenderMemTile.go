package draw


import (
  "mas/core"
  "image"
)


type JobRenderMemTile struct {
  m_X, m_Y, m_Z int
  m_World *core.World
  m_Theme map[byte]core.Block
  m_Chan chan *image.RGBA
}


func NewJobRenderMemTile(p_X, p_Y, p_Z int,
                            p_World *core.World,
                            p_Theme map[byte]core.Block,
                            p_Chan chan *image.RGBA) JobRenderMemTile {
  job := JobRenderMemTile{p_X, p_Y, p_Z, p_World, p_Theme, p_Chan}
  return job
}


func (j JobRenderMemTile) Do() {
  img := RenderTile(j.m_X, j.m_Y, j.m_Z, j.m_World, j.m_Theme)
  j.m_Chan <- img
}
