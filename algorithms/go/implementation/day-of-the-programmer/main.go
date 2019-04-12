package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the dayOfProgrammer function below.
const DAY_PROGRAMMER int32 = 256

func dayOfProgrammer(year int32) string {
	totalDays := sumDaysOfMonth(year)
	result := DAY_PROGRAMMER - totalDays

	if year == 1918 {
		result += 13
	}
	d := strconv.FormatInt(int64(result), 10)
	y := strconv.FormatInt(int64(year), 10)
	return d + ".09." + y

}

func sumDaysOfMonth(year int32) int32 {
	months := []int32{31, 28, 31, 30, 31, 30, 31, 31}
	julCalendar := year < 1918
	if isLeapYear(year, julCalendar) {
		months[1]++
	}
	var sum int32
	for _, v := range months {
		sum += v
	}
	return sum
}

func isLeapYear(year int32, julCalendar bool) bool {

	if julCalendar {
		return year%4 == 0
	}

	return year%400 == 0 || (year%4 == 0 && year%100 > 0)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	yearTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	year := int32(yearTemp)

	result := dayOfProgrammer(year)

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
