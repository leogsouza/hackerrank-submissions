package main

import "fmt"

// import "bufio"
// import "os"

func transformText(text string) string {
	var evenString, oddString string
	for k, v := range text {
		if k%2 == 0 {
			evenString += string(v)
		} else {
			oddString += string(v)
		}
	}

	return evenString + " " + oddString
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	// scanner = bufio.NewScanner(os.Stdin)
	var input int
	var input2 string
	fmt.Scanf("%d", &input)

	for i := 0; i < int(input); i++ {
		fmt.Scanf("%s", &input2)
		fmt.Println(transformText(input2))
	}
}
