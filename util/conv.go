package util

import (
	"encoding/base64"
	"fmt"
	"log"
)

func Base64(s string) {

	// 解码
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}

	// 输出解码后的结果
	fmt.Println(string(data))
}
