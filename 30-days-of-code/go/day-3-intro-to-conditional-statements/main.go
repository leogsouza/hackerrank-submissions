package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var input int
	fmt.Scanf("%d", &input)

	if input%2 != 0 {
		fmt.Println("Weird")
	} else if input >= 2 && input <= 5 {
		fmt.Println("Not Weird")
	} else if input >= 6 && input <= 20 {
		fmt.Println("Weird")
	} else {
		fmt.Println("Not Weird")
	}
}
