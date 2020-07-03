package main

import "fmt"

func concatenate(x string, y string) string {
  return x + y
}

func add(x int, y int) int {
  return x + y
}

// types can be defined at the end of arguments,
// provided all arguments are the same type

func swap(a, b string) (string, string) {
    return b, a
}

func main(){
  fmt.Println(concatenate("Jon", "athan"))
  fmt.Println(add(20, 39))
  a, b := swap("Engelbert", "Jonathan")
  fmt.Println(a, b)
}


