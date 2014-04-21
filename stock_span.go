package main

import "bufio"
import "container/list"
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
        span2(n, prices)
    }
}

// http://www.geeksforgeeks.org/the-stock-span-problem/
func span2(n int, prices []int) {
    var spans = make([]int, n)
    var stack = list.New()

    for i, price := range prices {
        var e = stack.Back()
        var value int

        for e != nil {
            value = e.Value.(int)
            if prices[value] < price {
                var f = e.Prev()
                stack.Remove(e)
                e = f
            } else {
                break
            }
        }

        if stack.Len() == 0 {
            spans[i] = i
        } else {
            spans[i] = i - value - 1
        }
        stack.PushBack(i)
    }
    debug(n, prices)
    debug(n, spans)
    debug()
}

// http://www.geeksforgeeks.org/the-stock-span-problem/
func span1(n int, prices []int) {
    var spans = make([]int, n)
    for i, price := range prices {
        var k int
        for k < i {
            if prices[i - k - 1] < price {
                k += 1
            } else {
                break
            }
        }
        spans[i] = k
    }
    debug(n, prices)
    debug(n, spans)
    debug()
}

func readline(in *bufio.Reader) string {
    var line, _ = in.ReadString('\n')
    line = strings.TrimSpace(line)
    return line
}
func debug(args ...interface{}) {
    fmt.Fprintln(os.Stderr, args)
}
