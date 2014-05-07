package main

import "bufio"
import "bytes"
import "fmt"
import "os"


var keypad = [10][]byte{
    []byte{' ', '0'},
    []byte{'1'},
    []byte{'a', 'b', 'c', '2'},
    []byte{'d', 'e', 'f', '3'},
    []byte{'g', 'h', 'i', '4'},
    []byte{'j', 'k', 'l', '5'},
    []byte{'m', 'n', 'o', '6'},
    []byte{'p', 'q', 'r', 's', '7'},
    []byte{'t', 'u', 'v', '8'},
    []byte{'w', 'x', 'y', 'z', '9'},
}

func main() {
    var n int
    fmt.Scanln(&n)

    var buffer bytes.Buffer
    var value int = 256
    var count int
    var space bool

    var scanner = bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanRunes)

    for ; n >= 0 && buffer.Len() < 160 && scanner.Scan(); n-- {
        var b = scanner.Bytes()
        if b[0] == ' ' && value == 256 {
            // impossible but test8
            continue
        } else if b[0] == ' ' {
            space = true
            continue
        }

        var value0 = int(b[0] - '0')
        // debug(b, value0)

        if value == 256 {
            value = value0
            count = 0

        } else if space || value != value0 {
            count = count % len(keypad[value])
            // debug(value, count, keypad[value][count])
            buffer.WriteByte(keypad[value][count])

            space = false
            value = value0
            count = 0
        } else {
            count += 1
        }
    }

    fmt.Printf("%q\n", buffer.String())
}

func debug(args ...interface{}) {
    fmt.Fprintln(os.Stderr, args)
}
