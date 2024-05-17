package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func hash(cmd string) int {
    curr_sum := 0
    for _,c := range cmd {
        ascii := int(c)
        curr_sum = ((curr_sum + ascii) * 17) % 256
    }
    return curr_sum
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    scanner.Scan()
    str := scanner.Text()
    sequence := strings.Split(str, ",") 

    sum := 0
    for _,cmd := range sequence {
        sum += hash(cmd)
    }

    fmt.Println(sum)
}
