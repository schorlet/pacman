package main

import "bufio"
import "fmt"
import "os"
import "strings"
import "strconv"

func main() {
    var in = bufio.NewReader(os.Stdin)

    var line = readline(in)
    var n, _ = strconv.Atoi(line)
    if n == 0 {
        fmt.Println(0)
        return
    }

    line = readline(in)
    var split = strings.Split(line, " ")

    var max, min, best int
    for i := 0; i < n; i++ {
        var val, _ = strconv.Atoi(split[i])
        if val > max {
            max = val
        }
        if i == 0 || val <= min {
            min = val
            best = min - max
        }
    }
    fmt.Println(best)
}

func readline(in *bufio.Reader) string {
    var line, _ = in.ReadString('\n')
    line = strings.TrimSpace(line)
    return line
}
