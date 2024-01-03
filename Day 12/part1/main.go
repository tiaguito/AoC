package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
    "log"
)

func validate(spring []int, target []int) bool {
    n := len(spring)
    curr := []int{}

    i := 0

    for i < n {
        for i < n && spring[i] == 0 {
            i += 1
        }
        if i == n {
            break
        }
        j := i
        c := 0
        for j < n && spring[j] != 0 {
            j++
            c++
        }
        curr = append(curr, c)
        i = j
    }

    if len(curr) == len(target) {
        for i := range curr {
            if curr[i] != target[i] {
                return false
            }
        }
        return true
    } else {
        return false
    }
}

func countArrangements(springs string, target []int, idx int) int {
    spring := []int{}
    idxs := []int{}
    for i, x := range springs {
        if x == '.' {
            spring = append(spring, 0)
        }
        if x == '?' {
            spring = append(spring, -1)
            idxs = append(idxs, i)
        }
        if x == '#' {
            spring = append(spring, 1)
        }
    }

    count := 0

    for mask := 0; mask < 1 << len(idxs); mask++ {
        springCopy := spring
        for i := range idxs {
            if mask & (1 << i) == 0 {
                springCopy[idxs[i]] = 0
            } else {
                springCopy[idxs[i]] = 1
            }
        }
        if validate(springCopy, target) {
            count += 1
        }
    }

    return count
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    sum := 0

    for scanner.Scan() {
        str := scanner.Text()

        coordinates := strings.Split(str, " ")
        springs := coordinates[0]
        targetStrs := strings.Split(coordinates[1], ",")
        target := []int{}

        for i := range targetStrs {
            num, err := strconv.Atoi(targetStrs[i])
            if err != nil {
                log.Panic("Error converting string number to int")
            }
            target = append(target, num)
        }

        sum += countArrangements(springs, target, 0)
    }
    fmt.Println(sum)
}
