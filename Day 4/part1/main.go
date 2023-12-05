package main

import (
    "fmt"
    "bufio"
    "os"
    "slices"
    "math"
    "regexp"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    sum := 0

    for scanner.Scan() {
        cnt := 0

        str := scanner.Text()

        re := regexp.MustCompile(":[  ]+")
        str = re.Split(str, -1)[1]

        regex := "[ ]+" + regexp.QuoteMeta("|") + "[ ]+"
        re = regexp.MustCompile(regex)

        numbers := re.Split(str, -1)[:2]
        leftSide := numbers[0]
        rightSide := numbers[1]
        
        re = regexp.MustCompile("[ ]+")
        winningNumbers := re.Split(leftSide, -1)
        playedNumbers := re.Split(rightSide, -1)

        slices.Sort(winningNumbers)
        slices.Sort(playedNumbers)

        fmt.Println(str)
        fmt.Println(winningNumbers)
        fmt.Println(playedNumbers)

        i, j := 0, 0
        for i < len(winningNumbers) && j < len(playedNumbers){
            if winningNumbers[i] == playedNumbers[j] {
                cnt++
                j++
            } else if winningNumbers[i] < playedNumbers[j] {
                i++
            } else {
                j++ 
            }
        }
        fmt.Println(cnt)
        if cnt > 0 {
            sum += int(math.Pow(2.0, float64(cnt - 1)))
        }
    }

    fmt.Println(sum)
}
