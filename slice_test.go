package gutil

import (
	"fmt"
	"testing"
)

func TestSplitIntArray(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	splitIntArray := SplitIntArray(s1, 3)
	fmt.Println(splitIntArray)
}

func TestSplitStringArray(t *testing.T) {
	s1 := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	splitStringArray := SplitStringArray(s1, 3)
	fmt.Println(splitStringArray)
}

func TestIntIntersect(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{4, 5, 6, 7, 8}
	intersect := IntIntersect(s1, s2)
	fmt.Println(intersect)
}

func TestIntDifference(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{4, 5, 6, 7, 8}
	difference := IntDifference(s2, s1)
	fmt.Println(difference)
}

func TestStringIntersect(t *testing.T) {
	s1 := []string{"1", "2", "3", "4", "5"}
	s2 := []string{"4", "5", "6", "7", "8"}
	intersect := StringIntersect(s1, s2)
	fmt.Println(intersect)
}

func TestStringDifference(t *testing.T) {
	s1 := []string{"1", "2", "3", "4", "5"}
	s2 := []string{"4", "5", "6", "7", "8"}
	difference := StringDifference(s2, s1)
	fmt.Println(difference)
}

func TestRemoveRepInt(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5, 5, 5}
	repInt := RemoveRepInt(s1)
	fmt.Println(repInt)
}

func TestRemoveRepStr(t *testing.T) {
	s1 := []string{"1", "2", "3", "4", "5", "2", "4"}
	repStr := RemoveRepStr(s1)
	fmt.Println(repStr)
}
