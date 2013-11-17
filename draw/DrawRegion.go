package draw


import (
  "fmt"
  "image"
  "image/png"
  "image/color"
  "mas/core"
  "mas/logger"
  "math"
  "os"
)


var s_Logger logger.Logger = logger.NewLogger(logger.DEBUG)


func CreateImage(p_SizeX, p_SizeZ int) *image.RGBA {
  return image.NewRGBA(image.Rect(0, 0, p_SizeX, p_SizeZ))
}


func Save(p_Path, p_FileName string, p_Img *image.RGBA) {
  os.MkdirAll(p_Path, 0700)
  file, err := os.Create(p_Path + "" + p_FileName)
  if err != nil {
    fmt.Print(err)
  }
  defer file.Close()
  png.Encode(file, p_Img)
}


func FillRect(p_Img *image.RGBA, p_X, p_Z, p_Width, p_Height int, p_Color color.Color) {
  if p_Width == 1 && p_Height == 1 {
    p_Img.Set(p_X, p_Z, p_Color)
    return
  }
  for i := p_X; i < p_X + p_Width; i++ {
    for j := p_Z; j < p_Z + p_Height; j++ {
      p_Img.Set(i, j, p_Color)
    }
  }
}


func GetRegionFromXYZ(x, y, z int) (int, int) {
  var regionX int = int(math.Floor(float64(x) / (math.Pow(2, float64(z)))))
  var regionZ int = int(math.Floor(float64(y) / (math.Pow(2, float64(z)))))
  return regionX, regionZ
}


func StartingChunk(x, z int) int {
  twoExpZ := int(math.Pow(2, float64(z)))
  mod := ((x % twoExpZ) + twoExpZ) % twoExpZ
  tmp := mod * int(32 / twoExpZ)
  return tmp
}


func NbChunk(z int) int {
  return int(32 / math.Pow(2, float64(z)))
}


func RenderTile(x, y, z int, p_World *core.World, p_Theme map[byte]core.Block) *image.RGBA {
  blockSize := 1
  chunkSize := 16 * blockSize
  regionX, regionZ := GetRegionFromXYZ(x, y, z)
  startingChunkX := StartingChunk(x, z)
  startingChunkZ := StartingChunk(y, z)
  nbChunk := NbChunk(z)
  scale := GetScale(z)
  skip := BlockToSkip(z)

  img := CreateImage(256, 256)
  region := p_World.RegionManager().GetRegion(regionX, regionZ)
  for chunkX := startingChunkX; chunkX < startingChunkX + nbChunk; chunkX++ {
    for chunkZ := startingChunkZ; chunkZ < startingChunkZ + nbChunk; chunkZ++ {
      chunk := region.GetChunk(chunkX, chunkZ)
      if chunk == nil {
        continue
      }

      heightmap := chunk.HeightMap()

      for block := 0; block < 256; block += skip {
        chunkY := uint8(heightmap[block])
        blockX := block % 16
        blockZ := block / 16
        blockId := chunk.BlockId(blockX, int(chunkY), blockZ)
        c := p_Theme[blockId]
        c2 := color.RGBA{c.Red, c.Green, c.Blue, c.Alpha}

        FillRect(img,
                 (block % 16 + (chunkX-startingChunkX) * chunkSize) * scale / skip,
                 (block / 16 + (chunkZ-startingChunkZ) * chunkSize) * scale / skip,
                 blockSize * scale,
                 blockSize * scale,
                 c2)
      }
    }
  }
  return img
}


func BlockToSkip(z int) int {
 if z == 0 {
  return 2
 } else {
  return 1
 }
}


func GetScale(z int) int {
  if z == 0 { return 1 }
  return int(math.Pow(2, float64((z - 1))))
}


// RenderRegionTile render a tile for a given region.
// p_Region the region to render.
// It returns an image tile.
func RenderRegionTile(p_Region *core.Region, p_Theme map[byte]core.Block) *image.RGBA {
  blockSize := 1
  chunkSize := 16 * blockSize
  regionSize := 32 * chunkSize

  img := CreateImage(regionSize, regionSize)

  if !p_Region.Exists() {
    return img
  }

  for chunkX := 0; chunkX < 32; chunkX++ {
    for chunkZ := 0; chunkZ < 32; chunkZ++ {
      chunk := p_Region.GetChunk(chunkX, chunkZ)
      if chunk == nil {
        continue
      }

      heightmap := chunk.HeightMap()

      for block := 0; block < 256; block++ {
        chunkY := uint8(heightmap[block])
        blockX := block % 16
        blockZ := block / 16
        blockId := chunk.BlockId(blockX, int(chunkY), blockZ)
        c := p_Theme[blockId]
        c2 := color.RGBA{c.Red, c.Green, c.Blue, c.Alpha}

        FillRect(img,
                 block % 16 + chunkX * chunkSize,
                 block / 16 + chunkZ * chunkSize,
                 blockSize,
                 blockSize,
                 c2)
      }
    }
  }
  return img
}
