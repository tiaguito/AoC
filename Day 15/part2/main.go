package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
    "strconv"
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
    var boxes = make([][]string, 256)

    sum := 0
    for _,cmd := range sequence {
        box := -1
        var sign string
        var info []string
        for i := range cmd {
            if cmd[i] == '=' || cmd[i] == '-' {
                sign = string(cmd[i])
                box = hash(cmd[:i])
                info = strings.Split(cmd, sign)
                break
            }
        }

        found := false
        i := 0
        for i = range boxes[box] {
            if strings.Contains(boxes[box][i], info[0]) {
                found = true
                break
            } 
        }

        if found {
            if sign == "-" {
                boxes[box] = append(boxes[box][:i], boxes[box][i+1:]...)
            } else {
                boxes[box][i] = info[0] + " " + info[1]
            }
        } else {
            if sign == "=" {
                boxes[box] = append(boxes[box], info[0] + " " + info[1])
            }
        }
    }

    for i := range boxes {
        for j := range boxes[i] {
            info := strings.Split(boxes[i][j], " ")
            focal_length, err := strconv.Atoi(info[1])
            if err == nil {
                power := (i + 1) * (j + 1) * focal_length
                sum += power
            }
        }
    }

    fmt.Println(sum)
}
