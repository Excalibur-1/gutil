package gutil

import (
	"bufio"
	"errors"
	"io/ioutil"
	"os"
)

// Write 写入文件
func Write(fileName, str string) (err error) {
	// 校验文件是否存在
	if checkFileExist(fileName) {
		err = errors.New("file already exists")
		return
	}

	//创建文件
	f, err := os.Create(fileName)
	if err != nil {
		return
	}
	defer func() { _ = f.Close() }()

	//创建新的 Writer 对象
	w := bufio.NewWriter(f)
	if _, err = w.WriteString(str); err != nil {
		return
	}

	// 刷新到磁盘
	err = w.Flush()

	return
}

// Read 读取到file中，再利用 ioutil 将 file 直接读取到 []byte 中
// 注意：此方法适合读取小文件，大文件读取最好是采用分片或者逐行读取的方式，减少内存使用
func Read(fileName string) (fileStr string, err error) {
	if !checkFileExist(fileName) {
		err = errors.New("file not found")
	}
	f, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer func() { _ = f.Close() }()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}

	fileStr = string(fd)
	return
}

// 检查文件是否存在
func checkFileExist(fileName string) (exists bool) {
	_, err := os.Stat(fileName)
	exists = os.IsExist(err)
	return
}
