package main

import "fmt"
import "os"
import "strings"

const sequence = "acdegilmnoprstuw"

func main() {
    // var source uint64 = 910897038977002
    // var i, m int = 0, 9
    var source uint64 = 680131659347
    var i, m int = 0, 7
    var indexes = make([]uint64, m)

    for i = m - 1; i >= 0; i-- {
        var q = source / 37
        indexes[i] = source - 37*q
        source = q
        debug(source, indexes[i])
    }

    var secret string
    for _, index := range indexes {
        secret += string(sequence[index])
    }
    debug(secret, hash(secret))
}

func hash(s string) uint64 {
    var h uint64 = 7
    for _, c := range s {
        h = h*37 + uint64(strings.IndexRune(sequence, c))
    }
    return h
}

func debug(args ...interface{}) {
    fmt.Fprintln(os.Stderr, args)
}
