package main

import "bufio"
import "container/heap"
import "fmt"
import "math"
import "os"
import "strings"
import "strconv"

const DegToRad = math.Pi / 180

type StopArea struct {
    id    string
    name  string
    lat   float64
    lon   float64
    dests []string
}

func (s StopArea) String() string {
    return fmt.Sprintf("{%s: %s, %s}", s.id, s.name, s.dests)
}

func distance(a, b *StopArea) float64 {
    var x = (b.lon - a.lon) * math.Cos((a.lat+b.lat)/2)
    var y = b.lat - a.lat
    return math.Sqrt(math.Pow(x, 2)+math.Pow(y, 2)) * 6371
}

var areas = make(map[string]*StopArea)

func main() {
    var in = bufio.NewReader(os.Stdin)

    var departure = readline(in)
    var arrival = readline(in)

    var line = readline(in)
    var n, _ = strconv.Atoi(line)
    if n == 0 {
        fmt.Println("IMPOSSIBLE")
        return
    }

    for i := 0; i < n; i++ {
        line = readline(in)
        var split = strings.Split(line, ",")

        var area = new(StopArea)
        area.id = split[0]
        area.name = strings.Trim(split[1], "\"")

        var lat, _ = strconv.ParseFloat(split[3], 64)
        area.lat = lat * DegToRad

        var lon, _ = strconv.ParseFloat(split[4], 64)
        area.lon = lon * DegToRad

        areas[area.id] = area
    }

    line = readline(in)
    var m, _ = strconv.Atoi(line)
    if m == 0 {
        fmt.Println("IMPOSSIBLE")
        return
    }

    for i := 0; i < m; i++ {
        line = readline(in)
        var split = strings.Split(line, " ")
        var from = split[0]
        var to = split[1]
        if from != to {
            areas[from].dests = append(areas[from].dests, to)
        }
    }

    print_best_path(departure, arrival)
}

type StopItem struct {
    stop   string
    parent string
    dist   float64
    // The index is needed by update and is maintained by the heap.Interface methods.
    index int // The index of the item in the heap.
}

type StopItems []*StopItem

func (pq StopItems) Len() int {
    return len(pq)
}
func (pq StopItems) Less(i, j int) bool {
    return pq[i].dist < pq[j].dist
}
func (pq StopItems) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
    pq[i].index = i
    pq[j].index = j
}
func (pq *StopItems) Push(x interface{}) {
    n := len(*pq)
    item := x.(*StopItem)
    item.index = n
    *pq = append(*pq, item)
}
func (pq *StopItems) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    item.index = -1 // for safety
    *pq = old[0 : n-1]
    return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *StopItems) update(item *StopItem, parent string, dist float64) {
    item.parent = parent
    item.dist = dist
    heap.Fix(pq, item.index)
}

func print_best_path(departure, arrival string) {
    var pq = StopItems{&StopItem{departure, departure, 0, 0}}
    heap.Init(&pq)

    var parents = make(map[string]string)
    var cache = make(map[string]*StopItem)

    for pq.Len() > 0 {
        var item = heap.Pop(&pq).(*StopItem)
        delete(cache, item.stop)

        parents[item.stop] = item.parent
        if item.stop == arrival {
            break
        }

        var stop = areas[item.stop]
        for _, dest := range stop.dests {
            if dest == item.parent {
                continue
            }

            var dist = item.dist + distance(stop, areas[dest])
            if next, ok := cache[dest]; ok {
                if dist < next.dist {
                    pq.update(next, item.stop, dist)
                }
            } else {
                next = &StopItem{
                    stop:   dest,
                    parent: item.stop,
                    dist:   dist}
                heap.Push(&pq, next)
                cache[dest] = next
            }
        }
    }

    delete(parents, departure)
    if parent, ok := parents[arrival]; ok {
        var path = []string{arrival}
        for i := 0; i < 40 && ok;  parent, ok = parents[parent] {
            path = append(path, parent)
            i += 1
        }
        for i := len(path) - 1; i >= 0; i-- {
            fmt.Println(areas[path[i]].name)
        }
    } else if departure == arrival {
        fmt.Println(areas[departure].name)
    } else {
        fmt.Println("IMPOSSIBLE")
    }
}

func readline(in *bufio.Reader) string {
    var line, _ = in.ReadString('\n')
    line = strings.TrimSpace(line)
    return line
}

func debug(args ...interface{}) {
    fmt.Fprintln(os.Stderr, args)
}
