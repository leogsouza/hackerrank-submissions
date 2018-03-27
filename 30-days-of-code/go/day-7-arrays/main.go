package main

import "fmt"

func scanIntLn(n int) ([]int, error) {
	x := make([]int, n)
	y := make([]interface{}, len(x))
	for i := range x {
		y[i] = &x[i]
	}
	n, err := fmt.Scanln(y...)
	x = x[:n]
	return x, err
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var countItems int

	fmt.Scanf("%d", &countItems)
	items, err := scanIntLn(countItems)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := len(items); i > 0; i-- {
		fmt.Printf("%d ", items[i-1])
	}
}
