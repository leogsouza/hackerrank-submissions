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
 * Complete the birthdayCakeCandles function below.
 */
func birthdayCakeCandles(n int32, ar []int32) int32 {

	tallest := tallestCandle(ar)

	total := countTallest(tallest, ar)

	return total
}

func tallestCandle(ar []int32) int32 {

	var tallest int32

	for _, value := range ar {

		if tallest < value {
			tallest = value
		}
	}

	return tallest
}

func countTallest(tallest int32, ar []int32) int32 {

	var count int32

	for _, value := range ar {
		if tallest == value {
			count++
		}
	}

	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer outputFile.Close()

	writer := bufio.NewWriterSize(outputFile, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arTemp := strings.Split(readLine(reader), " ")

	var ar []int32

	for arItr := 0; arItr < int(n); arItr++ {
		arItemTemp, err := strconv.ParseInt(arTemp[arItr], 10, 64)
		checkError(err)
		arItem := int32(arItemTemp)
		ar = append(ar, arItem)
	}

	result := birthdayCakeCandles(n, ar)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
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
