package main

import "fmt"

func main() {
    var tests int
    fmt.Scanf("%d", &tests)

    for ; tests > 0; tests-- {
        var number uint
        fmt.Scanf("%d", &number)

        if is_fibonacci(number) {
            fmt.Println("IsFibo")
        } else {
            fmt.Println("IsNotFibo")
        }
    }
}

func is_fibonacci(number uint) bool {
    var current uint = 0
    var next uint = 1
    for number > current {
        current, next = next, current + next
    }
    return number == current
}
