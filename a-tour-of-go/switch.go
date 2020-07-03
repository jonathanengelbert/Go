package main

import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Printf("This script is running on ")
    switch os := runtime.GOOS; os {
        case "darwin":
            fmt.Printf("OS X.")
        case "linux":
            fmt.Printf("Linux.")
        default:
            fmt.Printf("%s. \n", os)
    }
    fmt.Printf("\n\n")
}