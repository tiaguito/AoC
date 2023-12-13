package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func GCD(a, b int) int {
    for b != 0 {
        t := b
        b = a % b
        a = t
    }
    return a
}

func LCM(a, b int) int {
    return a * b / GCD(a, b)
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    lineNum := 1

    directions := ""
    m := make(map[string][]string, 0)
    itineraries := []string{}
    
    for scanner.Scan() {
        str := scanner.Text()
        if lineNum == 1 {
            directions = str
        } else {
            if len(str) > 0 {
                s := strings.Split(str, " = ")

                k := strings.Trim(s[0], " ")

                if k[2] == 'A' {
                    itineraries = append(itineraries, k)
                }

                vals := strings.Split(s[1], ",")

                m[k] = append(m[k], strings.Trim(vals[0][1:], " "))
                m[k] = append(m[k], strings.Trim(vals[1][:(len(vals[1])-1)], " "))

            }
        }
        lineNum++
    }

    steps := []int{}

    for _,v := range itineraries {
        cnt,ptr := 0,0
        for v[2] != 'Z' {
            if directions[ptr] == 'L' {
                v = m[v][0]
            } else {
                v = m[v][1]
            }
            ptr = (ptr + 1) % len(directions)
            cnt++
        }
        steps = append(steps, cnt)
    }

    res := steps[0]

    for i := 1; i < len(steps); i++ {
        res = LCM(res, steps[i])
    }

    fmt.Println(res)
}
