package main

// imports should be grouped as below, but can use
// several import statements
import (
    "fmt"
    "math"
)

func main() {
    fmt.Printf("Now you have %g problems.\n",
    math.Sqrt(7))
    // Println method handling one or more arguments
    fmt.Println("\nYo")
    fmt.Println("Stuff", "MOre stuff")
    // Using print on a slice
    items := []int{10, 20, 30}
    fmt.Println(items)
    items2 := []string{"Jon", "Max"}
    fmt.Println(items2)
}