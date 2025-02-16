package main

import "fmt"

func Max(nums []int) int { // 本来想给make设定容量。但是把数组和map混淆了。map不需要声明连续的长度。

	a := nums[0]
	i := 0
	for i < len(nums) {
		if nums[i] < nums[i+1] {
			a = nums[i+1]
			i++
		}
		if i >= len(nums)-1 {
			break
		}

	}
	return a
}

func twoSum(nums []int, target int) []int {

	hashmap := make(map[int]int)
	for i, num := range nums {
		hashmap[num] = i
	}
	for i, num := range nums {
		if val, ok := hashmap[target-num]; ok && i != val {
			return []int{i, val}
		}
	}
	return nil
}

func main() {
	nums := []int{3, 2, 4}
	target := 6
	fmt.Println(twoSum(nums, target))
}
