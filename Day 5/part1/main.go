package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
    "slices"
)

func stringToNums(str string) []int {
    nums := []int{}
    strs := strings.Split(str, " ")

    for _,v := range strs {
        if n,err := strconv.Atoi(v); err == nil {
            nums = append(nums, n)
        }
    }

    return nums
}

func seedToSoil(nums *[]int, maps *[][]int) {
    for i := 0; i < len(*nums); i++ {
        for _,v := range *maps {
            if ((*nums)[i] >= v[1]) &&
                ((*nums)[i] <= v[1] + v[2] - 1) {
                    diff := (*nums)[i] - v[1]
                    (*nums)[i] = v[0] + diff
                    break
            }
        }
    }
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    line := 1
    block := 1
    nums := []int{}
    maps := make([][]int, 0)

    for scanner.Scan() {
        str := scanner.Text()

        if block == 1 && line == 1 {
            s := strings.Split(str, ": ")[1]
            nums = stringToNums(s)
        } else if block > 1{
            if line > 1 && len(str) != 0 {
                maps = append(maps, stringToNums(str))
            } 
        }

        line++

        if len(str) == 0 {
            if block > 1 && line > 1 {
                seedToSoil(&nums, &maps)
                maps = make([][]int, 0)
            }

            block++
            line = 1
        }
    }

    seedToSoil(&nums, &maps)

    minVal := slices.Min(nums)

    fmt.Println(minVal)
}
