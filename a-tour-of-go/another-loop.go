package main

import (
    "fmt"
)

func counter(target, maxCount int) int {
    var result int
    for i := 0; i < maxCount; i++ {
        if i == target {
            return result
        }
        result = i
    }
    return result
}

func main() {
    fmt.Println(counter(90238100, 1200000323))
    fmt.Println("ENDED")
}