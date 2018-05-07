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
 * Complete the plusMinus function below.
 */
func plusMinus(arr []int32) {
	var neutral, positive, negative float32

	for _, number := range arr {

		switch {
		default:
			neutral++
		case number > 0:
			positive++
		case number < 0:
			negative++
		}
	}

	size := float32(len(arr))

	fmt.Printf("%f\n", positive/size)
	fmt.Printf("%f\n", negative/size)
	fmt.Printf("%f\n", neutral/size)

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for arrItr := 0; arrItr < int(n); arrItr++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[arrItr], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
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
