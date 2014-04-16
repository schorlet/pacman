package main

import "fmt"

func main() {
    var tests int
    fmt.Scanf("%d", &tests)

    var bill, price, wrappers int
    for ; tests > 0; tests-- {
        fmt.Scanf("%d %d %d", &bill, &price, &wrappers)
        var chocolates = bill / price
        var free = chocolates / wrappers

        var free2 = ((chocolates % wrappers) + free) / wrappers
        free += free2

        for free2 > wrappers {
            free2 = ((free2 % wrappers) + free2) / wrappers
            free += free2
        }
        fmt.Println(chocolates + free)
    }
}
