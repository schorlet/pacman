package main

import "fmt"
import "strings"

func main() {
    var tests int
    fmt.Scanf("%d", &tests)

    for ; tests > 0; tests-- {
        var n int
        fmt.Scanf("%d", &n)

        var m, o = n % 3, 0

        for n > 0 && m != 0 {
            n -= 5
            m = n % 3
            o += 5
        }
        if n > 0 && m == 0 || n >= 0 && o > 0 {
            fmt.Printf("%s%s\n", strings.Repeat("5", n),
                                 strings.Repeat("3", o))

        } else {
            fmt.Println("-1")
        }
    }
}
