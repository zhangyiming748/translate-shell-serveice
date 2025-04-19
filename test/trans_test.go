package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/zhangyiming748/translate-server/controller"
)

func setupRouter() *gin.Engine {
	r := gin.New()
	tc := controller.TranslateController{}
	r.GET("/trans/get", tc.GetSrc)
	r.POST("/trans/post", tc.PostSrcWithProxy)
	return r
}

func TestGetSrc(t *testing.T) {
	r := setupRouter()

	// 测试正常情况
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/trans/get?src=hello", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "dst is")

	// 测试空参数情况
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/trans/get", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
}

func TestPostSrcWithProxy(t *testing.T) {
	r := setupRouter()

	// 测试正常情况
	body := controller.RequestBody{
		Src:   "hello world",
		Proxy: "127.0.0.1:8889",
	}
	jsonData, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/trans/post", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var response controller.ResponseBody
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, body.Src, response.Src)
	assert.NotEmpty(t, response.Dst)

	// 测试无效JSON
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/trans/post", bytes.NewBufferString("invalid json"))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}
