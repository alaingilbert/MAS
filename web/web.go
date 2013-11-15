package web


import (
  "os"
  "io/ioutil"
  "fmt"
  "mas/core"
  "mas/logger"
  "mas/draw"
  "net/http"
  "image/png"
  "html/template"
  "strconv"
)

var world *core.World = core.NewWorld("/Users/agilbert/Desktop/minecraft/world")
var theme map[byte]core.Block = core.LoadTheme("default")

var s_Logger logger.Logger = logger.NewLogger(logger.DEBUG)


const PORT int = 8000


func TileHandler(w http.ResponseWriter, req *http.Request) {
  x, _ := strconv.Atoi(req.URL.Query()["x"][0])
  y, _ := strconv.Atoi(req.URL.Query()["y"][0])
  z, _ := strconv.Atoi(req.URL.Query()["z"][0])
  fileName := fmt.Sprintf("tiless/r.%s.%s.png", x ,z)
  s_Logger.Debug("Serve tile x:", x, "y:", y, "z:", z)
  file, err := os.Open(fileName)
  if err != nil {
    img := draw.RenderTile(x, y, z, world, theme)
    png.Encode(w, img)
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
