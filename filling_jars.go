package main

import "fmt"

func main() {
    var jars, operations uint
    fmt.Scanf("%d %d", &jars, &operations)

    var i, candys uint
    for i = 0; i < operations; i++ {
        var a, b, k uint
        fmt.Scanf("%d %d %d", &a, &b, &k)
        candys += (b + 1 - a) * k
    }
    fmt.Println(candys / jars)
}
