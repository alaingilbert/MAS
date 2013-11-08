package main


import (
  "fmt"
  "io"
  "log"
  "mas/Core"
  "net/http"
)


const PORT int = 8000

func HomeHandler(w http.ResponseWriter, req *http.Request) {
  fmt.Println(req.URL.Query()["x"])
  io.WriteString(w, "Hello\n")
}


func main() {
  log.Println("Start")

  world := Core.NewWorld("/Users/agilbert/Desktop/minecraft/world")
  region := world.RegionManager().GetRegion(-3, -5)
  fmt.Println(region.GetChunk(1, 1))

  log.Println("End")
  //http.HandleFunc("/", HomeHandler)
  //http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}
