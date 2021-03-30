package gutil

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
)

func Encrypt(src, key []byte) (dst []byte, err error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return
	}
	src = padding(src, block.BlockSize())
	dst = make([]byte, len(src))
	cipher.NewCBCEncrypter(block, key).CryptBlocks(dst, src)
	return

}

func Decrypt(src, key []byte) (dst []byte, err error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return
	}
	dst = make([]byte, len(src))
	cipher.NewCBCDecrypter(block, key).CryptBlocks(dst, src)
	dst = unPadding(dst)
	return
}

// 填充字符串（末尾）
func padding(text []byte, blockSize int) []byte {
	// 需要填充的数据长度
	paddingCount := blockSize - len(text)%blockSize
	// 填充数据为：paddingCount,填充的值为：paddingCount
	paddingText := bytes.Repeat([]byte{byte(paddingCount)}, paddingCount)
	return append(text, paddingText...)
}

// 去掉字符（末尾）
func unPadding(text []byte) []byte {
	textSize := len(text)
	return text[:textSize-int(text[textSize-1])]
}
