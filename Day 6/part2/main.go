package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "regexp"
    "strconv"
)

func concatStrings(str []string) string {
    s := "" 

    for _,v := range str {
        s += v
    }
    return s
}

func stringsToNums(str []string) []int {
    nums := []int{}
    s := concatStrings(str)

    if n,err := strconv.Atoi(s); err == nil {
        nums = append(nums, n)
    }
    return nums
}

func move(x, t int) int {
    return ((t - x) * x)
}

func lowerBound(t, d int) int {
    l, r := 1, t

    for (l + 1) < r {
        m := (r + l) / 2

        if move(m, t) <= d {
            l = m
        } else {
            r = m
        }
    }
    return r
}

func upperBound(t, d int) int {
    l, r := 1, t

    for (l + 1) < r {
        m := (r + l) / 2

        if move(m, t) > d {
            l = m
        } else {
            r = m
        }
    }
    return l
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)    

    lineNum := 1
    product := 1

    time := []int{}
    distance := []int{}

    for scanner.Scan() {
        str := scanner.Text()
        re := regexp.MustCompile(":[ ]+")

        if lineNum == 1 {
            time = stringsToNums(strings.Split(re.Split(str, -1)[1], " "))
        } else {
            distance = stringsToNums(strings.Split(re.Split(str, -1)[1], " "))
        }
        lineNum++
    }

    for i := 0; i < len(time); i++ {
        left := lowerBound(time[i], distance[i])
        right := upperBound(time[i], distance[i])
        product = product * (right - left + 1)
    }

    fmt.Println(product)
}
