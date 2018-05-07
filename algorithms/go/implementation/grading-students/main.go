package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const MULTIPLEOF int32 = 5

/*
 * Complete the gradingStudents function below.
 */
func gradingStudents(grades []int32) []int32 {

	var result []int32

	for _, grade := range grades {
		if grade >= 38 {
			nextMultiple := getNextMultiple(grade)

			diffGrade := nextMultiple - grade
			if diffGrade < 3 {
				result = append(result, nextMultiple)
			} else {
				result = append(result, grade)
			}
		} else {
			result = append(result, grade)
		}
	}

	return result

}

func getNextMultiple(grade int32) int32 {

	return grade + (MULTIPLEOF - grade%MULTIPLEOF)
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

	var grades []int32

	for gradesItr := 0; gradesItr < int(n); gradesItr++ {
		gradesItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		gradesItem := int32(gradesItemTemp)
		grades = append(grades, gradesItem)
	}

	result := gradingStudents(grades)

	for resultItr, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if resultItr != len(result)-1 {
			fmt.Fprintf(writer, "\n")
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
