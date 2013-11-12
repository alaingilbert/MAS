package draw


import (
  "fmt"
  "image"
  "image/png"
  "image/color"
  "mas/core"
  "mas/logger"
  "os"
)


var s_Logger logger.Logger = logger.NewLogger(logger.DEBUG)


func CreateImage(p_SizeX, p_SizeZ int) *image.RGBA {
  return image.NewRGBA(image.Rect(0, 0, p_SizeX, p_SizeZ))
}


func Save(p_FileName string, p_Img *image.RGBA) {
  file, err := os.Create(p_FileName)
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


// RenderRegionTile render a tile for a given region.
// p_Region the region to render.
// It returns an image tile.
func RenderRegionTile(p_Region *core.Region, theme map[byte]core.Block) *image.RGBA {
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
        c := theme[blockId]
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
