package main

import "fmt"
import "os"

func main() {
    var i, n, q int
    fmt.Scanf("%d", &n)
    fmt.Scanln(&q)

    var input string
    fmt.Scanln(&input)
    var sequence = []byte(input)

    for i = 0; i < q; i++ {
        var p int
        var c byte
        fmt.Scanf("%d", &p)
        fmt.Scanf("%c", &c)
        fmt.Scanln()
        sequence[p - 1] = c

        var set = map[string]bool{}
        var j, k int
        for j = 0; j < n; j++ {
            for k = j + 1; k <= n; k++ {
                set[string(sequence[j : k])] = true
            }
        }
        fmt.Println(len(set))
    }
}

func debug(args ...interface{}) {
    fmt.Fprintln(os.Stderr, args)
}
