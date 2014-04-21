package main

import "fmt"
import "os"

func main() {
    var l, r int
    fmt.Scanf("%d", &l)
    fmt.Scanf("%d", &r)

    var i, j, max int
    var length = len(fmt.Sprintf("%b", r))

    for i = r; i > l; i-- {
        if len(fmt.Sprintf("%b", i)) < length {
            break
        }
        for j = i - 1; j >= l; j -= 2 {
            var xor = i ^ j
            debug(i, j, xor)
            if xor > max {
                max = xor
            }
        }
        debug()
    }
    // debug(fmt.Sprintf("%b", max))
    fmt.Println(max)
}


func debug(args ...interface{}) {
    fmt.Fprintln(os.Stderr, args)
}
