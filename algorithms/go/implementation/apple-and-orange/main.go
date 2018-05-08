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
 * Complete the countApplesAndOranges function below.
 */
func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {

	var totalApples, totalOranges int32
	for _, apple := range apples {
		pos := a + apple
		if pos >= s && pos <= t {
			totalApples++
		}
	}

	for _, orange := range oranges {
		pos := b + orange
		if pos >= s && pos <= t {
			totalOranges++
		}
	}

	fmt.Println(totalApples)
	fmt.Println(totalOranges)

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	st := strings.Split(readLine(reader), " ")

	sTemp, err := strconv.ParseInt(st[0], 10, 64)
	checkError(err)
	s := int32(sTemp)

	tTemp, err := strconv.ParseInt(st[1], 10, 64)
	checkError(err)
	t := int32(tTemp)

	ab := strings.Split(readLine(reader), " ")

	aTemp, err := strconv.ParseInt(ab[0], 10, 64)
	checkError(err)
	a := int32(aTemp)

	bTemp, err := strconv.ParseInt(ab[1], 10, 64)
	checkError(err)
	b := int32(bTemp)

	mn := strings.Split(readLine(reader), " ")

	mTemp, err := strconv.ParseInt(mn[0], 10, 64)
	checkError(err)
	m := int32(mTemp)

	nTemp, err := strconv.ParseInt(mn[1], 10, 64)
	checkError(err)
	n := int32(nTemp)

	appleTemp := strings.Split(readLine(reader), " ")

	var apple []int32

	for appleItr := 0; appleItr < int(m); appleItr++ {
		appleItemTemp, err := strconv.ParseInt(appleTemp[appleItr], 10, 64)
		checkError(err)
		appleItem := int32(appleItemTemp)
		apple = append(apple, appleItem)
	}

	orangeTemp := strings.Split(readLine(reader), " ")

	var orange []int32

	for orangeItr := 0; orangeItr < int(n); orangeItr++ {
		orangeItemTemp, err := strconv.ParseInt(orangeTemp[orangeItr], 10, 64)
		checkError(err)
		orangeItem := int32(orangeItemTemp)
		orange = append(orange, orangeItem)
	}

	countApplesAndOranges(s, t, a, b, apple, orange)
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
