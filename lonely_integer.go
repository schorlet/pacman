package main

import "fmt"

func main() {
    var i, n, q int
    fmt.Scanf("%d", &n)

    var counter = make(map[int]int)
    for i = 0; i < n; i++ {
        fmt.Scanf("%d", &q)
        counter[q] += 1
    }
    for q, i := range counter {
        if i == 1 {
            fmt.Println(q)
            break
        }
    }
}

