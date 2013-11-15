package main


import (
  "fmt"
  "image/png"
  "net/http"
  "mas/core"
  "mas/license"
  "mas/logger"
  "mas/worker"
  "os"
  "runtime"
  "time"
  "strconv"
  "html/template"
  "io/ioutil"
  "mas/draw"
)


var s_Logger logger.Logger = logger.NewLogger(logger.INFO | logger.DEBUG)

var world *core.World = core.NewWorld("/Users/agilbert/Desktop/minecraft/world")
var theme map[byte]core.Block = core.LoadTheme("default")


const PORT int = 8000


func TileHandler(w http.ResponseWriter, req *http.Request) {
  x, _ := strconv.Atoi(req.URL.Query()["x"][0])
  y, _ := strconv.Atoi(req.URL.Query()["y"][0])
  z, _ := strconv.Atoi(req.URL.Query()["z"][0])
  path := fmt.Sprintf("tiles/%d/%d/", z, x)
  fileName := fmt.Sprintf("%d.png", y)
  s_Logger.Debug("Serve tile x:", x, "y:", y, "z:", z)
  file, err := os.Open(path + fileName)
  if err != nil {
    img := draw.RenderTile(x, y, z, world, theme)
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


func LeafletJsHandler(w http.ResponseWriter, req *http.Request) {
  file, _ := ioutil.ReadFile("static/js/leaflet.js")
  w.Header().Set("Content-Type", "application/x-javascript")
  w.Write(file)
}

func LeafletCssHandler(w http.ResponseWriter, req *http.Request) {
  file, _ := ioutil.ReadFile("static/css/leaflet.css")
  w.Header().Set("Content-Type", "text/css")
  w.Write(file)
}

func LeafletZoomInHandler(w http.ResponseWriter, req *http.Request) {
  file, _ := ioutil.ReadFile("static/css/images/zoom-in.png")
  w.Header().Set("Content-Type", "image/png")
  w.Write(file)
}

func LeafletZoomOutHandler(w http.ResponseWriter, req *http.Request) {
  file, _ := ioutil.ReadFile("static/css/images/zoom-out.png")
  w.Header().Set("Content-Type", "image/png")
  w.Write(file)
}


func Server() {
  s_Logger.Debug("Start web server")
  http.HandleFunc("/tile/", TileHandler)
  http.HandleFunc("/static/css/leaflet.css", LeafletCssHandler)
  http.HandleFunc("/static/css/images/zoom-in.png", LeafletZoomInHandler)
  http.HandleFunc("/static/css/images/zoom-out.png", LeafletZoomOutHandler)
  http.HandleFunc("/static/js/leaflet.js", LeafletJsHandler)
  http.HandleFunc("/", HomeHandler)
  http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
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
  license.PrintLicenseInfos()
  if !isLicenseValid {
    s_Logger.Error("License expired.")
    os.Exit(0)
  }

  s_Logger.Debug("Start")

  // Create worker pool
  workerPool := worker.NewWorkerPool(numCPU)

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

  workerPool.Wait()

  //s_Logger.Debug("End", time.Since(startTime))
}
