package main


import (
  "runtime"
  "strings"
  "fmt"
  "io"
  "log"
  "mas/core"
  "mas/draw"
  "net/http"
  "strconv"
)


const PORT int = 8000


func HomeHandler(w http.ResponseWriter, req *http.Request) {
  fmt.Println(req.URL.Query()["x"])
  io.WriteString(w, "Hello\n")
}


func main() {
  log.Println("Start")
  runtime.GOMAXPROCS(4)

  world := core.NewWorld("/Users/agilbert/Desktop/minecraft/world")

  nbThread := 4
  in := make(chan []int)

  for i := 0; i < nbThread; i++ {
    go Worker(i, world, in)
  }

  for index, fileName := range world.RegionManager().RegionFileNames() {
    if index > 10 {
      break
    }
    if !strings.HasSuffix(fileName, "mca") {
      continue
    }
    splits := strings.SplitN(fileName, ".", 4)
    regionX, _ := strconv.Atoi(splits[1])
    regionZ, _ := strconv.Atoi(splits[2])

    data := make([]int, 2)
    data[0] = regionX
    data[1] = regionZ
    in <- data
  }

  close(in)


  log.Println("End")
  //http.HandleFunc("/", HomeHandler)
  //http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}


func Worker(p_Id int, p_World *core.World, in chan []int) {
  for data := range in {
    regionX := data[0]
    regionZ := data[1]
    region := p_World.RegionManager().GetRegion(regionX, regionZ)
    fmt.Println("Region", regionX, regionZ)
    img := draw.RenderRegionTile(region)
    draw.Save(fmt.Sprintf("tiles/r%d.%d.png", regionX, regionZ), img)
  }
}
