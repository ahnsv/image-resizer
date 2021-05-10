package main

import (
	"bytes"
	"image/jpeg"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/nfnt/resize"
)

func resizeImage(c echo.Context) error {
	c.Logger().Info("resize")
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	image, err := file.Open()
	if err != nil {
		return err
	}
	defer image.Close()

	img, err := jpeg.Decode(image)
	if err != nil {
		log.Fatal(err)
	}

	m := resize.Resize(120, 120, img, resize.Lanczos3)

	buf := new(bytes.Buffer)
	jpeg.Encode(buf, m, nil)

	return c.Blob(http.StatusOK, "image/jpeg", buf.Bytes())
}

func resizeImageV2(c echo.Context) error {
	imageSrc := c.QueryParam("imageSrc")
	width, err := strconv.ParseUint(c.QueryParam("width"), 10, 64)
	height, err := strconv.ParseUint(c.QueryParam("height"), 10, 64)
	c.Logger().Infof("[%s] is going to be resized to [%d, %d]", imageSrc, width, height)

	res, err := http.Get(imageSrc)
	if err != nil {
		c.Logger().Errorf("[%s] is not valid imageSrc", imageSrc)
		return err
	}
	img, err := jpeg.Decode(res.Body)
	if err != nil {
		return err
	}
	m := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)
	buf := new(bytes.Buffer)
	jpeg.Encode(buf, m, nil)

	return c.Blob(http.StatusOK, "image/jpeg", buf.Bytes())
}

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.POST("/v1/api/images/resize", resizeImage)
	e.GET("/v2/api/images/resize", resizeImageV2)
	e.GET("/v1/api/healthcheckz", func(c echo.Context) error {
		c.Logger().Info("health check")
		return c.String(http.StatusOK, "healthy")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
