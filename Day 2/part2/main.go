package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const max_red   =   12
const max_green =   13 
const max_blue  =   14

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    
    gameId := 1
    sumPowers := 0

    min_red, min_green, min_blue := 1, 1, 1

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

                        if color == "red" {
                            min_red = int(math.Max(float64(min_red), float64(number)))
                        } else if color == "green" {
                            min_green = int(math.Max(float64(min_green), float64(number)))
                        } else if color == "blue" {
                            min_blue = int(math.Max(float64(min_blue), float64(number)))
                        } 

                    }
                } 
                l = r + 1
            } else if x[r] == ';' || x[r] == ',' {
                if s, err := strconv.Atoi(x[l:r]); err == nil {
                    number = s
                } else {
                    color = x[l:r]
                    
                    if color == "red" {
                        min_red = int(math.Max(float64(min_red), float64(number)))
                    } else if color == "green" {
                        min_green = int(math.Max(float64(min_green), float64(number)))
                    } else if color == "blue" {
                        min_blue = int(math.Max(float64(min_blue), float64(number)))
                    } 
                }

                l = r + 1
            }  else if r + 1 == len(x) {
                if s, err := strconv.Atoi(x[l:]); err == nil {
                    number = s
                } else {
                    color = x[l:]

                    if color == "red" {
                        min_red = int(math.Max(float64(min_red), float64(number)))
                    } else if color == "green" {
                        min_green = int(math.Max(float64(min_green), float64(number)))
                    } else if color == "blue" {
                        min_blue = int(math.Max(float64(min_blue), float64(number)))
                    } 
                }
            }
        }
        
        sumPowers += (min_red * min_green * min_blue)
        min_red, min_green, min_blue = 1, 1, 1
        gameId++
    }

    fmt.Println(sumPowers)

    if err := scanner.Err(); err != nil {
        log.Println(err)
    }
}

