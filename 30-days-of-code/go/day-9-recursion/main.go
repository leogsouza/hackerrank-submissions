package main

import "fmt"

func factorial(n int) int {
	if n == 1 {
		return n
	}

	return n * factorial(n-1)
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var number int
	fmt.Scan(&number)

	fmt.Println(factorial(number))

}
