package config

import (
	"gin/api"
	"github.com/gin-gonic/gin"
)

func Routes(e *gin.Engine) {
	e.GET("/indexs", api.FindAllIndex)
	e.POST("/index", api.AddOneIndex)

}
