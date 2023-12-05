package main

import (
	"bufio"
	"fmt"
	"os"
    "log"
    "unicode"
    "strconv"
)

var dx = [8]int{-1, 0, 1, 0, -1, 1, 1, -1}
var dy = [8]int{0, 1, 0, -1, 1, -1, 1, -1}

func isSymbol(b byte) bool {
    if b == '@' || b == '#' || b =='$' || b == '%' || b == '&' ||
        b == '*' || b == '-' || b == '=' || b == '+' || b == '/' {
        return true
    }     
    return false
}

func isValid(row, col, m, n int) bool {

    if (row >= 0 && row < m) && (col >= 0 && col< n) {
        return true
    }

    return false
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    schematic := []string{}

    for scanner.Scan() {
        str := scanner.Text()

        schematic = append(schematic, str)
    }

    str := ""
    foundSymbol := false
    sum := 0

    for i := 0; i < len(schematic); i++ {
        for j := 0; j < len(schematic[i]); j++ {
            char := schematic[i][j]
            if unicode.IsNumber(rune(char)) {
                str = str + string(char)
                if foundSymbol == false {
                    for k := 0; k < 8; k++ {
                        r := i + dx[k]
                        c := j + dy[k]
                        if isValid(r, c, len(schematic), len(schematic[i])) {
                            if isSymbol(schematic[r][c]) {
                                foundSymbol = true
                            }
                        }
                    }
                }
            } else {
                if foundSymbol {
                    n,_ := strconv.Atoi(str)
                    sum += n
                }
                str = ""
                foundSymbol = false
            }
        }
    }

    fmt.Printf("Sum: %v\n", sum)

    if err := scanner.Err(); err != nil {
        log.Println(err)
    }
}
