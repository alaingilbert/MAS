package draw

import (
  "image"
  "mas/core"
)

// JobRenderMemTile ...
type JobRenderMemTile struct {
  mX, mY, mZ int
  mWorld       *core.World
  mTheme       *core.Theme
  mChan        chan *image.RGBA
}

// NewJobRenderMemTile ...
func NewJobRenderMemTile(pX, pY, pZ int,
  pWorld *core.World,
  pTheme *core.Theme,
  pChan chan *image.RGBA) JobRenderMemTile {
  job := JobRenderMemTile{pX, pY, pZ, pWorld, pTheme, pChan}
  return job
}

// Do ...
func (j JobRenderMemTile) Do() {
  img := RenderTile(j.mX, j.mY, j.mZ, j.mWorld, j.mTheme)
  j.mChan <- img
}
