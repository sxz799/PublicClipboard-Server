package main

import (
	"PublicClipboard/router"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	router.RegRouter(r)
	r.Run(":9090")
}
