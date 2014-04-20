package main

import "fmt"
import "math"
import "os"
import "sort"

var prime_numbers = []int{2, 3, 5, 7}

// total number of ways in which the bricks can be arranged on the wall
// of size 4xnumber with bricks of size 4x1 and 1x4.
var bricks [42]int

func gen_bricks() {
    var i int
    for ; i < 5; i++ {
        bricks[i] = 1
    }
    for ; i < 42; i++ {
        bricks[i] = bricks[i - 1] + bricks[i - 4]
    }
}

func gen_primes(count int) {
    var i, p int
    var l = len(prime_numbers)
    var n = prime_numbers[l - 1]

    for i = n + 2; l < count; i += 2 {
        var sq = int(math.Sqrt(float64(i)) + 1)
        for _, p = range prime_numbers {
            if sq < p {
                prime_numbers = append(prime_numbers, i)
                l += 1
                break
            } else if i % p == 0 {
                break
            }
        }
    }
}

// number of prime numbers (P) up to M (i.e. <= M)
func primes(m int) int {
    var n int
    var l = len(prime_numbers)
    for {
        n = sort.SearchInts(prime_numbers, m)
        if n < l {
            if prime_numbers[n] == m {
                n += 1
            }
            break
        } else {
            l *= 2
            gen_primes(l)
        }
    }
    return n
}

func main() {
    var tests, number int
    gen_bricks()

    fmt.Scanf("%d", &tests)
    for ; tests > 0; tests-- {
        fmt.Scanf("%d", &number)
        var m = bricks[number + 1]
        var p = primes(m)
        // debug(number, m, p)
        fmt.Println(p)
    }
}

func debug(args ...interface{}) {
    fmt.Fprintln(os.Stderr, args)
}
