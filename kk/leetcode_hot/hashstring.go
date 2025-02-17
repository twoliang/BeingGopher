// package main
//
// import (
//
//	"fmt"
//	"sort"
//
// )
//
//	func groupAnagrams(strs []string) (ans [][]string) {
//		hashmap := make(map[string]([]string))
//		for _, val := range strs {
//			hashmap[sortLetters(val)] = append(hashmap[sortLetters(val)], val)
//		}
//
//		for _, val := range hashmap {
//			ans = append(ans, val)
//		}
//		return ans
//	}
//
//	func sortLetters(strs string) string {
//		s0 := []rune(strs)
//		sort.Slice(s0, func(i, j int) bool {
//			return s0[i] < s0[j]
//		})
//		return string(s0)
//	}
//
//	func main() {
//		strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
//		fmt.Println(groupAnagrams(strs))
//	}
package main

import "fmt"

func main() {
	// 包含 Unicode 字符的字符串
	str := "Hello, 世界"

	// 将字符串转换为 rune 切片
	runes := []rune(str)
	fmt.Println(runes)
	fmt.Println(string(runes))
	// 打印每个 rune 及其 Unicode 代码点
	for _, r := range runes {
		fmt.Printf("%c: %U\n", r, r)
	}
}
