package main

import "fmt"
import "bufio"
import "os"

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    t := scanner.Text()
    fmt.Println(t)
}

