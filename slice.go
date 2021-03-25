// Package gutil 封装常用go工具类
package gutil

// SplitStringArray 分割数组，根据传入的数组和分割大小，将数组分割为大小等于指定大小的多个数组，如果不够分，则最后一个数组元素小于其他数组
func SplitStringArray(arr []string, num int) [][]string {
	max := len(arr)
	//判断数组大小是否小于等于指定分割大小的值，是则把原数组放入二维数组返回
	if max <= num {
		return [][]string{arr}
	}
	//获取应该数组分割为多少份
	var quantity int
	if max%num == 0 {
		quantity = max / num
	} else {
		quantity = (max / num) + 1
	}
	//声明分割好的二维数组
	var segments = make([][]string, 0)
	//声明分割数组的截止下标
	var start, end, i int
	for i = 1; i <= quantity; i++ {
		end = i * num
		if i != quantity {
			segments = append(segments, arr[start:end])
		} else {
			segments = append(segments, arr[start:])
		}
		start = i * num
	}
	return segments
}

// SplitIntArray 分割数组，根据传入的数组和分割大小，将数组分割为大小等于指定大小的多个数组，如果不够分，则最后一个数组元素小于其他数组
func SplitIntArray(arr []int, num int) [][]int {
	max := len(arr)
	//判断数组大小是否小于等于指定分割大小的值，是则把原数组放入二维数组返回
	if max <= num {
		return [][]int{arr}
	}
	//获取应该数组分割为多少份
	var quantity int
	if max%num == 0 {
		quantity = max / num
	} else {
		quantity = (max / num) + 1
	}
	//声明分割好的二维数组
	var segments = make([][]int, 0)
	//声明分割数组的截止下标
	var start, end, i int
	for i = 1; i <= quantity; i++ {
		end = i * num
		if i != quantity {
			segments = append(segments, arr[start:end])
		} else {
			segments = append(segments, arr[start:])
		}
		start = i * num
	}
	return segments
}

// IntIntersect 求交集
func IntIntersect(slice1, slice2 []int) []int {
	m := make(map[int]int)
	nn := make([]int, 0)
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
func IntDifference(slice1, slice2 []int) []int {
	slice1 = RemoveRepInt(slice1)
	slice2 = RemoveRepInt(slice2)

	m := make(map[int]int)
	nn := make([]int, 0)
	inter := IntIntersect(slice1, slice2)
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
func RemoveRepInt(slc []int) []int {
	if len(slc) < 1024 {
		// 切片长度小于1024的时候，循环来过滤
		return removeRepIntByLoop(slc)
	} else {
		// 大于的时候，通过map来过滤
		return removeRepIntByMap(slc)
	}
}

// 通过两重循环过滤重复元素
func removeRepIntByLoop(slc []int) []int {
	// 存放结果
	var result []int
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
func removeRepIntByMap(slc []int) []int {
	var result []int
	// 存放不重复主键
	tempMap := map[int]byte{}
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
