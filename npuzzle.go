package main

import "container/heap"
import "fmt"
import "math"
import "os"
import "strings"

var (
    dimension int
    goal string
    moves = map[string]point{
            "UP":point{0, -1}, "LEFT":point{-1, 0},
            "RIGHT":point{1,  0}, "DOWN":point{ 0, 1}}
)

type (
    grid [][]int
    shift struct {
        move string
        board grid
    }
    point struct {
        x, y int
    }
    node struct {
        parent    *node
        children  nodes
        move      string
        board     grid
        level     int
    }
    nodes []*node
    cost struct {
        node *node
        priority, index int
    }
    costs []*cost
)

func newNode(parent *node, move string, board grid) *node {
    var child = &node{parent:parent, children:nodes{},
                      board:board, move:move}
    if parent != nil {
        child.level = parent.level + 1
        parent.children = append(parent.children, child)
    }
    return child
}

func (self point) String() string {
    return fmt.Sprintf("%d %d", self.y, self.x)
}
func (self point) distance(other point) int {
    return int(math.Abs(float64(self.x - other.x)) +
               math.Abs(float64(self.y - other.y)))
}

func (self point) neighbor(move point) point {
    return point{self.x + move.x, self.y + move.y}
}
func (self point) neighbors(board grid) []point {
    var neighbors = []point{}
    for _, move := range moves {
        var next = self.neighbor(move)

        if next.x >= 0 && next.x < dimension &&
            next.y >= 0 && next.y < dimension {
            neighbors = append(neighbors, next)
        }
    }
    return neighbors
}
func (self point) shifts(board grid) []shift {
    var list = []shift{}
    for move, direction := range moves {
        var next = self.neighbor(direction)

        if next.x >= 0 && next.x < dimension &&
            next.y >= 0 && next.y < dimension {

            var board0 = board.copy()
            board0[self.y][self.x] = board[next.y][next.x]
            board0[next.y][next.x] = 0
            list = append(list, shift{move, board0})
        }
    }
    return list
}

func (self grid) repr() string {
    return fmt.Sprintf("%v", self)
}
func (self grid) zero() point {
    var i, j int
    for j = 0; j < dimension; j++ {
        for i = 0; i < dimension; i++ {
            if self[j][i] == 0 {
                return point{i, j}
            }
        }
    }
    panic("can not find the zero point")
}
// return number of blocks out of place
func (self grid) hamming() int {
    var i, j, n int
    for j = 0; j < dimension; j++ {
        for i = 0; i < dimension; i++ {
            var k = j * dimension + i
            if k > 0 && self[j][i] != k {
                n += 1
            }
        }
    }
    return n
}
// return sum of Manhattan distances between blocks and goal
func (self grid) manhattan() int {
    var i, j, n int
    var points = make([]point, dimension * dimension)
    for j = 0; j < dimension; j++ {
        for i = 0; i < dimension; i++ {
            var v = self[j][i]
            points[v] = point{i, j}
        }
    }
    for j = 0; j < dimension; j++ {
        for i = 0; i < dimension; i++ {
            var k = j * dimension + i
            if k > 0 {
                n += points[k].distance(point{i, j})
            }
        }
    }
    return n
}
func (self grid) copy() grid {
    var j int
    var board = make(grid, dimension)
    for j = 0; j < dimension; j++ {
        board[j] = make([]int, dimension)
        copy(board[j], self[j])
    }
    return board
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

func find_food(board grid) {
    var visited = map[string]bool{}

    var root = newNode(nil, "INIT", board)
    var solution *node

    var costs = costs{}
    heap.Init(&costs)
    heap.Push(&costs, &cost{priority:0, node:root})

    for costs.Len() > 0 {
        var cost0 = heap.Pop(&costs).(*cost)
        var node0 = cost0.node

        var board0 = node0.board
        var repr0 = board0.repr()

        if repr0 == goal {
            solution = node0
            break
        }

        var zero = board0.zero()
        for _, shift := range zero.shifts(board0) {
            var board1 = shift.board
            var repr1 = board1.repr()
            if visited[repr1] {
                continue
            }
            visited[repr1] = true

            var k = board1.manhattan() + node0.level
            var node1 = newNode(node0, shift.move, board1)
            // debug("  ", node1.level, k, repr1)
            heap.Push(&costs, &cost{priority:k, node:node1})
        }
    }

    if solution == nil {
        fmt.Println("No solution possible")
        return
    }

    // debug(len(visited))
    var moves = []string{}
    for solution.parent != nil {
        moves = append(moves, solution.move)
        // debug(solution.level, solution.board.repr())
        solution = solution.parent
    }
    // debug(solution.level, solution.board.repr())

    // reverse moves
    var i, d = 0, len(moves)
    for i = 0; i < d/2; i++ {
        moves[i], moves[d-1-i] = moves[d-1-i], moves[i]
    }

    // moves between initial and goal
    fmt.Println(len(moves))
    for _, move := range moves {
        fmt.Println(move)
    }
}

func define_goal() {
    var board = grid{{0, 1, 2}, {3, 4, 5}, {6, 7, 8},}
    goal = board.repr()
}

func play() {
    fmt.Scanf("%d", &dimension)

    var i, j int
    var board = make(grid, dimension)
    for j = 0; j < dimension; j++ {
        board[j] = make([]int, dimension)

        for i = 0; i < dimension; i++ {
            fmt.Scanf("%d", &board[j][i])
        }
    }
    define_goal()
    find_food(board)
}

func main() {
    play()
    // test()
}

//  0  3  6     0  1  2     1  2  3  4  5  6  7  8    1  2  3  4  5  6  7  8
//  1  4  7     3  4  5     ----------------------    ----------------------
//  2  5  8     6  7  8     1  1  1  0  1  1  1  0    2  4  2  0  2  4  2  0
//
//  initial       goal         Hamming = 6              Manhattan = 16

func test() {
    dimension = 3
    var board = grid{
            {0, 3, 6},
            {1, 4, 7},
            {2, 5, 8},
        }
    debug_board(board)
    debug(board.repr())
    debug(board.hamming())
    debug(board.manhattan())
    var zero = board.zero()
    debug(zero)
    debug(zero.neighbors(board))

    define_goal()

    for _, shift := range zero.shifts(board) {
        debug(shift.move)
        debug_board(shift.board)
        debug(shift.board.repr())
        debug(shift.board.hamming())
        debug(shift.board.manhattan())
    }
    // find_food(board)
}

func debug(args ...interface{}) {
    fmt.Fprintln(os.Stderr, args)
}
func debug_node(node *node, level int) {
    debug(strings.Repeat("   ", level), node.board.repr())
    for _, child := range node.children {
        debug_node(child, level+1)
    }
}
func debug_board(board grid) {
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
            fmt.Fprintf(os.Stderr, "%3d", board[y][x])
        }
        fmt.Fprintln(os.Stderr)
    }
    fmt.Fprintln(os.Stderr, "----------------------")
}

