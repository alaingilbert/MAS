package license


import (
  "encoding/hex"
  "encoding/xml"
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


func Verify() bool {
  license := []byte(_DecryptLicense())
  var lic License
  xml.Unmarshal(license, &lic)
  expireDate, _ := time.Parse("2006-01-02 15:04", lic.Expired)
  isValid := expireDate.Sub(time.Now()) > 0
  return isValid
}


func _DecryptLicense() string {
  file, _ := ioutil.ReadFile("license.key")
  fileStr := strings.TrimSpace(string(file))
  d, _ := hex.DecodeString(fileStr)
  key := []byte("pd$5fK40sL!S?p048sCXmQ9%Z*oPa&ey")
  license := crypto.Decrypt(key, d)
  return license
}
