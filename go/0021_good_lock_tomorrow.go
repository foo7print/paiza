package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)

func main() {
    sc := bufio.NewScanner(os.Stdin)
    sc.Scan()
    n, _ := strconv.Atoi(sc.Text())
    for i := 0; i < n; i++ {
        sc.Scan()
        s := sc.Text()
        if s == "forward" {
            fmt.Println("Sunny")
        } else if s == "reverse" {
            fmt.Println("Rainy")
        } else if s == "sideways" {
            fmt.Println("Cloudy")
        }
    }
}

