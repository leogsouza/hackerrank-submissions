package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the isBalanced function below.
func isBalanced(s string) string {
	brackets := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	stack := []rune{}
	for _, c := range s {
		res, ok := brackets[c]

		if !ok {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return "NO"
			}
			top := stack[len(stack)-1]
			if res != top {
				return "NO"
			}
			stack = stack[:len(stack)-1]
		}
	}

	return "YES"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		s := readLine(reader)

		result := isBalanced(s)

		fmt.Fprintf(writer, "%s\n", result)
	}

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
