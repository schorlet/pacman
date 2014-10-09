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

    var best = 6000
    for i := 0; i < n; i++ {
        var temp, _ = strconv.Atoi(split[i])
        if abs(temp) < abs(best) {
            best = temp
        } else if abs(temp) == abs(best) && temp > 0 {
            best = temp
        }
    }
    fmt.Println(best)
}

func abs(x int) int {
    switch {
    case x < 0:
        return -x
    case x == 0:
        return 0 // return correctly abs(-0)
    }
    return x
}

func readline(in *bufio.Reader) string {
    var line, _ = in.ReadString('\n')
    line = strings.TrimSpace(line)
    return line
}
