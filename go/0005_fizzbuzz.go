package main

import "fmt"

func main(){
    var n int
    fmt.Scan(&n)
    for i := 1; i < n + 1; i++ {
        if i % 3 == 0 && i % 5 == 0 {
            fmt.Println("Fizz Buzz")
        } else if i % 3 == 0 {
            fmt.Println("Fizz")
        } else if i % 5 == 0 {
            fmt.Println("Buzz")
        } else {
            fmt.Println(i)
        }
    }
}

