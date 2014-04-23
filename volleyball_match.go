package main

import "fmt"
import "math"
import "math/big"

func main() {
    var a, b int64
    fmt.Scanf("%d", &a)
    fmt.Scanf("%d", &b)

    var d, e int64
    // var d, e = math.Min(a, b), math.Max(a, b)
    if a < b {
        d = a
        e = b
    } else {
        d = b
        e = a
    }

    if e < 25 {
        fmt.Println(0)

    } else if e == 25 && d >= 24 {
        fmt.Println(0)

    } else if e > 25 && e - d != 2 {
        fmt.Println(0)

    } else {
        // var g = factorial(c) / (factorial(d) * factorial(c-d))
        // http://fr.wikipedia.org/wiki/Loi_binomiale
        var c = a + b - 1
        if e > 25 {
            c -= 1
        }
        var f = big.NewInt(int64(math.Pow10(9) + 7))
        var z3 = big.NewInt(1).Binomial(c, d)
        if z3.Cmp(f) == 1 {
            fmt.Println(z3.Mod(z3, f))
        } else {
            fmt.Println(z3.Abs(z3))
        }
    }
}


// func factorial(n *big.Int) *big.Int {
    // var result = big.NewInt(1)
    // var one = big.NewInt(1)
    // for n.Cmp(one) == 1 {
        // result.Mul(result, n)
        // n.Sub(n, one)
    // }
    // return result
// }


// func factorial(a float64) float64 {
    // var b float64 = 1.
    // for ; a > 1; a-- {
        // b *= a
    // }
    // return b
// }

