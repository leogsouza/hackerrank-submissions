package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

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

func sum(values []int) int {
	val := 0
	for _, v := range values {
		val += v
	}

	return val
}

func median(numbers []int) int {
	size := len(numbers)
	if size%2 == 0 {
		fmt.Println(numbers[size/2])
		return sum(numbers[size/2-1:size/2+1]) / 2
	}
	return numbers[size/2]
}

func Quartiles(n int, numbers []int) (q1, q2, q3 int) {
	q1, q2, q3 = 0, 0, 0
	sort.Ints(numbers)
	q1 = median(numbers[:n/2])
	q2 = median(numbers)
	if n%2 == 0 {
		q3 = median(numbers[n/2:])
	} else {
		q3 = median(numbers[n/2+1:])
	}

	return
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer outputFile.Close()

	writer := bufio.NewWriterSize(outputFile, 1024*1024)

	numItems, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int(numItems)

	lineValues := strings.Split(readLine(reader), " ")
	var numbers []int
	for _, value := range lineValues {
		itemValue, err := strconv.ParseInt(value, 10, 64)
		checkError(err)
		number := int(itemValue)
		numbers = append(numbers, number)
	}

	if len(numbers) != int(int(n)) {
		panic("Bad input")
	}

	q1, q2, q3 := Quartiles(n, numbers)

	fmt.Fprintf(writer, "%d\n", q1)
	fmt.Fprintf(writer, "%d\n", q2)
	fmt.Fprintf(writer, "%d\n", q3)

	writer.Flush()
}
