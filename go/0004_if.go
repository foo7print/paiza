package main

import "fmt"

func main(){
    var a, b string
    fmt.Scan(&a)
    fmt.Scan(&b)
    if a == b {
        fmt.Println("OK")
    } else {
        fmt.Println("NG")
    }
}

