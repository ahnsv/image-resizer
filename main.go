package main

import (
	"image-resizer/handler"
	"image-resizer/router"
)

func main() {
	r := router.New()
	v1 := r.Group("/api")
	h := handler.NewHandler()
	h.Register(v1)
}

// func main() {
// 	e := echo.New()
// 	e.Logger.SetLevel(log.INFO)
// 	e.GET("/healthcheckz", func(c echo.Context) error {
// 		return c.String(http.StatusOK, "healthy")
// 	})
// 	e.POST("/api/images/resize", func(c echo.Context) error {
// 		c.Logger().Info("resize")
// 		file, err := c.FormFile("file")
// 		if err != nil {
// 			return err
// 		}

// 		image, err := file.Open()
// 		if err != nil {
// 			return err
// 		}
// 		defer image.Close()

// 		img, err := jpeg.Decode(image)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		m := resize.Resize(120, 120, img, resize.Lanczos3)

// 		buf := new(bytes.Buffer)
// 		jpeg.Encode(buf, m, nil)

// 		return c.Blob(http.StatusOK, "image/jpeg", buf.Bytes())
// 	})
// 	e.Logger.Fatal(e.Start(":1323"))
// }
