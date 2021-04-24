package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    var n, _ = strconv.Atoi(scanner.Text())
    for i := 0; i < n; i++ {
        scanner.Scan()
        fmt.Println(scanner.Text())
    }
}

