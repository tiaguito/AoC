package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "log"
    "strconv"
)

const max_red   =   12
const max_green =   13 
const max_blue  =   14

func isValidNumberOfCubes(cubes int, color string) bool {
    if color == "blue" && cubes > max_blue {
        return false
    } else if color == "green" && cubes > max_green {
        return false
    } else if color == "red" && cubes > max_red {
        return false
    }

    return true
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    
    gameId := 1
    sumGameIds := 0
    validGame := true

    for scanner.Scan() {
        str := scanner.Text()

        x := strings.Split(str, ": ")[1]

        l := 0
        number := 0
        color := ""

        for r := 0; r < len(x); r++ {
            if x[r] == ' ' {
                if l != r {
                    if s, err := strconv.Atoi(x[l:r]); err == nil {
                        number = s
                    } else {
                        color = x[l:r]
                        if validGame = isValidNumberOfCubes(number, color); validGame == false {
                            break;
                        }
                    }
                } 
                l = r + 1
            } else if x[r] == ';' || x[r] == ',' {
                if s, err := strconv.Atoi(x[l:r]); err == nil {
                    number = s
                } else {
                    color = x[l:r]
                    if validGame = isValidNumberOfCubes(number, color); validGame == false {
                        break;
                    }
                }

                l = r + 1
            }  else if r + 1 == len(x) {
                if s, err := strconv.Atoi(x[l:]); err == nil {
                    number = s
                } else {
                    color = x[l:]
                    if validGame = isValidNumberOfCubes(number, color); validGame == false {
                        break;
                    }
                }
            }
        }
        
        if validGame {
            sumGameIds += gameId
        } else {
            validGame = true
        }
        gameId++
    }

    fmt.Println(sumGameIds)

    if err := scanner.Err(); err != nil {
        log.Println(err)
    }
}
