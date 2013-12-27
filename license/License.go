package license

import (
  "encoding/hex"
  "encoding/xml"
  "errors"
  "fmt"
  "io/ioutil"
  "log"
  "mas/crypto"
  "strings"
  "time"
)

// License ...
type License struct {
  Created string
  Expired string
  Owner   Owner
}

// Owner ...
type Owner struct {
  FirstName string
  LastName  string
}

// PrintLicenseInfos ...
func PrintLicenseInfos() {
  license, _ := _DecryptLicense()
  licenseBytes := []byte(license)
  var lic License
  xml.Unmarshal(licenseBytes, &lic)
  expireDate, _ := time.Parse("2006-01-02 15:04", lic.Expired)
  fmt.Println("--------------------------------------------------")
  fmt.Println("--- LICENSE INFORMATIONS")
  fmt.Println("--- OWNER: " + lic.Owner.FirstName + " " + lic.Owner.LastName)
  fmt.Println("--- CREATION DATE: " + lic.Created)
  fmt.Println("--- EXPIRATION DATE: " + lic.Expired)
  fmt.Printf("--- LICENSE VALID: %t\n", expireDate.Sub(time.Now()) > 0)
  fmt.Println("--------------------------------------------------")
}

// IsValid ...
var IsValid = false

// LicenseVerifier ...
func LicenseVerifier() {
  Verify()
  c := time.Tick(1 * time.Hour)
  for _ = range c {
    Verify()
  }
}

// Verify will tell you if the license file is valid and not expired.
func Verify() bool {
  license, err := _DecryptLicense()
  if err != nil {
    log.Println(err)
    IsValid = false
    return false
  }
  licenseBytes := []byte(license)
  var lic License
  xml.Unmarshal(licenseBytes, &lic)
  expireDate, _ := time.Parse("2006-01-02 15:04", lic.Expired)
  isValid := expireDate.Sub(time.Now()) > 0
  IsValid = isValid
  return isValid
}

// Infos ...
func Infos() (map[string]string, error) {
  license, err := _DecryptLicense()
  if err != nil {
    return nil, errors.New("license file invalid")
  }
  licenseBytes := []byte(license)
  var lic License
  xml.Unmarshal(licenseBytes, &lic)
  res := make(map[string]string)
  res["Created"] = lic.Created
  res["Expired"] = lic.Expired
  res["FirstName"] = lic.Owner.FirstName
  res["LastName"] = lic.Owner.LastName
  return res, nil
}

// _DecryptLicense decrypt the license.key file.
// It returns the license xml string.
func _DecryptLicense() (string, error) {
  file, err := ioutil.ReadFile("license.key")
  if err != nil {
    return "", err
  }
  fileStr := strings.TrimSpace(string(file))
  fileStr = strings.Replace(fileStr, "\n", "", -1)
  d, _ := hex.DecodeString(fileStr)
  key := []byte("pd$5fK40sL!S?p048sCXmQ9%Z*oPa&ey")
  license, err := crypto.Decrypt(key, d)
  if err != nil {
    return "", err
  }
  return license, nil
}
