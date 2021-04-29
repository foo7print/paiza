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
    if n > 24 {
        i := n - 24
        fmt.Println(i)
    } else {
        fmt.Println(n)
    }
}

