package handler

import (
	"bytes"
	"image/jpeg"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nfnt/resize"
)

func (h *Handler) resize(c echo.Context) error {
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
