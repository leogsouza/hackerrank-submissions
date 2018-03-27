package main

import "fmt"
import "bufio"
import "os"
import "strconv"
import "math"

func round(f float64) int {
    if math.Abs(f) < 0.5 {
        return 0
    }
    return int(f + math.Copysign(0.5, f))
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()    
    mealCost, _ := strconv.ParseFloat(scanner.Text(), 64)
    
    scanner.Scan()
    tipValue, _ := strconv.ParseFloat(scanner.Text(), 64)
    tipPercent := mealCost * ( tipValue/ 100)
    
    scanner.Scan()
    taxValue, _ := strconv.ParseFloat(scanner.Text(), 64)
    taxPercent:= mealCost * ( taxValue / 100)
    
    totalCost := mealCost + tipPercent + taxPercent
    
    fmt.Printf("The total meal cost is %d dollars.", round(totalCost))
}