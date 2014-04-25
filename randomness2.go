package main

import "bytes"
import "fmt"
import "math/big"
import "os"
import "sort"

func main() {
    var i, n, q int
    fmt.Scanf("%d", &n)
    fmt.Scanln(&q)

    var input string
    fmt.Scanln(&input)
    var sequence = []byte(input)

    for i = 0; i < q; i++ {
        var p int
        var c byte
        fmt.Scanf("%d", &p)
        fmt.Scanf("%c", &c)
        fmt.Scanln()
        sequence[p - 1] = c
        distinct_substrings(n, sequence)
    }
}

func distinct_substrings(n int, sequence []byte) {
    var j, k int
    var m = n + 1

    var suffixes = make(suffixArray, m)
    suffixes[0] = []byte{36} // ord('$') == 36

    for j = 1; j < m; j++ {
        suffixes[j] = sequence[n - j : n]
    }
    sort.Sort(suffixes)

    // var prefixes = make([]int, m)
    var prefixes int

    for j = 2; j < m; j++ {
        var l = len(suffixes[j])
        for k = 1; k < l + 1; k++ {
            if !bytes.HasPrefix(suffixes[j - 1], suffixes[j][:k]) {
                break
            }
        }
        // prefixes[j] = k - 1
        prefixes += (k - 1)
    }

    // for j = 1; j < m; j++ {
        // fmt.Printf("%s %d\n", suffixes[j], prefixes[j])
    // }
    fmt.Println(binomial(n + 1, 2) - int64(prefixes))
}

type suffixArray [][]byte
func (self suffixArray) Len() int {
    return len(self)
}
func (self suffixArray) Swap(i, j int) {
    self[i], self[j] = self[j], self[i]
}
func (self suffixArray) Less(i, j int) bool {
    var suffixi, suffixj = self[i], self[j]
    return bytes.Compare(suffixi, suffixj) == -1
}

func binomial(n, k int) int64 {
    return big.NewInt(0).Binomial(int64(n), int64(k)).Int64()
}

func debug(args ...interface{}) {
    fmt.Fprintln(os.Stderr, args)
}
