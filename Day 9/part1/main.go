package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)

func stringToNums(str string) []int {
    nums := []int{}
    s := strings.Split(str, " ")

    for _,v := range s {
        if n,err := strconv.Atoi(v); err == nil {
            nums = append(nums, n)
        }
    }

    return nums
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    sum := 0
    
    for scanner.Scan() {
        str := scanner.Text()

        matrix := [][]int{}

        matrix = append(matrix, stringToNums(str))

        for i := 0; i < len(matrix); i++ {
            matrix = append(matrix, []int{})
            for j := 1; j < len(matrix[i]); j++ {
                matrix[i+1] = append(matrix[i+1], matrix[i][j] - matrix[i][j-1])
            }
            allZeros := true

            m := len(matrix) - 1
            for i := 0; i < len(matrix[m]); i++ {
                if matrix[m][i] != 0 {
                    allZeros = false
                    break
                }
            }

            if allZeros {
                break
            }
        }

        for i := len(matrix) - 2; i >= 0; i-- {
            nCurr := len(matrix[i]) - 1
            nNext := len(matrix[i+1]) - 1

            matrix[i] = append(matrix[i], matrix[i][nCurr] + matrix[i+1][nNext])
        }

        sum = sum + matrix[0][len(matrix[0])-1]
    }
    fmt.Println(sum)
}
