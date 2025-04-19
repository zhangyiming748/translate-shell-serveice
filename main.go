package main

import (
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"github.com/zhangyiming748/translate-server/bootstrap"
	mysql "github.com/zhangyiming748/translate-server/storage"
	"github.com/zhangyiming748/translate-server/util"
)

func testResponse(c *gin.Context) {
	c.JSON(http.StatusGatewayTimeout, gin.H{
		"code": http.StatusGatewayTimeout,
		"msg":  "timeout",
	})
}

func timeoutMiddleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(3000*time.Millisecond),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(testResponse),
	)
}
func init() {

	if runtime.GOOS == "linux" {
		util.SetLog("/app/translate-server.log")
	} else {
		util.SetLog("translate-server.log")
	}
	log.SetFlags(log.Ltime | log.Lshortfile)
	// 初始化mysql
	mysql.SetMysql()
}

func main() {
	// gin服务
	gin.SetMode(gin.DebugMode)
	engine := gin.New()
	engine.Use(timeoutMiddleware())
	bootstrap.InitTranslateService(engine)
	// 启动http服务
	engine.Run(":8192")
}
