package main


import (
  //"mas/core"
  //"mas/draw"
  "mas/license"
  "mas/logger"
  "mas/web"
  "mas/worker"
  "os"
  "runtime"
  "time"
)


var s_Logger logger.Logger = logger.NewLogger(logger.INFO | logger.DEBUG)


func main() {
  numCPU := runtime.NumCPU()
  runtime.GOMAXPROCS(numCPU)

  go func () {
    c := time.Tick(1 * time.Hour)
    for now := range c {
      isLicenseValid := license.Verify()
      if !isLicenseValid {
        s_Logger.Error("License expired.", now)
        os.Exit(0)
      }
    }
  }()

  // Load settings

  // Load license
  isLicenseValid := license.Verify()
  license.PrintLicenseInfos()
  if !isLicenseValid {
    s_Logger.Error("License expired.")
    os.Exit(0)
  }

  s_Logger.Debug("Start")

  // Create worker pool
  workerPool := worker.NewWorkerPool(numCPU)

  // start webserver
  web.Server()

  //startTime := time.Now()
  //world := core.NewWorld("/Users/agilbert/Desktop/minecraft/world")
  //theme := core.LoadTheme("default")

  //regionsCoordinates := world.RegionManager().RegionsCoordinates()
  //for index, regionCoord := range regionsCoordinates {
  //  if index > 0 {
  //    //break
  //  }
  //  job := draw.NewJobRenderRegionTile(regionCoord[0], regionCoord[1], world, theme)
  //  workerPool.Do(job)
  //}

  workerPool.Wait()

  //s_Logger.Debug("End", time.Since(startTime))
}
