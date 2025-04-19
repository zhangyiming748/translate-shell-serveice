package logic

import (
	"fmt"
	"github.com/zhangyiming748/basicGin/util"
	"testing"
)

func init() {
	util.SetLog("telegram.log")
}

// go test -v -run TestTrans
func TestTrans(t *testing.T) {
	src := "hello world"
	proxy := "127.0.0.1:8889"
	output, err := Trans(src, proxy)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(output)
	// Output: "你好，世界！
}
