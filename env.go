// Package gutil 封装常用go工具类
package gutil

import (
	"os"
	"strconv"
	"strings"
)

// GetIntEnv 获取 int 类型环境变量，没有返回默认值
func GetIntEnv(env string, defaultValue int) int {
	value := strings.Trim(os.Getenv(env), "")
	if value == "" {
		return defaultValue
	}
	v, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return v
}

// GetStringEnv 获取 string 类型环境变量，没有返回默认值
func GetStringEnv(envKey string, defaultValue string) string {
	if value := strings.TrimSpace(os.Getenv(envKey)); value != "" {
		return value
	}
	return defaultValue
}
