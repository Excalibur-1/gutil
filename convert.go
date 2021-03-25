package gutil

import (
	"errors"
	"math"
	"strings"
)

var tenToAny = map[int]string{
	0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9",
	10: "a", 11: "b", 12: "c", 13: "d", 14: "e", 15: "f", 16: "g", 17: "h", 18: "i",
	19: "j", 20: "k", 21: "l", 22: "m", 23: "n", 24: "o", 25: "p", 26: "q", 27: "r",
	28: "s", 29: "t", 30: "u", 31: "v", 32: "w", 33: "x", 34: "y", 35: "z", 36: ":",
	37: ";", 38: "<", 39: "=", 40: ">", 41: "?", 42: "@", 43: "[", 44: "]", 45: "^",
	46: "_", 47: "{", 48: "|", 49: "}", 50: "A", 51: "B", 52: "C", 53: "D", 54: "E",
	55: "F", 56: "G", 57: "H", 58: "I", 59: "J", 60: "K", 61: "L", 62: "M", 63: "N",
	64: "O", 65: "P", 66: "Q", 67: "R", 68: "S", 69: "T", 70: "U", 71: "V", 72: "W",
	73: "X", 74: "Y", 75: "Z",
}

var anyToTen = map[string]int{
	"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
	"a": 10, "b": 11, "c": 12, "d": 13, "e": 14, "f": 15, "g": 16, "h": 17, "i": 18,
	"j": 19, "k": 20, "l": 21, "m": 22, "n": 23, "o": 24, "p": 25, "q": 26, "r": 27,
	"s": 28, "t": 29, "u": 30, "v": 31, "w": 32, "x": 33, "y": 34, "z": 35, ":": 36,
	";": 37, "<": 38, "=": 39, ">": 40, "?": 41, "@": 42, "[": 43, "]": 44, "^": 45,
	"_": 46, "{": 47, "|": 48, "}": 49, "A": 50, "B": 51, "C": 52, "D": 53, "E": 54,
	"F": 55, "G": 56, "H": 57, "I": 58, "J": 59, "K": 60, "L": 61, "M": 62, "N": 63,
	"O": 64, "P": 65, "Q": 66, "R": 67, "S": 68, "T": 69, "U": 70, "V": 71, "W": 72,
	"X": 73, "Y": 74, "Z": 75,
}

// DecimalToBase 10进制转指定进制(仅支持2-36进制以内的转换)
func DecimalToBase(num, base int) (newNumStr string, err error) {
	if 2 > base && base > 36 {
		err = errors.New("进制范围输入错误，应该在2-36之间")
		return
	}
	var remainder int
	for num != 0 {
		remainder = num % base
		newNumStr = tenToAny[remainder] + newNumStr
		num = num / base
	}
	return
}

// BaseToDecimal 指定进制转10进制(仅支持2-36进制以内的转换)
func BaseToDecimal(numStr string, base int) (newNum float64, err error) {
	if 2 > base && base > 36 {
		err = errors.New("进制范围输入错误，应该在2-36之间")
		return
	}
	split := strings.Split(numStr, "")
	x := float64(base)
	y := float64(len(split) - 1)
	for _, value := range split {
		if ten, ok := anyToTen[value]; ok {
			newNum = newNum + (float64(ten) * math.Pow(x, y))
			y = y - 1
		} else {
			err = errors.New("转换失败，转换的字符串不合法")
			return
		}
	}
	return
}
