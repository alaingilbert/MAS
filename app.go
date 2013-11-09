package main


import (
  "strings"
  "fmt"
  "io"
  "log"
  "mas/core"
  "mas/draw"
  "net/http"
  "image/color"
  "strconv"
)


const PORT int = 8000


func HomeHandler(w http.ResponseWriter, req *http.Request) {
  fmt.Println(req.URL.Query()["x"])
  io.WriteString(w, "Hello\n")
}


func main() {
  log.Println("Start")

  world := core.NewWorld("/Users/agilbert/Desktop/minecraft/world")

  blockSize := 1
  chunkSize := 16 * blockSize
  regionSize := 32 * chunkSize

  for _, fileName := range world.RegionManager().RegionFileNames() {
    if !strings.HasSuffix(fileName, "mca") {
      continue
    }
    splits := strings.SplitN(fileName, ".", 4)
    regionX, _ := strconv.Atoi(splits[1])
    regionZ, _ := strconv.Atoi(splits[2])
    region := world.RegionManager().GetRegion(regionX, regionZ)
    fmt.Println("Region", regionX, regionZ)
    img := draw.CreateImage(regionSize, regionSize)
    if !region.Exists() {
      continue
    }
    for chunkX := 0; chunkX < 32; chunkX++ {
      for chunkZ := 0; chunkZ < 32; chunkZ++ {
        chunk := region.GetChunk(chunkX, chunkZ)
        heightmap := chunk.HeightMap()
        for block := 0; block < 256; block++ {
          c := uint8(heightmap[block])
          draw.FillRect(img,
                        block % 16 + chunkX * chunkSize,
                        block / 16 + chunkZ * chunkSize,
                        blockSize,
                        blockSize,
                        color.RGBA{c, c, c, 255})
        }
      }
    }
    draw.Save(fmt.Sprintf("tiles/r%d.%d.png", regionX, regionZ), img)
  }


  log.Println("End")
  //http.HandleFunc("/", HomeHandler)
  //http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}
