package logic

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func Trans(src, proxy string) (string, error) {
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
	return string(output), nil
}
