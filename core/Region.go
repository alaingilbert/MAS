package core


import (
  "bytes"
  "compress/zlib"
  "fmt"
  "log"
  "mas/nbt"
  "os"
  "path"
  "strings"
)


// Region represent a minecraft region.
type Region struct {
  m_X, m_Z int
  m_RegionManager *RegionManager
  m_Chunks [1024]*Chunk
  m_Data nbt.NbtTree
  m_File *os.File
}


// NewRegion instantiate a Region.
// p_RegionManager pointer to the region manager who is calling the function.
// p_X region X axis.
// p_Z region Z axis.
// It returns a pointer to the region.
func NewRegion(p_RegionManager *RegionManager, p_X, p_Z int) *Region {
  region := Region{}
  region.m_RegionManager = p_RegionManager
  region.m_X = p_X
  region.m_Z = p_Z
  file, err := os.Open(path.Join(region.FilePath(), region.FileName()))
  if err != nil {
    log.Println(err)
    return nil
  }
  region.m_File = file
  return &region
}


func (r *Region) Dispose() {
  r.m_File.Close()
}


// FileName get the file name for the region.
// It returns the file name for the region.
func (r *Region) FileName() string {
  return fmt.Sprintf("r.%d.%d.mca", r.m_X, r.m_Z)
}


// FilePath get the file path.
// It returns the file path.
func (r *Region) FilePath() string {
  return path.Join(r.m_RegionManager.RegionPath(), REGION_DIR)
}


func (r *Region) Exists() bool {
  path := path.Join(r.FilePath(), r.FileName())
  _, err := os.Stat(path)
  if err == nil { return true }
  if os.IsNotExist(err) { return false }
  return false
}


// GetChunk get the information for a specific chunk.
// p_LocalX X position of the chunk in the region.
// p_LocalZ Z position of the chunk in the region.
// It returns a pointer to the chunk.
func (r *Region) GetChunk(p_LocalX, p_LocalZ int) *Chunk {
  location := r.chunkCoordinate(p_LocalX, p_LocalZ)
  r.m_File.Seek(int64(location), 0)
  offsetBytes := make([]byte, 3)
  var offset int64
  r.m_File.Read(offsetBytes)
  for _, value := range offsetBytes {
    offset = offset << 8 + int64(value)
  }

  if offset > 0 {
    r.m_File.Seek(offset * 4096, 0)
    lengthBytes := make([]byte, 4)
    r.m_File.Read(lengthBytes)
    var length int64
    for _, value := range lengthBytes { length = length << 8 + int64(value) }
    versionByte := make([]byte, 1)
    r.m_File.Read(versionByte)
    version := int(versionByte[0])
    if version == 2 {
      compress := make([]byte, length - 1)
      r.m_File.Read(compress)
      reader, err := zlib.NewReader(bytes.NewReader(compress))
      if err != nil {
        log.Fatal(err)
      }
      defer reader.Close()
      tree := nbt.NewNbtTree()

      buf := new(bytes.Buffer)
      buf.ReadFrom(reader)
      s := buf.String()
      re := strings.NewReader(s)
      tree.Init(re)
      chunk := NewChunk(p_LocalX, p_LocalZ)
      chunk.SetData(tree)
      return chunk
    }
  }

  return nil
}


// ChunkCoordinate get the offset of the chunk informations in the file.
// It return the offset in bytes.
func (r *Region) chunkCoordinate(p_LocalX, p_LocalZ int) int {
  return (p_LocalX + p_LocalZ * 32) * 4
}
