package main

import (
	"PublicClipboard/gobalConfig"
	"PublicClipboard/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gobalConfig.GinMode)
	r := gin.Default()
	if gobalConfig.FrontMode {
		log.Println("已开启前后端整合模式！")
		r.LoadHTMLGlob("static/index.html")
		r.Static("/static", "static")
		r.GET("/", func(context *gin.Context) {
			context.HTML(200, "index.html", "")
		})
	}
	router.RegRouter(r)
	log.Println("路由注册完成！当前端口为:", gobalConfig.ServerPort)
	r.Run(":" + gobalConfig.ServerPort)
}
