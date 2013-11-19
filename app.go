package main


import (
  "fmt"
  "github.com/codegangsta/martini"
  "mas/api"
  "mas/app"
  "mas/app/middleware"
  "mas/core"
  "mas/license"
  "mas/logger"
  "net/http"
  "os"
  "runtime"
  "time"
)


var s_Logger logger.Logger = logger.NewLogger(logger.INFO | logger.DEBUG)


func LicenseVerifier() {
  c := time.Tick(1 * time.Hour)
  for now := range c {
    if !license.Verify() {
      s_Logger.Error("License expired.", now)
      os.Exit(0)
    }
  }
}


func main() {
  numCPU := runtime.NumCPU()
  runtime.GOMAXPROCS(numCPU)

  license.Verify()
  go LicenseVerifier()

  // Load license
  license.PrintLicenseInfos()

  // Load settings
  settings, err := core.LoadSettings()
  if err != nil {
    fmt.Println(err)
  }

  // Load theme
  theme := core.LoadTheme(settings.Theme)

  world := core.NewWorld(settings.WorldPath)

  // start webserver
  m := martini.Classic()
  m.Map(world)
  m.Map(settings)
  m.Map(theme)
  m.Use(martini.Static("public/static"))
  m.Use(middleware.LicenseMiddleware)
  m.Get("/", app.HomeHandler)
  m.Get("/tile/:z/:x/:y.png", app.TileHandler)
  m.Get("/license/", app.LicenseHandler)
  m.Get("/theme/", app.ThemeHandler)
  m.Get("/api/players/", api.PlayersHandler)
  m.Get("/renewtiles/", app.RenewTilesHandler)
  http.ListenAndServe(fmt.Sprintf(":%d", settings.WebServer.Port) , m)
}
