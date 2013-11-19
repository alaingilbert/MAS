package main


import (
  "github.com/codegangsta/martini"
  "fmt"
  "image/png"
  "net/http"
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
  "encoding/json"
)


var s_Logger logger.Logger = logger.NewLogger(logger.INFO | logger.DEBUG)

var m_World *core.World
var m_Settings *core.Settings
var m_Theme *core.Theme = core.LoadTheme("default")
var m_LicenseValid bool = false
var m_WorldPathValid = true

func TileHandler(w http.ResponseWriter, req *http.Request, params martini.Params) {
  x, _ := strconv.Atoi(params["x"])
  y, _ := strconv.Atoi(params["y"])
  z, _ := strconv.Atoi(params["z"])
  path := fmt.Sprintf("tiles/%d/%d/", z, x)
  fileName := fmt.Sprintf("%d.png", y)
  file, err := os.Open(path + fileName)
  if err != nil {
    img := draw.RenderTile(x, y, z, m_World, m_Theme)
    if img == nil {
      http.NotFound(w, req)
      return
    }
    png.Encode(w, img)
    draw.Save(path, fileName, img)
    return
  }
  defer file.Close()
  w.Header().Set("Content-type", "image/png")
  img, err := png.Decode(file)
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


func LicenseMiddleware(res http.ResponseWriter, req *http.Request) {
  if !m_WorldPathValid {
    io.WriteString(res, "World path invalid. Change your settings.xml")
    return
  }
  if !m_LicenseValid {
    if req.Method == "POST" {
      lic := req.PostFormValue("license")
      file, err := os.Create("license.key")
      if err != nil {
        fmt.Println(err)
      }
      defer file.Close()
      io.WriteString(file, lic)
      m_LicenseValid = license.Verify()
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


func ApiPlayersHandler(res http.ResponseWriter, req *http.Request) {
  players := m_World.PlayerManager().GetPlayers()
  var playersJson []core.PlayerJson
  for _, player := range players {
    playerJson := player.ToJson()
    playersJson = append(playersJson, playerJson)
  }
  b, _ := json.Marshal(playersJson)
  res.Write(b)
}


func RenewTilesHandler(res http.ResponseWriter, req *http.Request) {
  os.RemoveAll("./tiles/")
  m_Theme = core.LoadTheme(m_Settings.Theme)
}


func Server() {
  s_Logger.Debug("Start web server")
  m := martini.Classic()
  m.Use(martini.Static("public/static"))
  m.Use(LicenseMiddleware)
  m.Get("/", HomeHandler)
  m.Get("/tile/:z/:x/:y.png", TileHandler)
  m.Get("/license/", LicenseHandler)
  m.Get("/theme/", ThemeHandler)
  m.Get("/api/players/", ApiPlayersHandler)
  m.Get("/renewtiles/", RenewTilesHandler)
  http.ListenAndServe(fmt.Sprintf(":%d", m_Settings.WebServer.Port) , m)
}


func LicenseVerifier() {
  c := time.Tick(1 * time.Hour)
  for now := range c {
    isLicenseValid := license.Verify()
    if !isLicenseValid {
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
  isLicenseValid := license.Verify()
  m_LicenseValid = isLicenseValid
  license.PrintLicenseInfos()

  // Load settings
  settings, err := core.LoadSettings()
  if err != nil {
    fmt.Println(err)
  }
  m_Settings = settings

  _, err = os.Stat(settings.WorldPath)
  if err != nil {
    m_WorldPathValid = false
  }

  m_World = core.NewWorld(settings.WorldPath)

  // start webserver
  Server()
}
