package main

import "fmt"
import "math"
import "os"

func main() {
    var i, n, p int
    fmt.Scanf("%d", &n)
    fmt.Scanf("%d", &p)
    // debug(n , p)

    var elements = Range(1, p + 1)
    // debug("elements", elements)

    var productn = make([][]int, n)
    for i = 0; i < n; i++ {
        productn[i] = elements
    }
    // debug("productn", productn)

    var products = Product(p, productn...)
    // for i, product := range products {
        // if i % p == 0 {
            // fmt.Fprintln(os.Stderr)
        // }
        // fmt.Fprint(os.Stderr, product)
    // }
    // fmt.Fprintln(os.Stderr)
    // debug("products", len(products))
    // fmt.Println(len(products))
    fmt.Println(products)
    // 2 2: 3
    // 3 4: 4*4+6=22
    // 3 3: 3*3+2=11
    // 3 5: 5*5+7=32
    // 4 3: 3*5+6=21
    // 2 5: 10
}

func debug(args ...interface{}) {
    fmt.Fprintln(os.Stderr, args)
}

func Product(p int, args ...[]int) uint64 {
    pools := args
    npools := len(pools)
    indices := make([]int, npools)

    result := make([]int, npools)
    for i := range result {
        if len(pools[i]) == 0 {
            return 0
        }
        result[i] = pools[i][0]
    }

    // results := [][]int{result}
    var results uint64 = 1
    var f = uint64(math.Pow10(9) + 7)

    for {
        i := npools - 1
        for ; i >= 0; i -= 1 {
            pool := pools[i]
            indices[i] += 1

            if indices[i] == len(pool) {
                // debug("i:", i, "indices:", indices, result, ";")
                indices[i] = 0
                result[i] = pool[0]

            } else {
                result[i] = pool[indices[i]]
                // debug("i:", i, "indices:", indices, result)

                if i > 0 && result[i] * result[i-1] > p {
                    indices[i] = 0
                    result[i] = pool[0]
                    continue
                } else {
                    break
                }
            }
        }

        if i < 0 {
            return results % f
        }

        // newresult := make([]int, npools)
        // copy(newresult, result)
        // results = append(results, newresult)
        results += 1
    }

    return 0
}

func Range(start, stop int) []int {
    var list = make([]int, stop-start)
    for i := 0; i < stop-start; i++ {
        list[i] = start + i
    }
    return list
}
