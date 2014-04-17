package main

import "container/heap"
import "fmt"
import "os"
import "math"

var (
    width, height, max_moves int
    start, goal point
    moves = map[rune]point{
            'U':point{0, -1}, 'L':point{-1, 0},
            'R':point{1,  0}, 'D':point{ 0, 1}}
)

type (
    table [][]rune
    point struct {
        x, y int
    }
    node struct {
        index       int
        parent      *node
        point       point
        distance1   int
        distance2   int
        operations  int
    }
    nodes []*node
)

func (self point) String() string {
    return fmt.Sprintf("{%d,%d}", self.x, self.y)
}
func (self point) distance(other point) int {
    return int(math.Abs(float64(self.x - other.x)) +
               math.Abs(float64(self.y - other.y)))
}
func (self point) neighbor(move point) point {
    return point{self.x + move.x, self.y + move.y}
}
func (self point) neighbors(board table) map[rune]point {
    var neighbors = make(map[rune]point)
    for direction, move := range moves {
        var next = self.neighbor(move)

        if next.x >= 0 && next.x < width &&
            next.y >= 0 && next.y < height {
            neighbors[direction] = next
        }
    }
    return neighbors
}

func (self node) String() string {
    return fmt.Sprintf("%s:%d, %d+%d=%d", self.point, self.operations,
            self.distance1, self.distance2,
            self.distance1 + self.distance2)
}
func newNode(parent *node, point point) *node {
    var child = &node{parent:parent, point:point}
    if parent != nil {
        child.distance1 = point.distance(start)
        child.operations = parent.operations
    }
    child.distance2 = point.distance(goal)
    return child
}
func (self node) distance() int {
    return self.distance1 + self.distance2
}

func (self nodes) Len() int {
    return len(self)
}
func (self nodes) Less(i, j int) bool {
    var ni, nj = self[i], self[j]
    if ni.operations < nj.operations {
        return true
    } else if ni.operations == nj.operations {
        return ni.distance() <= nj.distance()
    }
    return false
}
func (self nodes) Swap(i, j int) {
    self[i], self[j] = self[j], self[i]
    self[i].index = i
    self[j].index = j
}
func (self *nodes) Push(x interface{}) {
    var n = len(*self)
    var item = x.(*node)
    item.index = n
    *self = append(*self, item)
}
func (self *nodes) Pop() interface{} {
    var old = *self
    var n = len(old)
    var item = old[n-1]
    item.index = -1 // for safety
    *self = old[0:n-1]
    return item
}
func (self *nodes) update(item *node, distance1, distance2 int) {
    heap.Remove(self, item.index)
    item.distance1 = distance1
    item.distance2 = distance2
    heap.Push(self, item)
}

func find_star(board table) int {
    var visited = map[point]bool{}

    var root = newNode(nil, start)
    var solution *node

    var nodes = nodes{}
    heap.Init(&nodes)
    heap.Push(&nodes, root)

    for nodes.Len() > 0 {
        var node0 = heap.Pop(&nodes).(*node)
        var point0 = node0.point
        visited[point0] = true

        if point0 == goal {
            solution = node0
            break
        }

        var dir0 = board[point0.y][point0.x]

        for dir1, point1 := range point0.neighbors(board) {
            if visited[point1] {
                continue
            }

            var node1 = newNode(node0, point1)
            if node1.distance() > max_moves {
                continue
            }
            if dir0 != dir1 {
                node1.operations += 1
            }
            heap.Push(&nodes, node1)
        }
    }

    if solution == nil {
        // debug("No solution possible")
        return -1
    }
    return solution.operations
    // debug(solution)
    // debug(len(visited))
    // for solution.parent != nil {
        // board[solution.point.y][solution.point.x] = '*'
        // solution = solution.parent
    // }
    // debug_board(board)
    // return solution.operations
}

func main() {
    fmt.Scanf("%d %d %d", &height, &width, &max_moves)

    var i, j int
    var board = make(table, height)

    for j = 0; j < height; j++ {
        var line string
        var _, err = fmt.Scanf("%s", &line)
        if err != nil {
            fmt.Scanf("%s", &line)
        }
        board[j] = make([]rune, width)

        for i = 0; i < width; i++ {
            board[j][i] = rune(line[i])
            if board[j][i] == '*' {
                goal = point{i, j}
            }
        }
    }
    start = point{}
    var operations = find_star(board)
    fmt.Println(operations)
}


func debug(args ...interface{}) {
    fmt.Fprintln(os.Stderr, args)
}
func debug_board(board table) {
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

