package main

import "fmt"

func main() {
    var n int
    fmt.Scanln(&n)

    var trie = make(map[rune]interface{})
    var total int

    for i := 0; i < n; i++ {
        var line string
        fmt.Scanln(&line)

        var root = trie
        for _, c := range line {
            if sub, ok := root[c]; ok {
                root = sub.(map[rune]interface{})
            } else {
                var nod = make(map[rune]interface{})
                root[c] = nod
                root = nod
                total += 1
            }
        }
    }

    fmt.Println(total)
}

