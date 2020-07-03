package main

import (
    "fmt"
    "time"
)

func greet() string {
   t := time.Now()

   switch {
    case t.Hour() < 12:
        return "Good Morning!"

    case t.Hour() < 17:
        return "Good Afternoon!"

    case t.Hour() > 17 && t.Hour() < 0:
        return "Good Evening!"

    default:
        return "Fuck off"
   }
}


func main() {
    fmt.Println(greet())
}