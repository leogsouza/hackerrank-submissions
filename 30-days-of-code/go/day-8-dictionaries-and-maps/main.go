package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PhoneBook struct {
	Name   string
	Number int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	var countInput string = scanner.Text()

	totalRecords, err := strconv.ParseInt(countInput, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	var phoneBook map[string]string
	phoneBook = make(map[string]string)
	for i := 0; i < int(totalRecords); i++ {
		scanner.Scan()
		var countInput []string = strings.Split(scanner.Text(), " ")
		if len(countInput) == 2 {
			var name, phone string = countInput[0], countInput[1]
			phoneBook[name] = phone
		}
	}
	for i := 0; i < int(totalRecords); i++ {
		scanner.Scan()
		var name string = scanner.Text()
		if name != "" {
			if phoneNumber, ok := phoneBook[name]; ok {
				fmt.Printf("%s=%s\n", name, phoneNumber)
			} else {
				fmt.Println("Not found")
			}
		}
	}
}
