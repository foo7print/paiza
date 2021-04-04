package main

import (
    "fmt"
    "sort"
)

func main() {
    var n_1, n_2, n_3, n_4, n_5 int
    fmt.Scan(&n_1)
    fmt.Scan(&n_2)
    fmt.Scan(&n_3)
    fmt.Scan(&n_4)
    fmt.Scan(&n_5)
    
    a := []int { n_1, n_2, n_3, n_4, n_5 }
    sort.Ints(a)
    fmt.Println(a[0])
}

