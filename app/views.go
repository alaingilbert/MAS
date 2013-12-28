package app

import (
  "encoding/xml"
  "fmt"
  "github.com/codegangsta/martini"
  "github.com/codegangsta/martini-contrib/sessions"
  "html/template"
  "image/png"
  "mas/core"
  "mas/draw"
  "mas/license"
  "net/http"
  "os"
  "strconv"
)

// HomeHandler ...
func HomeHandler(w http.ResponseWriter, req *http.Request) {
  tmpl, err := template.ParseFiles("public/templates/index.html")
  if err != nil {
    fmt.Println(err)
  }
  tmpl.Execute(w, map[string]string{"title": "Test title"})
}

// LicenseHandler ...
func LicenseHandler(w http.ResponseWriter, req *http.Request) {
  tmpl, err := template.ParseFiles("public/templates/license.html")
  if err != nil {
    fmt.Println(err)
  }
  licInfos, err := license.Infos()
  context := map[string]string{
    "title":            "License",
    "licenseCreated":   licInfos["Created"],
    "licenseExpired":   licInfos["Expired"],
    "licenseFirstName": licInfos["FirstName"],
    "licenseLastName":  licInfos["LastName"],
  }
  if err != nil {
    context["licenseErr"] = err.Error()
  }
  tmpl.Execute(w, context)
}

// SettingsPostHandler ...
func SettingsPostHandler(res http.ResponseWriter, req *http.Request, pSettings *core.Settings, pWorld *core.World, session sessions.Session) {
  err := req.ParseForm()
  if err != nil {
    fmt.Println(err)
  }

  settings := core.Settings{}
  webServer := core.WebServer{}

  settings.Theme = req.Form["Theme"][0]
  settings.NbtVersion = req.Form["NbtVersion"][0]
  settings.WorldPath = req.Form["WorldPath"][0]
  webServer.Host = req.Form["Host"][0]
  webServer.Port, err = strconv.Atoi(req.Form["Port"][0])
  if err != nil {
    fmt.Println(err)
    session.AddFlash("Invalid port.")
    http.Redirect(res, req, "/settings/", http.StatusFound)
    return
  }
  settings.WebServer = webServer

  bytes, err := xml.Marshal(settings)
  if err != nil {
    fmt.Println(err)
  }

  pSettings.Theme = settings.Theme
  pSettings.NbtVersion = settings.NbtVersion
  pSettings.WorldPath = settings.WorldPath
  pSettings.WebServer.Host = webServer.Host
  pSettings.WebServer.Port = webServer.Port
  pWorld.Path = settings.WorldPath

  file, err := os.Create("settings.xml")
  if err != nil {
    fmt.Println(err)
  }
  defer file.Close()
  file.Write(bytes)

  session.AddFlash("Settings saved")
  http.Redirect(res, req, "/settings/", http.StatusFound)
  return
}

// SettingsHandler ...
func SettingsHandler(res http.ResponseWriter, req *http.Request, pSettings *core.Settings, session sessions.Session) {
  err := req.ParseForm()
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(req.Form)
  context := map[string]interface{}{}
  context["Settings"] = pSettings
  context["Flash"] = session.Flashes()
  tmpl, _ := template.ParseFiles("public/templates/settings.html")
  tmpl.Execute(res, context)
}

// ThemeHandler ...
func ThemeHandler(res http.ResponseWriter, req *http.Request, theme *core.Theme) {
  if req.Method == "POST" {
    err := req.ParseForm()
    if err != nil {
      fmt.Println(err)
    }
    fmt.Println(req.PostForm)
    for id, value := range req.PostForm {
      fmt.Println(id, value)
    }
    // Should redirect...
    return
  }
  context := map[string]interface{}{}
  context["Theme"] = theme
  tmpl, _ := template.ParseFiles("public/templates/theme.html")
  tmpl.Execute(res, theme.GetMap())
}

// TileHandler ...
func TileHandler(w http.ResponseWriter, req *http.Request,
  params martini.Params, world *core.World,
  theme *core.Theme) {
  x, _ := strconv.Atoi(params["x"])
  y, _ := strconv.Atoi(params["y"])
  z, _ := strconv.Atoi(params["z"])
  path := fmt.Sprintf("tiles/%d/%d/", z, x)
  fileName := fmt.Sprintf("%d.png", y)
  file, err := os.Open(path + fileName)
  // No image found. Try to render it.
  if err != nil {
    if world.RegionManager().GetRegionFromXYZ(x, y, z).Exists() {
      img := draw.RenderTile(x, y, z, world, theme)
      png.Encode(w, img)
      draw.Save(path, fileName, img)
      return
    }
    http.NotFound(w, req)
    return
  }
  defer file.Close()
  w.Header().Set("Content-type", "image/png")
  img, _ := png.Decode(file)
  png.Encode(w, img)
}

// RenewTilesHandler ...
func RenewTilesHandler(res http.ResponseWriter, req *http.Request,
  pSettings *core.Settings, theme *core.Theme) {
  os.RemoveAll("./tiles/")
  theme.Reload()
}
