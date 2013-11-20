package core


import (
  "mas/logger"
  "mas/nbt"
)


var s_Logger logger.Logger = logger.NewLogger(logger.DEBUG)


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


func (c *Chunk) BlockId(p_X, p_Y, p_Z int) byte {
  sectionY := p_Y / 16
  sections := c.m_Data.Root().Entries["Level"].(nbt.TagNodeCompound).Entries["Sections"].(nbt.TagNodeList)

  if int32(sectionY) >= sections.Length() {
    return 0
  }

  var section nbt.TagNode
  var blocks nbt.TagNodeByteArray
  var index int
  var blockId byte

  validBlockId := false
  for !validBlockId {
    sectionY = p_Y / 16
    section = sections.Get(sectionY)
    blocks = section.(nbt.TagNodeCompound).Entries["Blocks"].(nbt.TagNodeByteArray)
    index = (p_Y % 16) * 16 * 16 + p_Z * 16 + p_X
    blockId = blocks.Data()[index]
    if p_Y == 0 {
      return blockId
    }
    if blockId != 38 && blockId != 37 && blockId != 59 && blockId != 0 &&
       blockId != 102 && blockId != 85 && blockId != 139 && blockId != 20 && blockId != 141 &&
       blockId != 142 && blockId != 106 && blockId != 66 && blockId != 55 && blockId != 115 &&
       blockId != 83 && blockId != 104 && blockId != 105 && blockId != 140 && blockId != 68 &&
       blockId != 64 && blockId != 65 && blockId != 107 && blockId != 132 && blockId != 69 &&
       blockId != 63 && blockId != 32 && blockId != 27 && blockId != 77 && blockId != 143 &&
       blockId != 71 && blockId != 117{
      validBlockId = true
    }
    p_Y -= 1
  }
  return blockId
}





func (c *Chunk) SectionY(p_SectionY int) {
}


func (c *Chunk) Sections() {
  
}


func (c *Chunk) HeightMap() []int32 {
  heightmap := make([]int32, 256)
  if c.m_Data.Root().Entries["Level"] != nil {
    heightmap = c.m_Data.Root().Entries["Level"].(nbt.TagNodeCompound).Entries["HeightMap"].(nbt.TagNodeIntArray).Data()
  }
  return heightmap
}
