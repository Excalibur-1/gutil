// Package gutil 封装常用go工具类
package gutil

import (
	"fmt"
)

// IntIntersect 求交集
func IntIntersect(slice1, slice2 []int64) []int64 {
	m := make(map[int64]int64)
	nn := make([]int64, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, ok := m[v]
		if ok {
			if times == 1 {
				nn = append(nn, v)
			}
		}
	}
	return nn
}

// IntDifference 求差集 (除slice1，2并集外slice1存在的)
func IntDifference(slice1, slice2 []int64) []int64 {
	slice1 = RemoveRepInt(slice1)
	slice2 = RemoveRepInt(slice2)

	m := make(map[int64]int64)
	nn := make([]int64, 0)
	inter := IntIntersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
		fmt.Println()
	}
	return nn
}

// StringIntersect 求交集
func StringIntersect(slice1, slice2 []string) []string {
	m := make(map[string]int64)
	nn := make([]string, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, ok := m[v]
		if ok {
			if times == 1 {
				nn = append(nn, v)
			}
		}
	}
	return nn
}

// StringDifference 求差集 (除slice1，2并集外slice1存在的)
func StringDifference(slice1, slice2 []string) []string {
	slice1 = RemoveRepStr(slice1)
	slice2 = RemoveRepStr(slice2)

	m := make(map[string]int64)
	nn := make([]string, 0)
	inter := StringIntersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}

// RemoveRepStr 字符串元素去重
func RemoveRepStr(slc []string) []string {
	if len(slc) < 1024 {
		// 切片长度小于1024的时候，循环来过滤
		return removeRepStrByLoop(slc)
	} else {
		// 大于的时候，通过map来过滤
		return removeRepStrByMap(slc)
	}
}

// RemoveRepInt 整形元素去重
func RemoveRepInt(slc []int64) []int64 {
	if len(slc) < 1024 {
		// 切片长度小于1024的时候，循环来过滤
		return removeRepIntByLoop(slc)
	} else {
		// 大于的时候，通过map来过滤
		return removeRepIntByMap(slc)
	}
}

// 通过两重循环过滤重复元素
func removeRepIntByLoop(slc []int64) []int64 {
	// 存放结果
	var result []int64
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				// 存在重复元素，标识为false
				flag = false
				break
			}
		}
		if flag {
			// 标识为false，不添加进结果
			result = append(result, slc[i])
		}
	}
	return result
}

// 通过map主键唯一的特性过滤重复元素
func removeRepIntByMap(slc []int64) []int64 {
	var result []int64
	// 存放不重复主键
	tempMap := map[int64]byte{}
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		// 加入map后，map长度变化，则元素不重复
		if len(tempMap) != l {
			result = append(result, e)
		}
	}
	return result
}

// 通过两重循环过滤重复元素
func removeRepStrByLoop(slc []string) []string {
	// 存放结果
	var result []string
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				// 存在重复元素，标识为false
				flag = false
				break
			}
		}
		if flag {
			// 标识为false，不添加进结果
			result = append(result, slc[i])
		}
	}
	return result
}

// 通过map主键唯一的特性过滤重复元素
func removeRepStrByMap(slc []string) []string {
	var result []string
	// 存放不重复主键
	tempMap := map[string]byte{}
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		// 加入map后，map长度变化，则元素不重复
		if len(tempMap) != l {
			result = append(result, e)
		}
	}
	return result
}
