package main

// http://www.geeksforgeeks.org/backtracking-set-1-the-knights-tour-problem/
// Backtracking is *not* the best solution for the Knight’s tour problem
// Plus, configuration of xMove and yMove is very important
// Plus, it seems this algo run fine when starting at [0][0] but *not* with other positions

import "fmt"
import "os"

const N = 8
var yMove = [N]int{ 2, 1, -1, -2, -2, -1,  1,  2 }
var xMove = [N]int{ 1, 2,  2,  1, -1, -2, -2, -1 }
var solution [N][N]int

func isSafe(x, y int) bool {
    if x >= 0 && x < N && y >= 0 && y < N && solution[y][x] == 0 {
        return true
    }
    return false
}

func solveKT(x, y int) bool {
    if solution[y][x] == N*N {
        return true
    }

    for k := 0; k < N; k++ {
        var kx = x + xMove[k]
        var ky = y + yMove[k]

        if isSafe(kx, ky) {
            solution[ky][kx] = solution[y][x] + 1
            if solveKT(kx, ky) {
                return true
            } else {
                // backtracking
                solution[ky][kx] = 0
            }
        }
    }
    return false
}

func solve(x, y int) {
    solution[y][x] = 1
    if !solveKT(x, y) {
        debug("Solution does not exist")
        return
    }
    for y := 0; y < N; y++ {
        fmt.Printf("%2d\n", solution[y])
    }
}

func main() {
    solve(0, 0)
}

func debug(args ...interface{}) {
    fmt.Fprintln(os.Stderr, args)
}
