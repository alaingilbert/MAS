package main


import (
  "mas/core"
  "mas/draw"
  "mas/license"
  "mas/logger"
  "mas/worker"
  "runtime"
  "time"
)


var s_Logger logger.Logger = logger.NewLogger(logger.INFO | logger.DEBUG)


func main() {
  runtime.GOMAXPROCS(4)
  s_Logger.Debug("Start")

  // Load settings
  // Load license
  license.Verify()
  // Create worker pool
  workerPool := worker.NewWorkerPool(4)
  // start webserver

  startTime := time.Now()
  world := core.NewWorld("/Users/agilbert/Desktop/minecraft/world")
  theme := core.LoadTheme("default")

  regionsCoordinates := world.RegionManager().RegionsCoordinates()
  for index, regionCoord := range regionsCoordinates {
    if index > 0 {
      //break
    }
    job := draw.NewJobRenderRegionTile(regionCoord[0], regionCoord[1], world, theme)
    workerPool.Do(job)
  }

  workerPool.Wait()

  s_Logger.Debug("End", time.Since(startTime))
}
