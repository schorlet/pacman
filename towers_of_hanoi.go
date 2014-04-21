package main

import "fmt"
import "os"

// The minimum number of moves required to solve a Tower of Hanoi puzzle is 2n - 1,
// where n is the number of disks.

func towers_of_hanoi(n, a, b, c int) {
    if n > 0 {
        towers_of_hanoi(n - 1, a, c, b);   //recursion
        fmt.Printf("> Move disk %d from tower %d to tower %d.\n", n, a, c)
        towers_of_hanoi(n - 1, b, a, c);   //recursion
    }
}

func main() {
    towers_of_hanoi(3, 1, 2, 3)
}

func debug(args ...interface{}) {
    fmt.Fprintln(os.Stderr, args)
}
