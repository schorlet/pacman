package main

import "fmt"
import "os"
import "strconv"

func main() {
    var i, n int
    fmt.Scanf("%d", &n)

    var board = make([]string, n)
    var start, end int
    for i = 0; i < n; i++ {
        fmt.Scanf("%s", &board[i])
        if board[i] == "S" {
            start = i
        } else if board[i] == "E" {
            end = i
        }
    }

    var counter = make([]int, n)
    var queue = []int{start}

    for len(queue) > 0 {
        var index = queue[0]
        queue = queue[1:]
        var count = counter[index]

        if board[index] == "S" || board[index] == "R" {
            var j = 1
            for ; j < 7; j++ {
                if index + j < 0 || index + j >= n {
                    continue
                } else if counter[index + j] > 0 &&
                        counter[index + j] < count + 1 {
                    continue
                }
                queue = append(queue, index + j)
                counter[index + j] = count + 1
            }

        } else if board[index] == "E" {
            break

        } else {
            var j, _ = strconv.Atoi(board[index])
            if index + j < 0 || index + j >= n {
                continue
            } else if counter[index + j] > 0 &&
                    counter[index + j] < count + 1 {
                continue
            }
            queue = append(queue, index + j)
            counter[index + j] = count + 1
        }

        // debug(index)
        // debug(counter)
        // debug(queue)
        // debug()
    }


    if counter[end] > 0 {
        fmt.Println(counter[end])
    } else {
        fmt.Println("impossible")
    }
}

func debug(args ...interface{}) {
    fmt.Fprintln(os.Stderr, args)
}
