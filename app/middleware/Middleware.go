package middleware

import (
  "fmt"
  "html/template"
  "io"
  "mas/core"
  "mas/license"
  "net/http"
  "os"
)

// LicenseMiddleware ...
func LicenseMiddleware(res http.ResponseWriter, req *http.Request, pWorld *core.World) {
  if !pWorld.PathValid() {
    io.WriteString(res, "World path invalid. Change your settings.xml")
    return
  }
  if !license.IsValid {
    if req.Method == "POST" {
      lic := req.PostFormValue("license")
      file, err := os.Create("license.key")
      if err != nil {
        fmt.Println(err)
      }
      defer file.Close()
      io.WriteString(file, lic)
      license.Verify()
      http.Redirect(res, req, "/", 302)
      return
    }

    licInfos, err := license.Infos()
    context := map[string]string{
      "title":            "Invalid License",
      "licenseCreated":   licInfos["Created"],
      "licenseExpired":   licInfos["Expired"],
      "licenseFirstName": licInfos["FirstName"],
      "licenseLastName":  licInfos["LastName"],
    }
    if err != nil {
      context["licenseErr"] = err.Error()
    }
    tmpl, _ := template.ParseFiles("public/templates/license.html")
    tmpl.Execute(res, context)
  }
}
