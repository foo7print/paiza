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
    i, _ := strconv.Atoi(sc.Text())
    if i >= 80 {
        fmt.Println("pass")
    } else {
        fmt.Println("fail")
    }
}

