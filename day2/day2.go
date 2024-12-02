package day2

import (
	"bytes"
	"strconv"
)

func checkIncrease(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		val := nums[i] - nums[i+1]
		if val < 0 {
			val *= -1
		}
		if nums[i] > nums[i+1] || val < 1 || val > 3 {
			return false
		}
	}

	return true
}

func checkDecrease(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		val := nums[i] - nums[i+1]
		if val < 0 {
			val *= -1
		}
		if nums[i] < nums[i+1] || val < 1 || val > 3 {
			return false
		}
	}
	return true
}

func Solve(content []byte) int {
	lines := bytes.Split(content, []byte("\n"))
	res := 0
	for _, line := range lines[:len(lines)-1] {
		curNums := bytes.Split(line, []byte(" "))
		nums := make([]int, len(curNums))
		for i, cur_n := range curNums {
			n, _ := strconv.Atoi(string(cur_n))
			nums[i] = n
		}
		if checkIncrease(nums) || checkDecrease(nums) {
			res++
		}
	}

	return res
}
