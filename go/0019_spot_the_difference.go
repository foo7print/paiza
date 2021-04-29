package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    sc := bufio.NewScanner(os.Stdin)
    sc.Scan()
    s_1 := sc.Text()
    sc.Scan()
    s_2 := sc.Text()
    if s_1 == s_2 {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}

