package main

import "fmt"

// keyword var is needed if statement is outside function body
var i, j = 1, "No"

func main() {
   // short assignment is available inside function body
   p := "Yo Dog"
   fmt.Println(i, j)
   fmt.Println(p)
}