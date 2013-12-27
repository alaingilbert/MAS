package core

import (
  "mas/nbt"
)

// Chunk ...
type Chunk struct {
  mLocalX, mLocalZ int
  mData            *nbt.NbtTree
}

// NewChunk ...
func NewChunk(pLocalX, pLocalZ int) *Chunk {
  chunk := new(Chunk)
  chunk.mLocalX = pLocalX
  chunk.mLocalZ = pLocalZ
  return chunk
}

// SetData ...
func (c *Chunk) SetData(pData *nbt.NbtTree) {
  c.mData = pData
}

// BlockID ...
func (c *Chunk) BlockID(pX, pY, pZ int) byte {
  sectionY := pY / 16
  sections := c.mData.Root().Entries["Level"].(*nbt.TagNodeCompound).Entries["Sections"].(*nbt.TagNodeList)

  if int32(sectionY) >= sections.Length() {
    return 0
  }

  var section nbt.ITagNode
  var blocks nbt.TagNodeByteArray
  var index int
  var blockID byte

  validBlockID := false
  for !validBlockID {
    sectionY = pY / 16
    section = sections.Get(sectionY)
    blocks = section.(*nbt.TagNodeCompound).Entries["Blocks"].(nbt.TagNodeByteArray)
    index = (pY%16)*16*16 + pZ*16 + pX
    blockID = blocks.Data()[index]
    if pY == 0 {
      return blockID
    }
    if blockID != 38 && blockID != 37 && blockID != 59 && blockID != 0 &&
      blockID != 102 && blockID != 85 && blockID != 139 && blockID != 20 && blockID != 141 &&
      blockID != 142 && blockID != 106 && blockID != 66 && blockID != 55 && blockID != 115 &&
      blockID != 83 && blockID != 104 && blockID != 105 && blockID != 140 && blockID != 68 &&
      blockID != 64 && blockID != 65 && blockID != 107 && blockID != 132 && blockID != 69 &&
      blockID != 63 && blockID != 32 && blockID != 27 && blockID != 77 && blockID != 143 &&
      blockID != 71 && blockID != 117 && blockID != 95 && blockID != 160 {
      validBlockID = true
    }
    pY -= 1
  }
  return blockID
}

// SectionY ...
func (c *Chunk) SectionY(pSectionY int) {
}

// Sections ...
func (c *Chunk) Sections() {
}

// HeightMap ...
func (c *Chunk) HeightMap() []int32 {
  heightmap := make([]int32, 256)
  if c.mData.Root().Entries["Level"] != nil {
    heightmap = c.mData.Root().Entries["Level"].(*nbt.TagNodeCompound).Entries["HeightMap"].(nbt.TagNodeIntArray).Data()
  }
  return heightmap
}
