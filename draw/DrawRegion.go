package draw


import (
  "fmt"
  "image"
  "image/png"
  "image/color"
  "mas/core"
  "os"
)


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
func RenderRegionTile(p_Region *core.Region) *image.RGBA {
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
        var c color.RGBA
        switch blockId {
          case 0: c = color.RGBA{255, 0, 0, 255}
          case 1: c = color.RGBA{50, 50, 50, 255} // Stone
          case 2: c = color.RGBA{0, 255, 0, 255} // Grass
          case 3: c = color.RGBA{0, 255, 0, 255} // Dirt
          case 4: c = color.RGBA{50, 50, 50, 255} // CobbleStone
          case 33: c = color.RGBA{50, 50, 50, 255} // piston
          case 5: c = color.RGBA{89, 71, 43, 255} // Planks
          case 6: c = color.RGBA{78, 198, 50, 255} // sapling
          case 8: c = color.RGBA{0, 0, 255, 255} // Flowing water
          case 9: c = color.RGBA{0, 0, 255, 255} // Water
          case 37: c = color.RGBA{253, 255, 106, 255} // yellow_flower
          case 38: c = color.RGBA{235, 96, 96, 255} // red_flower
          case 31: c = color.RGBA{0, 255, 0, 255} // Tall Grass
          case 59: c = color.RGBA{136, 230, 45, 255} // wheat
          case 170: c = color.RGBA{136, 230, 45, 255} // hay_block
          case 39: c = color.RGBA{0, 255, 0, 255} // brown_mushroom
          case 99: c = color.RGBA{0, 255, 0, 255} // brown_mushroom_block
          case 60: c = color.RGBA{0, 255, 0, 255} // FarmLand
          case 79: c = color.RGBA{90, 90, 255, 255} // Ice
          case 43: c = color.RGBA{50, 50, 50, 255} // Double stone slab
          case 155: c = color.RGBA{120, 120, 120, 255} // Quartz block
          case 156: c = color.RGBA{120, 120, 120, 255} // Quartz stair
          case 44: c = color.RGBA{50, 50, 50, 255} // Stone slab
          case 98: c = color.RGBA{50, 50, 50, 255} // Stone brick
          case 109: c = color.RGBA{50, 50, 50, 255} // Stone brick stair
          case 110: c = color.RGBA{50, 50, 50, 255} // mycelium
          case 118: c = color.RGBA{50, 50, 50, 255} // cauldron
          case 97: c = color.RGBA{50, 50, 50, 255} // monster_egg (stone)
          case 15: c = color.RGBA{50, 50, 50, 255} // coal_ore
          case 82: c = color.RGBA{50, 50, 50, 255} // clay
          case 48: c = color.RGBA{50, 50, 50, 255} // mossy_cobblestone
          case 16: c = color.RGBA{50, 50, 50, 255} // coal_ore
          case 61: c = color.RGBA{50, 50, 50, 255} // Furnace
          case 93: c = color.RGBA{50, 50, 50, 255} // unpowered_repeater
          case 94: c = color.RGBA{50, 50, 50, 255} // powered_repeater
          case 70: c = color.RGBA{50, 50, 50, 255} // stone_pressure_plate
          case 81: c = color.RGBA{78, 198, 50, 255} // cactus
          case 154: c = color.RGBA{50, 50, 50, 255} // hopper
          case 139: c = color.RGBA{50, 50, 50, 255} // Cobblestone_wall
          case 145: c = color.RGBA{50, 50, 50, 255} // anvil
          case 67: c = color.RGBA{50, 50, 50, 255} // Stone stairs
          case 13: c = color.RGBA{50, 50, 50, 255} // Gravel
          case 134: c = color.RGBA{89, 71, 43, 255} // Spruce stairs
          case 53: c = color.RGBA{89, 71, 43, 255} // Oak stairs
          case 136: c = color.RGBA{89, 71, 43, 255} // jungle_stairs
          case 135: c = color.RGBA{89, 71, 43, 255} // birch_stairs
          case 22: c = color.RGBA{76, 80, 255, 255} // lapis_block
          case 126: c = color.RGBA{89, 71, 43, 255} // Wooden slab
          case 84: c = color.RGBA{89, 71, 43, 255} // jukebox
          case 65: c = color.RGBA{89, 71, 43, 255} // ladder
          case 107: c = color.RGBA{89, 71, 43, 255} // fence_gate
          case 112: c = color.RGBA{103, 22, 22, 255} // nether_brick
          case 114: c = color.RGBA{103, 22, 22, 255} // nether_brick_stairs
          case 32: c = color.RGBA{89, 71, 43, 255} // deadbush
          case 63: c = color.RGBA{89, 71, 43, 255} // standing_sign
          case 72: c = color.RGBA{89, 71, 43, 255} // wooden_pressure_plate
          case 47: c = color.RGBA{89, 71, 43, 255} // bookshelf
          case 96: c = color.RGBA{89, 71, 43, 255} // trapdoor
          case 85: c = color.RGBA{89, 71, 43, 255} // fence
          case 58: c = color.RGBA{89, 71, 43, 255} // crafting_table
          case 54: c = color.RGBA{89, 71, 43, 255} // chest
          case 146: c = color.RGBA{89, 71, 43, 255} // trapped_chest
          case 106: c = color.RGBA{29, 72, 22, 255} // vine
          case 125: c = color.RGBA{89, 71, 43, 255} // double_wooden_slab
          case 88: c = color.RGBA{59, 44, 0, 255} // soul_sand
          case 128: c = color.RGBA{59, 44, 0, 255} // sandstone_stairs
          case 49: c = color.RGBA{0, 0, 0, 255} // obsidian
          case 35: c = color.RGBA{190, 190, 190, 255} // wool
          case 92: c = color.RGBA{190, 190, 190, 255} // cake
          case 171: c = color.RGBA{190, 190, 190, 255} // carpet
          case 17: c = color.RGBA{89, 71, 43, 255} // log
          case 18: c = color.RGBA{61, 96, 49, 255} // leaves
          case 78: c = color.RGBA{255, 255, 255, 255} // Snow
          case 138: c = color.RGBA{255, 255, 255, 255} // beacon
          case 30: c = color.RGBA{255, 255, 255, 255} // web
          case 12: c = color.RGBA{255, 213, 131, 255} // Sand
          case 83: c = color.RGBA{255, 213, 131, 255} // reeds
          case 24: c = color.RGBA{255, 213, 131, 255} // Sandstone
          case 103: c = color.RGBA{64, 135, 4, 255} // Melon block
          case 111: c = color.RGBA{45, 79, 16, 255} // Waterlily
          case 11: c = color.RGBA{210, 69, 0, 255} // Lava
          case 51: c = color.RGBA{210, 69, 0, 255} // fire
          case 10: c = color.RGBA{210, 69, 0, 255} // Flowing lava
          case 159: c = color.RGBA{232, 201, 184, 255} // Stained hardened clay
          case 50: c = color.RGBA{255, 255, 0, 255} // torch
          case 14: c = color.RGBA{255, 255, 0, 255} // gold_ore
          case 76: c = color.RGBA{255, 0, 0, 255} // Redstone torch
          case 100: c = color.RGBA{255, 0, 0, 255} // red_mushroom_block
          case 40: c = color.RGBA{255, 0, 0, 255} // red_mushroom
          case 55: c = color.RGBA{255, 0, 0, 255} // redstone_wire
          case 26: c = color.RGBA{255, 0, 0, 255} // bed
          case 86: c = color.RGBA{255, 108, 0, 255} // Pumpkin
          case 89: c = color.RGBA{255, 255, 0, 255} // Glowstone
          case 172: c = color.RGBA{193, 85, 48, 255} // hardened_clay
          case 87: c = color.RGBA{160, 55, 55, 255} // netherrack
          case 45: c = color.RGBA{235, 96, 96, 255} // brick_block
          case 108: c = color.RGBA{235, 96, 96, 255} // brick_stairs
          case 42: c = color.RGBA{183, 183, 183, 255} // iron_block
          case 148: c = color.RGBA{183, 183, 183, 255} // heavy_weighted_pressure_plate
          case 41: c = color.RGBA{255, 255, 0, 255} // gold_block
          case 57: c = color.RGBA{88, 219, 217, 255} // diamond_block
          case 130: c = color.RGBA{85, 39, 125, 255} // ender_chest
          case 152: c = color.RGBA{165, 39, 39, 255} // redstone_block
          default:
            fmt.Println(blockId)
            c = color.RGBA{chunkY, chunkY, chunkY, 255}
        }
        FillRect(img,
                 block % 16 + chunkX * chunkSize,
                 block / 16 + chunkZ * chunkSize,
                 blockSize,
                 blockSize,
                 c)
      }
    }
  }
  return img
}
