package api


import (
  "encoding/json"
  "fmt"
  "github.com/codegangsta/martini"
  "github.com/nfnt/resize"
  //"io"
  "image"
  "image/draw"
  "image/png"
  "mas/core"
  "net/http"
  "os"
)


func PlayersHandler(res http.ResponseWriter, req *http.Request, p_World *core.World) {
  players := p_World.PlayerManager().GetPlayers()
  var playersJson []core.PlayerJson
  for _, player := range players {
    playerJson := player.ToJson()
    playersJson = append(playersJson, playerJson)
  }
  b, _ := json.Marshal(playersJson)
  res.Write(b)
}


func PlayerIconHandler(w http.ResponseWriter, req *http.Request, params martini.Params) {
  name := params["name"]
  fmt.Println(name)

  resp, err := http.Get("http://s3.amazonaws.com/MinecraftSkins/" + name + ".png")
  if err != nil {
    fmt.Println(err)
  }

  scale := 4
  is := 8 * scale
  var img image.Image

  if resp.StatusCode != 200 {
    file, err := os.Open("./public/skins/default.png")
    if err != nil {
      fmt.Println(err)
    }
    defer file.Close()
    img, _ = png.Decode(file)
  } else {
    img, _ = png.Decode(resp.Body)
    resp.Body.Close()
  }

  w.Header().Set("Content-type", "image/png")
  img = resize.Resize(uint(64 * scale), 0, img, resize.NearestNeighbor)
  dest := image.NewRGBA(image.Rect(0, 0, is, is))
  draw.Draw(dest, image.Rect(0, 0, is, is), img, image.Point{scale * 8, scale * 8}, draw.Src)
  png.Encode(w, dest)
}
