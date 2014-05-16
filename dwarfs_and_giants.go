package main

import "fmt"
import "os"

var relations = make(map[int][]int)
var mem = make(map[int]int)

func main() {
    var i, n int
    fmt.Scanf("%d", &n)

    for i = 0; i < n; i++ {
        var a, b int
        fmt.Scanf("%d %d", &a, &b)
        relations[a] = append(relations[a], b)
    }
    debug(relations)

    var depth int
    for node, _ := range relations {
        depth = max(depth, max_depth(node))
    }
    debug(mem)
    fmt.Println(depth + 1)
}

func max_depth(node int) int {
    var depth, ok = mem[node]
    if ok {
        return depth
    }
    for _, rel := range relations[node] {
        depth = max(depth, 1 + max_depth(rel))
    }
    mem[node] = depth
    return depth
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
