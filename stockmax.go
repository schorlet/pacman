package main

import "bufio"
import "fmt"
import "os"
import "strings"
import "strconv"

func main() {
    var in = bufio.NewReader(os.Stdin)

    var line = readline(in)
    var tests, _ = strconv.Atoi(line)

    for ; tests > 0; tests-- {
        line = readline(in)
        var n, _ = strconv.Atoi(line)

        line = readline(in)
        var split = strings.Split(line, " ")

        var prices = make([]int, n)
        for i, el := range split {
            var price, _ = strconv.Atoi(el)
            prices[i] = price
        }
        stockmax(n, prices)
    }
}

func stockmax(n int, prices []int) {
    var spans = make([]int, n)
    var j = n - 1
    spans[j] = 0
    for i := j - 1; i >= 0; i-- {
        if prices[i] >= prices[j] {
            j = i
            spans[j] = 0
        } else {
            spans[i] = j - i
        }
    }
    var stock, gain int
    for i, el := range spans {
        if el == 0 && stock > 0 {
            gain += stock
            stock = 0
        } else if el > 0 {
            stock += (prices[i + el] - prices[i])
        }
    }
    // debug(n, prices)
    // debug(n, spans)
    fmt.Println(gain)
    // debug()
}

func readline(in *bufio.Reader) string {
    var line, _ = in.ReadString('\n')
    line = strings.TrimSpace(line)
    return line
}
func debug(args ...interface{}) {
    fmt.Fprintln(os.Stderr, args)
}
