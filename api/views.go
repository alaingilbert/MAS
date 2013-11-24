package api

import (
  "encoding/json"
  "fmt"
  "github.com/codegangsta/martini"
  "github.com/nfnt/resize"
  "image"
  "image/draw"
  "image/png"
  "mas/core"
  "net/http"
  "os"
)

// PlayersHandler ...
func PlayersHandler(res http.ResponseWriter, req *http.Request, pWorld *core.World) {
  players := pWorld.PlayerManager().GetPlayers()
  var playersJSON []core.PlayerJSON
  for _, player := range players {
    playerJSON := player.ToJSON()
    playersJSON = append(playersJSON, playerJSON)
  }
  b, _ := json.Marshal(playersJSON)
  res.Write(b)
}

// PlayerIconHandler ...
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
  img = resize.Resize(uint(64*scale), 0, img, resize.NearestNeighbor)
  dest := image.NewRGBA(image.Rect(0, 0, is, is))
  draw.Draw(dest, image.Rect(0, 0, is, is), img, image.Point{scale * 8, scale * 8}, draw.Src)
  png.Encode(w, dest)
}
