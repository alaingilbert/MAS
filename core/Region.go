package core

import (
  "bytes"
  "compress/zlib"
  "fmt"
  "log"
  "mas/nbt"
  "math"
  "os"
  "path"
  "strings"
)

// Region represent a minecraft region.
type Region struct {
  mX, mZ         int
  mRegionManager *RegionManager
  mChunks        [1024]*Chunk
  mData          nbt.NbtTree
  mFile          *os.File
}

// NewRegion instantiate a Region.
// pRegionManager pointer to the region manager who is calling the function.
// pX region X axis.
// pZ region Z axis.
// It returns a pointer to the region.
func NewRegion(pRegionManager *RegionManager, pRegionX, pRegionZ int) *Region {
  region := Region{}
  region.mRegionManager = pRegionManager
  region.mX = pRegionX
  region.mZ = pRegionZ
  region.mFile = nil
  return &region
}

// NewRegionFromXYZ ...
func NewRegionFromXYZ(pRegionManager *RegionManager, pX, pY, pZ int) *Region {
  region := Region{}
  regionX, regionZ := region.RegionCoordinatesFromXYZ(pX, pY, pZ)
  region.mRegionManager = pRegionManager
  region.mX = regionX
  region.mZ = regionZ
  region.mFile = nil
  return &region
}

// RegionCoordinatesFromXYZ ...
func (r *Region) RegionCoordinatesFromXYZ(x, y, z int) (int, int) {
  var regionX = int(math.Floor(float64(x) / (math.Pow(2, float64(z)))))
  var regionZ = int(math.Floor(float64(y) / (math.Pow(2, float64(z)))))
  return regionX, regionZ
}

// Dispose ...
func (r *Region) Dispose() {
  r.mFile.Close()
  r.mFile = nil
}

// FileName get the file name for the region.
// It returns the file name for the region.
func (r *Region) FileName() string {
  return fmt.Sprintf("r.%d.%d.mca", r.mX, r.mZ)
}

// FilePath get the file path.
// It returns the file path.
func (r *Region) FilePath() string {
  return path.Join(r.mRegionManager.RegionPath(), RegionDir)
}

// Exists ...
func (r *Region) Exists() bool {
  path := path.Join(r.FilePath(), r.FileName())
  _, err := os.Stat(path)
  if err == nil {
    return true
  }
  if os.IsNotExist(err) {
    return false
  }
  return false
}

// GetChunk get the information for a specific chunk.
// pLocalX X position of the chunk in the region.
// pLocalZ Z position of the chunk in the region.
// It returns a pointer to the chunk.
func (r *Region) GetChunk(pLocalX, pLocalZ int) *Chunk {
  if r.mFile == nil {
    file, err := os.Open(path.Join(r.FilePath(), r.FileName()))
    if err != nil {
      log.Println(err)
    }
    r.mFile = file
  }
  location := r.chunkCoordinate(pLocalX, pLocalZ)
  r.mFile.Seek(int64(location), 0)
  offsetBytes := make([]byte, 3)
  var offset int64
  r.mFile.Read(offsetBytes)
  for _, value := range offsetBytes {
    offset = offset<<8 + int64(value)
  }

  if offset > 0 {
    r.mFile.Seek(offset*4096, 0)
    lengthBytes := make([]byte, 4)
    r.mFile.Read(lengthBytes)
    var length int64
    for _, value := range lengthBytes {
      length = length<<8 + int64(value)
    }
    versionByte := make([]byte, 1)
    r.mFile.Read(versionByte)
    version := int(versionByte[0])
    if version == 2 {
      compress := make([]byte, length-1)
      r.mFile.Read(compress)
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
      chunk := NewChunk(pLocalX, pLocalZ)
      chunk.SetData(tree)
      return chunk
    }
  }

  return nil
}

// ChunkCoordinate get the offset of the chunk informations in the file.
// It return the offset in bytes.
func (r *Region) chunkCoordinate(pLocalX, pLocalZ int) int {
  return (pLocalX + pLocalZ*32) * 4
}
