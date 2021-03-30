package gutil

import (
	"encoding/base64"
	"fmt"
	"testing"
)

const key = "12345678"

func TestEncrypt(t *testing.T) {
	src := "{\"backend\":\"localhost:2181\",\"username\":\"guest\",\"password\":\"guest\"}"
	fmt.Println("原文：" + src)
	enc, err := Encrypt([]byte(src), []byte(key))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("密文：" + base64.URLEncoding.EncodeToString(enc))
}

func TestDecrypt(t *testing.T) {
	decodeString, err := base64.URLEncoding.DecodeString("vs13PY3sDZ4y40bWslbg5ExmTbXQQdyRo7AvpiMrIG-B-ukE8dLK5DrJOzS9dhmWjqhtIxWDPn8ynRvwjrDob2N25A0U_OZa")
	if err != nil {
		fmt.Println(err)
		return
	}
	dec, err := Decrypt(decodeString, []byte(key))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("解码：" + string(dec))
}
