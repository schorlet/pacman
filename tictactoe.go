package main

import "fmt"
import "math"
import "os"
import "strings"

const dimension = 3
const empty = '_'

var player, opponent rune
var ordered_moves [9]point

var _ = strings.Repeat

type (
    grid [][]rune
    point struct {
        x, y int
    }
)

func (self point) String() string {
    return fmt.Sprintf("(%d,%d)", self.x, self.y)
}
func (self point) repr() string {
    return fmt.Sprintf("%d %d", self.y, self.x)
}

func (self grid) repr() string {
    return fmt.Sprintf("%c", self)
}
func (self grid) copy() grid {
    var j int
    var board = make(grid, dimension)
    for j = 0; j < dimension; j++ {
        board[j] = make([]rune, dimension)
        copy(board[j], self[j])
    }
    return board
}
func (self grid) moves() []point {
    var moves = []point{}
    for _, move := range ordered_moves {
        if self[move.y][move.x] == empty {
            moves = append(moves, move)
        }
    }
    return moves
}
func (self grid) eval(row... point) int {
    var score int
    var counter = make(map[rune]int)
    for _, point := range row {
        var block = self[point.y][point.x]
        counter[block] += 1
    }

    if counter[player] == 3 {
        score = 1
    } else if counter[opponent] == 3 {
        score = -1
    }

    return score
}
func (self grid) score() int {
    var score int
    score = self.eval(point{0,0}, point{1,0}, point{2,0})
    if score != 0 { return score }
    score = self.eval(point{0,1}, point{1,1}, point{2,1})
    if score != 0 { return score }
    score = self.eval(point{0,2}, point{1,2}, point{2,2})
    if score != 0 { return score }

    score = self.eval(point{0,0}, point{0,1}, point{0,2})
    if score != 0 { return score }
    score = self.eval(point{1,0}, point{1,1}, point{1,2})
    if score != 0 { return score }
    score = self.eval(point{2,0}, point{2,1}, point{2,2})
    if score != 0 { return score }

    score = self.eval(point{0,0}, point{1,1}, point{2,2})
    if score != 0 { return score }
    score = self.eval(point{2,0}, point{1,1}, point{0,2})
    return score
}

func min_food(board grid, level, alpha, beta int) int {
    var moves = board.moves()
    var score = board.score()
    if score != 0 || level == 0 || len(moves) == 0 {
        return score * (level + 1)
    }

    var best_score = math.MaxInt32
    for _, move := range moves {
        var board0 = board.copy()
        board0[move.y][move.x] = opponent

        score, _ = max_food(board0, level - 1, alpha, beta)
        // debug(fmt.Sprintf("%q %s %s %-5d %s", opponent,
                // strings.Repeat("   ", 9 - level),
                // move, score, board0.repr()))

        if score < best_score {
            best_score = score
        }
        if score < beta {
            beta = score
            if beta <= alpha {
                break
            }
        }
    }
    return best_score
}
func max_food(board grid, level, alpha, beta int) (int, point) {
    var moves = board.moves()
    var score = board.score()
    if score != 0 || level == 0 || len(moves) == 0 {
        return score * (level + 1), point{-1,-1}
    }

    var best_score = math.MinInt32
    var best_move point
    for _, move := range moves {
        var board0 = board.copy()
        board0[move.y][move.x] = player

        score = min_food(board0, level - 1, alpha, beta)
        // debug(fmt.Sprintf("%q %s %s %-5d %s", player,
                // strings.Repeat("   ", 9 - level),
                // move, score, board0.repr()))

        if score > best_score {
            best_move = move
            best_score = score
        }
        if score > alpha {
            alpha = score
            if beta <= alpha {
                break
            }
        }
    }
    return best_score, best_move
}
func find_food(board grid, level int) point {
    var _, move = max_food(board, level, math.MinInt32, math.MaxInt32)
    // var score, move = max_food(board, level, math.MinInt32, math.MaxInt32)
    // debug(fmt.Sprintf("%q", player), move, score)
    return move
}

func play() {
    opponent = 'X'
    fmt.Scanf("%c", &player)
    fmt.Scanln()
    if player == 'X' {
        opponent = 'O'
    }

    var i, j int
    var level int

    var board = make(grid, dimension)
    for j = 0; j < dimension; j++ {
        board[j] = make([]rune, dimension)

        for i = 0; i < dimension; i++ {
            fmt.Scanf("%c", &board[j][i])
            if board[j][i] == empty {
                level += 1
            }
        }
        fmt.Scanln()
    }

    ordered_moves = [9]point{
        point{1,1},
        point{0,0}, point{0,2}, point{2,0}, point{2,2},
        point{1,0}, point{2,1}, point{1,2}, point{0,1},
    }
    // debug_board(board)
    // debug(level)

    if level == 9 {
        fmt.Println("1 1")
    } else if level == 8 && board[1][1] == empty {
        fmt.Println("1 1")
    } else {
        var move = find_food(board, level)
        fmt.Println(move.repr())
    }
}

func main() {
    // play()
    test()
}

func test() {
    ordered_moves = [9]point{
        point{1,1},
        point{0,0}, point{0,2}, point{2,0}, point{2,2},
        point{1,0}, point{2,1}, point{1,2}, point{0,1},
    }
    var board = grid{{'O', '_', '_'}, {'_', 'X', 'X'}, {'_', '_', '_'},}
    player, opponent = 'O', 'X'
    var level = 6

    for ;level > 0; level-- {
        var move = find_food(board, level)
        board[move.y][move.x] = player
        debug_board(board)
        if player == 'X' {
            player, opponent = 'O', 'X'
        } else {
            player, opponent = 'X', 'O'
        }
        if board.score() != 0 {
            break
        }
    }
}

func debug(args ...interface{}) {
    fmt.Fprintln(os.Stderr, args)
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
            fmt.Fprintf(os.Stderr, "%3c", board[y][x])
        }
        fmt.Fprintln(os.Stderr)
    }
    fmt.Fprintln(os.Stderr, "----------------------")
}
