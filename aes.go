package main

import (
  "crypto/aes"
  "crypto/cipher"
  "crypto/rand"
  "encoding/base64"
  "encoding/hex"
  "fmt"
  "io"
  "io/ioutil"
)

func main() {
  key := []byte("pd$5fK40sL!S?p048sCXmQ9%Z*oPa&ey") // 32 bytes
  plaintext, _ := ioutil.ReadFile("license.xml")
  ciphertext := encrypt(key, plaintext)
  s := hex.EncodeToString(ciphertext)
  fmt.Println(s)
}

func encodeBase64(b []byte) string {
  return base64.StdEncoding.EncodeToString(b)
}

func decodeBase64(s string) []byte {
  data, err := base64.StdEncoding.DecodeString(s)
  if err != nil {
    panic(err)
  }
  return data
}

func encrypt(key, text []byte) []byte {
  block, err := aes.NewCipher(key)
  if err != nil {
    panic(err)
  }
  b := encodeBase64(text)
  ciphertext := make([]byte, aes.BlockSize+len(b))
  iv := ciphertext[:aes.BlockSize]
  if _, err := io.ReadFull(rand.Reader, iv); err != nil {
    panic(err)
  }
  cfb := cipher.NewCFBEncrypter(block, iv)
  cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
  return ciphertext
}

func decrypt(key, text []byte) string {
  block, err := aes.NewCipher(key)
  if err != nil {
    panic(err)
  }
  if len(text) < aes.BlockSize {
    panic("ciphertext too short")
  }
  iv := text[:aes.BlockSize]
  text = text[aes.BlockSize:]
  cfb := cipher.NewCFBDecrypter(block, iv)
  cfb.XORKeyStream(text, text)
  return string(decodeBase64(string(text)))
}
