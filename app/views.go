package app


import (
  "fmt"
  "github.com/codegangsta/martini"
  "html/template"
  "image/png"
  "mas/core"
  "mas/draw"
  "net/http"
  "mas/license"
  "os"
  "strconv"
)


func HomeHandler(w http.ResponseWriter, req *http.Request) {
  tmpl, err := template.ParseFiles("public/templates/index.html")
  if err != nil {
    fmt.Println(err)
  }
  tmpl.Execute(w, map[string] string {"title": "Test title"})
}


func LicenseHandler(w http.ResponseWriter, req *http.Request) {
  tmpl, err := template.ParseFiles("public/templates/license.html")
  if err != nil {
    fmt.Println(err)
  }
  licInfos, err := license.Infos()
  context := map[string] string {
    "title": "License",
    "licenseCreated": licInfos["Created"],
    "licenseExpired": licInfos["Expired"],
    "licenseFirstName": licInfos["FirstName"],
    "licenseLastName": licInfos["LastName"],
  }
  if err != nil {
    context["licenseErr"] = err.Error()
  }
  tmpl.Execute(w, context)
}


func ThemeHandler(res http.ResponseWriter, req *http.Request) {
  context := map[string] string {
  }
  tmpl, _ := template.ParseFiles("public/templates/theme.html")
  tmpl.Execute(res, context)
}


func TileHandler(w http.ResponseWriter, req *http.Request,
                 params martini.Params, p_World *core.World,
                 p_Theme *core.Theme) {
  x, _ := strconv.Atoi(params["x"])
  y, _ := strconv.Atoi(params["y"])
  z, _ := strconv.Atoi(params["z"])
  path := fmt.Sprintf("tiles/%d/%d/", z, x)
  fileName := fmt.Sprintf("%d.png", y)
  file, err := os.Open(path + fileName)
  // No image found. Try to render it.
  if err != nil {
    if p_World.RegionManager().GetRegionFromXYZ(x, y, z).Exists() {
      img := draw.RenderTile(x, y, z, p_World, p_Theme)
      png.Encode(w, img)
      draw.Save(path, fileName, img)
      return
     } else {
      http.NotFound(w, req)
      return
    }
  }
  defer file.Close()
  w.Header().Set("Content-type", "image/png")
  img, _ := png.Decode(file)
  png.Encode(w, img)
}


func RenewTilesHandler(res http.ResponseWriter, req *http.Request,
                       p_Settings *core.Settings, p_Theme *core.Theme) {
  os.RemoveAll("./tiles/")
  p_Theme = core.LoadTheme(p_Settings.Theme)
}
