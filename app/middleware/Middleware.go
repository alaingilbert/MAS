package middleware

import (
  "fmt"
  "github.com/codegangsta/martini-contrib/sessions"
  "html/template"
  "io"
  "mas/core"
  "mas/license"
  "net/http"
  "os"
)

// LicenseMiddleware ...
func LicenseMiddleware(res http.ResponseWriter, req *http.Request, pWorld *core.World, session sessions.Session) {
  // World path invalid, change your settings.xml
  if !pWorld.PathValid() && req.URL.Path != "/settings/" {
    session.AddFlash("World path is invalid.")
    http.Redirect(res, req, "/settings/", http.StatusFound)
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
