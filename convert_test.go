package gutil

import (
	"fmt"
	"testing"
)

func TestDecimalToAny(t *testing.T) {
	numStr, err := DecimalToBase(10000, 36)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(numStr)
}

func TestAnyToDecimal(t *testing.T) {
	decimal, err := BaseToDecimal("7ps", 36)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(decimal)
}

