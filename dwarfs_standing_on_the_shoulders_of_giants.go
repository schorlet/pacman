package main

import "fmt"
import "os"

var dwarfs = make(map[int][]int)
var influences = make(map[int]int)

func main() {
    var n int
    fmt.Scanf("%d", &n)

    for i := 0; i < n; i++ {
        var x, y int
        fmt.Scanf("%d %d", &x, &y)
        dwarfs[x] = append(dwarfs[x], y)
    }

    var dist int
    for dwarf, _ := range dwarfs {
        dist = max(dist, 1 + count_influences(dwarf))
    }

    fmt.Println(dist)
}

func count_influences(dwarf int) int {
    var dist, ok = influences[dwarf]
    if ok {
        return dist
    }

    var others = dwarfs[dwarf]
    for _, other := range others {
        dist = max(dist, 1 + count_influences(other))
    }

    influences[dwarf] = dist
    return dist
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
