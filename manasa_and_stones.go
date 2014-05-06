package main

import "fmt"
import "os"
import "strings"
// import "sort"

func main() {
    var tests int
    fmt.Scanf("%d", &tests)

    for ; tests > 0; tests-- {
        var n, a, b int
        fmt.Scanln(&n)
        fmt.Scanln(&a)
        fmt.Scanln(&b)
        a, b = min(a, b)
        // totals = make(map[int]bool)
        // last_stones_recurse(0, n, a, b)
        // var sorted_stones = []int{}
        // for k, _ := range totals {
            // sorted_stones = append(sorted_stones, k)
        // }
        // sort.Ints(sorted_stones)
        // var out = fmt.Sprintf("%d ", sorted_stones)
        // fmt.Println(strings.Trim(out, "[ ]"))
        last_stones(n, a, b)
        debug()
    }
}

func last_stones(n, a, b int) {
    var i, j int
    var stones = make([][]int, n-1)
    stones[0] = []int{a, b}
    // debug(stones[0])
    for i = 1; i < n-1; i++ {
        stones[i] = make([]int, 2+i)
        for j = 0; j < 2+i-1; j++ {
            stones[i][j] = stones[i-1][j] + a
        }
        stones[i][j] = stones[i-1][j-1] + b
        // debug(stones[i])
    }
    var out = fmt.Sprintf("%d ", stones[n-2])
    out = strings.Trim(out, "[ ]")
    fmt.Println(out)
}

func min(a, b int) (int, int) {
    if a < b {
        return a, b
    }
    return b, a
}

var totals map[int]bool

func last_stones_recurse(node, n, a, b int) {
    if n == 0 {
        return
    } else if n == 1 {
        totals[node] = true
    }
    // debug(strings.Repeat(" ", 20-n), node)
    last_stones_recurse(node+a, n-1, a, b)
    last_stones_recurse(node+b, n-1, a, b)
}

func debug(args ...interface{}) {
    fmt.Fprintln(os.Stderr, args)
}
