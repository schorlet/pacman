package main

import "fmt"
import "os"

type set map[rune]bool

func main() {
    var i, n int
    fmt.Scanf("%d", &n)

    var all = make(set)
    var stones = make([]set, n)
    var line string

    for i = 0; i < n; i++ {
        fmt.Scanf("%s", &line)
        stones[i] = make(set)

        for _, r := range line {
            all[r] = true
            stones[i][r] = true
        }
    }

    var gems []rune
    all_range:
    for r, _ := range all {
        for _, stone := range stones {
            if !stone[r] {
                continue all_range
            }
        }
        gems = append(gems, r)
    }
    fmt.Println(len(gems))
}

func debug(args ...interface{}) {
    fmt.Fprintln(os.Stderr, args)
}
