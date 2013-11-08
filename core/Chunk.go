package core


type Chunk struct {
  m_LocalX, m_LocalZ int
}


func NewChunk(p_LocalX, p_LocalZ int) *Chunk {
  chunk := Chunk{}
  chunk.m_LocalX = p_LocalX
  chunk.m_LocalZ = p_LocalZ
  return &chunk
}
