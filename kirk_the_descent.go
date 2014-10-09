package main

import (
    "fmt"
    "os"
    "sort"
)

type (
    Mountain struct {
        x, height int
    }
    Mountains []*Mountain
)

func (mountains Mountains) Len() int {
    return len(mountains)
}
func (mountains Mountains) Swap(i, j int) {
    mountains[i], mountains[j] = mountains[j], mountains[i]
}
func (mountains Mountains) Less(i, j int) bool {
    var m0, m1 = mountains[i], mountains[j]
    return m0.height > m1.height
}

func main() {
    for {
        var sx, sy int
        fmt.Scanf("%d %d", &sx, &sy)

        var mountains = make(Mountains, 8)
        for i := 0; i < 8; i++ {
            var mountain = &Mountain{x: i}
            fmt.Scanf("%d", &mountain.height)
            mountains[i] = mountain
        }
        sort.Sort(mountains)

        var mountain = mountains[0]
        if mountain.x == sx {
            fmt.Println("FIRE")
        } else {
            fmt.Println("HOLD")
        }
    }
}

func debug(args ...interface{}) {
    fmt.Fprintln(os.Stderr, args)
}
