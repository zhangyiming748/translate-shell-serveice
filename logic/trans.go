package logic

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/zhangyiming748/translate-server/model"
)

func Trans(src, proxy string) (string, error) {
	h := new(model.History)
	h.Src = src
	if found, _ := h.FindBySrc(); found {
		return h.Dst, nil
	}
	var cmd *exec.Cmd
	if proxy == "" {
		cmd = exec.Command("trans", "-brief", "-engine", "bing", ":zh-CN", src)
	} else {
		cmd = exec.Command("trans", "-brief", "-engine", "google", "-proxy", proxy, ":zh-CN", src)
	}
	log.Println(cmd.String())
	output, err := cmd.CombinedOutput()
	if err != nil || strings.Contains(string(output), "u001b") || strings.Contains(string(output), "Didyoumean") || strings.Contains(string(output), "Connectiontimedout") {
		return "", fmt.Errorf("查询命令执行出错\t命令原文:%v\t错误原文:%v\n", cmd.String(), err.Error())
	}
	dst := string(output)
	dst = strings.ReplaceAll(dst, "\n", "") // 删除所有换行符
	dst = strings.ReplaceAll(dst, "\r", "") // 删除所有回车符
	h.Dst = dst
	h.InsertOne()
	return dst, nil
}
