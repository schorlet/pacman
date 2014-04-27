package main

import (
    "fmt"
    "os"
    "strings"
    "unicode"
)

const ROT = 13

func main() {
    debug(rot13("abcd xyz"))
}

func rot13(s string) string {
    return strings.Map(func(c rune) rune {
        switch {
        case unicode.IsLetter(c + ROT):
            return c + ROT
        case unicode.IsLetter(c - ROT):
            return c - ROT
        }
        return c
    }, s)
}

func debug(args ...interface{}) {
    fmt.Fprintln(os.Stderr, args)
}
