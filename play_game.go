package main

import "bufio"
import "fmt"
import "os"
import "strconv"
import "strings"

func main() {
    var in = bufio.NewReader(os.Stdin)
    var tests int
    fmt.Fscanf(in, "%d\n", &tests)

    for ; tests > 0; tests-- {
        var i, n int
        fmt.Fscanf(in, "%d\n", &n)

        var line, _ = in.ReadString('\n')
        line = strings.TrimSpace(line)

        var splits = strings.Split(line, " ")
        var values = make([]int, n)

        for i = 0; i < n; i++ {
            values[i], _ = strconv.Atoi(splits[i])
        }
        if n < 3 {
            fmt.Println(sum(values))
        } else {
            fmt.Println(max_range(values))
        }
    }
}

func sum(values []int) int {
    var c int
    for _, v := range values {
        c+= v
    }
    return c
}
func max_range(values []int) int {
    var n = len(values)
    var c = make([]int, n)

    c[n-1] = values[n-1]
    c[n-2] = values[n-1] + values[n-2]
    c[n-3] = values[n-1] + values[n-2] + values[n-3]

    var cc = c[n-3]
    for i := n - 4; i >= 0; i-- {
        cc += values[i]
        c[i] = cc - min3(c[i+1], c[i+2], c[i+3])
    }
    // debug(values)
    // debug(c)
    // debug()
    return c[0]
}

func min3(a, b, c int) int {
    var d = min2(a, b)
    var e = min2(b, c)
    return min2(d, e)
}

func min2(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func debug(args ...interface{}) {
    fmt.Fprintln(os.Stderr, args)
}
