package main

import "fmt"
import "sort"

func main() {
    var i, flowers, friends int
    fmt.Scanf("%d %d", &flowers, &friends)

    var prices = make([]int, flowers)
    for i = 0; i < flowers; i++ {
        fmt.Scanf("%d", &prices[i])
    }
    sort.Ints(prices)
    var costs = prices[0:flowers]
    for i = 0; i < flowers / 2; i++ {
        costs[i], costs[flowers-i-1] = costs[flowers-i-1], costs[i]
    }
    var total, f = 0, friends
    for i = 0; i < flowers; i++ {
        total += costs[i] * (f / friends)
        f += 1
    }
    fmt.Println(total)
}

