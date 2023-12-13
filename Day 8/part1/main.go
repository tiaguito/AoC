package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    lineNum := 1

    m := make(map[string][]string, 0)
    
    directions := ""
    for scanner.Scan() {
        str := scanner.Text()
        
        if lineNum == 1 {
            directions = str
        } else {
            if len(str) > 0 {
                strs := strings.Split(str, " = ")

                k := strs[0]

                values := strings.Split(strs[1], ", ")
                
                m[k] = append(m[k], values[0][1:])
                m[k] = append(m[k], values[1][:len(values[1])-1])
            }
        }

        lineNum++
    }

    key := "AAA"
    ptr,ctr := 0,0

    for key != "ZZZ" {
        if directions[ptr] == 'L' {
            key = m[key][0]
        } else {
            key = m[key][1]
        }

        ctr++
        ptr = (ptr + 1) % len(directions)
    }

    fmt.Println(ctr)
}
