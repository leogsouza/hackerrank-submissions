package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the csRush function below.
func csRush(n int32, m int32, css [][]int32, customers [][]int32, vacant_css []int32) int32 {

	// Creating map for customers
	mapCss := make(map[int32]int32)
	for _, cs := range css {
		mapCss[cs[0]] = cs[1]
	}

	// Remove customers on vacation
	for _, vc := range vacant_css {
		delete(mapCss, vc)
	}

	mapCustomers := make(map[int32]int32)

	for _, customer := range customers {
		mapCustomers[customer[0]] = customer[1]
	}

	mapCssAttend := make(map[int32]int32)
	mapCustomerAttended := make(map[int32]bool)

	for k, csCapacity := range mapCss {
		for kc, customerSize := range mapCustomers {
			_, ok := mapCustomerAttended[kc]
			if csCapacity >= customerSize && !ok {
				mapCssAttend[k]++
				mapCustomerAttended[kc] = true
			}
		}
	}

	var max, customerId int32

	for k, totalAttend := range mapCssAttend {
		if totalAttend > max {
			max = totalAttend
			customerId = k
		}
	}

	return customerId

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var css [][]int32
	for i := 0; i < int(n); i++ {
		cssRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var cssRow []int32
		for _, cssRowItem := range cssRowTemp {
			cssItemTemp, err := strconv.ParseInt(cssRowItem, 10, 64)
			checkError(err)
			cssItem := int32(cssItemTemp)
			cssRow = append(cssRow, cssItem)
		}

		if len(cssRow) != 2 {
			panic("Bad input")
		}

		css = append(css, cssRow)
	}

	mTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	m := int32(mTemp)

	var customers [][]int32
	for i := 0; i < int(m); i++ {
		customersRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var customersRow []int32
		for _, customersRowItem := range customersRowTemp {
			customersItemTemp, err := strconv.ParseInt(customersRowItem, 10, 64)
			checkError(err)
			customersItem := int32(customersItemTemp)
			customersRow = append(customersRow, customersItem)
		}

		if len(customersRow) != 2 {
			panic("Bad input")
		}

		customers = append(customers, customersRow)
	}

	vacant_cssCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	vacant_cssTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var vacant_css []int32

	for i := 0; i < int(vacant_cssCount); i++ {
		vacant_cssItemTemp, err := strconv.ParseInt(vacant_cssTemp[i], 10, 64)
		checkError(err)
		vacant_cssItem := int32(vacant_cssItemTemp)
		vacant_css = append(vacant_css, vacant_cssItem)
	}

	cs_distribution := csRush(n, m, css, customers, vacant_css)

	fmt.Fprintf(writer, "%d\n", cs_distribution)

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
