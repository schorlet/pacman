package main

import "fmt"

func main() {
    var n, m int
    fmt.Scanf("%d", &n)

    for ; n > 0; n-- {
        var height = 1
        fmt.Scanf("%d", &m)

        for i := 1; i <= m; i++ {
            if i % 2 == 0 {
                height += 1
            } else {
                height *= 2
            }
        }
        fmt.Println(height)
    }
}
