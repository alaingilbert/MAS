package core


import (
  "mas/nbt"
)


type Chunk struct {
  m_LocalX, m_LocalZ int
  m_Data nbt.NbtTree
}


func NewChunk(p_LocalX, p_LocalZ int) *Chunk {
  chunk := Chunk{}
  chunk.m_LocalX = p_LocalX
  chunk.m_LocalZ = p_LocalZ
  return &chunk
}


func (c *Chunk) SetData(p_Data nbt.NbtTree) {
  c.m_Data = p_Data
}


func (c *Chunk) Data() nbt.NbtTree {
  return c.m_Data
}


func (c *Chunk) HeightMap() []int32 {
  heightmap := make([]int32, 256)
  if c.m_Data.Root().Entries["Level"] != nil {
    heightmap = c.m_Data.Root().Entries["Level"].(nbt.TagNodeCompound).Entries["HeightMap"].(nbt.TagNodeIntArray).Data()
  }
  return heightmap
}
