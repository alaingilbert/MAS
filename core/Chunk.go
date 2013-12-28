package core

import (
  "mas/nbt"
)

const NbSection int = 16
const SectionHeight int = 16
const XDim int = 16
const YDim int = 256
const ZDim int = 16

// Chunk ...
type Chunk struct {
  localX, localZ int
  data           *nbt.NbtTree
}

// NewChunk ...
func NewChunk(localX, localZ int) *Chunk {
  chunk := new(Chunk)
  chunk.localX = localX
  chunk.localZ = localZ
  return chunk
}

// SetData ...
func (c *Chunk) SetData(data *nbt.NbtTree) {
  c.data = data
}

// BlockID ...
func (c *Chunk) BlockID(x, y, z int) byte {
  sectionY := y / SectionHeight
  sections := c.data.Root().Entries["Level"].(*nbt.TagNodeCompound).Entries["Sections"].(*nbt.TagNodeList)

  if int32(sectionY) >= sections.Length() {
    return Air
  }

  section := sections.Get(sectionY)
  blocks := section.(*nbt.TagNodeCompound).Entries["Blocks"].(*nbt.TagNodeByteArray)
  index := (y%NbSection)*ZDim*XDim + z*XDim + x
  blockID := blocks.Data()[index]

  return blockID
}

// HeightMap ...
func (c *Chunk) HeightMap() []int32 {
  heightmap := make([]int32, XDim*ZDim)
  if c.data.Root().Entries["Level"] != nil {
    heightmap = c.data.Root().Entries["Level"].(*nbt.TagNodeCompound).Entries["HeightMap"].(*nbt.TagNodeIntArray).Data()
  }
  return heightmap
}
