package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangyiming748/ollamaTranslate/controller"
)

func InitTranslate(engine *gin.Engine) {
	routeGroup := engine.Group("/api")
	{
		c := new(controller.TransController)
		routeGroup.GET("/v1/trans/health", c.GetHealth)
		routeGroup.POST("/v1/trans/chinese", c.PostChinese)
	}
}
