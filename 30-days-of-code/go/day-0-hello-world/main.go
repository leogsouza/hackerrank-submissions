package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    inputreader := bufio.NewReader(os.Stdin)
        input, _ := inputreader.ReadString('\n')
    
    fmt.Println("Hello, World.")
    
    fmt.Println(input)
}