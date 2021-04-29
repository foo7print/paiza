package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func main() {
    sc := bufio.NewScanner(os.Stdin)
    sc.Scan()
    list := strings.Split(sc.Text(), " ")
    fmt.Println(list[0])
    fmt.Println(list[1])
    fmt.Println(list[2])
}

