package gutil

import (
	"fmt"
	"testing"
)

func TestGetIntEnv(t *testing.T) {
	env := GetIntEnv("limit", 10)
	fmt.Println(env)
}

func TestGetStringEnv(t *testing.T) {
	env := GetStringEnv("profile", "dev")
	fmt.Println(env)
}
