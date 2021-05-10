package main

import (
	"github.com/ahnsv/image-resizer/handler"
	"github.com/ahnsv/image-resizer/router"
)

func main() {
	r := router.New()
	v1 := r.Group("/api")
	h := handler.NewHandler()
	h.Register(v1)
}
