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

    // for i = 0; i < m; i++ {
        // debug(C[i])
    // }

    return C[m-1][n-1]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func debug(args ...interface{}) {
    fmt.Fprintln(os.Stderr, args)
}
