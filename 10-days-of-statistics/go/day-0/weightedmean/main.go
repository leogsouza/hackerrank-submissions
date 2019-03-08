package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func WeightedMean(x []int, w []int) float64 {
	sumx := 0.0
	sumw := 0.0
	size := len(x)

	for i := 0; i < size; i++ {
		sumx += float64(x[i] * w[i])
		sumw += float64(w[i])
	}

	return roundTo(sumx/sumw, 10)
}

func roundTo(n float64, decimals uint32) float64 {
	return math.Round(n*float64(decimals)) / float64(decimals)
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
	n := int(nTemp)

	aNumTemp := strings.Split(readLine(reader), " ")
	var x []int
	for _, aNumItem := range aNumTemp {
		aNumItemTemp, err := strconv.ParseInt(aNumItem, 10, 64)
		checkError(err)
		aItem := int(aNumItemTemp)
		x = append(x, aItem)
	}

	if len(x) != n {
		panic("Bad input")
	}

	weights := strings.Split(readLine(reader), " ")
	var w []int
	for _, wItem := range weights {
		wItemTemp, err := strconv.ParseInt(wItem, 10, 64)
		checkError(err)
		wItem := int(wItemTemp)
		w = append(w, wItem)
	}

	if len(w) != n {
		panic("Bad input")
	}

	result := WeightedMean(x, w)

	fmt.Fprintf(writer, "%.1f\n", result)

	writer.Flush()
}
