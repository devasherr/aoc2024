package day3

import (
	"bytes"
	"regexp"
	"strconv"
)

func regexMatch(line string) int {
	curMul := 0
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`) // i do love me some regex
	i := 0

	for i < len(line) {
		substr := line[i:]
		loc := re.FindStringSubmatchIndex(substr)

		if loc != nil {
			num1 := substr[loc[2]:loc[3]]
			num2 := substr[loc[4]:loc[5]]

			a, _ := strconv.Atoi(num1)
			b, _ := strconv.Atoi(num2)

			curMul += (a * b)
			i += loc[1]
		} else {
			i++
		}
	}

	return curMul
}

func Solve(content []byte) int {
	res := 0
	lines := bytes.Split(content, []byte("\n"))
	for _, line := range lines {
		res += regexMatch(string(line))
	}
	return res
}
