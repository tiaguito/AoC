package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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
    for i := 0; i < len(*nums); i=i+2 {
        for _,v := range *maps {
            a := v[1]
            b := v[1] + v[2] - 1
            c := (*nums)[i]
            d := (*nums)[i] + (*nums)[i+1] - 1

            if (c >= a && c <= b) && (d <= b) {
                diff := (*nums)[i] - v[1]
                (*nums)[i] = v[0] + diff
                break
            } else if (c >= a && c <= b) && (d > b) {
                diff := d - b + 1
                *nums = append(*nums, b + 1)
                *nums = append(*nums, diff)

                (*nums)[i+1] = b - (*nums)[i]
                diff2 := (*nums)[i] - v[1]
                (*nums)[i] = v[0] + diff2
                break
            } else if (c < a) && (d >= a && d <= b) {
                diff := a - c
                *nums = append(*nums, c)
                *nums = append(*nums, diff)

                (*nums)[i+1] = d - a
                (*nums)[i] = v[0]
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

    minVal := math.MaxInt

    for i := 0; i < len(nums); i=i+2 {
        if nums[i] < minVal {
            minVal = nums[i]
        }
    }

    fmt.Println(minVal)
}
