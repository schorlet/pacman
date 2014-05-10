package main

import "container/heap"
import "bufio"
import "fmt"
import "os"

type road struct {
    src, dest int
    fare int64
}
type roads []road

func (self road) String() string {
    return fmt.Sprintf("%d:%d", self.dest, self.fare)
}

func (pq roads) Len() int {
    return len(pq)
}
func (pq roads) Less(i, j int) bool {
    return pq[i].fare < pq[j].fare
}
func (pq roads) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}
func (pq *roads) Push(x interface{}) {
    var item = x.(road)
    *pq = append(*pq, item)
}
func (pq *roads) Pop() interface{} {
    var old = *pq
    var n = len(old)
    var item = old[n-1]
    *pq = old[0 : n-1]
    return item
}

func main() {
    var in = bufio.NewReader(os.Stdin)
    var i, n, e int
    fmt.Fscanf(in, "%d %d\n", &n, &e)

    // roads starting from city
    var sources = make(map[int]roads)

    for ; i < e; i++ {
        var a, b int
        var c int64
        fmt.Fscanf(in, "%d %d %d\n", &a, &b, &c)
        sources[a] = append(sources[a], road{a, b, c})
        sources[b] = append(sources[b], road{b, a, c})
    }
    // for src, dests := range sources {
        // debug(src, dests)
    // }
    find_city(n, sources)
}

func find_city(goal int, sources map[int]roads) {
    var nodes = roads{}
    heap.Init(&nodes)
    heap.Push(&nodes, road{0, 1, 0})
    var visited = make(map[int]bool)

    for nodes.Len() > 0 {
        var road0 = heap.Pop(&nodes).(road)
        visited[road0.src] = true
        debug(road0)

        if road0.dest == goal {
            fmt.Println(road0.fare)
            return
        }

        var dests = sources[road0.dest]
        for _, road1 := range dests {
            if visited[road1.src] {
                continue
            }
            var cost = road0.fare + max(0, road1.fare - road0.fare)
            var road2 = road{road1.src, road1.dest, cost}
            heap.Push(&nodes, road2)
        }
    }
    fmt.Println("NO PATH EXISTS")
}

func max(a, b int64) int64 {
    if a > b {
        return a
    }
    return b
}

func debug(args ...interface{}) {
    fmt.Fprintln(os.Stderr, args)
}
