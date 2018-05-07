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
 * Complete the timeConversion function below.
 */
func timeConversion(s string) string {

	var strConverted, suffix string

	if strings.HasSuffix(s, "PM") {
		suffix = "PM"
	} else {
		suffix = "AM"
	}

	strConverted = strings.Replace(s, suffix, "", 1)
	splitTime := strings.Split(strConverted, ":")

	if splitTime[0] == "12" && suffix == "AM" {
		splitTime[0] = "00"
	} else if suffix == "PM" && splitTime[0] != "12" {
		convHours, _ := strconv.Atoi(splitTime[0])
		convHours += 12

		splitTime[0] = strconv.Itoa(convHours)
	}

	return strings.Join(splitTime, ":")

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer outputFile.Close()

	writer := bufio.NewWriterSize(outputFile, 1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

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
