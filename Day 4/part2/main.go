package main

import (
    "fmt"
    "bufio"
    "os"
    "slices"
    "regexp"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    sum := 0

    cardId := 1

    m := make(map[int]int, 0)

    for scanner.Scan() {
        m[cardId]++

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

        for i := cardId + 1; i <= cardId + cnt; i++ {
            m[i] += m[cardId]
        }

        cardId++
    }

    for _,v := range m {
        sum += v
    }

    fmt.Println(sum)
}
