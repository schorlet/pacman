package main

import "fmt"
import "os"

func main() {
    var a, b string
    fmt.Scanf("%s", &a)
    fmt.Scanf("%s", &b)
    fmt.Println(lcs(a, b))
}

func lcs(a, b string) int {
    var m, n = len(a) + 1, len(b) + 1
    var i, j int

    var C = make([][]int, m)
    for i = 0; i < m; i++ {
        C[i] = make([]int, n)
    }

    for i = 1; i < m; i++ {
        for j = 1; j < n; j++ {
            if a[i-1] == b[j-1] {
                C[i][j] = C[i-1][j-1] + 1
            } else {
                C[i][j] = max(C[i][j-1], C[i-1][j])
            }
        }
    }

    for i = 0; i < m; i++ {
        debug(C[i])
    }

    var s = make([]uint8, C[m-1][n-1])
    var k int

    for i = m - 1; i > 0; {
        for j = n - 1; i > 0 && j > 0; {
            debug(i, j)

            if C[i][j] == C[i-1][j] {
                i -= 1
            } else if C[i][j] == C[i][j-1] {
                j -= 1
            } else {
                i -= 1
                j -= 1
                s[k] = a[i]
                k += 1
            }
        }
    }
    reverse(s)
    debug(fmt.Sprintf("%s", s))
    return C[m-1][n-1]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func reverse(list []uint8) {
    var l = len(list)
    for i := 0; i < l/2; i++ {
        list[i], list[l-1-i] = list[l-1-i], list[i]
    }
}

func debug(args ...interface{}) {
    fmt.Fprintln(os.Stderr, args)
}
