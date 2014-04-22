package main

import "fmt"
import "math"

func main() {
    var a, b float64
    fmt.Scanf("%f", &a)
    fmt.Scanf("%f", &b)

    var c = a + b - 1
    var d, e = math.Min(a, b), math.Max(a, b)

    if e < 25 {
        fmt.Println(0)

    } else if e == 25 && d == 24 {
        fmt.Println(0)

    } else if e > 25 && e - d != 2 {
        fmt.Println(0)

    } else {
        var f = math.Pow10(9) + 7.
        var g = factorial(c) / (factorial(d) * factorial(c-d))

        if g > f {
            fmt.Println(uint(math.Mod(g, f)))
        } else {
            fmt.Println(uint(g))
        }
    }
}

func factorial(a float64) float64 {
    var b float64 = 1.
    for ; a > 1; a-- {
        b *= a
    }
    return b
}

