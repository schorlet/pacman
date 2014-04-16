package main

import "fmt"
import "sort"

type pair struct {
    a, b int
}

func main() {
    var i, j, n int
    fmt.Scanf("%d", &n)

    var numbers = make([]int, n)
    for i = 0; i < n; i++ {
        fmt.Scanf("%d", &numbers[i])
    }
    sort.Ints(numbers)

    var closest int = 1 << 31 - 1
    var pairs = []pair{}

    for i = 0; i < n - 1; i++ {
        for j = i + 1; j < n; j++ {
            var diff = numbers[j] - numbers[i]
            if diff < closest {
                pairs = []pair{pair{numbers[i], numbers[j]}}
                closest = diff
            } else if diff == closest {
                pairs = append(pairs, pair{numbers[i], numbers[j]})
            } else {
                break
            }
        }
    }
    for _, pair := range pairs {
        fmt.Printf("%d %d ", pair.a, pair.b)
    }
    fmt.Println()
}
