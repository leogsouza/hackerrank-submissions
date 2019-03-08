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

func Mean(n []int) float64 {
	size := float64(len(n))
	t := 0.0
	for _, v := range n {
		t += float64(v)
	}

	return t / size
}

func Median(n []int) float64 {
	size := len(n)
	sort.Ints(n)
	median := Mean(n[size/2-1 : size/2+1])

	return median
}

func Mode(n []int) int {
	size := len(n)
	sort.Ints(n)
	smallest := n[0]
	m := make(map[int]int, size)

	for _, v := range n {
		m[v]++
	}

	max := 0
	num := 0
	for _, v := range n {
		if m[v] > max {
			max = m[v]
			num = v
		}
	}

	if smallest == max {
		return smallest
	}

	return num
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

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer outputFile.Close()

	writer := bufio.NewWriterSize(outputFile, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	aRowTemp := strings.Split(readLine(reader), " ")
	var aRow []int
	for _, aRowItem := range aRowTemp {
		aItemTemp, err := strconv.ParseInt(aRowItem, 10, 64)
		checkError(err)
		aItem := int(aItemTemp)
		aRow = append(aRow, aItem)
	}

	if len(aRow) != int(int(n)) {
		panic("Bad input")
	}

	mean := Mean(aRow)
	median := Median(aRow)
	mode := Mode(aRow)
	fmt.Fprintf(writer, "%f\n", mean)
	fmt.Fprintf(writer, "%f\n", median)
	fmt.Fprintf(writer, "%d\n", mode)

	writer.Flush()
}
