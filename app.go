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
  "runtime"
)


var s_Logger logger.Logger = logger.NewLogger(logger.INFO | logger.DEBUG)


func main() {
  numCPU := runtime.NumCPU()
  runtime.GOMAXPROCS(numCPU)

  // Load license
  license.Verify()
  go license.LicenseVerifier()
  license.PrintLicenseInfos()

  // Load settings
  settings, _ := core.LoadSettings()

  // Load theme
  theme := core.NewTheme(settings.Theme)

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
  m.Get("/api/players/icon/:name.png", api.PlayerIconHandler)
  m.Get("/renewtiles/", app.RenewTilesHandler)
  http.ListenAndServe(fmt.Sprintf("%s:%d", settings.WebServer.Host, settings.WebServer.Port) , m)
}
