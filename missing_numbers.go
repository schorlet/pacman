package main

import "bufio"
import "fmt"
import "os"
import "sort"
import "strconv"
import "strings"

func main() {
    var numbers = make(map[string]int)
    var missing = []int{}
    var line string

    var in = bufio.NewReader(os.Stdin)
    in.ReadString('\n')
    line, _ = in.ReadString('\n')

    var scanner = bufio.NewScanner(strings.NewReader(line))
    scanner.Split(bufio.ScanWords)
    for scanner.Scan() {
        var c = scanner.Text()
        numbers[c] += 1
    }

    in.ReadString('\n')
    line, _ = in.ReadString('\n')

    scanner = bufio.NewScanner(strings.NewReader(line))
    scanner.Split(bufio.ScanWords)
    for scanner.Scan() {
        var c = scanner.Text()
        numbers[c] -= 1
        if numbers[c] == -1 {
            var d, _ = strconv.Atoi(c)
            missing = append(missing, d)
        }
    }

    sort.Ints(missing)
    for _, v := range missing {
        fmt.Printf("%d ", v)
    }
    fmt.Println()
}
