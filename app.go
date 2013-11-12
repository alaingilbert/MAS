package main


import (
  "mas/core"
  "mas/draw"
  "mas/license"
  "mas/logger"
  "mas/worker"
  "os"
  "runtime"
  "time"
)


var s_Logger logger.Logger = logger.NewLogger(logger.INFO | logger.DEBUG)


func main() {
  runtime.GOMAXPROCS(4)
  s_Logger.Debug("Start")

  // Load settings
  // Load license
  isLicenseValid := license.Verify()
  s_Logger.Debug("License valide:", isLicenseValid)
  if !isLicenseValid {
    s_Logger.Error("License expired.")
    os.Exit(0)
  }
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
