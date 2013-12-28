package main

import (
  "fmt"
  "github.com/codegangsta/martini"
  "github.com/codegangsta/martini-contrib/sessions"
  "mas/api"
  "mas/app"
  "mas/app/middleware"
  "mas/core"
  "mas/license"
  "net/http"
  "runtime"
)

func main() {
  numCPU := runtime.NumCPU()
  runtime.GOMAXPROCS(numCPU)

  // Load license
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

  store := sessions.NewCookieStore([]byte("di$SjdCs@abZ7#y26K3t"))
  m.Use(sessions.Sessions("my_session", store))

  m.Use(middleware.LicenseMiddleware)
  m.Get("/", app.HomeHandler)
  m.Get("/tile/:z/:x/:y.png", app.TileHandler)
  m.Get("/license/", app.LicenseHandler)
  m.Get("/theme/", app.ThemeHandler)
  m.Post("/theme/", app.ThemeHandler)
  m.Get("/settings/", app.SettingsHandler)
  m.Post("/settings/", app.SettingsPostHandler)
  m.Get("/api/players/", api.PlayersHandler)
  m.Get("/api/players/icon/:name.png", api.PlayerIconHandler)
  m.Get("/renewtiles/", app.RenewTilesHandler)

  fmt.Println(fmt.Sprintf("Start listening on %s:%d", settings.WebServer.Host, settings.WebServer.Port))
  http.ListenAndServe(fmt.Sprintf("%s:%d", settings.WebServer.Host, settings.WebServer.Port), m)
}
