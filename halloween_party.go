package main

import "fmt"

func main() {
    var tests int
    fmt.Scanf("%d", &tests)

    var cuts int
    for ; tests > 0; tests-- {
        fmt.Scanf("%d", &cuts)
        var h, v int
        h = cuts / 2
        v = h + cuts % 2
        fmt.Println(h * v)
    }
}
