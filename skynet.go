package main

import "fmt"
import "os"

type (
    edge struct {
        src, dst int
    }
    edges []edge
)

func (self edge) String() string {
    return fmt.Sprintf("%d %d", self.src, self.dst)
}

// edges starting from node
var nodes = make(map[int]edges)

func main() {
    var i, n, l, e int
    fmt.Scanf("%d %d %d", &n, &l, &e)

    for ; i < l; i++ {
        var a, b int
        fmt.Scanf("%d %d", &a, &b)
        nodes[a] = append(nodes[a], edge{a, b})
        nodes[b] = append(nodes[b], edge{b, a})
    }

    var exits = make(map[int]bool)
    for i = 0; i < e; i++ {
        var f int
        fmt.Scanf("%d", &f)
        exits[f] = true
    }

    for {
        var agent int
        fmt.Scanf("%d", &agent)
        bfs(agent, exits)
    }
}

func bfs(agent int, exits map[int]bool) {
    var queue = edges{edge{agent, agent}}
    var visited = make(map[edge]bool)

    for {
        var link = queue[0]
        queue = queue[1:]
        visited[link] = true
        // debug(link)

        if exits[link.dst] {
            fmt.Println(link)
            break
        }

        for _, next := range nodes[link.dst] {
            if visited[next] {
                continue
            }
            queue = append(queue, next)
        }
    }
}

func debug(args ...interface{}) {
    fmt.Fprintln(os.Stderr, args)
}
