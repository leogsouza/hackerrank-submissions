package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	buffer := bufio.NewReader(os.Stdin)
	arr2d := make([][]int, 6)

	for i := 0; i < len(arr2d); i++ {
		//fmt.Println(scanner.Text())
		arr2d[i] = convertStrToIntArray(buffer)
	}

	fmt.Println(getMaxSumHourglass(arr2d))

}

func convertStrToIntArray(buff *bufio.Reader) []int {

	inputString, _ := buff.ReadString('\n')

	inputArray := strings.Split(inputString, " ")

	converted := make([]int, 6)

	for i := 0; i < 6; i++ {
		converted[i], _ = strconv.Atoi(strings.TrimSpace(inputArray[i]))
	}

	return converted
}

func getSumHourglass(matrix [][]int, row, column int) int {

	return matrix[row][column] + matrix[row][column+1] + matrix[row][column+2] +
		matrix[row+1][column+1] +
		matrix[row+2][column] + matrix[row+2][column+1] + matrix[row+2][column+2]
}

func getMaxSumHourglass(matrix [][]int) int {
	var maxSumHourglass int

	for row := 0; row < len(matrix)-2; row++ {
		for column := 0; column < len(matrix)-2; column++ {
			sumHourglass := getSumHourglass(matrix, row, column)

			if sumHourglass > maxSumHourglass || (row == 0 && column == 0) {
				maxSumHourglass = sumHourglass
			}
		}
	}

	return maxSumHourglass
}
