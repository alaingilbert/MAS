package main


import (
  "fmt"
  "mas/core"
  "mas/draw"
  "mas/logger"
  "runtime"
  "strconv"
  "strings"
  "sync"
  "time"
)


var s_Logger logger.Logger = logger.NewLogger(logger.INFO | logger.DEBUG)


func main() {

  s_Logger.Debug("Start")
  startTime := time.Now()
  runtime.GOMAXPROCS(4)

  world := core.NewWorld("/Users/agilbert/Desktop/minecraft/world")

  nbThread := 4
  in := make(chan []int)
  waitGroup := new(sync.WaitGroup)

  for i := 0; i < nbThread; i++ {
    waitGroup.Add(1)
    go Worker(i, world, in, waitGroup)
  }

  files := world.RegionManager().RegionFileNames()
  files[0] = "r.-1.1.mca"
  for index, fileName := range files {
    if index > 0 {
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
  waitGroup.Wait()

  s_Logger.Debug("End", time.Since(startTime))
}


func Worker(p_Id int, p_World *core.World, p_In chan []int, p_WaitGroup *sync.WaitGroup) {
  defer p_WaitGroup.Done()

  for data := range p_In {
    regionX := data[0]
    regionZ := data[1]
    s_Logger.Debug("Start drawing region", regionX, regionZ)
    region := p_World.RegionManager().GetRegion(regionX, regionZ)
    img := draw.RenderRegionTile(region)
    draw.Save(fmt.Sprintf("tiles/r%d.%d.png", regionX, regionZ), img)
    s_Logger.Debug("End drawing region", regionX, regionZ)
  }
}
