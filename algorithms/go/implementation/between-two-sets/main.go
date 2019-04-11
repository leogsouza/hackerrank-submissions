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
 * Complete the getTotalX function below.
 */
func getTotalX(a []int32, b []int32) int32 {
	var count, it int32 = 0, 1
	var lcmNum = a[0]
	var gcdNum = b[0]

	for i := 1; i < len(a); i++ {
		lcmNum = lcm(lcmNum, a[i])
	}

	for i := 1; i < len(b); i++ {
		gcdNum = gcd(gcdNum, b[i])
	}

	for it*lcmNum <= gcdNum {
		if gcdNum%(lcmNum*it) == 0 {
			count++
		}
		it++
	}
	return count

}

func lcm(a, b int32) int32 {

	var max, min int32
	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}

	for i := max; ; i += max {
		if i%min == 0 {
			return i
		}
	}
}

func gcd(a, b int32) int32 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer outputFile.Close()

	writer := bufio.NewWriterSize(outputFile, 1024*1024)

	nm := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(nm[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	aTemp := strings.Split(readLine(reader), " ")

	var a []int32

	for aItr := 0; aItr < int(n); aItr++ {
		aItemTemp, err := strconv.ParseInt(aTemp[aItr], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	bTemp := strings.Split(readLine(reader), " ")

	var b []int32

	for bItr := 0; bItr < int(m); bItr++ {
		bItemTemp, err := strconv.ParseInt(bTemp[bItr], 10, 64)
		checkError(err)
		bItem := int32(bItemTemp)
		b = append(b, bItem)
	}

	total := getTotalX(a, b)

	fmt.Fprintf(writer, "%d\n", total)

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
