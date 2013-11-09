package draw


import (
  "fmt"
  "image"
  "image/png"
  "image/color"
  "os"
)


func CreateImage(p_SizeX, p_SizeZ int) *image.RGBA {
  return image.NewRGBA(image.Rect(0, 0, p_SizeX, p_SizeZ))
}


func Save(p_FileName string, p_Img *image.RGBA) {
  file, err := os.Create(p_FileName)
  if err != nil {
    fmt.Print(err)
  }
  defer file.Close()
  png.Encode(file, p_Img)
}


func FillRect(p_Img *image.RGBA, p_X, p_Z, p_Width, p_Height int, p_Color color.Color) {
  for i := p_X; i < p_X + p_Width; i++ {
    for j := p_Z; j < p_Z + p_Height; j++ {
      p_Img.Set(i, j, p_Color)
    }
  }
}
