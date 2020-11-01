package main

/**
TODO Golang中单元测试的文件名必须以_test结尾
	测试用例必须以Test开头，且参数必须为(*testing.T)
*/
import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	s := "abba"
	bytes := []byte(s)
	result := 0
	/**
	  使用两个头尾指针来处理
	  **/
	left := 0
	right := 0
	windows := make(map[byte]int, len(bytes))
	for left < len(s) && right < len(s) {
		value := bytes[right]
		if deleteIndex, ok := windows[value]; ok {
			for left <= deleteIndex {
				delete(windows, bytes[left])
				left++
			}
		} else {
			windows[value] = right
			right = right + 1
			if result <= (right - left) {
				result = (right - left)
			}
		}
	}
	fmt.Println(result)
}

func TestWorld(t *testing.T) {
	fmt.Println("World")
}
