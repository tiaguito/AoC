package main

import (
    "fmt"
    "bufio"
    "log"
    "unicode"
    "os"
    "strconv"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    sum := 0
    for scanner.Scan() {
        str := scanner.Text()
        m, n := 0, 0
        first_found := false
        for _,v := range str {
            if unicode.IsNumber(v) {
                if first_found == false {
                    m = int(v - '0')
                    first_found = true
                }
                n = int(v - '0')
            }
        }
        num, err := strconv.Atoi(fmt.Sprintf("%v%v", m, n))
        if err != nil {
            log.Println(err)
        }
        sum += num
    }

    fmt.Println(sum)

    if err := scanner.Err(); err != nil {
        log.Println(err)
    }
}
