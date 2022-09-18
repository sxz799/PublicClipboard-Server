package main

import (
	"PublicClipboard/router"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	router.RegRouter(r)
	r.Run(":9090")
}
