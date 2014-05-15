package main

import "container/heap"
import "bufio"
import "fmt"
import "os"
// import "sort"

type road struct {
    src, dest int
    fare      int64
}
type roads []road

type city struct {
    id    int
    cost  int64
    index int
}
type cities []*city

type byCost struct {
    dests []int
    costs map[int]int64
}

func (bc byCost) Len() int {
    return len(bc.dests)
}
func (bc byCost) Swap(i, j int) {
    bc.dests[i], bc.dests[j] = bc.dests[j], bc.dests[i]
}
func (bc byCost) Less(i, j int) bool {
    var dest0, dest1 = bc.dests[i], bc.dests[j]
    return bc.costs[dest0] < bc.costs[dest1]
}

func (self road) String() string {
    return fmt.Sprintf("%d:%d=%d", self.src, self.dest, self.fare)
}
func (self city) String() string {
    return fmt.Sprintf("%d:%d (%d)", self.id, self.cost, self.index)
}
func newCity(id int, cost int64) *city {
    return &city{id: id, cost: cost}
}

func (pq cities) Len() int {
    return len(pq)
}
func (pq cities) Less(i, j int) bool {
    return pq[i].cost < pq[j].cost
}
func (pq cities) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
    pq[i].index = i
    pq[j].index = j
}
func (pq *cities) Push(x interface{}) {
    var n = len(*pq)
    var item = x.(*city)
    item.index = n
    *pq = append(*pq, item)
}
func (pq *cities) Pop() interface{} {
    var old = *pq
    var n = len(old)
    var item = old[n-1]
    item.index = -1 // for safety
    *pq = old[0 : n-1]
    return item
}
func (pq *cities) update(item *city) {
    heap.Remove(pq, item.index)
    heap.Push(pq, item)
}

func main() {
    var in = bufio.NewReader(os.Stdin)
    var i, goal, e int
    fmt.Fscanf(in, "%d %d\n", &goal, &e)

    // roads starting from city
    var sources = make(map[int]roads)

    for ; i < e; i++ {
        var a, b int
        var c int64
        fmt.Fscanf(in, "%d %d %d\n", &a, &b, &c)
        sources[a] = append(sources[a], road{a, b, c})
        sources[b] = append(sources[b], road{b, a, c})
    }
    for src, dests := range sources {
        debug(src, dests)
    }
    prim_mst(1, goal, sources)
}

// prim's minimum spanning tree
func prim_mst(start, goal int, sources map[int]roads) {
    var nodes = cities{}
    heap.Init(&nodes)

    // cumulated costs
    var costs = make(map[int]int64)
    var visited = make(map[int]bool)
    var edges = make(map[int]int)

    for src, _ := range sources {
        if visited[src] {
            continue
        }

        heap.Push(&nodes, newCity(src, 0))
        for nodes.Len() > 0 {
            var city0 = heap.Pop(&nodes).(*city)
            visited[city0.id] = true

            for _, road0 := range sources[city0.id] {
                if visited[road0.dest] {
                    continue
                }

                var cost0, ok = costs[road0.dest]
                var cost = road0.fare

                if cost < cost0 || !ok {
                    costs[road0.dest] = cost
                    edges[road0.dest] = road0.src

                    if !ok {
                        /// add new city with cost
                        var next = newCity(road0.dest, cost)
                        heap.Push(&nodes, next)
                    } else {
                        /// update city cost
                        for _, city1 := range nodes {
                            if city1.id == road0.dest {
                                city1.cost = cost
                                nodes.update(city1)
                                break
                            }
                        }
                    }
                }
            }
        }
    }

    debug(edges)
}

func debug(args ...interface{}) {
    fmt.Fprintln(os.Stderr, args)
}
