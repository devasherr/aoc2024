package day4

import (
	"bytes"
)

func checkHorizontal(i, j int, mat [][]byte) bool {
	return string(mat[i][j:j+4]) == "XMAS" || string(mat[i][j:j+4]) == "SAMX"
}

func checkVertical(i, j int, mat [][]byte) bool {
	var curStr bytes.Buffer
	count := 0

	for count < 4 {
		curStr.WriteString(string(mat[i+count][j]))
		count++
	}

	return curStr.String() == "XMAS" || curStr.String() == "SAMX"
}

func checkDiagonal(i, j int, mat [][]byte) bool {
	var curStr bytes.Buffer
	count := 0

	for count < 4 {
		curStr.WriteString(string(mat[i+count][j+count]))
		count++
	}

	return curStr.String() == "XMAS" || curStr.String() == "SAMX"
}

func checkDiagonalRev(i, j int, mat [][]byte) bool {
	var curStr bytes.Buffer
	count := 0

	for count < 4 {
		curStr.WriteString(string(mat[i+count][j-count]))
		count++
	}

	return curStr.String() == "XMAS" || curStr.String() == "SAMX"
}

func checkX_MAS(i, j int, mat [][]byte) bool {
	if string(mat[i][j]) != "A" {
		return false
	}

	// diagonal cannot be A
	if string(mat[i-1][j-1]) == "A" || string(mat[i-1][j+1]) == "A" || string(mat[i+1][j-1]) == "A" || string(mat[i+1][j+1]) == "A" ||
		string(mat[i-1][j-1]) == "X" || string(mat[i-1][j+1]) == "X" || string(mat[i+1][j-1]) == "X" || string(mat[i+1][j+1]) == "X" {
		return false
	}
	if mat[i-1][j-1] != mat[i+1][j+1] && mat[i+1][j-1] != mat[i-1][j+1] {
		return true
	}
	return false
}

func Solve(content []byte) int {
	matrix := bytes.Split(content, []byte("\n"))
	res := 0
	for i := 0; i < len(matrix)-1; i++ {
		line := matrix[i]
		for j := 0; j < len(line); j++ {
			// if j+3 < len(line) {
			// 	if checkHorizontal(i, j, matrix) {
			// 		res++
			// 	}
			// }
			// if i+3 < len(matrix)-1 {
			// 	if checkVertical(i, j, matrix) {
			// 		res++
			// 	}
			// }
			// if i+3 < len(matrix)-1 && j+3 < len(line) {
			// 	if checkDiagonal(i, j, matrix) {
			// 		res++
			// 	}
			// }
			// if i+3 < len(matrix)-1 && j >= 3 {
			// 	if checkDiagonalRev(i, j, matrix) {
			// 		res++
			// 	}
			// }
			if (i > 0 && i < len(matrix)-2) && (j > 0 && j < len(line)-1) {
				if checkX_MAS(i, j, matrix) {
					res++
				}
			}
		}
	}
	return res
}
