package main


import (
  "os"
  "io"
  "fmt"
  "net/http"
  "image/png"
)


const PORT int = 8000


func TileHandler(w http.ResponseWriter, req *http.Request) {
  x := req.URL.Query()["x"][0]
  y := req.URL.Query()["y"][0]
  z := req.URL.Query()["z"][0]
  fileName := fmt.Sprintf("tiles/r%s.%s.png", x ,z)
  fmt.Println("Serve:", x, y, z, fileName)
  file, err := os.Open(fileName)
  if err != nil {
    io.WriteString(w, "FileNotFound")
    return
  }
  defer file.Close()
  w.Header().Set("Content-type", "image/png")
  img, err := png.Decode(file)
  png.Encode(w, img)
}


func HomeHandler(w http.ResponseWriter, req *http.Request) {
  io.WriteString(w, "Hello\n")
}


func main() {
  http.HandleFunc("/tile/", TileHandler)
  http.HandleFunc("/", HomeHandler)
  http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}
