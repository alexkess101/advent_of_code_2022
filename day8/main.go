package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	block := ParseBlock()
	//var block = [][]int{
	//	{3, 0, 3, 7, 3},
	//	{2, 5, 5, 1, 2},
	//	{6, 5, 3, 3, 2},
	//	{3, 3, 5, 4, 9},
	//	{3, 5, 3, 9, 0},
	//}
	area := 0
	for rowIndex, row := range block {
		for columnIndex := range row {
			checkArea := CheckLeft(block, rowIndex, columnIndex) * CheckUp(block, rowIndex, columnIndex) * CheckRight(block, rowIndex, columnIndex) * CheckDown(block, rowIndex, columnIndex)
			if checkArea > area {
				area = checkArea
			}
		}
	}

	log.Println(area)
}

func ParseBlock() [][]int {
	var block [][]int
	content, _ := os.ReadFile("input.txt")
	rows := strings.Split(string(content), "\n")

	for _, row := range rows {
		rawColumns := strings.Split(row, "")
		var columns []int
		for _, column := range rawColumns {
			height, _ := strconv.Atoi(column)
			columns = append(columns, height)
		}
		block = append(block, columns)
	}

	return block
}

func GetPerimeterAmount(block [][]int) int {
	width := len(block[0])
	height := len(block)
	return (width * 2) + (height-2)*2
}

func CheckLeft(block [][]int, row, column int) int {
	height := block[row][column]
	count := 0

	for i := column - 1; i >= 0; i-- {
		if height <= block[row][i] {
			count++
			return count
		}
		count++
	}
	return count
}

func CheckRight(block [][]int, row, column int) int {
	height := block[row][column]
	count := 0

	for i := column + 1; i <= len(block[0])-1; i++ {
		if height <= block[row][i] {
			count++
			return count
		}
		count++
	}
	return count
}

func CheckUp(block [][]int, row, column int) int {
	height := block[row][column]
	count := 0

	for i := row - 1; i >= 0; i-- {
		if height <= block[i][column] {
			count++
			return count
		}
		count++
	}
	return count
}

func CheckDown(block [][]int, row, column int) int {
	height := block[row][column]
	count := 0

	for i := row + 1; i <= len(block)-1; i++ {
		if height <= block[i][column] {
			count++
			return count
		}
		count++
	}
	return count
}

func IsEdge(block [][]int, row, column int) bool {
	rowMax := len(block) - 1
	columnMax := len(block[0]) - 1

	if row == 0 || row == rowMax {
		return true
	}
	if column == 0 || column == columnMax {
		return true
	}

	return false
}
