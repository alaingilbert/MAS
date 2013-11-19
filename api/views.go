package api


import (
  "encoding/json"
  "mas/core"
  "net/http"
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
