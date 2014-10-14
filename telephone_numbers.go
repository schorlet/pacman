package main

import "fmt"

func main() {
    var n int
    fmt.Scanln(&n)

    type node map[rune]interface{}
    var trie = make(node)

    var total int

    for i := 0; i < n; i++ {
        var line string
        fmt.Scanln(&line)

        var root = trie
        for _, c := range line {
            if sub, ok := root[c]; ok {
                root = sub.(node)
            } else {
                var nod = make(node)
                root[c] = nod
                root = nod
                total += 1
            }
        }
    }

    fmt.Println(total)
}

