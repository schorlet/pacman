package main

import "bytes"
import "fmt"

func main() {
    var tests int
    fmt.Scanf("%d", &tests)

    var input string
    for ; tests > 0; tests-- {
        fmt.Scanf("%s", &input)
        arrange(input)
    }
}

func arrange(input string) {
    var l = len(input)
    if l%2 == 1 {
        fmt.Println(-1)
        return
    }

    var a = []byte(input[:l/2])
    var b = []byte(input[l/2:])

    var counter = make(map[byte]int)
    var commons int

    for _, c := range b {
        var i = counter[c]
        var d = bytes.IndexByte(a[i:], c)
        if d > -1 {
            counter[c] = d + i + 1
            commons += 1
        }
    }
    fmt.Println(l/2 - commons)
}
