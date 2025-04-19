package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangyiming748/translate-server/controller"
)

func InitTranslateService(engine *gin.Engine) {
	routeGroup := engine.Group("/api")
	{
		c := new(controller.TranslateController)
		routeGroup.GET("/v1/GetTrans", c.GetSrc)
		routeGroup.POST("/v1/PostTrans", c.PostSrcWithProxy)
	}
}
