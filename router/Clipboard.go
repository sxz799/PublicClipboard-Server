package router

import (
	"PublicClipboard/model"
	"github.com/gin-gonic/gin"
	"time"
)

func upd(c *gin.Context) {
	json := make(map[string]string)
	c.BindJSON(&json)
	msg := json["content"]
	ok, err := model.UpdClipboard(msg)
	if ok {
		c.JSON(200, model.Result{Code: 200, Success: true, Message: "更新成功"})
		model.AddLog(model.Log{
			Msg:  msg,
			Date: time.Now().Format("2006-01-02 15:04:05"),
		})
	} else {
		c.JSON(200, model.Result{Success: false, Message: "更新失败" + err.Error()})
	}

}
func get(c *gin.Context) {
	cl := model.GetClipboard()
	c.JSON(200, model.Result{Code: 200, Success: true, Message: "获取成功", Clipboard: cl})
}
func listLog(c *gin.Context) {
	c.JSON(200, model.ListLog())
}
func Clipboard(e *gin.Engine) {
	g := e.Group("/clipboard")
	{
		g.POST("/upd", upd)
		g.GET("/get", get)
		g.GET("/listLog", listLog)
	}

}
