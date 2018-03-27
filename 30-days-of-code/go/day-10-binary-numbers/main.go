package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var input int64
	fmt.Scanln(&input)

	binaryNumber := strconv.FormatInt(input, 2)
	var consecutiveNum, maxConsecutiveNum int

	for i := 0; i < len(binaryNumber); i++ {
		if string([]rune(binaryNumber)[i]) == "1" {
			consecutiveNum++
			maxConsecutiveNum = int(math.Max(float64(consecutiveNum), float64(maxConsecutiveNum)))
		} else {
			consecutiveNum = 0
		}
	}

	fmt.Print(maxConsecutiveNum)

}
