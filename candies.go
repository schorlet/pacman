package main

import "fmt"
import "os"

func main() {
    var i, n int
    fmt.Scanf("%d", &n)

    var ratings = make([]int, n)
    var candies = make([]int, n)

    for ; i < n; i++ {
        fmt.Scanf("%d", &ratings[i])
        candies[i] = 1
    }

    for i = n - 2; i >= 0; i-- {
        if ratings[i] > ratings[i + 1] {
            // candies[i] = max(candies[i], 1 + candies[i + 1])
            candies[i] = 1 + candies[i + 1]
        }
    }

    var total = candies[0]
    for i = 1; i < n; i++ {
        if ratings[i] > ratings[i - 1] {
            candies[i] = max(candies[i], 1 + candies[i - 1])
        }
        total += candies[i]
    }

    // debug(candies)
    fmt.Println(total)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func debug(args ...interface{}) {
    fmt.Fprintln(os.Stderr, args)
}
