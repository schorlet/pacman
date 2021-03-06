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
        var _, err = fmt.Scanf("%d", &agent)
        if err == nil {
            breadth_fs(agent, exits)
        } else {
            break
        }
    }
}

func breadth_fs(agent int, exits map[int]bool) {
    var queue = edges{edge{agent, agent}}
    var visited = make(map[edge]bool)

    for len(queue) > 0 {
        var link = queue[0]
        queue = queue[1:]
        visited[link] = true
        // debug(link)

        if exits[link.dst] {
            remove(link)
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

func remove(link edge) {
    var links = edges{}
    for _, link0 := range nodes[link.src] {
        if link != link0 {
            links = append(links, link0)
        }
    }
    nodes[link.src] = links
}

func debug(args ...interface{}) {
    fmt.Fprintln(os.Stderr, args)
}
