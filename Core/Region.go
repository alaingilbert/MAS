package Core


import (
  "bytes"
  "compress/zlib"
  "fmt"
  "log"
  "mas/Nbt"
  "os"
  "path"
)


const REGION_DIR = "region"


type Region struct {
  m_X, m_Z int
  m_RegionManager *RegionManager
}


func NewRegion(p_RegionManager *RegionManager, p_X, p_Z int) *Region {
  region := Region{}
  region.m_RegionManager = p_RegionManager
  region.m_X = p_X
  region.m_Z = p_Z
  return &region
}


func (r *Region) FileName() string {
  return fmt.Sprintf("r.%d.%d.mca", r.m_X, r.m_Z)
}


func (r *Region) FilePath() string {
  return path.Join(r.m_RegionManager.RegionPath(), REGION_DIR)
}


func (r *Region) GetChunk(p_LocalX, p_LocalZ int) *Chunk {
  chunk := NewChunk(p_LocalX, p_LocalZ)
  file, err := os.Open(path.Join(r.FilePath(), r.FileName()))
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  location := r.chunkCoordinate(p_LocalX, p_LocalZ)
  file.Seek(int64(location), 0)
  offsetBytes := make([]byte, 3)
  var offset int64
  file.Read(offsetBytes)
  for _, value := range offsetBytes {
    offset = offset << 8 + int64(value)
  }

  if offset > 0 {
    file.Seek(offset * 4096, 0)
    lengthBytes := make([]byte, 4)
    file.Read(lengthBytes)
    var length int64
    for _, value := range lengthBytes { length = length << 8 + int64(value) }
    versionByte := make([]byte, 1)
    file.Read(versionByte)
    version := int(versionByte[0])
    if version == 2 {
      compress := make([]byte, length - 1)
      file.Read(compress)
      reader, err := zlib.NewReader(bytes.NewReader(compress))
      if err != nil {
        log.Fatal(err)
      }
      defer reader.Close()
      tree := Nbt.NbtTree{}
      tree.Init(reader)
      fmt.Println(tree)
    }
  }

  return chunk
}


func (r *Region) chunkCoordinate(p_LocalX, p_LocalZ int) int {
  return (p_LocalX + p_LocalZ * 32) * 4
}
