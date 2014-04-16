package main

import "fmt"
import "sort"

func main() {
    var i, packets, children int
    fmt.Scanf("%d", &packets)
    fmt.Scanf("%d", &children)

    var candies = make([]int, packets)
    for i = 0; i < packets; i++ {
        fmt.Scanf("%d", &candies[i])
    }
    sort.Ints(candies)

    var unfairness = candies[children - 1] - candies[0]
    for i = 1; i < packets - children; i++ {
        if unfairness > candies[i + children - 1] - candies[i] {
            unfairness = candies[i + children - 1] - candies[i]
        }
    }
    fmt.Println(unfairness)
}
