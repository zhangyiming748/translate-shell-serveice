package main

import (
	"log"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/zhangyiming748/translate-server/bootstrap"
	mysql "github.com/zhangyiming748/translate-server/storage"
	"github.com/zhangyiming748/translate-server/util"
)

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
	bootstrap.InitTranslateService(engine)
	// 启动http服务
	engine.Run(":8192")
}
