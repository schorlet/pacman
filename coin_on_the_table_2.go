package main

import (
    "fmt"
    "strings"
)

// input takes whitespace seperated values from STDIN and returns a matrix and stepMax.
func input() (matrix [][]string, stepMax int) {
    // fmt.Print("Enter the size of n×m matrix and the step limit: ")
    var n, m int
    _, err := fmt.Scanf("%d %d %d", &n, &m, &stepMax)
    if err != nil {
        panic(err)
    }

    // Initialize a 2D slice of n rows and m columns.
    matrix = make([][]string, n)
    for i, _ := range matrix {
        matrix[i] = make([]string, m)
    }

    // fmt.Print("Enter the matrix strings: ")
    var direction string
    for i := 0; i < n; i++ {
        fmt.Scan(&direction)
        directionList := strings.Split(direction, "")
        for j := 0; j < m; j++ {
            matrix[i][j] = directionList[j]
        }
    }

    return matrix, stepMax
}

func main() {
    matrix, stepMax := input()
    nMax, mMax := len(matrix), len(matrix[0])
    nEnd, mEnd := findEnd(nMax, mMax, matrix)

    costMatrix := make([][]int, len(matrix))
    // Initialize all the inner slices.
    for i := range matrix {
        costMatrix[i] = make([]int, len(matrix[0]))
    }
    for i := range costMatrix {
        for j := range costMatrix[0] {
            costMatrix[i][j] = 1000000
        }
    }

    minPath(0, nMax, 0, mMax, 0, stepMax, 0, matrix[0][0], matrix[0][0], &matrix, &costMatrix)

    if costMatrix[nEnd][mEnd] == 1000000 {
        fmt.Println(-1)
    } else {
        fmt.Println(costMatrix[nEnd][mEnd])
    }
}

func findEnd(nMax, mMax int, matrix [][]string) (nEnd, mEnd int) {
    for i := 0; i < nMax; i++ {
        for j := 0; j < mMax; j++ {
            if matrix[i][j] == "*" {
                return i, j
            }
        }
    }
    return -1, -1
}

// Change tally to a pointer.
func minPath(n, nMax, m, mMax, step, stepMax, opCount int, element, matrixElement string, matrix *[][]string, costMatrix *[][]int) {
    if n >= 0 && n < nMax && m >= 0 && m < mMax && step <= stepMax && opCount < (*costMatrix)[n][m] {
        if element != matrixElement {
            opCount++
        }

        (*costMatrix)[n][m] = opCount

        minPath(n+1, nMax, m, mMax, step+1, stepMax, opCount, "D", (*matrix)[n][m], matrix, costMatrix)
        minPath(n, nMax, m+1, mMax, step+1, stepMax, opCount, "R", (*matrix)[n][m], matrix, costMatrix)
        minPath(n-1, nMax, m, mMax, step+1, stepMax, opCount, "U", (*matrix)[n][m], matrix, costMatrix)
        minPath(n, nMax, m-1, mMax, step+1, stepMax, opCount, "L", (*matrix)[n][m], matrix, costMatrix)
    }
}
