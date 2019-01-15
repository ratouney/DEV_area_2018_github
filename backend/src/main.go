package main

import (
	"yolopack"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	yolopack.Logique()
}
