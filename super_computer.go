package main

import "fmt"
import "os"
import "sort"

type calcul struct {
    day int
    duration int
}
type calculs []calcul

func (c calcul) end() int {
    return c.day + c.duration
}

func (c calcul) String() string {
    return fmt.Sprintf("%d,%d", c.day, c.duration)
}

func (c calculs) Len() int {
    return len(c)
}
func (c calculs) Swap(i, j int) {
    c[i], c[j] = c[j], c[i]
}
func (c calculs) Less(i, j int) bool {
    var c0, c1 = c[i], c[j]
    return c0.end() < c1.end()
}

func main() {
    var n int
    fmt.Scanf("%d\n", &n)

    var calcs = make(calculs, n)
    for i:= 0; i < n; i++ {
        fmt.Scanf("%d %d\n", &calcs[i].day, &calcs[i].duration)
    }

    sort.Sort(calcs)

    var count int
    var start int

    for _, c := range calcs {
        if c.day >= start {
            start = c.end()
            count += 1
        }
    }

    fmt.Println(count)
}

func debug(args ...interface{}) {
    fmt.Fprintln(os.Stderr, args)
}
