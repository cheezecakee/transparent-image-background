package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var FileName string

type Templates struct {
  templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
  return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
  return &Templates {
    templates: template.Must(template.ParseGlob("ui/html/*.html")),
  }
}

func index(c echo.Context) error {
  return c.Render(http.StatusOK, "index", nil)
}

func upload(c echo.Context) error {
  file, err := c.FormFile("file")
  if err != nil{
    return err
  }

  src, err := file.Open()
  if err != nil{
    return err
  }
  defer src.Close()

  FileName = file.Filename

  dstPath := "./internal/img/" + FileName
  dst, err := os.Create(dstPath)
  if err != nil{
    return nil
  }
  defer dst.Close()
  
  if _, err = io.Copy(dst, src); err != nil{
    return err
  }

  log.Printf("[ File %s uploaded successfully. ]", file.Filename)
  
  return c.Render(http.StatusOK, "download", nil)
}

// Add something to preview the image
func process(c echo.Context) error {
  return nil
}

func download(c echo.Context) error{
  path := "./internal/img/" + FileName
  return c.Attachment(path, FileName)
}

func main() {
  e := echo.New()
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  e.Renderer = newTemplate()

  e.GET("/", index)
  e.POST("/upload", upload)
  e.GET("/download", download)

  e.Logger.Fatal(e.Start(":42069"))
}

