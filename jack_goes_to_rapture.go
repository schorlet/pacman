package main

import "container/heap"
import "bufio"
import "fmt"
import "os"

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
    find_city(1, n, sources)
}

// kind of dijkstra algo (without init)
func find_city(start, goal int, sources map[int]roads) {
    var nodes = cities{}
    heap.Init(&nodes)
    heap.Push(&nodes, newCity(start, 0))

    // cumulated costs
    var costs = make(map[int]int64)
    var previous = make(map[int]int)

    for nodes.Len() > 0 {
        var city0 = heap.Pop(&nodes).(*city)
        // debug(city0)

        if city0.id == goal {
            fmt.Println(city0.cost)
            var path []int
            for goal != start {
                path = append(path, goal)
                goal = previous[goal]
            }
            path = append(path, start)
            debug(path)
            return
        }

        for _, road0 := range sources[city0.id] {
            var cost0, ok = costs[road0.dest]
            var cost = city0.cost + max(0, road0.fare-city0.cost)
            // var cost = city0.cost + road0.fare

            if cost < cost0 || !ok {
                costs[road0.dest] = cost
                previous[road0.dest] = road0.src
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
