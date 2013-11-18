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
  "io/ioutil"
  "os"
  "runtime"
  "time"
  "strconv"
  "html/template"
  "mas/draw"
  "encoding/json"
  "encoding/xml"
)


type Settings struct {
  Theme string
  WorldPath string
  NbtVersion string
  WebServer WebServer
}

type WebServer struct {
  Host string
  Port int
}


var s_Logger logger.Logger = logger.NewLogger(logger.INFO | logger.DEBUG)

var m_World *core.World
var m_Settings Settings
var theme map[byte]core.Block = core.LoadTheme("default")
var m_LicenseValid bool = false

func TileHandler(w http.ResponseWriter, req *http.Request, params martini.Params) {
  x, _ := strconv.Atoi(params["x"])
  y, _ := strconv.Atoi(params["y"])
  z, _ := strconv.Atoi(params["z"])
  path := fmt.Sprintf("tiles/%d/%d/", z, x)
  fileName := fmt.Sprintf("%d.png", y)
  file, err := os.Open(path + fileName)
  if err != nil {
    img := draw.RenderTile(x, y, z, m_World, theme)
    png.Encode(w, img)
    draw.Save(path, fileName, img)
    //http.NotFound(w, req)
    return
  }
  defer file.Close()
  w.Header().Set("Content-type", "image/png")
  img, err := png.Decode(file)
  png.Encode(w, img)
}


func HomeHandler(w http.ResponseWriter, req *http.Request) {
  tmpl, err := template.ParseFiles("templates/index.html")
  if err != nil {
    fmt.Println(err)
  }
  tmpl.Execute(w, map[string] string {"title": "Test title"})
}


func LicenseHandler(w http.ResponseWriter, req *http.Request) {
  tmpl, err := template.ParseFiles("templates/license.html")
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
    tmpl, _ := template.ParseFiles("templates/license.html")
    tmpl.Execute(res, context)
  }
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


func Server() {
  s_Logger.Debug("Start web server")
  m := martini.Classic()
  m.Use(martini.Static("static"))
  m.Use(LicenseMiddleware)
  m.Get("/", HomeHandler)
  m.Get("/tile/:z/:x/:y.png", TileHandler)
  m.Get("/license/", LicenseHandler)
  m.Get("/api/players/", ApiPlayersHandler)
  http.ListenAndServe(fmt.Sprintf(":%d", m_Settings.WebServer.Port) , m)
}


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
  m_LicenseValid = isLicenseValid
  license.PrintLicenseInfos()

  s_Logger.Debug("Start")

  settingsFile, _ := ioutil.ReadFile("settings.xml")
  var settings Settings
  xml.Unmarshal(settingsFile, &settings)
  m_Settings = settings
  s_Logger.Debug(settings.WorldPath)
  m_World = core.NewWorld(settings.WorldPath)

  player := m_World.PlayerManager().GetPlayer("alaingilbert")
  s_Logger.Debug(player.X(), player.Y(), player.Z())

  // Create worker pool
  //workerPool := worker.NewWorkerPool(numCPU)

  // start webserver
  Server()

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

  //workerPool.Wait()

  //s_Logger.Debug("End", time.Since(startTime))
}
