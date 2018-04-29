package main

import (
	"github.com/labstack/echo"
	"net/http"
)

type Image struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func main() {
	e := echo.New()

	e.GET("/images", func(c echo.Context) error {
		return c.JSON(http.StatusOK, getImages())
	})

	e.Logger.Fatal(e.Start(":8080"))
}

func getImages() *Image {
	i := &Image{
		Name: "hoge",
		URL: "https://s3-ap-northeast-1.amazonaws.com/public-test-20180429/sample.jpg",
	}
	return i
}