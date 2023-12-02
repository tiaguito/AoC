package main

import (
    "fmt"
    "bufio"
    "os"
    "unicode"
    "strings"
    "log"
    "strconv"
)

func containsNum(s string) int {
    nums := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
    nums_map := map[string]int{
        "one": 1,
        "two": 2,
        "three": 3,
        "four": 4,
        "five": 5,
        "six": 6,
        "seven": 7,
        "eight": 8,
        "nine": 9,
    }
    
    for _, v := range nums {
        if strings.Contains(s, v) {
            return nums_map[v]
        }
    }

    return 0
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    sum := 0

    for scanner.Scan() {
        str := scanner.Text()
        m, n := 0, 0
        first_found := false

        l := 0
        for r := 0; r < len(str); r++ {
            if unicode.IsNumber(rune(str[r])) {
                num := int(str[r] - '0')

                if first_found == false {
                    m = num
                    first_found = true
                }
                n = num
                l = r
                continue
            }

            if num := containsNum(str[l:r+1]); num != 0 {
                if first_found == false {
                    m = num
                    first_found = true
                }
                n = num

                for num != 0 && l <= r {
                    num = containsNum(str[l+1:r+1])
                    l++
                }
            } 
        }

        calibration, err := strconv.Atoi(fmt.Sprintf("%v%v", m, n))
        if err != nil {
            log.Println(err)
        }
        sum += calibration
    }

    fmt.Printf("%v\n", sum)
}
