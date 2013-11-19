package main


import (
  "github.com/codegangsta/martini"
  "fmt"
  "image/png"
  "net/http"
  "mas/api"
  "mas/core"
  "mas/license"
  "mas/logger"
  //"mas/worker"
  "io"
  "os"
  "runtime"
  "time"
  "strconv"
  "html/template"
  "mas/draw"
)


var s_Logger logger.Logger = logger.NewLogger(logger.INFO | logger.DEBUG)


func TileHandler(w http.ResponseWriter, req *http.Request,
                 params martini.Params, p_World *core.World,
                 p_Theme *core.Theme) {
  x, _ := strconv.Atoi(params["x"])
  y, _ := strconv.Atoi(params["y"])
  z, _ := strconv.Atoi(params["z"])
  path := fmt.Sprintf("tiles/%d/%d/", z, x)
  fileName := fmt.Sprintf("%d.png", y)
  file, err := os.Open(path + fileName)
  // No image found. Try to render it.
  if err != nil {
    if p_World.RegionManager().GetRegionFromXYZ(x, y, z).Exists() {
      img := draw.RenderTile(x, y, z, p_World, p_Theme)
      png.Encode(w, img)
      draw.Save(path, fileName, img)
      return
     } else {
      http.NotFound(w, req)
      return
    }
  }
  defer file.Close()
  w.Header().Set("Content-type", "image/png")
  img, _ := png.Decode(file)
  png.Encode(w, img)
}


func HomeHandler(w http.ResponseWriter, req *http.Request) {
  tmpl, err := template.ParseFiles("public/templates/index.html")
  if err != nil {
    fmt.Println(err)
  }
  tmpl.Execute(w, map[string] string {"title": "Test title"})
}


func LicenseHandler(w http.ResponseWriter, req *http.Request) {
  tmpl, err := template.ParseFiles("public/templates/license.html")
  if err != nil {
    fmt.Println(err)
  }
  licInfos, err := license.Infos()
  context := map[string] string {
    "title": "License",
    "licenseCreated": licInfos["Created"],
    "licenseExpired": licInfos["Expired"],
    "licenseFirstName": licInfos["FirstName"],
    "licenseLastName": licInfos["LastName"],
  }
  if err != nil {
    context["licenseErr"] = err.Error()
  }
  tmpl.Execute(w, context)
}


func LicenseMiddleware(res http.ResponseWriter, req *http.Request, p_World *core.World) {
  if !p_World.PathValid() {
    io.WriteString(res, "World path invalid. Change your settings.xml")
    return
  }
  if !license.IsValid {
    if req.Method == "POST" {
      lic := req.PostFormValue("license")
      file, err := os.Create("license.key")
      if err != nil {
        fmt.Println(err)
      }
      defer file.Close()
      io.WriteString(file, lic)
      license.Verify()
      http.Redirect(res, req, "/", 302)
      return
    }

    licInfos, err := license.Infos()
    context := map[string] string {
      "title": "Invalid License",
      "licenseCreated": licInfos["Created"],
      "licenseExpired": licInfos["Expired"],
      "licenseFirstName": licInfos["FirstName"],
      "licenseLastName": licInfos["LastName"],
    }
    if err != nil {
      context["licenseErr"] = err.Error()
    }
    tmpl, _ := template.ParseFiles("public/templates/license.html")
    tmpl.Execute(res, context)
  }
}


func ThemeHandler(res http.ResponseWriter, req *http.Request) {
  context := map[string] string {
  }
  tmpl, _ := template.ParseFiles("public/templates/theme.html")
  tmpl.Execute(res, context)
}


func RenewTilesHandler(res http.ResponseWriter, req *http.Request,
                       p_Settings *core.Settings, p_Theme *core.Theme) {
  os.RemoveAll("./tiles/")
  p_Theme = core.LoadTheme(p_Settings.Theme)
}


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
  m.Use(LicenseMiddleware)
  m.Get("/", HomeHandler)
  m.Get("/tile/:z/:x/:y.png", TileHandler)
  m.Get("/license/", LicenseHandler)
  m.Get("/theme/", ThemeHandler)
  m.Get("/api/players/", api.PlayersHandler)
  m.Get("/renewtiles/", RenewTilesHandler)
  http.ListenAndServe(fmt.Sprintf(":%d", settings.WebServer.Port) , m)
}
