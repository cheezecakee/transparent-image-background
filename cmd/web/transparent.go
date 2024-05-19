package main

import (
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"log"
	"os"
)

func Decode(fileName string) image.Image{
  path := "./internal/img/" + fileName
  reader, err := os.Open(path)
  if err != nil {
    log.Fatal(err)
  }
  defer reader.Close()

  img, _, err := image.Decode(reader)
  if err != nil {
    log.Fatal(err)
  }

  return img
}

func Tensor(img image.Image) [][]color.Color {
  bounds := img.Bounds()
  width, height := bounds.Max.X, bounds.Max.Y

  // Create a 2D slice to store color data
  pixels := make([][]color.Color, width)
  for x := 0; x < width; x++{
    pixels[x] = make([]color.Color, height)
    for y := 0; y < height; y++{
      pixels[x][y] = img.At(x, y)
    }
  }

  return pixels
} 

func TransparentBackground(pixels [][]color.Color) [][]color.Color{
  for x := range pixels{
    for y := range pixels[x]{
      r, g, b, a := pixels[x][y].RGBA()
      if r == 65535 && g == 65535 && b == 65535 && a == 65535 {
        pixels[x][y] = color.RGBA{0, 0, 0, 0}
      }
    }
  }
  return pixels
}

func Convert(pixels [][]color.Color) image.Image {
  width := len(pixels)
  height := len(pixels[0])
  newImg := image.NewRGBA(image.Rect(0, 0, width, height))

  for x := range pixels{
    for y := range pixels[x]{
      newImg.Set(x, y, pixels[x][y])
    }
  }

  return newImg
}

func SaveImage(img image.Image, FileName string) error {
  file, err := os.Create(FileName)
  if err != nil{
    return err
  }
  defer file.Close()

  err = png.Encode(file, img)
  if err != nil{
    return err
  }

  return nil
}

