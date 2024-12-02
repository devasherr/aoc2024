package day1

import (
	"bytes"
	"sort"
	"strconv"
)

func compare(first, second []int) int {
	sort.Ints(first)
	sort.Ints(second)

	res := 0
	for i := 0; i < len(first)-1; i++ {
		cur := first[i] - second[i]
		if cur < 0 {
			cur *= -1
		}

		res += cur
	}
	return res
}

func compareWithCount(first, second []int) int {
	count := map[int]int{}
	for _, num := range second {
		count[num]++
	}

	res := 0
	for _, num := range first {
		res += (num * count[num])
	}

	return res
}

func Solve(content []byte) int {
	lines := bytes.Split(content, []byte("\n"))
	first := make([]int, len(lines)-1)
	second := make([]int, len(lines)-1)
	for i, line := range lines {
		if i == len(lines)-1 {
			break
		}
		inputs := bytes.Split(line, []byte("   "))
		int_a, _ := strconv.Atoi(string(inputs[0]))
		int_b, _ := strconv.Atoi(string(inputs[1]))

		first[i] = int_a
		second[i] = int_b
	}

	// first version
	// return compare(first, second)

	// second version
	return compareWithCount(first, second)
}
