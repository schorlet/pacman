package main

import "fmt"
import "math"

func main() {
    var message string
    fmt.Scanf("%s", &message)

    var length = len(message)
    var row = int(math.Sqrt(float64(length)))
    var col = row
    if row * col < length {
        col += 1
        if row * col < length {
            row += 1
        }
    }
    var i, j int
    // var rectangle = make([][]int, row)
    // for j = 0; j < row; j++ {
        // rectangle[j] = make([]int, col)
        // for i = 0; i < col; i++ {
            // if j * col + i < length {
                // rectangle[j][i] = int(message[j * col + i])
            // }
        // }
    // }
    // for i = 0; i < col; i++ {
        // for j = 0; j < row; j++ {
            // if rectangle[j][i] > 0 {
                // fmt.Printf("%c", rectangle[j][i])
            // }
        // }
        // fmt.Print(" ")
    // }
    // fmt.Println()

    var encoded string
    for i = 0; i < col; i++ {
        for j = 0; j < row; j++ {
            if j * col + i < length {
                encoded += fmt.Sprintf("%c", message[j * col + i])
            }
        }
        encoded += " "
    }
    fmt.Println(encoded)

}
