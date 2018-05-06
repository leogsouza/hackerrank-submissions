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
 * Complete the solve function below.
 */
func solve(a0 int32, a1 int32, a2 int32, b0 int32, b1 int32, b2 int32) []int32 {
	var scoreA, scoreB int32

	calculate(a0, b0, &scoreA, &scoreB)
	calculate(a1, b1, &scoreA, &scoreB)
	calculate(a2, b2, &scoreA, &scoreB)

	result := []int32{scoreA, scoreB}
	return result

}

func calculate(ratingA, ratingB int32, scoreA, scoreB *int32) {

	if ratingA > ratingB {
		*scoreA++
	} else if ratingA < ratingB {
		*scoreB++
	}

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer outputFile.Close()

	writer := bufio.NewWriterSize(outputFile, 1024*1024)

	a0A1A2 := strings.Split(readLine(reader), " ")

	a0Temp, err := strconv.ParseInt(a0A1A2[0], 10, 64)
	checkError(err)
	a0 := int32(a0Temp)

	a1Temp, err := strconv.ParseInt(a0A1A2[1], 10, 64)
	checkError(err)
	a1 := int32(a1Temp)

	a2Temp, err := strconv.ParseInt(a0A1A2[2], 10, 64)
	checkError(err)
	a2 := int32(a2Temp)

	b0B1B2 := strings.Split(readLine(reader), " ")

	b0Temp, err := strconv.ParseInt(b0B1B2[0], 10, 64)
	checkError(err)
	b0 := int32(b0Temp)

	b1Temp, err := strconv.ParseInt(b0B1B2[1], 10, 64)
	checkError(err)
	b1 := int32(b1Temp)

	b2Temp, err := strconv.ParseInt(b0B1B2[2], 10, 64)
	checkError(err)
	b2 := int32(b2Temp)

	result := solve(a0, a1, a2, b0, b1, b2)

	for resultItr, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if resultItr != len(result)-1 {
			fmt.Fprintf(writer, " ")
		}
	}

	fmt.Fprintf(writer, "\n")

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
