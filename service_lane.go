package main

import "fmt"

func main() {
    var i, j int
    var length, tests int
    fmt.Scanf("%d %d", &length, &tests)

    var freeway = make([]int, length)
    for i = 0; i < length; i++ {
        fmt.Scanf("%d", &freeway[i])
    }

    var begin, end int
    for i = 0; i < tests; i++ {
        fmt.Scanf("%d %d", &begin, &end)
        var largest = 3
        for j = begin; j <= end; j++ {
            if freeway[j] < largest {
                largest = freeway[j]
                if largest == 1 {
                    break
                }
            }
        }
        fmt.Println(largest)
    }
}
