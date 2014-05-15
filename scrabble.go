package main

import "fmt"
import "os"

func main() {
    var scoring = map[int]string {
        1: "eaionrtlsu",
        2: "dg",
        3: "bcmp",
        4: "fhvwy",
        5: "k",
        8: "jx",
        10: "qz",
    }
    var reverse_scoring = make(map[rune]int)
    for score, letters := range scoring {
        for _, letter := range letters {
            reverse_scoring[letter] = score
        }
    }

    var i, n int
    fmt.Scanf("%d", &n)

    var words = make([]string, n)
    for ; i < n; i ++ {
        fmt.Scanf("%s", &words[i])
    }

    var letters string
    fmt.Scanf("%s", &letters)
    debug(letters)

    var counter = make(map[rune]int)
    for _, c := range letters {
        counter[c] += 1
    }

    var max_score int
    var max_word string

    for _, word := range words {
        if len(word) > len(letters) {
            continue
        }

        var cword = make(map[rune]int)
        var stop bool
        for _, c := range word {
            if counter[c] == 0 {
                stop = true
                break
            }
            cword[c] += 1
        }
        if stop {
            continue
        }

        var score int
        for c, j := range cword {
            if counter[c] >= j {
                score += reverse_scoring[c] * j
            }
        }
        if score > max_score {
            max_score = score
            max_word = word
        }
    }

    fmt.Println(max_word)
}

func debug(args ...interface{}) {
    fmt.Fprintln(os.Stderr, args)
}
