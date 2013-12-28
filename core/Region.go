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
  x, z          int
  regionManager *RegionManager
  file          *os.File
}

// NewRegion instantiate a Region.
// regionManager pointer to the region manager who is calling the function.
// regionX region X axis.
// regionZ region Z axis.
// It returns a pointer to the region.
func NewRegion(regionManager *RegionManager, regionX, regionZ int) *Region {
  region := new(Region)
  region.regionManager = regionManager
  region.x = regionX
  region.z = regionZ
  region.file = nil
  return region
}

// NewRegionFromXYZ ...
func NewRegionFromXYZ(regionManager *RegionManager, x, y, z int) *Region {
  region := new(Region)
  regionX, regionZ := region.RegionCoordinatesFromXYZ(x, y, z)
  region.regionManager = regionManager
  region.x = regionX
  region.z = regionZ
  region.file = nil
  return region
}

// RegionCoordinatesFromXYZ ...
func (r *Region) RegionCoordinatesFromXYZ(x, y, z int) (int, int) {
  var regionX = int(math.Floor(float64(x) / (math.Pow(2, float64(z)))))
  var regionZ = int(math.Floor(float64(y) / (math.Pow(2, float64(z)))))
  return regionX, regionZ
}

// Dispose ...
func (r *Region) Dispose() {
  r.file.Close()
  r.file = nil
}

// FileName get the file name for the region.
// It returns the file name for the region.
func (r *Region) FileName() string {
  return fmt.Sprintf("r.%d.%d.mca", r.x, r.z)
}

// FilePath get the file path.
// It returns the file path.
func (r *Region) FilePath() string {
  return path.Join(r.regionManager.RegionPath(), RegionDir)
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
// localX X position of the chunk in the region.
// localZ Z position of the chunk in the region.
// It returns a pointer to the chunk.
func (r *Region) GetChunk(localX, localZ int) *Chunk {
  if r.file == nil {
    file, err := os.Open(path.Join(r.FilePath(), r.FileName()))
    if err != nil {
      log.Println(err)
    }
    r.file = file
  }
  location := r.chunkCoordinate(localX, localZ)
  r.file.Seek(int64(location), 0)
  offsetBytes := make([]byte, 3)
  var offset int64
  r.file.Read(offsetBytes)
  for _, value := range offsetBytes {
    offset = offset<<8 + int64(value)
  }

  if offset > 0 {
    r.file.Seek(offset*4096, 0)
    lengthBytes := make([]byte, 4)
    r.file.Read(lengthBytes)
    var length int64
    for _, value := range lengthBytes {
      length = length<<8 + int64(value)
    }
    versionByte := make([]byte, 1)
    r.file.Read(versionByte)
    version := int(versionByte[0])
    if version == 2 {
      compress := make([]byte, length-1)
      r.file.Read(compress)
      reader, err := zlib.NewReader(bytes.NewReader(compress))
      if err != nil {
        log.Fatal(err)
      }
      defer reader.Close()

      buf := new(bytes.Buffer)
      buf.ReadFrom(reader)
      s := buf.String()
      re := strings.NewReader(s)
      tree := nbt.NewNbtTree(re)
      chunk := NewChunk(localX, localZ)
      chunk.SetData(tree)
      return chunk
    }
  }

  return nil
}

// ChunkCoordinate get the offset of the chunk informations in the file.
// It return the offset in bytes.
func (r *Region) chunkCoordinate(localX, localZ int) int {
  return (localX + localZ*32) * 4
}
