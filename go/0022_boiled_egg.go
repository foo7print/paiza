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
    m, _ := strconv.Atoi(sc.Text())
    if 0 <= m && m < 6 {
        fmt.Println("raw")
    } else if 6 <= m  && m < 8 {
        fmt.Println("soft boiled")
    } else if 8 <= m && m <= 15 {
        fmt.Println("hard boiled")
    }
}

