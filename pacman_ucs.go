package main

import "container/heap"
import "fmt"
import "os"
import "strings"

type (
    point struct {
        x, y int
    }
    points []point

    node struct {
        parent    *node
        children  nodes
        point     point
    }
    nodes []*node

    cost struct {
        priority  int
        node      *node
        index     int
    }
    costs []*cost
)

var width, height int
var pacman, food point

// UP, LEFT, RIGHT, DOWN
var moves = []point {point{0, -1}, point{-1, 0},
                     point{1,  0}, point{ 0, 1}, }


func newNode(parent *node, point point) *node {
    var child = &node{parent:parent, point:point, children:nodes{}}
    if parent != nil {
        parent.children = append(parent.children, child)
    }
    return child
}

func (pq costs) Len() int {
    return len(pq)
}
func (pq costs) Less(i, j int) bool {
    return pq[i].priority < pq[j].priority
}
func (pq costs) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
    pq[i].index = i
    pq[j].index = j
}
func (pq *costs) Push(x interface{}) {
    var n = len(*pq)
    var item = x.(*cost)
    item.index = n
    *pq = append(*pq, item)
}
func (pq *costs) Pop() interface{} {
    var old = *pq
    var n = len(old)
    var item = old[n-1]
    item.index = -1 // for safety
    *pq = old[0:n-1]
    return item
}
func (pq *costs) update(item *cost, priority int) {
    heap.Remove(pq, item.index)
    item.priority = priority
    heap.Push(pq, item)
}

func (self point) String() string {
    return fmt.Sprintf("%d %d", self.y, self.x)
}
func (self point) neighbor(move point) point {
    return point{self.x + move.x, self.y + move.y}
}
func (self point) neighbors(board [][]rune) points {
    var neighbors = points{}
    for _, move := range moves {
        var next = self.neighbor(move)

        if board[next.y][next.x] == '-' {
            neighbors = append(neighbors, next)

        } else if board[next.y][next.x] == '.' {
            neighbors = append(neighbors, next)
        }
    }
    return neighbors
}

func find_food(board [][]rune) {
    var visited = map[point]bool{}

    var root = newNode(nil, pacman)
    var food_node *node

    var costs = costs{}
    heap.Init(&costs)
    heap.Push(&costs, &cost{priority:0, node:root})

    for costs.Len() > 0 {
        var cost0 = heap.Pop(&costs).(*cost)
        var node0 = cost0.node

        var point0 = node0.point
        visited[point0] = true

        if point0 == food {
            food_node = node0
            break
        }
        neighborhood:
        for _, next := range point0.neighbors(board) {
            if visited[next] {
                continue
            }
            var k = cost0.priority
            if next != food {
                k += 1
            }
            for _, cost1 := range costs {
                if cost1.node.point == next {
                    if cost1.priority > k {
                        costs.update(cost1, k)
                    }
                    continue neighborhood
                }
            }
            heap.Push(&costs, &cost{priority:k, node:newNode(node0, next)})
        }
    }

    // debug(len(visited))

    var path = points{}
    for food_node.parent != nil {
        path = append(path, food_node.point)
        food_node = food_node.parent
    }

    // distance between source and the destination
    // debug(len(path))
    fmt.Println(len(path))
    path = append(path, pacman)

    // reverse path
    var i, d = 0, len(path)
    for i = 0; i < d/2; i++ {
        path[i], path[d-1-i] = path[d-1-i], path[i]
    }

    // path between source and the destination
    for _, explore := range path {
        // debug(explore)
        fmt.Println(explore)
    }
}

func play() {
    pacman = point{}
    food = point{}

    fmt.Scanf("%d %d", &pacman.y, &pacman.x)
    fmt.Scanf("%d %d", &food.y, &food.x)
    fmt.Scanf("%d %d", &height, &width)
    // debug(pacman, food, width, height)

    var i, j int
    var board = make([][]rune, height)
    for j = 0; j < height; j++ {
        board[j] = make([]rune, width)

        for i = 0; i < width; i++ {
            fmt.Scanf("%c", &board[j][i])
            if board[j][i] == '\n' {
                fmt.Scanf("%c", &board[j][i])
            }
        }
    }
    // debug_board(board)
    find_food(board)
}

func main() {
    play()
}


func debug(args ...interface{}) {
    fmt.Fprintln(os.Stderr, args)
}
func debug_node(node *node, level int) {
    debug(strings.Repeat("   ", level), node.point)
    for _, child := range node.children {
        debug_node(child, level+1)
    }
}
func debug_board(board [][]rune) {
    fmt.Fprint(os.Stderr, "   |")
    for x := range board[0] {
        fmt.Fprintf(os.Stderr, "%3d", x)
    }
    fmt.Fprintln(os.Stderr)
    fmt.Fprint(os.Stderr, "---+")
    for _ = range board[0] {
        fmt.Fprint(os.Stderr, "---")
    }
    fmt.Fprintln(os.Stderr)
    for y := range board {
        fmt.Fprintf(os.Stderr, "%2d |", y)
        for x := range board[y] {
            fmt.Fprintf(os.Stderr, "%3c", board[y][x])
        }
        fmt.Fprintln(os.Stderr)
    }
    fmt.Fprintln(os.Stderr, "----------------------")
}

