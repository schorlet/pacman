package main

import "bufio"
import "fmt"
import "os"
import "sort"
import "strconv"
import "strings"

var nodes []int
var edges [][]int
var weights []int

func dfs(q int, visited []bool) int {
    if visited[q] {
        return 0
    }
    visited[q] = true
    // leaf node
    if len(edges[q]) == 1 {
        weights[q] = nodes[q]
        return weights[q]
    }
    // union join node
    var weight = nodes[q]
    for _, node := range edges[q] {
        weight += dfs(node, visited)
    }
    weights[q] = weight
    return weights[q]
}

func main() {
    var i, n int
    var in = bufio.NewReader(os.Stdin)
    fmt.Fscanf(in, "%d\n", &n)
    n += 1

    var total int
    nodes = make([]int, n)
    edges = make([][]int, n)
    weights = make([]int, n)

    var line = readline(in)
    var split = strings.Split(line, " ")

    for i = 1; i < n; i++ {
        nodes[i], _ = strconv.Atoi(split[i-1])
        total += nodes[i]
    }

    // n vertices, n-1 edges
    for i = 1; i < n-1; i++ {
        var a, b int
        fmt.Fscanf(in, "%d %d\n", &a, &b)
        edges[a] = append(edges[a], b)
        edges[b] = append(edges[b], a)
    }

    var root int
    for i, rel := range edges {
        if len(rel) > 1 {
            root = i
            break
        }
    }

    var visited = make([]bool, n)
    dfs(root, visited)
    // debug_node(0, root, 0)

    for i = 1; i < n; i++ {
        // diff between two trees:
        var t1, t2 = total - weights[i], weights[i]
        weights[i] = abs(t1 - t2)
    }
    // debug_node(0, root, 0)
    sort.Ints(weights)
    fmt.Println(weights[1])
}

func abs(a int) int {
    if a < 0 {
        a = -a
    }
    return a
}

func readline(in *bufio.Reader) string {
    var line, _ = in.ReadString('\n')
    line = strings.TrimSpace(line)
    return line
}

func debug_node(p, q int, level int) {
    debug(strings.Repeat("  ", level),
        fmt.Sprintf("%d:%d,%d", q, nodes[q], weights[q]))
    for _, node := range edges[q] {
        if p != node {
            debug_node(q, node, level+1)
        }
    }
}

func debug(args ...interface{}) {
    fmt.Fprintln(os.Stderr, args)
}
