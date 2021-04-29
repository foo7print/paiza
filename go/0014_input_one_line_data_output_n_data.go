package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)

func main(){
    sc := bufio.NewScanner(os.Stdin)
    sc.Scan()
    var n, _ = strconv.Atoi(sc.Text())
    sc.Scan()
    list := strings.Split(sc.Text(), " ")
    for i := 0; i < n; i++ {
        fmt.Println(list[i])
    }
}

