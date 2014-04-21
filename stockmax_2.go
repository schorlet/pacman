package main

import "fmt"

func doit(arr []int) int {
  max := 0
  sell := make([]bool, len(arr))
  for i := len(arr) - 1; i >= 0; i-- {
    if arr[i] > max {
      max = arr[i]
      sell[i] = true
    }
  }
  var shares, profit int
  for i := range arr {
    if sell[i] {
      profit += arr[i] * shares
      shares = 0
    } else {
      shares++
      profit -= arr[i]
    }
  }
  return profit
}

func main() {
  var n, t int
  fmt.Scanln(&t)
  for i := 0; i < t; i++ {
    fmt.Scan(&n)
    arr := make([]int, n)
    for j := range arr {
      fmt.Scan(&arr[j])
    }
    fmt.Println(doit(arr))
  }
}
