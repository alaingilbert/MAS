package main


import (
  "fmt"
  "io"
  "net/http"
  "log"
  "os"
  "compress/zlib"
  "bytes"
)


const PORT int = 8000


func GetRegion(p_ChunkX, p_ChunkZ int) (regionX, regionZ int) {
  regionX = p_ChunkX >> 5
  regionZ = p_ChunkZ >> 5
  return
}


func GetChunkColor(p_ChunkX, p_ChunkZ int) string {
  regionX, regionZ := GetRegion(p_ChunkX, p_ChunkZ)
  data := make([]byte, 100)
  file, err := os.Open(fmt.Sprintf("/Users/agilbert/Desktop/minecraft/world/region/r.%d.%d.mca", regionX, regionZ))
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  count, err := file.Read(data)
  fmt.Println(count)
  fmt.Println(data)
  if err != nil {
    log.Fatal(err)
  }

  chunkLocation := 4 * ((p_ChunkX % 32) + (p_ChunkZ % 32) * 32)
  fmt.Println(chunkLocation)

  file.Seek(8 * 1024, 0)
  data = make([]byte, 4)
  file.Read(data)
  fmt.Println(data)
  return "blue"
}


func ReadNbt(p_FileName string) {
  file, err := os.Open(p_FileName)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  for i := 0; i < 4096; i += 4 {
    data := make([]byte, 4)
    file.Seek(int64(i), 0)
    file.Read(data)
    var offset int64
    for j:=0; j<3; j++ { offset = offset << 8 + int64(data[j]) }
    var length int64
    for j:=3; j<4; j++ { length = length << 8 + int64(data[j]) }
    if offset > 0 {
      file.Seek(offset * 4096, 0)
      lengthBytesData := make([]byte, 4)
      file.Read(lengthBytesData)
      var lengthBytes int64
      for j:=0; j<4; j++ { lengthBytes = lengthBytes << 8 + int64(lengthBytesData[j]) }
      versionData := make([]byte, 1)
      file.Read(versionData)
      version := int64(versionData[0])

      if version == 2 {
        compress := make([]byte, lengthBytes - 1)
        file.Read(compress)


        var out bytes.Buffer
        r, err := zlib.NewReader(bytes.NewReader(compress))
        if err != nil {
          log.Fatal(err)
        }
        io.Copy(&out, r)
        defer r.Close()
        //fmt.Println(out.Bytes())
      }
    }
  }

}


func HomeHandler(w http.ResponseWriter, req *http.Request) {
  fmt.Println(req.URL.Query()["x"])
  io.WriteString(w, "Hello\n")
}


func main() {
  //color := GetChunkColor(70, 30)
  //fmt.Printf("%s\n", color)

  ReadNbt("/Users/agilbert/Desktop/minecraft/world/region/r.-3.-5.mca")

  log.Println(fmt.Sprintf("Start listening on port %d", PORT))
  //http.HandleFunc("/", HomeHandler)
  //http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}
