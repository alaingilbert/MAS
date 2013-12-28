package draw

import (
  "fmt"
  "image"
  "image/color"
  "image/png"
  "mas/core"
  "math"
  "os"
)

const TileSize int = 256

// CreateImage ...
func CreateImage(pSizeX, pSizeZ int) *image.RGBA {
  return image.NewRGBA(image.Rect(0, 0, pSizeX, pSizeZ))
}

// Save ...
func Save(pPath, pFileName string, pImg *image.RGBA) {
  os.MkdirAll(pPath, 0700)
  file, err := os.Create(pPath + "" + pFileName)
  if err != nil {
    fmt.Print(err)
  }
  defer file.Close()
  png.Encode(file, pImg)
}

// FillRect ...
func FillRect(pImg *image.RGBA, pX, pZ, pWidth, pHeight int, pColor color.Color) {
  if pWidth == 1 && pHeight == 1 {
    pImg.Set(pX, pZ, pColor)
    return
  }
  for i := pX; i < pX+pWidth; i++ {
    for j := pZ; j < pZ+pHeight; j++ {
      pImg.Set(i, j, pColor)
    }
  }
}

// StartingChunk ...
func StartingChunk(x, z int) int {
  twoExpZ := int(math.Pow(2, float64(z)))
  mod := ((x % twoExpZ) + twoExpZ) % twoExpZ
  tmp := mod * int(32/twoExpZ)
  return tmp
}

// NbChunk ...
func NbChunk(z int) int {
  return int(32 / math.Pow(2, float64(z)))
}

// IsWater ...
func IsWater(blockID byte) bool {
  return blockID == core.Water ||
    blockID == core.FlowingWater ||
    blockID == core.Ice
}

// IsBlockIdDrawable ...
func IsBlockIdDrawable(blockID byte) bool {
  return blockID != core.RedFlower &&
    blockID != core.YellowFlower &&
    blockID != core.Wheat &&
    blockID != core.Air &&
    blockID != core.GlassPane &&
    blockID != core.Fence &&
    blockID != core.CobblestoneWall &&
    blockID != core.Glass &&
    blockID != core.Carrots &&
    blockID != core.Potatoes &&
    blockID != core.Vine &&
    blockID != core.Rail &&
    blockID != core.RedstoneWire &&
    blockID != core.NetherWart &&
    blockID != core.Reeds &&
    blockID != core.PumpkinStem &&
    blockID != core.MelonStem &&
    blockID != core.FlowerPot &&
    blockID != core.WallSign &&
    blockID != core.WoodenDoor &&
    blockID != core.Ladder &&
    blockID != core.FenceGate &&
    blockID != core.Tripwire &&
    blockID != core.Lever &&
    blockID != core.StandingSign &&
    blockID != core.DeadBush &&
    blockID != core.GoldenRail &&
    blockID != core.StoneButton &&
    blockID != core.WoodenButton &&
    blockID != core.IronDoor &&
    blockID != core.BrewingStand &&
    blockID != core.StainedGlass &&
    blockID != core.StainedGlassPane
}

// RenderTile ...
func RenderTile(x, y, z int, world *core.World, theme *core.Theme) *image.RGBA {
  blockSize := 1
  chunkSize := 16 * blockSize
  startingChunkX := StartingChunk(x, z)
  startingChunkZ := StartingChunk(y, z)
  nbChunk := NbChunk(z)
  scale := GetScale(z)
  skip := BlockToSkip(z)

  region := world.RegionManager().GetRegionFromXYZ(x, y, z)
  if !region.Exists() {
    return nil
  }
  img := CreateImage(TileSize, TileSize)
  for chunkX := startingChunkX; chunkX < startingChunkX+nbChunk; chunkX++ {
    for chunkZ := startingChunkZ; chunkZ < startingChunkZ+nbChunk; chunkZ++ {
      chunk := region.GetChunk(chunkX, chunkZ)
      if chunk == nil {
        continue
      }

      heightmap := chunk.HeightMap()

      for heightmapBlockIndex := 0; heightmapBlockIndex < 256; heightmapBlockIndex += skip {
        blockY := int(uint8(heightmap[heightmapBlockIndex]))
        blockX := heightmapBlockIndex % 16
        blockZ := heightmapBlockIndex / 16
        blockID := chunk.BlockID(blockX, blockY, blockZ)

        for !IsBlockIdDrawable(blockID) {
          blockY--
          blockID = chunk.BlockID(blockX, blockY, blockZ)
        }

        block := theme.GetByID(blockID)
        blockColor := color.RGBA{block.Red, block.Green, block.Blue, block.Alpha}

        if IsWater(blockID) {
          for IsWater(blockID) || !IsBlockIdDrawable(blockID) {
            blockY--
            blockID = chunk.BlockID(blockX, blockY, blockZ)
          }
          block = theme.GetByID(blockID)

          alpha1 := 1.0
          alpha2 := 0.2

          alpha := alpha2 + alpha1*(1.0-alpha2)
          red := uint8((float64(block.Red)*alpha2 + float64(blockColor.R)*alpha1*(1.0-alpha2)) / alpha)
          green := uint8((float64(block.Green)*alpha2 + float64(blockColor.G)*alpha1*(1.0-alpha2)) / alpha)
          blue := uint8((float64(block.Blue)*alpha2 + float64(blockColor.B)*alpha1*(1.0-alpha2)) / alpha)

          blockColor = color.RGBA{red, green, blue, 255}
        }

        FillRect(img,
          (heightmapBlockIndex%16+(chunkX-startingChunkX)*chunkSize)*scale/skip,
          (heightmapBlockIndex/16+(chunkZ-startingChunkZ)*chunkSize)*scale/skip,
          blockSize*scale,
          blockSize*scale,
          blockColor)
      }
    }
  }
  region.Dispose()
  return img
}

// BlockToSkip ...
func BlockToSkip(z int) int {
  if z == 0 {
    return 2
  }
  return 1
}

// GetScale ...
func GetScale(z int) int {
  if z == 0 {
    return 1
  }
  return int(math.Pow(2, float64((z - 1))))
}

// RenderRegionTile render a tile for a given region.
// region the region to render.
// It returns an image tile.
func RenderRegionTile(region *core.Region, theme *core.Theme) *image.RGBA {
  blockSize := 1
  chunkSize := 16 * blockSize
  regionSize := 32 * chunkSize

  img := CreateImage(regionSize, regionSize)

  if !region.Exists() {
    return img
  }

  for chunkX := 0; chunkX < 32; chunkX++ {
    for chunkZ := 0; chunkZ < 32; chunkZ++ {
      chunk := region.GetChunk(chunkX, chunkZ)
      if chunk == nil {
        continue
      }

      heightmap := chunk.HeightMap()

      for block := 0; block < 256; block++ {
        chunkY := uint8(heightmap[block])
        blockX := block % 16
        blockZ := block / 16
        blockID := chunk.BlockID(blockX, int(chunkY), blockZ)
        c := theme.GetByID(blockID)
        c2 := color.RGBA{c.Red, c.Green, c.Blue, c.Alpha}

        FillRect(img,
          block%16+chunkX*chunkSize,
          block/16+chunkZ*chunkSize,
          blockSize,
          blockSize,
          c2)
      }
    }
  }
  return img
}
