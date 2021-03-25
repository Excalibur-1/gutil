package gutil

import (
	"fmt"
	"testing"
)

func TestWrite(t *testing.T) {
	err := Write("1.txt", "hello world")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestRead(t *testing.T) {
	fileStr, err := Read("1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fileStr)
}
