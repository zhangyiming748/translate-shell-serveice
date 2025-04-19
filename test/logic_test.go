package test

import (
	"fmt"
	"testing"

	"github.com/zhangyiming748/basicGin/util"
	"github.com/zhangyiming748/translate-server/logic"
	mysql "github.com/zhangyiming748/translate-server/storage"
)

func init() {
	util.SetLog("TransService.log")
}

// go test -v -run TestTrans
func TestTrans(t *testing.T) {
	mysql.SetMysql()
	src := "hello world"
	proxy := "192.168.2.10:8889"
	output, err := logic.Trans(src, proxy)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(output)
	// Output: "你好，世界！
}
