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
    y, _ := strconv.Atoi(sc.Text())

    // （400で割り切れる）または（100で割り切れずかつ4で割り切ることができる）ならば閏年、そうでなければ平年
    if y % 400 == 0 || (y % 100 != 0 && y % 4 == 0) {
        fmt.Println("Yes")
    } else {
        fmt.Println("No")
    }
}

