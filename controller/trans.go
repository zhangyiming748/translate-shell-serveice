package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangyiming748/translate-server/logic"
	"log"
)

type TranslateController struct{}

// 结构体必须大写 否则找不到
type RequestBody struct {
	Src   string `json:"src"`
	Proxy string `json:"proxy,omitempty"`
}

type ResponseBody struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}

/*
curl --location --request GET 'http://127.0.0.1:8192/api/v1/GetTrans?src=hello'
*/
func (ts TranslateController) GetSrc(ctx *gin.Context) {
	log.Printf("received src: %s", ctx.Query("src"))
	src := ctx.Query("src")
	proxy := ctx.Query("proxy")
	dst, err := logic.Trans(src, proxy)
	if err != nil {
		log.Println(err)
		ctx.String(500, "服务器出错")
		return
	}
	log.Printf("received src: %s, dst: %s", src, dst)
	var rep ResponseBody
	rep.Src = src
	rep.Dst = dst
	ctx.JSON(200, rep)
}

/*
curl --location --request POST 'http://127.0.0.1:8192/api/v1/PostTrans' \
--header 'Content-Type: application/json' \

	--data-raw '{
	    "src": "string",
	    "proxy": "string"
	}'
*/
func (ts TranslateController) PostSrcWithProxy(ctx *gin.Context) {
	var requestBody RequestBody
	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	} else {
		log.Printf("received src: %s, proxy: %s", requestBody.Src, requestBody.Proxy)
	}
	var rep ResponseBody
	result, err := logic.Trans(requestBody.Src, requestBody.Proxy)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	rep.Src = requestBody.Src
	rep.Dst = result
	log.Printf("received src: %s, dst: %s", rep.Src, rep.Dst)
	ctx.JSON(200, rep)
}
