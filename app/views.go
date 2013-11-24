package app

import (
  "fmt"
  "github.com/codegangsta/martini"
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

// ThemeHandler ...
func ThemeHandler(res http.ResponseWriter, req *http.Request) {
  context := map[string]string{}
  tmpl, _ := template.ParseFiles("public/templates/theme.html")
  tmpl.Execute(res, context)
}

// TileHandler ...
func TileHandler(w http.ResponseWriter, req *http.Request,
  params martini.Params, pWorld *core.World,
  pTheme *core.Theme) {
  x, _ := strconv.Atoi(params["x"])
  y, _ := strconv.Atoi(params["y"])
  z, _ := strconv.Atoi(params["z"])
  path := fmt.Sprintf("tiles/%d/%d/", z, x)
  fileName := fmt.Sprintf("%d.png", y)
  file, err := os.Open(path + fileName)
  // No image found. Try to render it.
  if err != nil {
    if pWorld.RegionManager().GetRegionFromXYZ(x, y, z).Exists() {
      img := draw.RenderTile(x, y, z, pWorld, pTheme)
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
  pSettings *core.Settings, pTheme *core.Theme) {
  os.RemoveAll("./tiles/")
  pTheme.Reload()
}
