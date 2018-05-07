package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the miniMaxSum function below.
 */
func miniMaxSum(arr []int32) {

	var exc int32
	var size = int32(len(arr))
	var arrSum []int64
	var sum int64
	for exc = 0; exc < size; exc++ {
		sum = 0
		for i := int32(0); i < size; i++ {
			if i != exc {
				sum += int64(arr[i])
			}
		}
		arrSum = append(arrSum, sum)
	}

	fmt.Println(getMin(arrSum), getMax(arrSum))
}

func getMax(arr []int64) int64 {
	var max int64 = arr[0]

	for _, value := range arr {
		if max < value {
			max = value
		}
	}

	return max
}

func getMin(arr []int64) int64 {
	var min int64 = arr[0]

	for _, value := range arr {
		if min > value {
			min = value
		}
	}

	return min
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for arrItr := 0; arrItr < 5; arrItr++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[arrItr], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
