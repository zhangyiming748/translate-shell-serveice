package controller

import (
	"fmt"
	"log"
"github.com/zhangyiming748/translate-server/logic"
	"github.com/gin-gonic/gin"
)

type TranslateController struct{}

// 结构体必须大写 否则找不到
type RequestBody struct {
	Src string `json:"src"`
	Proxy string `json:"proxy,omitempty"`
}

type ResponseBody struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}

/*
使用 GET 请求获取翻译结果:
   curl --location --request GET 'http://127.0.0.1:8080/trans/get?src=hello'*/
func (ts TranslateController) GetSrc(ctx *gin.Context) {
	src := ctx.Query("src")
	dst,err:=logic.Trans(src, "")
	if err!= nil {
		log.Println(err)
		ctx.String(500, "服务器出错")
		return
	}
	ctx.String(200, fmt.Sprintf("dst is %s\n", dst))
}



/*
使用 POST 请求并指定代理获取翻译结果:
   curl -X POST -H "Content-Type: application/json" -d '{"src": "hello world", "proxy": "IP_ADDRESS:8889"}' http://localhost:8080/trans/post
 */
func (ts TranslateController) PostSrcWithProxy(ctx *gin.Context) {
	fmt.Println("get")
	var requestBody RequestBody
	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	} else {
		fmt.Println(requestBody)
	}
	fmt.Println(requestBody.Src, requestBody.Proxy)
	var rep ResponseBody
	result , err :=logic.Trans(requestBody.Src, requestBody.Proxy)
	if err!= nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	rep.Src = requestBody.Src
	rep.Dst =result
	ctx.JSON(200, rep)
}
