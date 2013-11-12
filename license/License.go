package license


import (
  "encoding/hex"
  "mas/crypto"
  "mas/logger"
  "io/ioutil"
  "strings"
)


var s_Logger logger.Logger = logger.NewLogger(logger.DEBUG)


func Verify() {
  s_Logger.Debug("VerifyLicense")
  file, _ := ioutil.ReadFile("license.key")
  fileStr := strings.TrimSpace(string(file))
  d, _ := hex.DecodeString(fileStr)
  key := []byte("a very very very very secret key")
  license := crypto.Decrypt(key, d)
  s_Logger.Debug("CRISS", license)
}
