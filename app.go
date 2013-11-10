package main


import (
  "fmt"
  "mas/core"
  "mas/draw"
  "mas/logger"
  "runtime"
  "strconv"
  "strings"
  "time"
)


var s_Logger logger.Logger = logger.NewLogger(logger.INFO)


func main() {

  s_Logger.Info("Start")
  startTime := time.Now()
  runtime.GOMAXPROCS(4)

  world := core.NewWorld("/Users/agilbert/Desktop/minecraft/world")

  nbThread := 4
  in := make(chan []int)

  for i := 0; i < nbThread; i++ {
    go Worker(i, world, in)
  }

  files := world.RegionManager().RegionFileNames()
  for index, fileName := range files {
    if index > 50 {
      break
    }
    if !strings.HasSuffix(fileName, "mca") {
      continue
    }
    splits := strings.SplitN(fileName, ".", 4)
    regionX, _ := strconv.Atoi(splits[1])
    regionZ, _ := strconv.Atoi(splits[2])

    data := make([]int, 2)
    data[0] = regionX
    data[1] = regionZ
    in <- data
  }

  close(in)


  s_Logger.Info("End", time.Since(startTime))
}


func Worker(p_Id int, p_World *core.World, in chan []int) {
  for data := range in {
    regionX := data[0]
    regionZ := data[1]
    region := p_World.RegionManager().GetRegion(regionX, regionZ)
    s_Logger.Info("Drawing region", regionX, regionZ)
    img := draw.RenderRegionTile(region)
    draw.Save(fmt.Sprintf("tiles/r%d.%d.png", regionX, regionZ), img)
  }
}
