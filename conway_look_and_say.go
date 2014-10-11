package main

import "bufio"
import "bytes"
import "fmt"
import "os"
import "strings"
import "strconv"

func main() {
    var r int
    fmt.Scanln(&r)

    var l int
    fmt.Scanln(&l)

    var line = strconv.Itoa(r)
    for i := 1; i < l; i++ {
        var tokens = NewTokens(line)
        line = tokens.String()
    }

    fmt.Println(line)
}

type Tokens struct {
    key  []string
    val  []int
    last int
}

func NewTokens(line string) *Tokens {
    var tokens = new(Tokens)
    tokens.last = -1

    var scanner = bufio.NewScanner(strings.NewReader(line))
    scanner.Split(bufio.ScanWords)
    for scanner.Scan() {
        var c = scanner.Text()
        tokens.put(c)
    }
    return tokens
}

func (t *Tokens) put(s string) {
    if t.last == -1 || t.key[t.last] != s {
        t.key = append(t.key, s)
        t.val = append(t.val, 1)
        t.last += 1
    } else {
        t.val[t.last] += 1
    }
}

func (t *Tokens) get(i int) (string, int) {
    return t.key[i], t.val[i]
}

func (t *Tokens) String() string {
    var buffer bytes.Buffer
    for i := 0; i < t.last; i++ {
        var key, val = t.get(i)
        buffer.WriteString(fmt.Sprintf("%d %s ", val, key))
    }
    if t.last > -1 {
        var key, val = t.get(t.last)
        buffer.WriteString(fmt.Sprintf("%d %s", val, key))
    }
    return buffer.String()
}

func debug(args ...interface{}) {
    fmt.Fprintln(os.Stderr, args)
}
