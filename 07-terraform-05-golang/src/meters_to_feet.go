package main

import "fmt"

func main() {
    fmt.Print("Enter value in meters: ")
    var input float64
    var multiplier = 3.28084
    fmt.Scanf("%f", &input)
    
    output := input * multiplier
    
    fmt.Println("Result in feet:", output)
}