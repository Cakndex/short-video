package main

import (
	"compress/gzip"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	g.Use(gzip.Gzip(gzip.DefaultCompression))

}
