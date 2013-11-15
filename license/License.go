package license


import (
  "encoding/hex"
  "encoding/xml"
  "fmt"
  "mas/crypto"
  "mas/logger"
  "io/ioutil"
  "strings"
  "time"
)


var s_Logger logger.Logger = logger.NewLogger(logger.DEBUG)


type License struct {
  Created string
  Expired string
  Owner Owner
}
type Owner struct {
  FirstName string
  LastName string
}


func PrintLicenseInfos() {
  license := []byte(_DecryptLicense())
  var lic License
  xml.Unmarshal(license, &lic)
  expireDate, _ := time.Parse("2006-01-02 15:04", lic.Expired)
  fmt.Println("--------------------------------------------------")
  fmt.Println("--- LICENSE INFORMATIONS")
  fmt.Println("--- OWNER: " + lic.Owner.FirstName + " " + lic.Owner.LastName)
  fmt.Println("--- CREATION DATE: " + lic.Created)
  fmt.Println("--- EXPIRATION DATE: " + lic.Expired)
  fmt.Printf("--- LICENSE VALID: %t\n", expireDate.Sub(time.Now()) > 0)
  fmt.Println("--------------------------------------------------")
}


// Verify will tell you if the license file is valid and not expired.
func Verify() bool {
  license := []byte(_DecryptLicense())
  var lic License
  xml.Unmarshal(license, &lic)
  expireDate, _ := time.Parse("2006-01-02 15:04", lic.Expired)
  isValid := expireDate.Sub(time.Now()) > 0
  return isValid
}


// _DecryptLicense decrypt the license.key file.
// It returns the license xml string.
func _DecryptLicense() string {
  file, _ := ioutil.ReadFile("license.key")
  fileStr := strings.TrimSpace(string(file))
  fileStr = strings.Replace(fileStr, "\n", "", -1)
  d, _ := hex.DecodeString(fileStr)
  key := []byte("pd$5fK40sL!S?p048sCXmQ9%Z*oPa&ey")
  license := crypto.Decrypt(key, d)
  return license
}
