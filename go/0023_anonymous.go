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
    fmt.Println(list[0][0:1] + "." + list[1][0:1] + ".")
}

