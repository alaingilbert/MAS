package main

import (
	"io/ioutil"
	"bytes"
	"encoding/binary"
  "fmt"
  "github.com/codegangsta/martini"
  "github.com/bearbin/mcgorcon"
  "github.com/codegangsta/martini-contrib/sessions"
  "mas/api"
  "mas/app"
  "mas/app/middleware"
  "mas/core"
  "mas/license"
  "net/http"
  "runtime"
	"net"
)

func main() {
  numCPU := runtime.NumCPU()
  runtime.GOMAXPROCS(numCPU)

  // Load license
  go license.LicenseVerifier()
  license.PrintLicenseInfos()

  // Load minecraft client
	mcgorcon.Dial("localhost", 25565, "test")
	conn, err := net.Dial("tcp", "localhost:25565")
	if err != nil {
		fmt.Println("ERR")
	}
	var buf bytes.Buffer
	username := [64]byte{}
	key := [64]byte{}
	for i := 0; i < 64; i++ {
		username[i] = 0x20
		key[i] = 0x20
	}
	copy(username[:], []byte("alaingilbert"))
	copy(key[:], []byte("verifkey"))
	binary.Write(&buf, binary.LittleEndian, []byte{0x00})
	binary.Write(&buf, binary.LittleEndian, []byte{0x07})
	binary.Write(&buf, binary.LittleEndian, username)
	binary.Write(&buf, binary.LittleEndian, key)
	binary.Write(&buf, binary.LittleEndian, []byte{0x00})
	fmt.Println(buf.Bytes())
	conn.Write(buf.Bytes())

	result, err := ioutil.ReadAll(conn)
	//var tmp = make([]byte, 1024)
	//conn.Read(tmp)
	fmt.Println(result, err)

  // Load settings
  settings, _ := core.LoadSettings()

  // Load theme
  theme := core.NewTheme(settings.Theme)

  world := core.NewWorld(settings.WorldPath)

  // start webserver
  m := martini.Classic()
  m.Map(world)
  m.Map(settings)
  m.Map(theme)
  m.Use(martini.Static("public/static"))

  store := sessions.NewCookieStore([]byte("di$SjdCs@abZ7#y26K3t"))
  m.Use(sessions.Sessions("my_session", store))

  m.Use(middleware.LicenseMiddleware)
  m.Get("/", app.HomeHandler)
  m.Get("/tile/:z/:x/:y.png", app.TileHandler)
  m.Get("/license/", app.LicenseHandler)
  m.Get("/theme/", app.ThemeHandler)
  m.Post("/theme/", app.ThemeHandler)
  m.Get("/settings/", app.SettingsHandler)
  m.Post("/settings/", app.SettingsPostHandler)
  m.Get("/api/players/", api.PlayersHandler)
  m.Get("/api/players/icon/:name.png", api.PlayerIconHandler)
  m.Get("/renewtiles/", app.RenewTilesHandler)

  fmt.Println(fmt.Sprintf("Start listening on %s:%d", settings.WebServer.Host, settings.WebServer.Port))
  http.ListenAndServe(fmt.Sprintf("%s:%d", settings.WebServer.Host, settings.WebServer.Port), m)
}
